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
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// RightScale API 1.5 client
// Instances of this struct should be created through `New`.
type Api15 struct {
	AccountId             int           // Account in which client is currently operating
	Auth                  Authenticator // Authenticator, signs requests for auth
	Logger                *log.Logger   // Optional logger, if specified requests and responses get logged
	Host                  string        // API host, e.g. "us-3.rightscale.com"
	Client                HttpClient    // Underlying http client
	DumpRequestResponse   bool          // Whether to dump HTTP requests and responses to STDOUT
	FetchLocationResource bool          // Whether to fetch resource pointed by Location header
}

// Use interface instead of raw http.Client to ease testing
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// New returns a API 1.5 client that uses User oauth authentication.
// logger and client are optional.
// host may be blank in which case client attempts to resolve it using auth.
// If no HTTP client is specified then the default client is used.
func New(accountId int, refreshToken string, host string, logger *log.Logger,
	client HttpClient) (*Api15, error) {
	if client == nil {
		client = http.DefaultClient
	}
	var auth = OAuthAuthenticator{
		RefreshToken: refreshToken,
		RefreshAt:    time.Now().Add(-1 * time.Hour),
		Client:       client,
	}
	if host == "" {
		if resolved, err := auth.ResolveHost(accountId); err != nil {
			return nil, err
		} else {
			host = resolved
		}
	}
	return &Api15{
		AccountId: accountId,
		Auth:      &auth,
		Logger:    logger,
		Host:      host,
		Client:    client,
	}, nil
}

// NewRL10 returns a API 1.5 client that uses the information stored in /var/run/rll-secret to do
// auth and configure the host. The client behaves identically to the new returned by New in
// all other regards.
func NewRL10(logger *log.Logger, client HttpClient) (*Api15, error) {
	var rllConfig, err = ioutil.ReadFile("/var/run/rll-secret")
	if err != nil {
		return nil, fmt.Errorf("Failed to load RLL config: %s", err.Error())
	}
	var port string
	var secret string
	for _, line := range strings.Split(string(rllConfig), "\n") {
		var elems = strings.Split(line, "=")
		if len(elems) != 2 {
			return nil, fmt.Errorf("Invalid RLL configuration line '%s'", line)
		}
		switch elems[0] {
		case "RS_RLL_PORT":
			port = elems[1]
			if _, err := strconv.Atoi(elems[1]); err != nil {
				return nil, fmt.Errorf("Invalid port value '%s'", port)
			}
		case "RS_RLL_SECRET":
			secret = elems[1]
		}
	}
	var auth = RL10Authenticator{secret}
	var host = "localhost:" + port
	return &Api15{
		Auth:   &auth,
		Logger: logger,
		Host:   host,
		Client: client,
	}, nil
}

// Generic API parameters type
type ApiParams map[string]interface{}

// Do is a generic client method and is meant for command line tools and other higher level clients.
// It accepts a resource or resource collection href (e.g. "/api/servers"), the name of an action
// (e.g. "create") and the request parameters (payload or query strings).
// The method makes the request and returns the raw HTTP response or an error.
// The LoadResponse method can be used to load the response body if needed.
func (a *Api15) Do(href, action string, params ApiParams) (*http.Response, error) {
	// Allow "servers" instead of "/api/servers"
	if !strings.HasPrefix(href, "/api/") {
		href = "/api/" + href
	}

	// First figure out action verb and uri
	var uri = href
	var method string
	switch action {
	case "show", "index":
		method = "GET"
	case "create":
		method = "POST"
	case "update":
		method = "PUT"
	case "destroy":
		method = "DELETE"
	default:
		if pair, ok := actionMap[action]; ok {
			uri = href + pair[0]
			method = pair[1]
		}
	}

	// Now call corresponding low-level request method
	switch method {
	case "GET":
		return a.GetRaw(uri, params)
	case "POST":
		return a.PostRaw(uri, params)
	case "PUT":
		return nil, a.Put(uri, params)
	case "DELETE":
		return nil, a.Delete(uri)
	default:
		return nil, fmt.Errorf("Unknown href '%s' or action '%s'", href, action)
	}
}

// Low-level GET request that loads response JSON into generic object
func (a *Api15) Get(uri string, params ApiParams) (interface{}, error) {
	resp, err := a.GetRaw(uri, params)
	if err != nil {
		return nil, err
	}
	return a.LoadResponse(resp)
}

// Low-level GET request
func (a *Api15) GetRaw(uri string, params ApiParams) (*http.Response, error) {
	return a.makeRequest("GET", uri, params)
}

// Low-level POST request that loads response JSON into generic object
// Any "Location" header present in the HTTP response is returned in a map under the "Location" key.
func (a *Api15) Post(uri string, body ApiParams) (interface{}, error) {
	resp, err := a.PostRaw(uri, body)
	if err != nil {
		return nil, err
	}
	return a.LoadResponse(resp)
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

// Deserialize JSON response into generic object.
// If the response has a "Location" header then the returned object is a map with one key "Location"
// containing the value of the header.
func (a *Api15) LoadResponse(resp *http.Response) (interface{}, error) {
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

// Helper function that signs, makes and logs HTTP request
func (a *Api15) makeRequest(verb, uri string, params ApiParams) (*http.Response, error) {
	var body = bytes.NewBuffer([]byte{})
	if (verb == "POST" || verb == "PUT") && params != nil {
		var jsonBytes, err = json.Marshal(params)
		if err != nil {
			return nil, fmt.Errorf("Failed to serialize response (%s)", err.Error())
		}
		body = bytes.NewBuffer([]byte(jsonBytes))
	}
	var u = url.URL{
		Scheme: "https",
		Host:   a.Host,
		Path:   uri,
	}
	if verb == "GET" && params != nil {
		for n, p := range params {
			switch t := p.(type) {
			case string:
				u.Query().Set(n, t)
			case []string:
				for _, e := range t {
					u.Query().Add(n, e)
				}
			default:
				return nil, fmt.Errorf("Invalid param value <%+v>, value must be a string or an array of strings", p)
			}
		}
	}
	var sUrl = u.String()
	var req, err = http.NewRequest(verb, sUrl, body)
	if err != nil {
		return nil, err
	}
	if err = a.Auth.Sign(req, a.Host); err != nil {
		return nil, err
	}
	req.Header.Set("X-API-Version", "1.5")
	if a.AccountId > 0 {
		req.Header.Set("X-Account", strconv.Itoa(a.AccountId))
	}
	var id string
	var startedAt time.Time
	if a.Logger != nil {
		startedAt = time.Now()
		b := make([]byte, 6)
		io.ReadFull(rand.Reader, b)
		id = base64.StdEncoding.EncodeToString(b)
		a.Logger.Printf("[%s] %s %s", id, req.Method, sUrl)
	}
	if a.DumpRequestResponse {
		var b, err = httputil.DumpRequest(req, true)
		if err != nil {
			fmt.Printf("REQUEST\n-------\n%v\n", b)
		}
	}
	resp, err := a.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if a.Logger != nil {
		d := time.Since(startedAt)
		a.Logger.Printf("[%s] %s in %s", id, resp.Status, d.String())
	}
	if a.DumpRequestResponse {
		var b, err = httputil.DumpResponse(resp, true)
		if err != nil {
			fmt.Printf("RESPONSE\n--------\n%v\n", b)
		}
	}
	if a.FetchLocationResource {
		var loc = resp.Header.Get("Location")
		if loc != "" {
			resp, err = a.makeRequest("GET", loc, ApiParams{})
		}
	}
	return resp, err
}
