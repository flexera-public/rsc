package ssc // import "gopkg.in/rightscale/rsc.v3/ss/ssc"

import (
	"net/http"

	"gopkg.in/rightscale/rsc.v3/rsapi"
)

// Self-Service catalog client
type Api struct {
	*rsapi.Api
}

// New returns a Self-Service catalog API client.
// It makes a test API request and returns an error if authentication fails.
func New(h string, a rsapi.Authenticator) *Api {
	api := rsapi.New(h, a)
	api.Metadata = GenMetadata
	ssApi := Api{Api: api}
	return &ssApi
}

// BuildHTTPRequest wraps the underlying rsapi implementation and simply prefixes the path with
// the service catalog path.
func (a *Api) BuildHTTPRequest(verb, path, version string, params, payload rsapi.ApiParams) (*http.Request, error) {
	return a.Api.BuildHTTPRequest(verb, "/catalog"+path, version, params, payload)
}
