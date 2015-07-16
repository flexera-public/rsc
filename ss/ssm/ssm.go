package ssm // import "gopkg.in/rightscale/rsc.v3/ss/ssm"

import (
	"net/http"

	"gopkg.in/rightscale/rsc.v3/dispatch"
	"gopkg.in/rightscale/rsc.v3/rsapi"
)

// Self-Service manager client
type Api struct {
	*rsapi.Api
}

// New returns a Self-Service catalog API client.
// It makes a test API request and returns an error if authentication fails.
func New(h string, a rsapi.Authenticator) *Api {
	api := rsapi.New(h, a)
	api.Metadata = GenMetadata
	ssApi := Api{Api: api}
	return &ssApi
}

// Dispatch request to appropriate low-level method
func (a *Api) Dispatch(method, actionUrl string, params, payload rsapi.ApiParams) (*http.Response, error) {
	details := dispatch.RequestDetails{
		HttpMethod:            method,
		Host:                  a.Host,
		Url:                   "/manager" + actionUrl,
		Params:                params,
		Payload:               payload,
		FetchLocationResource: a.FetchLocationResource,
	}
	return dispatch.Dispatch(&details, a)
}
