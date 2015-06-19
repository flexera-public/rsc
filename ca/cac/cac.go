package cac

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

// New returns a client that uses User oauth authentication.
// logger and client are optional.
// host may be blank in which case client attempts to resolve it using auth.
// If no HTTP client is specified then the default client is used.
func New(accountId int, host string, auth rsapi.Authenticator, logger *log.Logger,
	client rsapi.HttpClient) (*Api, error) {
	api, err := rsapi.New(accountId, host, auth, logger, client)
	if err != nil {
		return nil, err
	}
	api.Metadata = GenMetadata
	return &Api{api}, nil
}

// Dispatch request to appropriate low-level method
func (a *Api) Dispatch(method, actionUrl string, params, payload rsapi.ApiParams) (*http.Response, error) {
	details := dispatch.RequestDetails{
		HttpMethod:            method,
		Host:                  a.Host,
		Url:                   actionUrl,
		Params:                params,
		Payload:               payload,
		AccountId:             a.AccountId,
		FetchLocationResource: a.FetchLocationResource,
	}
	return dispatch.Dispatch(&details, a)
}
