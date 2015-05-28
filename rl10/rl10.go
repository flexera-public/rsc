package rl10

import (
	"log"

	"github.com/rightscale/rsc/cmd"
	"github.com/rightscale/rsc/rsapi"
)

// RightLink 10 client
// Just a vanilla RightScale API client.
type Api struct {
	*rsapi.Api
}

// New returns a client that uses RL10 authentication.
// logger and client are optional.
// accountId, host and auth arguments are not used.
// If no HTTP client is specified then the default client is used.
func New(host string, auth rsapi.Authenticator, logger *log.Logger,
	client rsapi.HttpClient) (*Api, error) {
	var err error
	if auth != nil {
		auth.SetHost(host)
		err = auth.CanAuthenticate()
	}
	return fromApi(rsapi.New(host, auth, logger, client), err)
}

// NewRL10 returns a RL10 client that uses the information stored in /var/run/rightlink/secret to do
// auth and configure the host. The client behaves identically to the one returned by New in
// all other aspects.
func NewRL10(logger *log.Logger, client rsapi.HttpClient) (*Api, error) {
	return fromApi(rsapi.NewRL10(logger, client))
}

// Build client from command line
func FromCommandLine(cmdLine *cmd.CommandLine) (*Api, error) {
	return fromApi(rsapi.FromCommandLine(cmdLine))
}

// Wrap generic client into RL10 client
func fromApi(api *rsapi.Api, err error) (*Api, error) {
	if err != nil {
		return nil, err
	}
	api.Metadata = GenMetadata
	return &Api{api}, nil
}
