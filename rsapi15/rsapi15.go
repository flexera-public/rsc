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
	AccountId int           // Account in which client is currently operating
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
		client = http.DefaultClient
	}
	auth := OAuthAuthenticator{
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
		Auth:      &auth,
		Logger:    logger,
		Endpoint:  endpoint,
		Client:    client,
	}, nil
}

// Generic API parameters type
type ApiParams map[string]interface{}

// Low-level GET request that loads response JSON into generic object
func (a *Api15) Get(uri string) (interface{}, error) {
	resp, err := a.GetRaw(uri)
	if err != nil {
		return nil, err
	}
	return a.loadResponse(resp)
}

// Low-level GET request
func (a *Api15) GetRaw(uri string) (*http.Response, error) {
	return a.makeRequest("GET", uri, nil)
}

// Low-level POST request that loads response JSON into generic object
// Any "Location" header present in the HTTP response is returned in a map under the "Location" key.
func (a *Api15) Post(uri string, body ApiParams) (interface{}, error) {
	resp, err := a.PostRaw(uri, body)
	if err != nil {
		return nil, err
	}
	return a.loadResponse(resp)
}

// Low-level POST request
func (a *Api15) PostRaw(uri string, body ApiParams) (*http.Response, error) {
	return a.makeRequest("POST", uri, body)
}

// Low-level PUT request
func (a *Api15) Put(uri string, body ApiParams) error {
	_, err := a.makeRequest("PUT", uri, body)
	return err
}

// Low-level DELETE request
func (a *Api15) Delete(uri string) error {
	_, err := a.makeRequest("DELETE", uri, nil)
	return err
}

// Helper function that signs, makes and logs HTTP request
func (a *Api15) makeRequest(verb, uri string, body ApiParams) (*http.Response, error) {
	url := fmt.Sprintf("https://%s/%s", a.Endpoint, uri)
	jsonStr, err := json.Marshal(body)
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
	var startedAt = time.Now()
	if a.Logger != nil {
		b := make([]byte, 6)
		io.ReadFull(rand.Reader, b)
		id = base64.StdEncoding.EncodeToString(b)
		a.Logger.Printf("[%s] %s %s", id, r.Method, r.URL.String())
	}
	resp, err := a.Client.Do(r)
	if err != nil {
		return nil, err
	}
	if a.Logger != nil {
		d := time.Since(startedAt)
		a.Logger.Printf("[%s] %s in %s", id, resp.Status, d.String())
	}
	return resp, err
}

// Deserialize JSON response into generic object.
// If the response has a "Location" header then the returned object is a map with one key "Location"
// containing the value of the header.
func (a *Api15) loadResponse(resp *http.Response) (interface{}, error) {
	defer resp.Body.Close()
	var respBody interface{}
	jsonResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response (%s)", err.Error())
	}
	if len(jsonResp) > 0 {
		err = json.Unmarshal(jsonResp, &respBody)
		if err != nil {
			return nil, fmt.Errorf("Failed to load response (%s)", err.Error())
		}
	}
	// Special case for "Location" header, assume that if there is a location there is no body
	loc := resp.Header.Get("Location")
	if len(loc) > 0 {
		var bodyMap = make(map[string]interface{})
		bodyMap["Location"] = loc
		respBody = interface{}(bodyMap)
	}
	return respBody, err
}
