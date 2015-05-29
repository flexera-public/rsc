package ssm // import "gopkg.in/rightscale/rsc.v2/ss/ssm"

import (
	"log"
	"net/http"

	"gopkg.in/rightscale/rsc.v2/dispatch"
	"gopkg.in/rightscale/rsc.v2/rsapi"
)

// Self-Service manager client
type Api struct {
	*rsapi.Api
}

// New returns a Self-Service catalog API client.
// It makes a test API request and returns an error if authentication fails.
// logger is optional.
// If no HTTP client is specified then the default client is used.
func New(h string, a rsapi.Authenticator, l *log.Logger, c rsapi.HttpClient) (*Api, error) {
	if a != nil {
		a.SetHost(h)
		if err := a.CanAuthenticate(); err != nil {
			return nil, err
		}
	}
	api := rsapi.New(h, a, l, c)
	api.Metadata = GenMetadata
	ssApi := Api{Api: api}
	return &ssApi, nil
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
