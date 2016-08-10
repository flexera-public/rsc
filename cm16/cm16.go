package cm16

import (
	"github.com/rightscale/rsc/cmd"
	"github.com/rightscale/rsc/rsapi"
)

// API 1.6 client
// Just a vanilla RightScale API client.
type API struct {
	*rsapi.API
}

// New returns a API 1.6 client.
// It makes a test request to API 1.5 and returns an error if authentication fails.
// host may be blank in which case client attempts to resolve it using auth.
func New(host string, auth rsapi.Authenticator) *API {
	return fromAPI(rsapi.New(host, auth))
}

// NewRL10 returns a API 1.6 client that uses the information stored in /var/run/rightlink/secret to do
// auth and configure the host. The client behaves identically to the new returned by New in
// all other regards.
func NewRL10() (*API, error) {
	raw, err := rsapi.NewRL10()
	if err != nil {
		return nil, err
	}
	return fromAPI(raw), nil
}

// FromCommandLine builds a client from the command line.
func FromCommandLine(cmdLine *cmd.CommandLine) (*API, error) {
	raw, err := rsapi.FromCommandLine(cmdLine)
	if err != nil {
		return nil, err
	}
	return fromAPI(raw), nil
}

// Wrap generic client into API 1.6 client
func fromAPI(api *rsapi.API) *API {
	api.Metadata = GenMetadata
	return &API{api}
}
