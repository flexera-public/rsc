package cm16 // import "gopkg.in/rightscale/rsc.v3/cm16"

import (
	"gopkg.in/rightscale/rsc.v3/cmd"
	"gopkg.in/rightscale/rsc.v3/rsapi"
)

// Api 1.6 client
// Just a vanilla RightScale API client.
type Api struct {
	*rsapi.Api
}

// New returns a API 1.6 client.
// It makes a test request to API 1.5 and returns an error if authentication fails.
// host may be blank in which case client attempts to resolve it using auth.
func New(host string, auth rsapi.Authenticator) *Api {
	return fromApi(rsapi.New(host, auth))
}

// NewRL10 returns a API 1.6 client that uses the information stored in /var/run/rightlink/secret to do
// auth and configure the host. The client behaves identically to the new returned by New in
// all other regards.
func NewRL10() (*Api, error) {
	raw, err := rsapi.NewRL10()
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

// Wrap generic client into API 1.6 client
func fromApi(api *rsapi.Api) *Api {
	api.Metadata = GenMetadata
	return &Api{api}
}
