package ssm

import "github.com/rightscale/rsc/rsapi"

// API is the Self-Service manager client.
type API struct {
	*rsapi.API
}

// New returns a Self-Service catalog API client.
// It makes a test API request and returns an error if authentication fails.
func New(h string, a rsapi.Authenticator) *API {
	api := rsapi.New(h, a)
	api.Metadata = GenMetadata
	ssAPI := API{API: api}
	return &ssAPI
}
