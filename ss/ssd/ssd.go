package ssd

import (
	"log"
	"net/http"

	"github.com/rightscale/rsc/dispatch"
	"github.com/rightscale/rsc/rsapi"
)

// Self-Service designer client
type Api struct {
	*rsapi.Api
}

// New returns a client.
// auth must be created with rsapi.NewSSAuthenticator.
// logger and client are optional.
// host may be blank in which case client attempts to resolve it using auth.
// If no HTTP client is specified then the default client is used.
func New(host string, auth rsapi.Authenticator, logger *log.Logger,
	client rsapi.HttpClient) *Api {
	api := rsapi.New(host, auth, logger, client)
	api.Metadata = GenMetadata
	return &Api{api}
}

// Dispatch request to appropriate low-level method
func (a *Api) Dispatch(method, actionUrl string, params, payload rsapi.ApiParams) (*http.Response, error) {
	details := dispatch.RequestDetails{
		HttpMethod:            method,
		Host:                  a.Host,
		Url:                   "/designer" + actionUrl,
		Params:                params,
		Payload:               payload,
		FetchLocationResource: a.FetchLocationResource,
	}
	return dispatch.Dispatch(&details, a)
}
