package rsapi

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
}

// NewBasicAuthenticator returns a authenticator that uses email and password to create sessions.
// This function creates an initial session and returns an error if the corresponding API request
// fails (for example if the credentials are not valid).
// The returned authenticator takes care of refreshing the RightScale session as needed.
// The host is optional, if empty then "us-3.rightscale.com" is used initially.
// The correct host is retrieved from any redirect response when creating the session.
func NewBasicAuthenticator(username, password, host string, accountID int) (Authenticator, error) {
	builder := basicLoginRequestBuilder{
		username:  username,
		password:  password,
		accountID: accountID,
		host:      host,
	}
	h, resp, err := resolveHost(accountID, &builder)
	if err != nil {
		return nil, err
	}
	builder.host = h
	signer := cookieSigner{
		host:      host,
		builder:   &builder,
		cookies:   resp.Cookies(),
		refreshAt: time.Now().Add(2 * time.Hour),
		client:    http.DefaultClient,
	}
	return &signer, nil
}

// NewOAuthAuthenticator returns a authenticator that uses a oauth refresh token to create access
// tokens.
// The refresh token can be found in the CM dashboard under Settings > Account Settings > API Credentials.
func NewOAuthAuthenticator(token, host string) (Authenticator, error) {
	s := oAuthSigner{
		refreshToken: token,
		host:         host,
		refreshAt:    time.Now().Add(-2 * time.Minute),
		client:       http.DefaultClient,
	}
	if err := testCM15Auth(&s, host); err != nil {
		return nil, err
	}
	return &s, nil
}

// NewTokenAuthenticator returns a authenticator that use a oauth access token to do authentication.
// This is useful if the oauth handshake has already happened.
// Use the OAuthAuthenticator to use a refresh token and have the authenticator do the handshake.
func NewTokenAuthenticator(token, host string) Authenticator {
	return &tokenAuthenticator{token: token}
}

// NewInstanceAuthenticator returns an authenticator that uses the instance facing API token to
// create sessions. This is the token found on RightLink instances under the RS_API_TOKEN
// environment variable.
// The host is optional, if empty then "us-3.rightscale.com" is used initially.
// The correct host is retrieved from any redirect response when creating the session.
// Note: Use of rsc made from RightLink10 instances can use the RL10 authenticator instead.
func NewInstanceAuthenticator(token, host string, accountID int) (Authenticator, error) {
	builder := instanceLoginRequestBuilder{token: token, host: host, accountID: accountID}
	h, resp, err := resolveHost(accountID, &builder)
	if err != nil {
		return nil, err
	}
	builder.host = h
	signer := cookieSigner{
		host:      host,
		builder:   &builder,
		cookies:   resp.Cookies(),
		refreshAt: time.Now().Add(2 * time.Hour),
		client:    http.DefaultClient,
	}
	return &signer, nil
}

// NewSSAuthenticator returns an authenticator that wraps another one and adds the logic needed to
// create sessions in Self-Service.
// It returns an error if Validate() does.
func NewSSAuthenticator(auther Authenticator, accountID int) (Authenticator, error) {
	if _, ok := auther.(*ssAuthenticator); ok {
		// Only wrap if not wrapped already
		return auther, nil
	}
	var host string
	if c, ok := auther.(*cookieSigner); ok {
		host = c.host
	} else if a, ok := auther.(*oAuthSigner); ok {
		host = a.host
	}
	if host == "" {
		return nil, fmt.Errorf("invalid core authenticator")
	}
	a := &ssAuthenticator{
		host:      host,
		auther:    auther,
		accountID: accountID,
		refreshAt: time.Now().Add(-2 * time.Minute),
		client:    http.DefaultClient,
	}
	return a, nil
}

// NewRL10Authenticator returns an authenticator that proxies all requests through the RightLink 10
// agent.
// It returns an error if Validate() does.
func NewRL10Authenticator(secret, host string) Authenticator {
	return &rl10Authenticator{secret: secret}
}

// loginRequestBuilder is a generic login request factory.
type loginRequestBuilder interface {
	BuildLoginRequest() (*http.Request, error)
}

// cookieSigner signs requests using adding a global session cookie.
// Used by both the basic and instance session managers.
type cookieSigner struct {
	builder   loginRequestBuilder
	cookies   []*http.Cookie
	host      string
	refreshAt time.Time
	client    HttpClient
}

// Sign adds the username and password authorization cookies to the *http.Request
// Checks the freshness of the session and creates a new one if needed.
func (s *cookieSigner) Sign(req *http.Request) error {
	if time.Now().After(s.refreshAt) {
		authReq, authErr := s.builder.BuildLoginRequest()
		if authErr != nil {
			return authErr
		}
		resp, err := s.client.Do(authReq)
		if err != nil {
			return fmt.Errorf("Authentication failed: %s", err)
		}
		if resp.StatusCode != 204 {
			return fmt.Errorf("Authentication failed: %s", resp.Status)
		}
		s.cookies = resp.Cookies()
		s.refreshAt = time.Now().Add(time.Duration(2) * time.Hour)
	}
	for _, c := range s.cookies {
		req.AddCookie(c)
	}
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
		authReq, authErr := s.BuildLoginRequest()
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

// BuildLoginRequest returns a new *http.Request that can refresh the access token
func (s *oAuthSigner) BuildLoginRequest() (*http.Request, error) {
	jsonStr := fmt.Sprintf(`{"grant_type":"refresh_token","refresh_token":"%s"}`, s.refreshToken)
	authReq, err := http.NewRequest("POST", endpoint(s.host, "api/oauth2"), bytes.NewBufferString(jsonStr))
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
func (t *tokenAuthenticator) Sign(r *http.Request) error {
	r.Header.Set("Authorization", "Bearer "+t.token)
	return nil
}

// RightLink 10 authenticator
type rl10Authenticator struct {
	secret string
}

// RL10 authenticator uses special header
func (s *rl10Authenticator) Sign(r *http.Request) error {
	r.Header.Set("X-RLL-Secret", s.secret)
	return nil
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
			endpoint(ssHostFromLogin(a.host),
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
	r.Host = ssHostFromLogin(a.host)

	return nil
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

// basicLoginRequestBuilder builds login requests from users email and password.
type basicLoginRequestBuilder struct {
	username  string
	password  string
	accountID int
	host      string
}

// BuildLoginRequest builds session create requests from users email and password.
func (b *basicLoginRequestBuilder) BuildLoginRequest() (*http.Request, error) {
	host := b.host
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
	host      string
	accountID int
}

// BuildLoginRequest builds session create requests from users email and password.
func (b *instanceLoginRequestBuilder) BuildLoginRequest() (*http.Request, error) {
	host := b.host
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
	authReq, authErr := b.BuildLoginRequest()
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

// testCM15Auth makes a GET /api/sessions CM 1.5 request using the given authenticator and returns
// an error if it failed, nil otherwise.
func testCM15Auth(auth Authenticator, host string) error {
	req, err := http.NewRequest("GET", endpoint(host, "api/sessions"), nil)
	req.Header.Set("X-Api-Version", "1.5")
	if err != nil {
		return err
	}
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
