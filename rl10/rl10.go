package rl10

import (
	"github.com/rightscale/rsc/cmd"
	"github.com/rightscale/rsc/rsapi"
)

// API is the RightLink 10 client.
type API struct {
	*rsapi.API
}

// New returns a client that uses RL10 authentication.
// accountId, host and auth arguments are not used.
// If no HTTP client is specified then the default client is used.
func New(host string, auth rsapi.Authenticator) *API {
	return fromAPI(rsapi.New(host, auth))
}

// NewRL10 returns a RL10 client that uses the information stored in /var/run/rightlink/secret to do
// auth and configure the host. The client behaves identically to the one returned by New in
// all other aspects.
func NewRL10() (*API, error) {
	raw, err := rsapi.NewRL10()
	if err != nil {
		return nil, err
	}
	return fromAPI(raw), nil
}

// FromCommandLine builds the API client from the command line.
func FromCommandLine(cmdLine *cmd.CommandLine) (*API, error) {
	raw, err := rsapi.FromCommandLine(cmdLine)
	if err != nil {
		return nil, err
	}
	return fromAPI(raw), nil
}

// Wrap generic client into RL10 client
func fromAPI(api *rsapi.API) *API {
	api.Metadata = GenMetadata
	return &API{api}
}
