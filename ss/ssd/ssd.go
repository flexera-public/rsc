package ssd // import "gopkg.in/rightscale/rsc.v4/ss/ssd"

import (
	"net/http"

	"gopkg.in/rightscale/rsc.v4/rsapi"
)

// API is the Self-Service designer client.
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

// BuildHTTPRequest wraps the underlying rsapi implementation and simply prefixes the path with
// the service designer path.
func (a *API) BuildHTTPRequest(verb, path, version string, params, payload rsapi.APIParams) (*http.Request, error) {
	return a.API.BuildHTTPRequest(verb, "/designer"+path, version, params, payload)
}
