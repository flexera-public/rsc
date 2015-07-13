package rsapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rightscale/rsc/metadata"
)

// Log request, dump its content if required then make request and log response and dump it.
func (a *Api) PerformRequest(req *http.Request) (*http.Response, error) {
	// Sign last so auth headers don't get printed or logged
	if a.Auth != nil {
		if err := a.Auth.Sign(req); err != nil {
			return nil, err
		}
	}
	resp, err := a.Client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}

// IdentifyParams organizes the given params in two groups: the payload params and the query params.
func IdentifyParams(a *metadata.Action, params ApiParams) (payloadParams ApiParams, queryParams ApiParams) {
	payloadParamNames := a.PayloadParamNames()
	payloadParams = make(ApiParams)
	for _, n := range payloadParamNames {
		if p, ok := params[n]; ok {
			payloadParams[n] = p
		}
	}
	queryParamNames := a.QueryParamNames()
	queryParams = make(ApiParams)
	for _, n := range queryParamNames {
		if p, ok := params[n]; ok {
			queryParams[n] = p
		}
	}
	return payloadParams, queryParams
}

// Deserialize JSON response into generic object.
// If the response has a "Location" header then the returned object is a map with one key "Location"
// containing the value of the header.
func (a *Api) LoadResponse(resp *http.Response) (interface{}, error) {
	defer resp.Body.Close()
	var respBody interface{}
	jsonResp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response (%s)", err)
	}
	if len(jsonResp) > 0 {
		err = json.Unmarshal(jsonResp, &respBody)
		if err != nil {
			return nil, fmt.Errorf("Failed to load response (%s)", err)
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
