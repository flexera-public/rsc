package rsapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Authenticator interface
type Authenticator interface {
	// Sign signs the given http Request (adds the auth headers).
	Sign(req *http.Request, host string, accountID int) error
	// ResolveHost returns the RightScale API endpoint for the given account.
	ResolveHost(host string, accountID int) (string, error)
}

// NewBasicAuthenticator returns a authenticator that uses email and password to create sessions
func NewBasicAuthenticator(username, password string) Authenticator {
	builder := basicLoginRequestBuilder{username: username, password: password}
	return newSessionManager(newCookieSigner(&builder))
}

// NewOAuthAuthenticator returns a authenticator that uses a oauth refresh token to create access
// tokens.
// The refresh token can be found in the CM dashboard under Settings > Account Settings > API Credentials.
func NewOAuthAuthenticator(token string) Authenticator {
	return newSessionManager(newOAuthSigner(token))
}

// NewTokenAuthenticator returns a authenticator that use a oauth access token to do authentication
// This is useful if the oauth handshake has already happened. Use the OAuthAuthenticator to use
// a refresh token and have the authenticator do the handshake.
func NewTokenAuthenticator(token string) Authenticator {
	return &tokenAuthenticator{token: token}
}

// NewInstanceAuthenticator returns an authenticator that uses the instance facing API token to
// create sessions. This is the token found on RightLink instances under the RS_API_TOKEN
// environment variable.
// Note: Use of rsc made from RightLink10 instances can use the RL10 authenticator instead.
func NewInstanceAuthenticator(token string) Authenticator {
	builder := instanceLoginRequestBuilder{token: token}
	return newSessionManager(newCookieSigner(&builder))
}

// NewSSAuthenticator returns an authenticator that wraps another one and adds the logic needed to
// create sessions in Self-Service.
func NewSSAuthenticator(auther Authenticator) Authenticator {
	return &ssAuthenticator{
		auther:    auther,
		refreshAt: time.Now().Add(-2 * time.Minute),
		client:    http.DefaultClient,
	}
}

// NewRL10Authenticator returns an authenticator that proxies all requests through the RightLink 10
// agent.
func NewRL10Authenticator(secret string) Authenticator {
	return &rl10Authenticator{secret: secret}
}

// newSessionManager returns a session manager that takes care of creating and refreshing sessions
// as needed. It implements the Authenticator interface.
// The signer does the actual work of creating sessions and checking session freshness.
// Different signers may be plugged in for different types of sessions.
func newSessionManager(signer sessionSigner) *sessionManager {
	return &sessionManager{signer: signer}
}

// sessionManager contains the logic to create a session and maintain it.
type sessionManager struct {
	signer sessionSigner // Session creator / signer
}

// Sign signs the given http Request (adds the auth headers).
func (m *sessionManager) Sign(r *http.Request, host string, accountID int) error {
	return m.signer.Sign(r, host, accountID)
}

// ResolveHost returns the RightScale API endpoint for the given account.
// It does that by catching any redirect returned by the API when attempting to use the provided
// host.
func (m *sessionManager) ResolveHost(host string, accountID int) (string, error) {
	authReq, authErr := m.signer.BuildLoginRequest(host, accountID)
	if authErr != nil {
		return host, authErr
	}
	redirectHost := host
	client := http.Client{
		CheckRedirect: func(nextRequest *http.Request, via []*http.Request) error {
			viaCount := len(via)
			if viaCount > 10 {
				return fmt.Errorf("Too many redirects.")
			}
			lastRequest := via[viaCount-1]
			if nextRequest.Header == nil {
				nextRequest.Header = make(http.Header)
			}
			for k, v := range lastRequest.Header {
				k = http.CanonicalHeaderKey(k)
				if !omitHeaders[k] {
					nextRequest.Header[k] = v
				}
			}
			redirectHost = nextRequest.URL.Host

			return nil
		},
	}
	_, err := client.Do(authReq)
	return redirectHost, err
}

// sessionSigner contains the logic to create and refresh sessions.
type sessionSigner interface {
	BuildLoginRequest(host string, accountID int) (*http.Request, error)
	Sign(req *http.Request, host string, accountID int) error
}

// cookieSigner signs requests using adding a global session cookie.
// Used by both the basic and instance session managers.
type cookieSigner struct {
	builder   loginRequestBuilder
	cookies   map[int][]*http.Cookie // cookied indexed by account id
	refreshAt time.Time
	client    HttpClient
}

// newCookieSigner returns a new cookie signer
func newCookieSigner(builder loginRequestBuilder) sessionSigner {
	return &cookieSigner{
		builder:   builder,
		cookies:   make(map[int][]*http.Cookie),
		refreshAt: time.Now().Add(-2 * time.Minute),
		client:    http.DefaultClient,
	}
}

