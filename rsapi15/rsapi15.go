//go:generate api15gen ./api_data.json
package rsapi15

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// RightScale API 1.5 client
// Instances of this struct should be created through `New`.
type Api15 struct {
	AccountId string        // Account in which client is currently operating
	Auth      Authenticator // Authenticator, signs requests for auth
	Logger    *log.Logger   // Optional logger, if specified requests and responses get logged
	Endpoint  string        // API endpoint, e.g. "us-3.rightscale.com"
	Client    *http.Client  // Underlying http client
}

// New returns a API 1.5 client that uses User oauth authentication.
// logger and client are optional.
// If no HTTP client is specified then the default client is used.
func New(accountId int, refreshToken string, logger *log.Logger, client *http.Client) (*Api15, error) {
	if client == nil {
		client := http.DefaultClient
	}
	auth := AccountAuthenticator{
		RefreshToken: refreshToken,
		RefreshAt:    time.Now().Add(-1 * time.Hour),
		Client:       client,
	}
	endpoint, err := auth.ResolveEndpoint(accountId)
	if err != nil {
		return nil, err
	}
	return &Api15{
		AccountId: accountId,
		Auth:      auth,
		Logger:    logger,
		Endpoint:  endpoint,
		Client:    client,
	}, nil
}

// Low-level GET request
func (a *Api15) Get(uri string) (map[string]interface{}, error) {
	return sendRequest("GET", uri, nil)
}

// Low-level POST request
// Any "Location" header present in the HTTP response is returned in the map under the "Location" key.
func (a *Api15) Post(uri string, body map[string]interface{}) (map[string]interface{}, error) {
	return sendRequest("POST", uri, body)
}

// Low-level PUT request
func (a *Api15) Put(uri string, body map[string]interface{}) error {
	_, err := sendRequest("PUT", uri, body)
	return err
}

// Low-level DELETE request
func (a *Api15) Delete(uri string) error {
	_, err := sendRequest("DELETE", uri, nil)
	return err
}

// Helper function that signs, sends and logs HTTP request
func (a *Api15) sendRequest(verb, uri string, body map[string]interface{}) (map[string]interface{}, error) {
	url := fmt.Sprintf("https://%s/%s", a.Endpoint, uri)
	err, jsonStr := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("Failed to serialize response (%s)", err.Error())
	}
	r, err := http.NewRequest(verb, url, bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		return nil, err
	}
	if err := a.Auth.Sign(r, a.Endpoint); err != nil {
		return nil, err
	}
	var id string
	if a.logger != nil {
		b := make([]byte, 6)
		io.ReadFull(rand.Reader, b)
		id = base64.StdEncoding.EncodeToString(b)
		a.logger.Printf("[%s] %s %s", id, r.Method, r.URL.String())
	}
	resp, err := a.Client.Do(r)
	if err != nil {
		return nil, err
	}
	if a.logger != nil {
		d := time.Since(startedAt)
		a.logger.Printf("[%s] %s in %s", id, resp.Status, d.String())
	}
	defer resp.Body.Close()
	var respBody map[string]interface{}
	jsonResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response (%s)", err.Error())
	}
	if len(jsonResp) > 0 {
		err = json.Unmarshal(string(jsonResp), &respBody)
		if err != nil {
			return nil, fmt.Errorf("Failed to load response (%s)", err.Error())
		}
	}
	// Special case for "Location" header
	if verb == "POST" {
		loc := resp.Header.Get("Location")
		if len(loc) > 0 {
			if respBody == nil {
				respBody := make(map[string]interface{})
			}
			respBody["Location"] = loc
		}
	}

	return respBody, err
}
