package rl10

import (
	"fmt"
	"net/http"

	"github.com/rightscale/rsc/rsapi"
)

// BuildRequest builds a HTTP request from a resource name and href and an action name and
// parameters.
// It is intended for generic clients that need to consume APIs in a generic maner.
// The method builds an HTTP request that can be fed to PerformRequest.
func (a *API) BuildRequest(resource, action, href string, params rsapi.APIParams) (*http.Request, error) {
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
	actionURL, err := act.URL(vars)
	if err != nil {
		return nil, err
	}
	payload, params := rsapi.IdentifyParams(act, params)
	return a.BuildHTTPRequest(actionURL.HTTPMethod, actionURL.Path, "", params, payload)
}
