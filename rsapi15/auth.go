package rsapi15

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
	Sign(req *http.Request, endpoint string) error // Sign http Request (add auth headers)
}

// OAuth authenticator uses the user oauth refresh token
type OAuthAuthenticator struct {
	RefreshToken string
	AccessToken  string
	RefreshAt    time.Time
	Client       *http.Client
}

// Account authenticator uses RS oauth
func (a *OAuthAuthenticator) Sign(r *http.Request, endpoint string) error {
	if time.Now().After(a.RefreshAt) {
		jsonStr := []byte(fmt.Sprintf(`{"grant_type":"refresh_token","refresh_token":"%s"}`, a.RefreshToken))
		authReq, err := http.NewRequest("POST", fmt.Sprintf("https://%s/api/oauth2", endpoint), bytes.NewBuffer(jsonStr))
		if err != nil {
			return fmt.Errorf("Authentication failed (failed to build request): %s", err.Error())
		}
		authReq.Header.Set("X-API-Version", "1.5")
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
		a.AccessToken = session["access_token"].(string)
		d, err := time.ParseDuration(session["expires_in"].(string) + "s")
		if err != nil {
			return fmt.Errorf("Authentication failed (failed to parse token duration): %s", err.Error())
		}
		a.RefreshAt = time.Now().Add(d / 2)
	}
	r.Header.Add("Authorization", fmt.Sprintf("Bearer %s", a.AccessToken))

	return nil
}

// Return endpoint that hosts given account (e.g. "us-3.rightscale.com")
func (a *OAuthAuthenticator) ResolveEndpoint(accountId int) (string, error) {
	return "my.rightscale.com", nil // TBD
}
