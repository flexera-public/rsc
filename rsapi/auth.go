package rsapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Authenticator interface
type Authenticator interface {
	Sign(req *http.Request, host string) error // Sign http Request (add auth headers)
}

// OAuth authenticator uses the user oauth refresh token
type OAuthAuthenticator struct {
	RefreshToken string
	AccessToken  string
	RefreshAt    time.Time
	Client       HttpClient
}

// Account authenticator uses RS oauth
func (a *OAuthAuthenticator) Sign(r *http.Request, host string) error {
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
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.AccessToken))

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
func (a *RL10Authenticator) Sign(r *http.Request, host string) error {
	r.Header.Set("X-RLL-Secret", a.Secret)

	return nil
}
