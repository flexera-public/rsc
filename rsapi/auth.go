package rsapi // import "gopkg.in/rightscale/rsc.v2/rsapi"

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

// Authenticator interface
type Authenticator interface {
	// Sign signs the given http Request (adds the auth headers).
	Sign(req *http.Request) error
	// SetHost updates the host used by the authenticator to create sessions.
	// This method is called internally by the various API clients upon creation.
	SetHost(host string)
	// CanAuthenticate returns nil if the authenticator is able to sign requests successfully
	// or an error with additional information otherwise.
	// It makes a test request to CM 1.5 to validate the provided credentials.
	CanAuthenticate() error
}

// NewBasicAuthenticator returns a authenticator that uses email and password to create sessions.
// The returned authenticator takes care of refreshing the RightScale session as needed.
func NewBasicAuthenticator(username, password string, accountID int) Authenticator {
	builder := basicLoginRequestBuilder{username: username, password: password, accountID: accountID}
	return newCookieSigner(&builder, accountID)
}

// NewInstanceAuthenticator returns an authenticator that uses the instance facing API token to
// create sessions. This is the token found on RightLink instances under the RS_API_TOKEN
// environment variable.
// The returned authenticator takes care of refreshing the RightScale session as needed.
// Note: Use of rsc made from RightLink10 instances can use the RL10 authenticator instead.
func NewInstanceAuthenticator(token string, accountID int) Authenticator {
	builder := instanceLoginRequestBuilder{token: token, accountID: accountID}
	return newCookieSigner(&builder, accountID)
}

// NewOAuthAuthenticator returns a authenticator that uses a oauth refresh token to create access
// tokens.
// The refresh token can be found in the CM dashboard under Settings > Account Settings > API Credentials.
func NewOAuthAuthenticator(token string) Authenticator {
	return &oAuthSigner{
		refreshToken: token,
		refreshAt:    time.Now().Add(-2 * time.Minute),
		client:       http.DefaultClient,
	}
}

// NewTokenAuthenticator returns a authenticator that use a oauth access token to do authentication.
// This is useful if the oauth handshake has already happened.
// Use the OAuthAuthenticator to use a refresh token and have the authenticator do the handshake.
func NewTokenAuthenticator(token string) Authenticator {
	return &tokenAuthenticator{token: token}
}

// NewSSAuthenticator returns an authenticator that wraps another one and adds the logic needed to
// create sessions in Self-Service.
func NewSSAuthenticator(auther Authenticator, accountID int) Authenticator {
	if _, ok := auther.(*ssAuthenticator); ok {
		// Only wrap if not wrapped already
		return auther
	}
	return &ssAuthenticator{
		auther:    auther,
		accountID: accountID,
		refreshAt: time.Now().Add(-2 * time.Minute),
		client:    http.DefaultClient,
	}
}

// NewRL10Authenticator returns an authenticator that proxies all requests through the RightLink 10
// agent.
func NewRL10Authenticator(secret string) Authenticator {
	return &rl10Authenticator{secret: secret}
}

// loginRequestBuilder is a generic login request factory.
type loginRequestBuilder interface {
	BuildLoginRequest(host string) (*http.Request, error)
}

// cookieSigner signs requests using adding a global session cookie.
// Used by both the basic and instance authenticators.
type cookieSigner struct {
	builder   loginRequestBuilder
	accountID int
	cookies   []*http.Cookie
	host      string
	refreshAt time.Time
	client    HttpClient
}

// newCookieSigner returns a cookie signer that uses the given builder to build login requests.
func newCookieSigner(builder loginRequestBuilder, accountID int) Authenticator {
	return &cookieSigner{
		builder:   builder,
		accountID: accountID,
		refreshAt: time.Now().Add(-2 * time.Minute),
		client:    http.DefaultClient,
	}
}

// Sign adds the username and password authorization cookies to the request.
// Checks the freshness of the session and creates a new one if needed.
func (s *cookieSigner) Sign(req *http.Request) error {
	if time.Now().After(s.refreshAt) {
		authReq, authErr := s.builder.BuildLoginRequest(s.host)
		if authErr != nil {
			return authErr
		}
		resp, err := s.client.Do(authReq)
		if err != nil {
			return fmt.Errorf("Authentication failed: %s", err)
		}
		if err := s.refresh(resp); err != nil {
			return err
		}
	}
	for _, c := range s.cookies {
		req.AddCookie(c)
	}
	return nil
}

// SetHost sets the host used to create sessions.
func (s *cookieSigner) SetHost(host string) {
	s.host = host
}

