package rsapi15

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"time"
)

// Generic API parameters type
type ApiParams map[string]interface{}

// Use interface instead of raw http.Client to ease testing
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// Do is a generic client method and is meant for command line tools and other higher level clients.
// It accepts a resource or resource collection href (e.g. "/api/servers"), the name of an action
// (e.g. "create") and the request parameters.
// The method makes the request and returns the raw HTTP response or an error.
// The LoadResponse method can be used to load the response body if needed.
func (a *Api15) Do(actionUrl, action string, params ApiParams) (*http.Response, error) {
	// First figure out action verb and uri
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
		for _, r := range api_metadata {
			for _, a := range r.Actions {
				if a.Name == action {
					method = a.HttpMethod
					break
				}
			}
		}
	}

	// Now call corresponding low-level request method
	switch method {
	case "GET":
		return a.GetRaw(actionUrl, params)
	case "POST":
		return a.PostRaw(actionUrl, params)
	case "PUT":
		return nil, a.Put(actionUrl, params)
	case "DELETE":
		return nil, a.Delete(actionUrl)
	}
	panic("Unknown API 1.5 HTTP method " + method)
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
	var u = url.URL{
		Scheme: "https",
		Host:   a.Host,
		Path:   uri,
	}
	if params != nil {
		var values = u.Query()
		for n, p := range params {
			switch t := p.(type) {
			case string:
				values.Set(n, t)
			case []string:
				for _, e := range t {
					values.Add(n, e)
				}
			case map[string]string:
				for pn, e := range t {
					values.Add(pn, e)
				}
			default:
				return nil, fmt.Errorf("Invalid param value <%+v>, value must be a string or an array of strings", p)
			}
		}
		u.RawQuery = values.Encode()
	}
	var sUrl = u.String()
	var req, err = http.NewRequest(verb, sUrl, nil)
	if err != nil {
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
		if err == nil {
			fmt.Printf("REQUEST\n-------\n%s\n", b)
		}
	}
	// Sign last so auth headers don't get printed or logged
	if err = a.Auth.Sign(req, a.Host); err != nil {
		return nil, err
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
		var b, err = httputil.DumpResponse(resp, false)
		if err == nil {
			fmt.Printf("RESPONSE\n--------\n%s", b)
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
