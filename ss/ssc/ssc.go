package ssc

import (
	"log"
	"net/http"

	"github.com/rightscale/rsc/cmd"
	"github.com/rightscale/rsc/rsapi"
	"github.com/rightscale/rsc/ss/dispatch"
)

// Self-Service catalog client
type Api struct {
	*rsapi.Api
}

// New returns a client that uses User oauth authentication.
// logger and client are optional.
// host may be blank in which case client attempts to resolve it using auth.
// If no HTTP client is specified then the default client is used.
func New(accountId int, refreshToken string, host string, logger *log.Logger,
	client rsapi.HttpClient) (*Api, error) {
	var base, err = rsapi.New(accountId, refreshToken, host, logger, client)
	if err != nil {
		return nil, err
	}
	return fromBase(base), nil
}

// Build client from command line
func FromCommandLine(cmdLine *cmd.CommandLine) (*Api, error) {
	var base, err = rsapi.FromCommandLine(cmdLine)
	if err != nil {
		return nil, err
	}
	return fromBase(base), nil
}

// Wrap generic client
func fromBase(api *rsapi.Api) *Api {
	api.Metadata = GenMetadata
	api.Auth = &rsapi.SSAuthenticator{api.Auth.(*rsapi.OAuthAuthenticator), api.AccountId}
	return &Api{api}
}

// Dispatch request to appropriate low-level method
func (a *Api) Dispatch(method, actionUrl string, params, payload rsapi.ApiParams) (*http.Response, error) {
	var details = dispatch.RequestDetails{
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