// CanAuthenticate makes a test request to CM 1.5 and returns true if it is successful.
func (s *cookieSigner) CanAuthenticate() error {
	// A cookie signer is able to resolve the host an account belongs to by
	// levaraging the redirect response sent by the API when creating a new session.
	// So first resolve the host if the signer doesn't have one already.
	if s.host == "" {
		h, r, err := resolveHost(s.accountID, s.builder)
		if err != nil {
			return err
		}
		if err := s.refresh(r); err != nil {
			return err
		}
		s.host = h
	}
	// Now that we have a host actually test the credentials.
	_, instance := s.builder.(*instanceLoginRequestBuilder)
	return testAuth(s, s.host, instance)
}

// refresh updates the cookie and expiration used to sign requests from a successful session
// creation API response.
func (s *cookieSigner) refresh(resp *http.Response) error {
	if resp.StatusCode != 204 {
		return fmt.Errorf("Authentication failed: %s", resp.Status)
	}
	s.cookies = resp.Cookies()
	s.refreshAt = time.Now().Add(time.Duration(2) * time.Hour)
	return nil
}

// oAuthSigner contains the logic to create new session using OAuth tokens
type oAuthSigner struct {
	refreshToken string
	accessToken  string
	host         string
	refreshAt    time.Time
	client       HttpClient
}

