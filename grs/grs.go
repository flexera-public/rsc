package grs

import (
	"regexp"
	"strings"

	"github.com/rightscale/rsc/cmd"
	"github.com/rightscale/rsc/metadata"
	"github.com/rightscale/rsc/rsapi"
)

// Metadata synthesized from all GRS APIs metadata
var genMetadata map[string]*metadata.Resource

// API is the GRS 2.0 common client to all cloud analytics APIs.
type API struct {
	*rsapi.API
}

// FromCommandLine builds a client from the command line.
func FromCommandLine(cmdLine *cmd.CommandLine) (*API, error) {
	api, err := rsapi.FromCommandLine(cmdLine)
	if err != nil {
		return nil, err
	}
	setupMetadata()
	api.Host = apiHostFromLogin(cmdLine.Host)
	api.Metadata = GenMetadata
	return &API{api}, nil
}

// New returns a GRS API client.
func New(h string, a rsapi.Authenticator) *API {
	api := rsapi.New(h, a)
	setupMetadata()
	api.Metadata = GenMetadata
	return &API{API: api}
}

func apiHostFromLogin(host string) string {
	staging, _ := regexp.MatchString("^moo", host)

	prefix := ""

	switch {
	case staging:
		prefix = "moo-grs"
	default:
		prefix = "grs"
	}

	urlElems := strings.Split(host, ".")
	urlElems[0] = prefix
	return strings.Join(urlElems, ".")
}

// Initialize GenMetadata from each GRS API generated metadata
func setupMetadata() {
	genMetadata = map[string]*metadata.Resource{}
	for n, r := range GenMetadata {
		genMetadata[n] = r
	}
}
