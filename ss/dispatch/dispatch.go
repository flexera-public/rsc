// Common dispatcher to all ss clients
package dispatch // import "gopkg.in/rightscale/rsc.v1-unstable/ss/dispatch"

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"gopkg.in/rightscale/rsc.v1-unstable/rsapi"
)

// Interface to API client
type ApiClient interface {
	PerformRequest(*http.Request) (*http.Response, error)
}

// Information needed to do the dispatch
type RequestDetails struct {
	HttpMethod            string
	Host                  string
	Url                   string
	Params                rsapi.ApiParams
	Payload               rsapi.ApiParams
	AccountId             int
	FetchLocationResource bool
}

// Dispatch request to appropriate low-level method
func Dispatch(details *RequestDetails, client ApiClient) (*http.Response, error) {
	var u = url.URL{Host: details.Host, Path: details.Url}
	var params = details.Params
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
	var payload = details.Payload
	var jsonBytes []byte
	if payload != nil {
		var err error
		if jsonBytes, err = json.Marshal(payload); err != nil {
			return nil, fmt.Errorf("Failed to serialize request body - %s", err.Error())
		}
	}
	var req, err = http.NewRequest(details.HttpMethod, u.String(), bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-API-Version", "1.0")
	req.Header.Set("Content-Type", "application/json")
	if details.AccountId > 0 {
		req.Header.Set("X-Account", strconv.Itoa(details.AccountId))
	}
	resp, err := client.PerformRequest(req)
	if details.FetchLocationResource {
		var loc = resp.Header.Get("Location")
		if loc != "" {
			details.HttpMethod = "GET"
			details.Params = rsapi.ApiParams{}
			details.Payload = rsapi.ApiParams{}
			details.FetchLocationResource = false
			resp, err = Dispatch(details, client)
		}
	}
	return resp, err
}
