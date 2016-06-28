package ssm // import "gopkg.in/rightscale/rsc.v6/ss/ssm"

import (
	"net/http"

	"gopkg.in/rightscale/rsc.v6/rsapi"
)

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

// BuildHTTPRequest wraps the underlying rsapi implementation and simply prefixes the path with
// the service manager path.
func (a *API) BuildHTTPRequest(verb, path, version string, params, payload rsapi.APIParams) (*http.Request, error) {
	return a.API.BuildHTTPRequest(verb, "/manager"+path, version, params, payload)
}
