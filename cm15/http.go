package cm15

import (
	"fmt"
	"net/http"

	"github.com/rightscale/rsc/rsapi"
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
	payloadParams, queryParams := rsapi.IdentifyParams(act, params)
	return a.BuildHTTPRequest(actionUrl.HttpMethod, actionUrl.Path, "1.5", queryParams, payloadParams)
}
