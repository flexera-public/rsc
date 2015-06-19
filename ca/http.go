package ca

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/rightscale/rsc/rsapi"
	"net/http"
	"net/url"
	"strconv"
)

// BuildRequest builds a HTTP request from a resource name and href and an action name and
// parameters.
// It is intended for generic clients that need to consume APIs in a generic maner.
// The method builds an HTTP request that can be fed to PerformRequest.
func (a *Api) BuildRequest(resource, action, href string, params rsapi.ApiParams) (*http.Request, error) {
	// First lookup metadata
	res, ok := GenMetadata[resource]
	if !ok {
		return nil, fmt.Errorf("No resource with name '%s'", resource)
	}
	act := res.GetAction(action)
	if act == nil {
		return nil, fmt.Errorf("No action with name '%s' on %s", action, resource)
	}

	// Now lookup action request HTTP method, url, params and payload.
	vars, err := res.ExtractVariables(href)
	if err != nil {
		return nil, err
	}
	actionUrl, err := act.Url(vars)
	if err != nil {
		return nil, err
	}
	payloadParams := make(rsapi.ApiParams, len(act.PayloadParamNames))
	for _, n := range act.PayloadParamNames {
		payloadParams[n] = params[n]
	}
	queryParams := make(rsapi.ApiParams, len(act.QueryParamNames))
	for _, n := range act.QueryParamNames {
		queryParams[n] = params[n]
	}
	return a.buildHttpRequest(actionUrl.HttpMethod, actionUrl.Path, queryParams, payloadParams)
}

// Helper function that signs, makes and logs HTTP request.
// Used by generated client code.
func (a *Api) Dispatch(verb, uri string, params, payload rsapi.ApiParams) (*http.Response, error) {
	req, err := a.buildHttpRequest(verb, uri, params, payload)
	if err != nil {
		return nil, err
	}
	resp, err := a.PerformRequest(req)
	if a.FetchLocationResource {
		loc := resp.Header.Get("Location")
		if loc != "" {
			resp, err = a.Dispatch("GET", loc, rsapi.ApiParams{}, rsapi.ApiParams{})
		}
	}
	return resp, err
}

// Helper function that puts together and HTTP request from its verb, uri and params.
func (a *Api) buildHttpRequest(verb, uri string, params rsapi.ApiParams, payload rsapi.ApiParams) (*http.Request, error) {
	u := url.URL{
		Host: a.Host,
		Path: uri,
	}
	if params != nil {
		values := u.Query()
		for n, p := range params {
			switch t := p.(type) {
			case string:
				values.Set(n, t)
			case int:
				values.Set(n, strconv.Itoa(t))
			case bool:
				values.Set(n, strconv.FormatBool(t))
			case []string:
				for _, e := range t {
					values.Add(n, e)
				}
			case []int:
				for _, e := range t {
					values.Add(n, strconv.Itoa(e))
				}
			case []bool:
				for _, e := range t {
					values.Add(n, strconv.FormatBool(e))
				}
			case map[string]string:
				for pn, e := range t {
					values.Add(fmt.Sprintf("%s[%s]", n, pn), e)
				}
			default:
				return nil, fmt.Errorf("Invalid param value <%+v>, value must be a string, an integer, a bool, an array of these types of a map of strings", p)
			}
		}
		u.RawQuery = values.Encode()
	}
	var jsonBytes []byte
	if payload != nil && len(payload) > 0 {
		var err error
		if jsonBytes, err = json.Marshal(payload); err != nil {
			return nil, fmt.Errorf("Failed to serialize request body - %s", err)
		}
	}
	req, err := http.NewRequest(verb, u.String(), bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-API-Version", "1.0")
	req.Header.Set("Content-Type", "application/json")
	if a.AccountId > 0 {
		req.Header.Set("X-Account", strconv.Itoa(a.AccountId))
	}
	return req, nil
}
