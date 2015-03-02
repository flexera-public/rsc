package rsapi16

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/rightscale/rsc/rsapi"
)

// Dispatch request, used by generated code
func (a *Api16) Dispatch(method, actionUrl string, queryParams, payloadParams rsapi.ApiParams) (*http.Response, error) {
	return a.GetRaw(actionUrl, queryParams)
}

// Low-level GET request that loads response JSON into generic object
func (a *Api16) Get(uri string, params rsapi.ApiParams) (interface{}, error) {
	resp, err := a.GetRaw(uri, params)
	if err != nil {
		return nil, err
	}
	return a.LoadResponse(resp)
}

// Low-level GET request
func (a *Api16) GetRaw(uri string, params rsapi.ApiParams) (*http.Response, error) {
	return a.makeRequest("GET", uri, params)
}

// Helper function that signs, makes and logs HTTP request
func (a *Api16) makeRequest(verb, uri string, params rsapi.ApiParams) (*http.Response, error) {
	u := url.URL{
		Scheme: "https",
		Host:   a.Host,
		Path:   uri,
	}
	if params != nil {
		values := u.Query()
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
	req, err := http.NewRequest(verb, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-API-Version", "1.6")
	if a.AccountId > 0 {
		req.Header.Set("X-Account", strconv.Itoa(a.AccountId))
	}
	return a.PerformRequest(req)
}
