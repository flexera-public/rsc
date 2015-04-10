package cm15

import (
	"log"

	"gopkg.in/rightscale/rsc.v1-unstable/cmd" // import "gopkg.in/rightscale/rsc.v1-unstable/cmd"
	"gopkg.in/rightscale/rsc.v1-unstable/rsapi" // import "gopkg.in/rightscale/rsc.v1-unstable/rsapi"
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
func New(accountId int, host string, auth rsapi.Authenticator, logger *log.Logger,
	client rsapi.HttpClient) (*Api, error) {
	return fromApi(rsapi.New(accountId, host, auth, logger, client))
}

// NewRL10 returns a API 1.5 client that uses the information stored in /var/run/rightlink/secret to do
// auth and configure the host. The client behaves identically to the new returned by New in
// all other regards.
func NewRL10(logger *log.Logger, client rsapi.HttpClient) (*Api, error) {
	return fromApi(rsapi.NewRL10(logger, client))
}

// Build client from command line
func FromCommandLine(cmdLine *cmd.CommandLine) (*Api, error) {
	return fromApi(rsapi.FromCommandLine(cmdLine))
}

// Wrap generic client into API 1.5 client
func fromApi(api *rsapi.Api, err error) (*Api, error) {
	if err != nil {
		return nil, err
	}
	api.Metadata = GenMetadata
	return &Api{api}, nil
}
