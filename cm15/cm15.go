package cm15

import (
	"log"

	"github.com/rightscale/rsc/cmd"
	"github.com/rightscale/rsc/rsapi"
)

// Api 1.5 client
// Just a vanilla RightScale API client.
type Api struct {
	*rsapi.Api
}

// New returns a API 1.5 client that uses User oauth authentication.
// logger and client are optional.
// host may be blank in which case client attempts to resolve it using auth.
// If no HTTP client is specified then the default client is used.
func New(host string, auth rsapi.Authenticator, logger *log.Logger,
	client rsapi.HttpClient) *Api {
	return fromApi(rsapi.New(host, auth, logger, client))
}

// NewRL10 returns a API 1.5 client that uses the information stored in /var/run/rightlink/secret to do
// auth and configure the host. The client behaves identically to the new returned by New in
// all other regards.
func NewRL10(logger *log.Logger, client rsapi.HttpClient) (*Api, error) {
	api, err := rsapi.NewRL10(logger, client)
	if err != nil {
		return nil, err
	}
	return fromApi(api), nil
}

// Build client from command line
func FromCommandLine(cmdLine *cmd.CommandLine) (*Api, error) {
	api, err := rsapi.FromCommandLine(cmdLine)
	if err != nil {
		return nil, err
	}
	return fromApi(api), nil
}

// Wrap generic client into API 1.5 client
func fromApi(api *rsapi.Api) *Api {
	api.Metadata = GenMetadata
	return &Api{api}
}
