package rsapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
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
	CanAuthenticate(host string) error
	// EnableDump sets the format used by the authenticator to dump requests.
	// The format flag must have the Verbose bit enabled to have any effect.
	EnableDump(format Format)
	// Enable insecure mode: all auth requests are made using HTTP instead of HTTPS.
	// Used by tests.
	Insecure()
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
		client:       NewHttpClient(NoDump),
	}
}

// NewTokenAuthenticator returns a authenticator that use a oauth access token to do authentication.
// This is useful if the oauth handshake has already happened.
// Use the OAuthAuthenticator to use a refresh token and have the authenticator do the handshake.
func NewTokenAuthenticator(token string) Authenticator {
	return &tokenAuthenticator{token: token, dump: NoDump}
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
		client:    NewHttpClient(NoDump),
	}
}

// NewRL10Authenticator returns an authenticator that proxies all requests through the RightLink 10
// agent.
func NewRL10Authenticator(secret string) Authenticator {
	return &rl10Authenticator{secret: secret}
}

// NewHttpClient returns a http client that handles redirect in a way that's compatible with the
// RightScale APIs, namely: it copies over the headers that need to be copied over (e.g.
// X-Api-Version).
func NewHttpClient(dumpFormat Format) HttpClient {
	return &dumpClient{
		RoundTripper: &http.Transport{ResponseHeaderTimeout: 20 * time.Second},
		Format:       dumpFormat,
	}
}

// loginRequestBuilder is a generic login request factory.
type loginRequestBuilder interface {
	BuildLoginRequest(host string, insecure bool) (*http.Request, error)
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
	dump      Format
	insecure  bool
}

// newCookieSigner returns a cookie signer that uses the given builder to build login requests.
func newCookieSigner(builder loginRequestBuilder, accountID int) Authenticator {
	return &cookieSigner{
		builder:   builder,
		accountID: accountID,
		refreshAt: time.Now().Add(-2 * time.Minute),
		client:    NewHttpClient(NoDump),
	}
}

