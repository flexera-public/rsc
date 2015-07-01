package rl10

import (
	"github.com/rightscale/rsc/cmd"
	"github.com/rightscale/rsc/rsapi"
)

// RightLink 10 client
// Just a vanilla RightScale API client.
type Api struct {
	*rsapi.Api
}

// New returns a client that uses RL10 authentication.
// accountId, host and auth arguments are not used.
// If no HTTP client is specified then the default client is used.
func New(host string, auth rsapi.Authenticator, client rsapi.HttpClient) *Api {
	return fromApi(rsapi.New(host, auth, client))
}

// NewRL10 returns a RL10 client that uses the information stored in /var/run/rightlink/secret to do
// auth and configure the host. The client behaves identically to the one returned by New in
// all other aspects.
func NewRL10(client rsapi.HttpClient) (*Api, error) {
	raw, err := rsapi.NewRL10(client)
	if err != nil {
		return nil, err
	}
	return fromApi(raw), nil
}

// Build client from command line
func FromCommandLine(cmdLine *cmd.CommandLine) (*Api, error) {
	raw, err := rsapi.FromCommandLine(cmdLine)
	if err != nil {
		return nil, err
	}
	return fromApi(raw), nil
}

// Wrap generic client into RL10 client
func fromApi(api *rsapi.Api) *Api {
	api.Metadata = GenMetadata
	return &Api{api}
}
