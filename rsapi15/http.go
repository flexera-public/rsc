package rsapi15

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/rightscale/rsc/rsapi"
)

// Do is a generic client method and is meant for command line tools and other higher level clients.
// It accepts a resource or resource collection href (e.g. "/api/servers"), the name of an action
// (e.g. "create") and the request parameters.
// The method makes the request and returns the raw HTTP response or an error.
// The LoadResponse method can be used to load the response body if needed.
func (a *Api15) Do(actionUrl, action string, params rsapi.ApiParams) (*http.Response, error) {
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
					method = a.PathPatterns[0].HttpMethod
					break
				}
			}
		}
	}
	return a.Dispatch(method, actionUrl, params)
}

// Dispatch request to appropriate low-level method
func (a *Api15) Dispatch(method, actionUrl string, params rsapi.ApiParams) (*http.Response, error) {
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
	return nil, fmt.Errorf("Unsupported HTTP method %s", method)
}

// Low-level GET request that loads response JSON into generic object
func (a *Api15) Get(uri string, params rsapi.ApiParams) (interface{}, error) {
	resp, err := a.GetRaw(uri, params)
	if err != nil {
		return nil, err
	}
	return a.LoadResponse(resp)
}

// Low-level GET request
func (a *Api15) GetRaw(uri string, params rsapi.ApiParams) (*http.Response, error) {
	return a.makeRequest("GET", uri, params)
}

// Low-level POST request that loads response JSON into generic object
// Any "Location" header present in the HTTP response is returned in a map under the "Location" key.
func (a *Api15) Post(uri string, body rsapi.ApiParams) (interface{}, error) {
	resp, err := a.PostRaw(uri, body)
	if err != nil {
		return nil, err
	}
	return a.LoadResponse(resp)
}

// Low-level POST request
func (a *Api15) PostRaw(uri string, body rsapi.ApiParams) (*http.Response, error) {
	return a.makeRequest("POST", uri, body)
}

// Low-level PUT request
func (a *Api15) Put(uri string, body rsapi.ApiParams) error {
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
func (a *Api15) makeRequest(verb, uri string, params rsapi.ApiParams) (*http.Response, error) {
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
	var req, err = http.NewRequest(verb, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-API-Version", "1.5")
	if a.AccountId > 0 {
		req.Header.Set("X-Account", strconv.Itoa(a.AccountId))
	}
	resp, err := a.PerformRequest(req)
	if a.FetchLocationResource {
		var loc = resp.Header.Get("Location")
		if loc != "" {
			resp, err = a.makeRequest("GET", loc, rsapi.ApiParams{})
		}
	}
	return resp, err
}
