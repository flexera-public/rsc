package rsapi16

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/rightscale/rsc/rsapi"
)

// Dispatch request, used by generated code
func (a *Api16) Dispatch(method, actionUrl string, params rsapi.ApiParams) (*http.Response, error) {
	return a.GetRaw(actionUrl, params)
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

// Deserialize JSON response into generic object.
func (a *Api16) LoadResponse(resp *http.Response) (interface{}, error) {
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
	return respBody, err
}

// Helper function that signs, makes and logs HTTP request
func (a *Api16) makeRequest(verb, uri string, params rsapi.ApiParams) (*http.Response, error) {
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
	req.Header.Set("X-API-Version", "1.6")
	if a.AccountId > 0 {
		req.Header.Set("X-Account", strconv.Itoa(a.AccountId))
	}
	return a.PerformRequest(req)
}