// BuildLoginRequest returns a http.Request that creates a new session
func (r *cookieSigner) BuildLoginRequest(host string, accountID int) (*http.Request, error) {
	return r.builder.BuildLoginRequest(host, accountID)
}

// Sign adds the username and password authorization cookies to the *http.Request
// Checks the freshness of the session and creates a new one if needed.
func (r *cookieSigner) Sign(req *http.Request, host string, accountID int) error {
	if time.Now().After(r.refreshAt) {
		authReq, authErr := r.builder.BuildLoginRequest(host, accountID)
		if authErr != nil {
			return authErr
		}
		resp, err := r.client.Do(authReq)
		if err != nil {
			return fmt.Errorf("Authentication failed: %s", err)
		}
		if resp.StatusCode != 204 {
			return fmt.Errorf("Authentication failed: %s", resp.Status)
		}
		r.cookies[accountID] = resp.Cookies()
		r.refreshAt = time.Now().Add(time.Duration(2) * time.Hour)
	}
	for _, c := range r.cookies[accountID] {
		req.AddCookie(c)
	}
	return nil
}

// oAuthSigner contains the logic to create new session using OAuth tokens
type oAuthSigner struct {
	refreshToken string
	accessToken  string
	refreshAt    time.Time
	client       HttpClient
}

// newOAuthSigner returns a new oauth signer
func newOAuthSigner(token string) sessionSigner {
	return &oAuthSigner{
		refreshToken: token,
		refreshAt:    time.Now().Add(-2 * time.Minute),
		client:       http.DefaultClient,
	}
}

// Sign adds the OAuth bearer header to the *http.Request
func (r *oAuthSigner) Sign(req *http.Request, host string, accountID int) error {
	if time.Now().After(r.refreshAt) {
		authReq, authErr := r.BuildLoginRequest(host, accountID)
		if authErr != nil {
			return fmt.Errorf("Authentication failed: %s", authErr)
		}
		resp, err := r.client.Do(authReq)
		if err != nil {
			return fmt.Errorf("Authentication failed: %s", err) // TBD RETRY A FEW TIMES
		}
		defer resp.Body.Close()
		var session map[string]interface{}
		jsonBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("Authentication failed (failed to read response): %s", err)
		}
		if resp.StatusCode != 200 {
			return fmt.Errorf("Authentication failed: %s", resp.Status)
		}
		err = json.Unmarshal(jsonBytes, &session)
		if err != nil {
			return fmt.Errorf("Authentication failed (failed to load response JSON): %s", err)
		}
		accessToken, ok := session["access_token"].(string)
		if !ok {
			return fmt.Errorf("Unexpected auth response: %s", jsonBytes)
		}
		r.accessToken = accessToken
		d, err := time.ParseDuration(fmt.Sprintf("%vs", session["expires_in"]))
		if err != nil {
			return fmt.Errorf("Authentication failed (failed to parse token duration): %s", err)
		}
		r.refreshAt = time.Now().Add(d / 2)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", r.accessToken))
	return nil
}

// BuildLoginRequest returns a new *http.Request that can refresh the access token
func (r *oAuthSigner) BuildLoginRequest(host string, _ int) (*http.Request, error) {
	jsonStr := fmt.Sprintf(`{"grant_type":"refresh_token","refresh_token":"%s"}`, r.refreshToken)
	authReq, err := http.NewRequest("POST", endpoint(host, "api/oauth2"), bytes.NewBufferString(jsonStr))
	if err != nil {
		return nil, fmt.Errorf("Authentication failed (failed to build request): %s", err)
	}
	authReq.Header.Set("X-API-Version", "1.5")
	authReq.Header.Set("Content-Type", "application/json")
	return authReq, nil
}

// OAuth access token authenticator
type tokenAuthenticator struct {
	token string
}

// Sign sets the OAuth authorization header
func (t *tokenAuthenticator) Sign(r *http.Request, host string, accountID int) error {
	r.Header.Set("Authorization", "Bearer "+t.token)
	return nil
}

// Can't resolve host using an access token, host must be correct
func (t *tokenAuthenticator) ResolveHost(host string, accountID int) (string, error) {
	return host, nil
}

// RightLink 10 authenticator
type rl10Authenticator struct {
	secret string
}

// RL10 authenticator uses special header
func (s *rl10Authenticator) Sign(r *http.Request, host string, accountID int) error {
	r.Header.Set("X-RLL-Secret", s.secret)
	return nil
}

