package cac

import "github.com/rightscale/rsc/rsapi"

// API is the Cloud Analytics API client.
type API struct {
	*rsapi.API
}

// New returns a client that uses User oauth authentication.
// logger and client are optional.
// host may be blank in which case client attempts to resolve it using auth.
// If no HTTP client is specified then the default client is used.
func New(h string, a rsapi.Authenticator) *API {
	api := rsapi.New(h, a)
	api.Metadata = GenMetadata
	return &API{API: api}
}
