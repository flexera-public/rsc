package cm15 // import "gopkg.in/rightscale/rsc.v2/cm15"

import (
	"log"

	"gopkg.in/rightscale/rsc.v2/cmd"
	"gopkg.in/rightscale/rsc.v2/rsapi"
)

// Api 1.5 client
// Just a vanilla RightScale API client.
type Api struct {
	*rsapi.Api
}

// New returns a API 1.5 client.
// It makes a test request to API 1.5 and returns an error if authentication fails.
// host may be blank in which case client attempts to resolve it using auth.
// logger is optional.
// If client is nil then the default HTTP client is used.
func New(host string, auth rsapi.Authenticator, logger *log.Logger,
	client rsapi.HttpClient) (*Api, error) {
	var err error
	if auth != nil {
		auth.SetHost(host)
		err = auth.CanAuthenticate()
	}
	return fromApi(rsapi.New(host, auth, logger, client), err)
}

// NewRL10 returns a API 1.5 client that uses the information stored in /var/run/rightlink/secret to do
// auth and configure the host. The client behaves identically to the one returned by New in
// all other aspects.
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
