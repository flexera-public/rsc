package rsapi15

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/rightscale/rsc/metadata"
	"github.com/rightscale/rsc/rsapi"
)

// Do is a generic client method and is meant for command line tools and other higher level clients.
// It accepts a resource href, the name of an action and resource and the action parameters.
// The method makes the request and returns the raw HTTP response or an error.
// The LoadResponse method can be used to load the response body if needed.
func (a *Api15) Do(resource, action, href string, params rsapi.ApiParams) (*http.Response, error) {
	// First lookup metadata
	var res, ok = GenMetadata[resource]
	if !ok {
		return nil, fmt.Errorf("No resource with name '%s'", resource)
	}
	var act *metadata.Action
	if act == nil {
		return nil, fmt.Errorf("No action with name '%s' on %s", action, resource)
	}

	// Now lookup action request HTTP method, url, params and payload.
	var vars, err = res.ExtractVariables(href)
	if err != nil {
		return nil, err
	}
	actionUrl, err := act.Url(vars)
	if err != nil {
		return nil, err
	}
	var payloadParams = make(rsapi.ApiParams, len(act.PayloadParamNames))
	for _, n := range act.PayloadParamNames {
		payloadParams[n] = params[n]
	}
	var queryParams = make(rsapi.ApiParams, len(act.QueryParamNames))
	for _, n := range act.QueryParamNames {
		queryParams[n] = params[n]
	}
	return a.Dispatch(actionUrl.HttpMethod, actionUrl.Path, queryParams, payloadParams)
}

// Dispatch request to appropriate low-level method
func (a *Api15) Dispatch(method, actionUrl string, params, payload rsapi.ApiParams) (*http.Response, error) {
	switch method {
	case "GET":
		return a.GetRaw(actionUrl, params)
	case "POST":
		return a.PostRaw(actionUrl, params, payload)
	case "PUT":
		return nil, a.Put(actionUrl, params, payload)
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
	return a.makeRequest("GET", uri, params, nil)
}

// Low-level POST request that loads response JSON into generic object
// Any "Location" header present in the HTTP response is returned in a map under the "Location" key.
func (a *Api15) Post(uri string, params rsapi.ApiParams, payload rsapi.ApiParams) (interface{}, error) {
	resp, err := a.PostRaw(uri, params, payload)
	if err != nil {
		return nil, err
	}
	return a.LoadResponse(resp)
}

// Low-level POST request
func (a *Api15) PostRaw(uri string, params rsapi.ApiParams, payload rsapi.ApiParams) (*http.Response, error) {
	return a.makeRequest("POST", uri, params, payload)
}

// Low-level PUT request
func (a *Api15) Put(uri string, params rsapi.ApiParams, payload rsapi.ApiParams) error {
	_, err := a.makeRequest("PUT", uri, params, payload)
	return err
}

// Low-level DELETE request
func (a *Api15) Delete(uri string) error {
	_, err := a.makeRequest("DELETE", uri, nil, nil)
	return err
}

// Helper function that signs, makes and logs HTTP request
func (a *Api15) makeRequest(verb, uri string, params rsapi.ApiParams, payload rsapi.ApiParams) (*http.Response, error) {
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
	var jsonBytes []byte
	if payload != nil {
		var err error
		if jsonBytes, err = json.Marshal(payload); err != nil {
			return nil, fmt.Errorf("Failed to serialize request body - %s", err.Error())
		}
	}
	var req, err = http.NewRequest(verb, u.String(), bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-API-Version", "1.5")
	req.Header.Set("Content-Type", "application/json")
	if a.AccountId > 0 {
		req.Header.Set("X-Account", strconv.Itoa(a.AccountId))
	}
	resp, err := a.PerformRequest(req)
	if a.FetchLocationResource {
		var loc = resp.Header.Get("Location")
		if loc != "" {
			resp, err = a.makeRequest("GET", loc, rsapi.ApiParams{}, rsapi.ApiParams{})
		}
	}
	return resp, err
}
