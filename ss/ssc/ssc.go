package ssc // import "gopkg.in/rightscale/rsc.v2/ss/ssc"

import (
	"log"
	"net/http"

	"gopkg.in/rightscale/rsc.v2/dispatch"
	"gopkg.in/rightscale/rsc.v2/rsapi"
)

// Self-Service catalog client
type Api struct {
	*rsapi.Api
}

// New returns a Self-Service catalog API client.
// It makes a test API request and returns an error if authentication fails.
// logger is optional.
// If no HTTP client is specified then the default client is used.
func New(h string, a rsapi.Authenticator, l *log.Logger, c rsapi.HttpClient) *Api {
	api := rsapi.New(h, a, l, c)
	api.Metadata = GenMetadata
	ssApi := Api{Api: api}
	return &ssApi
}

// Dispatch request to appropriate low-level method
func (a *Api) Dispatch(method, actionUrl string, params, payload rsapi.ApiParams) (*http.Response, error) {
	details := dispatch.RequestDetails{
		HttpMethod:            method,
		Host:                  a.Host,
		Url:                   "/catalog" + actionUrl,
		Params:                params,
		Payload:               payload,
		FetchLocationResource: a.FetchLocationResource,
	}
	return dispatch.Dispatch(&details, a)
}