// A stub to satisfy the interface. RL10 will always use host "localhost" and
// will not be redirected
func (a *rl10Authenticator) ResolveHost(host string, accountID int) (string, error) {
	return host, nil
}

// SS authenticator
type ssAuthenticator struct {
	auther    Authenticator // Authentication against core
	refreshAt time.Time
	client    HttpClient
}

// Account authenticator uses RS oauth
func (a *ssAuthenticator) Sign(r *http.Request, host string, accountID int) error {
	if time.Now().After(a.refreshAt) {
		authReq, err := http.NewRequest("GET",
			endpoint(ssHostFromLogin(host),
				fmt.Sprintf("api/catalog/new_session?account_id=%d", accountID)),
			nil)
		if err != nil {
			return err
		}
		a.auther.Sign(authReq, host, accountID)
		authReq.Header.Set("Content-Type", "application/json")
		_, err = a.client.Do(authReq)
		if err != nil {
			return fmt.Errorf("Authentication failed: %s", err)
		}
	}
	a.auther.Sign(r, host, accountID)
	r.Header.Set("X-Api-Version", "1.0")
	r.Host = ssHostFromLogin(host)

	return nil
}

// Determine the correct SS host based on the CM/Core host. Shard redirection
// will already have occurred since the rsapi.Api for the core has already been
// initialized and (if necessary) redirected.
func (a *ssAuthenticator) ResolveHost(host string, accountID int) (string, error) {
	core, err := a.auther.ResolveHost(host, accountID)
	if err != nil {
		return "", err
	}
	return ssHostFromLogin(core), nil
}

// Return Self-service endpoint from login endpoint
// The following isn't great but seems better than having to enter by hand
func ssHostFromLogin(host string) string {
	urlElems := strings.Split(host, ".")
	hostPrefix := urlElems[0]
	elems := strings.Split(hostPrefix, "-")
	if len(elems) < 2 {
		return host
	}
	elems[len(elems)-2] = "selfservice"
	ssLoginHostPrefix := strings.Join(elems, "-")
	return strings.Join(append([]string{ssLoginHostPrefix}, urlElems[1:]...), ".")
}

// loginRequestBuilder is a generic login request factory.
type loginRequestBuilder interface {
	BuildLoginRequest(host string, accountID int) (*http.Request, error)
}

// basicLoginRequestBuilder builds login requests from users email and password.
type basicLoginRequestBuilder struct {
	username string
	password string
}

// BuildLoginRequest builds session create requests from users email and password.
func (b *basicLoginRequestBuilder) BuildLoginRequest(host string, accountID int) (*http.Request, error) {
	jsonStr := fmt.Sprintf(`{"email":"%s","password":"%s","account_href":"/api/accounts/%d"}`,
		b.username, b.password, accountID)
	authReq, err := http.NewRequest("POST", endpoint(host, "api/sessions"),
		bytes.NewBufferString(jsonStr))
	if err != nil {
		return authReq, fmt.Errorf("Authentication failed (failed to build request): %s", err.Error())
	}
	authReq.Header.Set("X-API-Version", "1.5")
	authReq.Header.Set("Content-Type", "application/json")
	return authReq, nil
}

// instanceLoginRequestBuilder builds session create requests from instance API tokens.
type instanceLoginRequestBuilder struct {
	token string
}

// BuildLoginRequest builds session create requests from users email and password.
func (b *instanceLoginRequestBuilder) BuildLoginRequest(host string, accountID int) (*http.Request, error) {
	accountHref := fmt.Sprintf("/api/accounts/%d", accountID)
	jsonStr := fmt.Sprintf(`{"instance_token":"%s","account_href":"%s"}`, b.token, accountHref)
	authReq, err := http.NewRequest("POST", endpoint(host, "api/session/instance"), bytes.NewBufferString(jsonStr))
	if err != nil {
		return nil, fmt.Errorf("Authentication failed (failed to build request): %s", err)
	}
	authReq.Header.Set("X-API-Version", "1.5")
	authReq.Header.Set("Content-Type", "application/json")
	return authReq, nil
}

// Compute API endpoint given a hostname and a path
func endpoint(host, suffix string) string {
	if !strings.HasPrefix(host, "http") {
		// Be nice to tests
		if !strings.HasPrefix(host, "localhost") && !strings.HasPrefix(host, "127.0.0.1") {
			host = "https://" + host
		} else {
			host = "http://" + host
		}
	}
	if !strings.HasSuffix(host, "/") {
		host += "/"
	}
	return host + suffix
}

// Headers that should be copied when creating the redirect request
var omitHeaders map[string]bool = map[string]bool{
	"Content-Type":   true,
	"Content-Length": true,
}
