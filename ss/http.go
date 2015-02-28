package ss

import (
	"fmt"
	"net/http"

	"github.com/rightscale/rsc/metadata"
	"github.com/rightscale/rsc/rsapi"
	"github.com/rightscale/rsc/ss/ssc"
	"github.com/rightscale/rsc/ss/ssd"
	"github.com/rightscale/rsc/ss/ssm"
)

// Do is a generic client method and is meant for command line tools and other higher level clients.
// It accepts the service name (one of "designer", "catalog" or "manager"), a resource href, the
// name of an action and resource and the action parameters.
// The method makes the request and returns the raw HTTP response or an error.
// The LoadResponse method can be used to load the response body if needed.
func (a *Api) Do(service, resource, action, href string, params rsapi.ApiParams) (*http.Response, error) {
	var md map[string]*metadata.Resource
	var dispatch DispatchFunc
	switch service {
	case "designer":
		md = ssd.GenMetadata
		dispatch = a.designerClient.Dispatch
	case "catalog":
		md = ssc.GenMetadata
		dispatch = a.catalogClient.Dispatch
	case "manager":
		md = ssm.GenMetadata
		dispatch = a.managerClient.Dispatch
	default:
		return nil, fmt.Errorf("Unknown service %s, must be one of 'designer', 'catalog' or 'manager'", service)
	}

	// First lookup metadata
	var res, ok = md[resource]
	if !ok {
		return nil, fmt.Errorf("No resource with name '%s' in service %s", resource, service)
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
	return dispatch(actionUrl.HttpMethod, actionUrl.Path, queryParams, payloadParams)
}
