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
	Sign(req *http.Request, host string, accountId int) error // Sign http Request (add auth headers)
	ResolveHost(host string, accountId int) (string, error)
}

func LoginAndOAuthResolveHost(authReq *http.Request, host string, accountId int) (string, error) {
	redirectHost := host
	client := http.Client{
		// Pretty much a direct copy/paste from;
		// https://github.com/rightscale/go-lib/blob/42736168b5205993699ecc52391c84fa3f8520d2/net/clients/httpClient.go#L345-L375
		CheckRedirect: func(nextRequest *http.Request, via []*http.Request) error {
			omitHeaders := map[string]bool {
				"Content-Type": true,
				"Content-Length": true,
			}
			viaCount := len(via)
			// Just go with the standard redirect count
			maxRedirects := 10
			if viaCount > maxRedirects {
				return fmt.Errorf("Stopped after %d redirects.", maxRedirects)
			}

			// ensure next request has same headers as last request. the net/http default
			// behavior is to pass the empty header to the redirected URL, which does not
			// work for RS APIs because they require X-Api-Version, etc.
			//
			// note that a redirected POST or PUT is changed to a GET and the original
			// body is discarded per 302 logic. we should therefore also omit body-related
			// headers from the next request. we will use a configurable header whitelist.
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

// Login authenticator uses username/password
type LoginAuthenticator struct {
	Username  string
	Password  string
	Cookies   map[int][]*http.Cookie
	RefreshAt time.Time
	Client    HttpClient
}

func (a *LoginAuthenticator) Sign(r *http.Request, host string, accountId int) error {
	if _,err := a.Refresh(host, accountId); err != nil {
		return err
	}
	for _, c := range a.Cookies[accountId] {
		r.AddCookie(c)
	}
	return nil
}

// Return a new *http.Request with the specified host, and username/password
func (a *LoginAuthenticator) newLoginRequest(host string, accountId int) (*http.Request, error) {
	jsonStr := fmt.Sprintf(`{"email":"%s","password":"%s","account_href":"/api/accounts/%d"}`,
		a.Username, a.Password, accountId)
	authReq, err := http.NewRequest("POST", fmt.Sprintf("https://%s/api/sessions", host),
		bytes.NewBufferString(jsonStr))
	if err != nil {
		return authReq, fmt.Errorf("Authentication failed (failed to build request): %s", err.Error())
	}
	authReq.Header.Set("X-API-Version", "1.5")
	authReq.Header.Set("Content-Type", "application/json")
	return authReq,nil
}

// Make sure global session cookie is up-to-date
func (a *LoginAuthenticator) Refresh(host string, accountId int) (string, error) {
	if time.Now().After(a.RefreshAt) {
		authReq, authErr := a.newLoginRequest(host,accountId)
		if authErr != nil {
			return host, authErr
		}
		resp, err := a.Client.Do(authReq)
		if err != nil {
			return host, fmt.Errorf("Authentication failed: %s", err.Error()) // TBD RETRY A FEW TIMES
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
	return host, nil
}

// To be called from rsapi.Api to verify credentials, and (re)set host if redirected
func (a *LoginAuthenticator) ResolveHost(host string, accountId int) (string, error) {
	authReq, authErr := a.newLoginRequest(host,accountId)
	if authErr != nil {
		return host, authErr
	}
	return LoginAndOAuthResolveHost(authReq, host,accountId)
}

// OAuth authenticator uses the user oauth refresh token
type OAuthAuthenticator struct {
	RefreshToken string
	AccessToken  string
	RefreshAt    time.Time
	Client       HttpClient
}

// Account authenticator uses RS oauth
func (a *OAuthAuthenticator) Sign(r *http.Request, host string, accountId int) error {
	if _, err := a.Refresh(host); err != nil {
		return err
	}
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.AccessToken))

	return nil
}

// Return a new *http.Request with the specified host, and oauth refresh token
func (a *OAuthAuthenticator) newLoginRequest(host string) (*http.Request, error) {
	jsonStr := fmt.Sprintf(`{"grant_type":"refresh_token","refresh_token":"%s"}`, a.RefreshToken)
	authReq, err := http.NewRequest("POST", fmt.Sprintf("https://%s/api/oauth2", host), bytes.NewBufferString(jsonStr))
	if err != nil {
		return nil, fmt.Errorf("Authentication failed (failed to build request): %s", err)
	}
	authReq.Header.Set("X-API-Version", "1.5")
	authReq.Header.Set("Content-Type", "application/json")
	return authReq, nil
}

// Make sure access token is up-to-date
func (a *OAuthAuthenticator) Refresh(host string) (string, error) {
	if time.Now().After(a.RefreshAt) {
		authReq, authErr := a.newLoginRequest(host)
		if authErr != nil {
			return host, fmt.Errorf("Authentication failed: %s", authErr)
		}
		resp, err := a.Client.Do(authReq)
		if err != nil {
			return host, fmt.Errorf("Authentication failed: %s", err) // TBD RETRY A FEW TIMES
		}
		defer resp.Body.Close()
		var session map[string]interface{}
		jsonBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return host, fmt.Errorf("Authentication failed (failed to read response): %s", err)
		}
		if resp.StatusCode != 200 {
			return fmt.Errorf("Authentication failed: %s", resp.Status)
		}
		err = json.Unmarshal(jsonBytes, &session)
		if err != nil {
			return host, fmt.Errorf("Authentication failed (failed to load response JSON): %s", err)
		}
		accessToken, ok := session["access_token"].(string)
		if !ok {
			return host, fmt.Errorf("Unexpected auth response: %s", jsonBytes)
		}
		a.AccessToken = accessToken
		d, err := time.ParseDuration(fmt.Sprintf("%vs", session["expires_in"]))
		if err != nil {
			return host, fmt.Errorf("Authentication failed (failed to parse token duration): %s", err)
		}
		a.RefreshAt = time.Now().Add(d / 2)
	}
	return host, nil
}

// To be called from rsapi.Api to verify credentials, and (re)set host if redirected
//
// Note: As of 3-19-2015, the login request will never be redirected. Instead a
// 422 will be returned by the API since the oauth token does not exist in all
// shards, but only in the shard where the account lives. Still implementing
// redirect handling in anticipation that oauth may be more cross shard friendly
// in the future.
func (a *OAuthAuthenticator) ResolveHost(host string, accountId int) (string, error) {
	authReq, authErr := a.newLoginRequest(host)
	if authErr != nil {
		return host, authErr
	}
	return LoginAndOAuthResolveHost(authReq,host,accountId)
}

// RightLink 10 authenticator
type RL10Authenticator struct {
	Secret string
}

// RL10 authenticator uses special header
func (a *RL10Authenticator) Sign(r *http.Request, host string, accountId int) error {
	r.Header.Set("X-RLL-Secret", a.Secret)

	return nil
}

// A stub to satisfy the interface. RL10 will always use host "localhost" and
// will not be redirected
func (a *RL10Authenticator) ResolveHost(host string, accountId int) (string, error) {
	return host, nil
}

// SS authenticator
type SSAuthenticator struct {
	CoreAuth  Authenticator // Authentication against core
	AccountId int
	RefreshAt time.Time
	Client    HttpClient
}

// Account authenticator uses RS oauth
func (a *SSAuthenticator) Sign(r *http.Request, host string, accountId int) error {
	refreshAt := a.RefreshAt
	if time.Now().After(refreshAt) {
		authReq, err := http.NewRequest("GET",
			fmt.Sprintf("https://%s/api/catalog/new_session?account_id=%d",
				ssHost(host), accountId), nil)
		if err != nil {
			return err
		}
		a.CoreAuth.Sign(authReq, host, accountId)
		authReq.Header.Set("Content-Type", "application/json")
		_, err = a.Client.Do(authReq)
		if err != nil {
			return fmt.Errorf("Authentication failed: %s", err) // TBD RETRY A FEW TIMES
		}
	}
	a.CoreAuth.Sign(r, host, accountId)
	r.Host = ssHost(host)

	return nil
}

// Determine the correct SS host based on the CM/Core host. Shard redirection
// will already have occurred since the rsapi.Api for the core has already been
// initialized and (if necessary) redirected.
func (a *SSAuthenticator) ResolveHost(host string, accountId int) (string, error) {
	return ssHost(host),nil
}

// Return Self-service endpoint from login endpoint
// The following isn't great but seems better than having to enter by hand
func ssHost(host string) string {
	urlElems := strings.Split(host, ".")
	hostPrefix := urlElems[0]
	elems := strings.Split(hostPrefix, "-")
	elems[len(elems)-2] = "selfservice"
	ssLoginHostPrefix := strings.Join(elems, "-")
	return strings.Join(append([]string{ssLoginHostPrefix}, urlElems[1:]...), ".")
}
