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
}

// Login authenticator uses username/password
type LoginAuthenticator struct {
	Username  string
	Password  string
	Cookies   map[int][]*http.Cookie
	RefreshAt time.Time
	Client    HttpClient
}

// Account authenticator uses RS oauth
func (a *LoginAuthenticator) Sign(r *http.Request, host string, accountId int) error {
	if err := a.Refresh(host, accountId); err != nil {
		return err
	}
	for _, c := range a.Cookies[accountId] {
		r.AddCookie(c)
	}
	return nil
}

// Make sure global session cookie is up-to-date
func (a *LoginAuthenticator) Refresh(host string, accountId int) error {
	if time.Now().After(a.RefreshAt) {
		jsonStr := fmt.Sprintf(`{"email":"%s","password":"%s","account_href":"/api/accounts/%d"}`,
			a.Username, a.Password, accountId)
		authReq, err := http.NewRequest("POST", fmt.Sprintf("https://%s/api/sessions", host),
			bytes.NewBufferString(jsonStr))
		if err != nil {
			return fmt.Errorf("Authentication failed (failed to build request): %s", err.Error())
		}
		authReq.Header.Set("X-API-Version", "1.5")
		authReq.Header.Set("Content-Type", "application/json")
		resp, err := a.Client.Do(authReq)
		if err != nil {
			return fmt.Errorf("Authentication failed: %s", err.Error()) // TBD RETRY A FEW TIMES
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

// OAuth authenticator uses the user oauth refresh token
type OAuthAuthenticator struct {
	RefreshToken string
	AccessToken  string
	RefreshAt    time.Time
	Client       HttpClient
}

// Account authenticator uses RS oauth
func (a *OAuthAuthenticator) Sign(r *http.Request, host string, accountId int) error {
	if err := a.Refresh(host); err != nil {
		return err
	}
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.AccessToken))

	return nil
}

// Make sure access token is up-to-date
func (a *OAuthAuthenticator) Refresh(host string) error {
	if time.Now().After(a.RefreshAt) {
		jsonStr := fmt.Sprintf(`{"grant_type":"refresh_token","refresh_token":"%s"}`, a.RefreshToken)
		authReq, err := http.NewRequest("POST", fmt.Sprintf("https://%s/api/oauth2", host), bytes.NewBufferString(jsonStr))
		if err != nil {
			return fmt.Errorf("Authentication failed (failed to build request): %s", err.Error())
		}
		authReq.Header.Set("X-API-Version", "1.5")
		authReq.Header.Set("Content-Type", "application/json")
		resp, err := a.Client.Do(authReq)
		if err != nil {
			return fmt.Errorf("Authentication failed: %s", err.Error()) // TBD RETRY A FEW TIMES
		}
		defer resp.Body.Close()
		var session map[string]interface{}
		jsonBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("Authentication failed (failed to read response): %s", err.Error())
		}
		if resp.StatusCode != 200 {
			return fmt.Errorf("Authentication failed: %s", resp.Status)
		}
		err = json.Unmarshal(jsonBytes, &session)
		if err != nil {
			return fmt.Errorf("Authentication failed (failed to load response JSON): %s", err.Error())
		}
		accessToken, ok := session["access_token"].(string)
		if !ok {
			return fmt.Errorf("Unexpected auth response: %s", jsonBytes)
		}
		a.AccessToken = accessToken
		d, err := time.ParseDuration(fmt.Sprintf("%vs", session["expires_in"]))
		if err != nil {
			return fmt.Errorf("Authentication failed (failed to parse token duration): %s", err.Error())
		}
		a.RefreshAt = time.Now().Add(d / 2)
	}
	return nil
}

// Return host that hosts given account (e.g. "us-3.rightscale.com")
func (a *OAuthAuthenticator) ResolveHost(accountId int) (string, error) {
	return "my.rightscale.com", nil // TBD
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
			return fmt.Errorf("Authentication failed: %s", err.Error()) // TBD RETRY A FEW TIMES
		}
	}
	a.CoreAuth.Sign(r, host, accountId)
	r.Host = ssHost(host)

	return nil
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