// Sign adds the username and password authorization cookies to the request.
// Checks the freshness of the session and creates a new one if needed.
func (s *cookieSigner) Sign(req *http.Request) error {
	if time.Now().After(s.refreshAt) {
		authReq, authErr := s.builder.BuildLoginRequest(s.host, s.insecure)
		if authErr != nil {
			return authErr
		}
		resp, err := s.client.Do(authReq)
		if err != nil {
			return err
		}
		url, err := extractRedirectURL(resp)
		if err != nil {
			return err
		}
		if url != nil {
			authReq, authErr = s.builder.BuildLoginRequest(url.Host, s.insecure)
			if authErr != nil {
				return authErr
			}
			s.host = url.Host
			req.Host = url.Host
			req.URL.Host = url.Host
			resp, err = s.client.Do(authReq)
		}
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
	req.Header.Set("X-Account", strconv.Itoa(s.accountID))
	return nil
}

// SetHost sets the host used to create sessions.
func (s *cookieSigner) SetHost(host string) {
	s.host = host
}

// CanAuthenticate makes a test request to CM 1.5 and returns true if it is successful.
func (s *cookieSigner) CanAuthenticate(host string) error {
	_, instance := s.builder.(*instanceLoginRequestBuilder)
	return testAuth(s, s.client, host, instance, s.insecure)
}

// EnableDump sets the dump format.
func (s *cookieSigner) EnableDump(format Format) {
	if c, ok := s.client.(*dumpClient); ok {
		c.EnableDump(format)
	}
}

// Insecure forces the use of HTTP
func (s *cookieSigner) Insecure() {
	s.insecure = true
	if c, ok := s.client.(*dumpClient); ok {
		c.Insecure()
	}
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
	insecure     bool
}

// Sign adds the OAuth bearer header to the *http.Request
func (s *oAuthSigner) Sign(req *http.Request) error {
	if time.Now().After(s.refreshAt) {
		jsonStr := fmt.Sprintf(`{"grant_type":"refresh_token","refresh_token":"%s"}`, s.refreshToken)
		authReq, err := http.NewRequest("POST", buildURL(s.host, "api/oauth2", s.insecure),
			bytes.NewBufferString(jsonStr))
		if err != nil {
			return fmt.Errorf("Authentication failed (failed to build request): %s", err)
		}
		authReq.Header.Set("X-API-Version", "1.5")
		authReq.Header.Set("Content-Type", "application/json")
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
			body, err := ioutil.ReadAll(resp.Body)
			var msg string
			if err != nil {
				msg = " - <failed to read body>"
			}
			if len(body) > 0 {
				msg = " - " + string(body)
			}
			return fmt.Errorf("Authentication failed: %s%s", resp.Status, msg)
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
func (s *oAuthSigner) CanAuthenticate(host string) error {
	return testAuth(s, s.client, host, false, s.insecure)
}

// EnableDump sets the dump format.
func (s *oAuthSigner) EnableDump(format Format) {
	s.client.(*dumpClient).EnableDump(format)
}

// Insecure forces the use of HTTP
func (s *oAuthSigner) Insecure() {
	s.client.(*dumpClient).Insecure()
	s.insecure = true
}

// OAuth access token authenticator
type tokenAuthenticator struct {
	token    string
	host     string // Only used by CanAuthenticate
	dump     Format
	insecure bool
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
func (t *tokenAuthenticator) CanAuthenticate(host string) error {
	client := NewHttpClient(t.dump)
	return testAuth(t, client, host, false, t.insecure)
}

// EnableDump sets the dump format.
func (t *tokenAuthenticator) EnableDump(format Format) {
	t.dump = format
}

// Insecure forces the use of HTTP
func (t *tokenAuthenticator) Insecure() {
	t.insecure = true
}

// RightLink 10 authenticator
type rl10Authenticator struct {
	secret   string
	host     string
	dump     Format
	insecure bool
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
func (a *rl10Authenticator) CanAuthenticate(host string) error {
	client := NewHttpClient(a.dump)
	return testAuth(a, client, host, true, a.insecure)
}

// EnableDump sets the dump format.
func (a *rl10Authenticator) EnableDump(format Format) {
	a.dump = format
}

// Insecure forces the use of HTTP
func (a *rl10Authenticator) Insecure() {
	a.insecure = true
}

// SS authenticator
type ssAuthenticator struct {
	auther    Authenticator // Authentication against core
	accountID int           // Account used to create SS local session
	host      string        // Login (core) host
	refreshAt time.Time     // SS local session refresh deadline
	client    HttpClient
	insecure  bool
}

// Self-Service authenticator first creates a global session with the core then creates a local
// session with self-service.
func (a *ssAuthenticator) Sign(r *http.Request) error {
	if time.Now().After(a.refreshAt) {
		u := buildURL(a.host, "api/catalog/new_session", a.insecure)
		u += "?account_id=" + strconv.Itoa(a.accountID)
		authReq, err := http.NewRequest("GET", u, nil)
		if err != nil {
			return err
		}
		if err := a.auther.Sign(authReq); err != nil {
			return err
		}

		// A bit tricky: if the auther is the cookie signer it could have updated the
		// host after being redirected.
		if ca, ok := a.auther.(*cookieSigner); ok {
			a.SetHost(ca.host)
			authReq.Host = a.host
			authReq.URL.Host = a.host
		}

		authReq.Header.Set("Content-Type", "application/json")
		resp, err := a.client.Do(authReq)
		if err != nil {
			return fmt.Errorf("Authentication failed: %s", err)
		}
		if resp.StatusCode != 303 {
			body, err := ioutil.ReadAll(resp.Body)
			var msg string
			if err != nil {
				msg = " - <failed to read body>"
			}
			if len(body) > 0 {
				msg = " - " + string(body)
			}
			return fmt.Errorf("Authentication failed: %s%s", resp.Status, msg)
		}
		a.refreshAt = time.Now().Add(2 * time.Hour)
	}
	a.auther.Sign(r)
	r.Header.Set("X-Api-Version", "1.0")
	r.Host = a.host
	r.URL.Host = a.host

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
	elems[len(elems)-2] = strings.Replace(elems[len(elems)-2], "us", "selfservice", 1)
	ssLoginHostPrefix := strings.Join(elems, "-")
	a.host = strings.Join(append([]string{ssLoginHostPrefix}, urlElems[1:]...), ".")
}

// CanAuthenticate makes a test request to SS and returns true if it is successful.
func (a *ssAuthenticator) CanAuthenticate(host string) error {
	url := fmt.Sprintf("api/catalog/accounts/%d/user_preferences", a.accountID)
	req, err := http.NewRequest("GET", buildURL(host, url, a.insecure), nil)
	if err != nil {
		return err
	}
	req.Header.Set("X-Api-Version", "1.0")
	if err := a.Sign(req); err != nil {
		return err
	}
	resp, err := a.client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		var body string
		if b, err := ioutil.ReadAll(resp.Body); err != nil {
			body = ": " + string(b)
		}
		return fmt.Errorf("%s%s", resp.Status, body)
	}
	return nil
}

// EnableDump sets the dump format.
func (a *ssAuthenticator) EnableDump(format Format) {
	a.client.(*dumpClient).EnableDump(format)
	a.auther.EnableDump(format)
}

// Insecure forces the use of HTTP
func (a *ssAuthenticator) Insecure() {
	a.client.(*dumpClient).Insecure()
	a.auther.Insecure()
	a.insecure = true
}

// basicLoginRequestBuilder builds login requests from users email and password.
type basicLoginRequestBuilder struct {
	username  string
	password  string
	accountID int
}

// BuildLoginRequest builds session create requests from users email and password.
func (b *basicLoginRequestBuilder) BuildLoginRequest(host string, insecure bool) (*http.Request, error) {
	if host == "" {
		host = "us-3.rightscale.com"
	}
	jsonStr := fmt.Sprintf(`{"email":"%s","password":"%s","account_href":"/api/accounts/%d"}`,
		b.username, b.password, b.accountID)
	authReq, err := http.NewRequest("POST", buildURL(host, "api/sessions", insecure),
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
func (b *instanceLoginRequestBuilder) BuildLoginRequest(host string, insecure bool) (*http.Request, error) {
	if host == "" {
		host = "us-3.rightscale.com"
	}
	accountHref := fmt.Sprintf("/api/accounts/%d", b.accountID)
	jsonStr := fmt.Sprintf(`{"instance_token":"%s","account_href":"%s"}`, b.token, accountHref)
	authReq, err := http.NewRequest("POST", buildURL(host, "api/session/instance", insecure),
		bytes.NewBufferString(jsonStr))
	if err != nil {
		return nil, fmt.Errorf("Authentication failed (failed to build request): %s", err)
	}
	authReq.Header.Set("X-API-Version", "1.5")
	authReq.Header.Set("Content-Type", "application/json")
	return authReq, nil
}

// HTTP client that optionally dumps requests and responses.
// This client also disables the default http client redirect handling.
type dumpClient struct {
	RoundTripper http.RoundTripper
	Format       Format
	UseHTTP      bool
}

// Do dumps the request, makes the request and dumps the response as specified by the format.
func (d *dumpClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Set("User-Agent", UA)
	var b []byte
	if d.Format.IsVerbose() {
		b = dumpRequest(d.Format, req)
	}
	resp, err := d.RoundTripper.RoundTrip(req)
	if err != nil {
		return nil, err
	}
	if d.Format.IsVerbose() {
		dumpResponse(d.Format, resp, req, b)
	}
	return resp, nil
}

// EnableDump sets the dump format
func (d *dumpClient) EnableDump(format Format) {
	d.Format = format
}

// Insecure forces the use of HTTP
func (d *dumpClient) Insecure() {
	d.UseHTTP = true
}

// extractRedirectURL is a helper function that extracts the Location header from a redirect
// response. It returns nil if the header is missing, an error if it's malformed.
func extractRedirectURL(resp *http.Response) (*url.URL, error) {
	var u *url.URL
	if resp.StatusCode > 299 && resp.StatusCode < 399 {
		loc := resp.Header.Get("Location")
		if loc != "" {
			var err error
			u, err = url.Parse(loc)
			if err != nil {
				return nil, fmt.Errorf("invalid Location header '%s': %s", loc, err)
			}
		}
	}
	return u, nil
}

// Compute API URL given a scheme, hostname and a path
func buildURL(host, path string, insecure bool) string {
	scheme := "https"
	if insecure {
		scheme = "http"
	}
	u := url.URL{
		Scheme: scheme,
		Host:   host,
		Path:   path,
	}
	return u.String()
}

// testAuth makes a GET /api/sessions CM 1.5 request using the given authenticator and returns
// an error if it failed, nil otherwise.
// The instance flag specifies whether an instance or an account facing API request should be made.
func testAuth(auth Authenticator, client HttpClient, host string, instance, insecure bool) error {
	if host == "" {
		return fmt.Errorf("missing host information")
	}
	var req *http.Request
	var err error
	if instance {
		req, err = http.NewRequest("GET", buildURL(host, "api/user_data", insecure), nil)
	} else {
		req, err = http.NewRequest("GET", buildURL(host, "api/sessions", insecure), nil)
	}
	if err != nil {
		return err
	}
	req.Header.Set("X-Api-Version", "1.5")
	if err = auth.Sign(req); err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		var body string
		if b, err := ioutil.ReadAll(resp.Body); err != nil {
			body = ": " + string(b)
		}
		return fmt.Errorf("%s%s", resp.Status, body)
	}
	return nil
}

// Headers that should be copied when creating the redirect request
var omitHeaders map[string]bool = map[string]bool{
	"Content-Type":   true,
	"Content-Length": true,
}

// CA authenticator
type CAAuthenticator struct {
	Host      string
	Username  string
	Password  string
	Cookies   map[int][]*http.Cookie
	RefreshAt time.Time
	Client    HttpClient
}

// Add username/password authorization cookies to the *http.Request
func (a *CAAuthenticator) Sign(r *http.Request, host string, accountId int) error {
	if err := a.Refresh(a.Host, accountId); err != nil {
		return err
	}
	for _, c := range a.Cookies[accountId] {
		r.AddCookie(c)
	}

	return nil
}

// Make sure global session cookie is up-to-date
func (a *CAAuthenticator) Refresh(host string, accountId int) error {
	if time.Now().After(a.RefreshAt) {
		authReq, authErr := a.newLoginRequest(host, accountId)
		if authErr != nil {
			return authErr
		}
		resp, err := a.Client.Do(authReq)

		fmt.Println(resp)
		if err != nil {
			return fmt.Errorf("Authentication failed: %s", err) // TBD RETRY A FEW TIMES
		}
		if resp.StatusCode != 204 {
			return fmt.Errorf("Authentication failed: %s", resp.Status)
		}
		if a.Cookies == nil {
			a.Cookies = make(map[int][]*http.Cookie)
		}
		a.Cookies[accountId] = resp.Cookies()
		a.RefreshAt = time.Now().Add(time.Duration(2) * time.Hour)
	}
	return nil
}

// To be called from rsapi.Api to verify credentials, and (re)set host if redirected
func (a *CAAuthenticator) ResolveHost(host string, accountId int) (string, error) {
	authReq, authErr := a.newLoginRequest(host, accountId)
	if authErr != nil {
		return host, authErr
	}
	return resolveHost(authReq, host, accountId)
}

// Return a new *http.Request with the specified host, and username/password
func (a *CAAuthenticator) newLoginRequest(host string, accountId int) (*http.Request, error) {
	jsonStr := fmt.Sprintf(`{"email":"%s","password":"%s","account_href":"/api/accounts/%d"}`,
		a.Username, a.Password, accountId)
	authReq, err := http.NewRequest("POST", fmt.Sprintf("https://%s/api/sessions", host),
		bytes.NewBufferString(jsonStr))
	if err != nil {
		return authReq, fmt.Errorf("Authentication failed (failed to build request): %s", err.Error())
	}
	authReq.Header.Set("X-API-Version", "1.5")
	authReq.Header.Set("Content-Type", "application/json")
	return authReq, nil
}