// Sign adds the OAuth bearer header to the *http.Request
func (s *oAuthSigner) Sign(req *http.Request) error {
	if time.Now().After(s.refreshAt) {
		authReq, authErr := s.BuildLoginRequest(s.host)
		if authErr != nil {
			return fmt.Errorf("Authentication failed: %s", authErr)
		}
		resp, err := s.client.Do(authReq)
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
		s.accessToken = accessToken
		d, err := time.ParseDuration(fmt.Sprintf("%vs", session["expires_in"]))
		if err != nil {
			return fmt.Errorf("Authentication failed (failed to parse token duration): %s", err)
		}
		s.refreshAt = time.Now().Add(d / 2)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", s.accessToken))
	return nil
}

// SetHost sets the host used to create sessions.
func (s *oAuthSigner) SetHost(host string) {
	s.host = host
}

// CanAuthenticate makes a test request to CM 1.5 and returns nil if it is successful.
func (s *oAuthSigner) CanAuthenticate() error {
	return testAuth(s, s.host, false)
}

// BuildLoginRequest returns a new *http.Request that can refresh the access token
func (s *oAuthSigner) BuildLoginRequest(host string) (*http.Request, error) {
	jsonStr := fmt.Sprintf(`{"grant_type":"refresh_token","refresh_token":"%s"}`, s.refreshToken)
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
	host  string // Only used by CanAuthenticate
}

// Sign sets the OAuth authorization header
func (t *tokenAuthenticator) Sign(r *http.Request) error {
	r.Header.Set("Authorization", "Bearer "+t.token)
	return nil
}

// SetHost is not used by the token authenticator as it does not actually create sessions.
func (t *tokenAuthenticator) SetHost(h string) {
	t.host = h
}

// CanAuthenticate makes a test request to CM 1.5 and returns nil if it is successful.
func (t *tokenAuthenticator) CanAuthenticate() error {
	return testAuth(t, t.host, false)
}

// RightLink 10 authenticator
type rl10Authenticator struct {
	secret string
	host   string
}

// RL10 authenticator uses special header
func (s *rl10Authenticator) Sign(r *http.Request) error {
	r.Header.Set("X-RLL-Secret", s.secret)
	return nil
}

// SetHost sets the host used to proxy requests.
func (a *rl10Authenticator) SetHost(h string) {
	a.host = h
}

// CanAuthenticate makes a test request to CM 1.5 and returns nil if it is successful.
func (a *rl10Authenticator) CanAuthenticate() error {
	return testAuth(a, a.host, true)
}

// SS authenticator
type ssAuthenticator struct {
	auther    Authenticator // Authentication against core
	accountID int           // Account used to create SS local session
	host      string        // Login (core) host
	refreshAt time.Time     // SS local session refresh deadline
	client    HttpClient
}

// Account authenticator uses RS oauth
func (a *ssAuthenticator) Sign(r *http.Request) error {
	if time.Now().After(a.refreshAt) {
		authReq, err := http.NewRequest("GET",
			endpoint(a.host,
				fmt.Sprintf("api/catalog/new_session?account_id=%d", a.accountID)),
			nil)
		if err != nil {
			return err
		}
		a.auther.Sign(authReq)
		authReq.Header.Set("Content-Type", "application/json")
		_, err = a.client.Do(authReq)
		if err != nil {
			return fmt.Errorf("Authentication failed: %s", err)
		}
		a.refreshAt = time.Now().Add(2 * time.Hour)
	}
	a.auther.Sign(r)
	r.Header.Set("X-Api-Version", "1.0")
	r.Host = a.host

	return nil
}

// SetHost sets the host used to create Self-Service session.
// Pass in the CM 1.5 host, this method computes the Self-Service host from it.
func (a *ssAuthenticator) SetHost(host string) {
	a.auther.SetHost(host)
	urlElems := strings.Split(host, ".")
	hostPrefix := urlElems[0]
	elems := strings.Split(hostPrefix, "-")
	if len(elems) < 2 {
		a.host = host
		return
	}
	elems[len(elems)-2] = "selfservice"
	ssLoginHostPrefix := strings.Join(elems, "-")
	a.host = strings.Join(append([]string{ssLoginHostPrefix}, urlElems[1:]...), ".")
}

// CanAuthenticate makes a test request to SS and returns true if it is successful.
func (a *ssAuthenticator) CanAuthenticate() error {
	if a.host == "" {
		return fmt.Errorf("missing host information")
	}
	url := fmt.Sprintf("api/catalog/accounts/%d/user_preferences", a.accountID)
	req, err := http.NewRequest("GET", endpoint(a.host, url), nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-Api-Version", "1.0")
	a.Sign(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		var body string
		if b, err := ioutil.ReadAll(resp.Body); err != nil {
			body = string(b)
		}
		return fmt.Errorf("%s: %s", resp.Status, body)
	}
	return nil
}

// basicLoginRequestBuilder builds login requests from users email and password.
type basicLoginRequestBuilder struct {
	username  string
	password  string
	accountID int
}

// BuildLoginRequest builds session create requests from users email and password.
func (b *basicLoginRequestBuilder) BuildLoginRequest(host string) (*http.Request, error) {
	if host == "" {
		host = "us-3.rightscale.com"
	}
	jsonStr := fmt.Sprintf(`{"email":"%s","password":"%s","account_href":"/api/accounts/%d"}`,
		b.username, b.password, b.accountID)
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
	token     string
	accountID int
}

// BuildLoginRequest builds session create requests from users email and password.
func (b *instanceLoginRequestBuilder) BuildLoginRequest(host string) (*http.Request, error) {
	if host == "" {
		host = "us-3.rightscale.com"
	}
	accountHref := fmt.Sprintf("/api/accounts/%d", b.accountID)
	jsonStr := fmt.Sprintf(`{"instance_token":"%s","account_href":"%s"}`, b.token, accountHref)
	authReq, err := http.NewRequest("POST", endpoint(host, "api/session/instance"), bytes.NewBufferString(jsonStr))
	if err != nil {
		return nil, fmt.Errorf("Authentication failed (failed to build request): %s", err)
	}
	authReq.Header.Set("X-API-Version", "1.5")
	authReq.Header.Set("Content-Type", "application/json")
	byt, _ := httputil.DumpRequest(authReq, true)
	fmt.Printf(string(byt))
	return authReq, nil
}

// resolveHost returns the RightScale API endpoint for the given account.
// It does that by catching any redirect returned by the API when attempting to use the provided
// host.
func resolveHost(accountID int, b loginRequestBuilder) (string, *http.Response, error) {
	authReq, authErr := b.BuildLoginRequest("us-3.rightscale.com")
	if authErr != nil {
		return "", nil, authErr
	}
	var redirectHost string
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
	redirectHost = authReq.Host
	resp, err := client.Do(authReq)
	if err == nil && resp.StatusCode > 299 {
		err = fmt.Errorf(resp.Status)
	}
	return redirectHost, resp, err
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

// testAuth makes a GET /api/sessions CM 1.5 request using the given authenticator and returns
// an error if it failed, nil otherwise.
// The instance flag specifies whether an instance or an account facing API request should be made.
func testAuth(auth Authenticator, host string, instance bool) error {
	if host == "" {
		return fmt.Errorf("missing host information")
	}
	var req *http.Request
	var err error
	if instance {
		req, err = http.NewRequest("GET", endpoint(host, "api/user_data"), nil)
	} else {
		req, err = http.NewRequest("GET", endpoint(host, "api/sessions"), nil)
	}
	if err != nil {
		return err
	}
	req.Header.Set("X-Api-Version", "1.5")
	auth.Sign(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		var body string
		if b, err := ioutil.ReadAll(resp.Body); err != nil {
			body = string(b)
		}
		return fmt.Errorf("%s: %s", resp.Status, body)
	}
	return nil
}

// Headers that should be copied when creating the redirect request
var omitHeaders map[string]bool = map[string]bool{
	"Content-Type":   true,
	"Content-Length": true,
}
