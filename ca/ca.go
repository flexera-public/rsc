package ca

import (
	"regexp"
	"strings"

	"github.com/rightscale/rsc/ca/cac"
	"github.com/rightscale/rsc/cmd"
	"github.com/rightscale/rsc/metadata"
	"github.com/rightscale/rsc/rsapi"
)

// Metadata synthetized from all CA APIs metadata; setup once
var GenMetadata = setupMetadata()

// API is the CA 1.0 common client to all cloud analytics APIs.
type API struct {
	*rsapi.API
}

// FromCommandLine builds a client from the command line.
func FromCommandLine(cmdLine *cmd.CommandLine) (*API, error) {
	api, err := rsapi.FromCommandLine(cmdLine)
	if err != nil {
		return nil, err
	}
	api.Host = apiHostFromLogin(cmdLine.Host)
	api.Metadata = GenMetadata
	return &API{api}, nil
}

// New returns a CA API client.
func New(h string, a rsapi.Authenticator) *API {
	api := rsapi.New(h, a)
	api.Metadata = GenMetadata
	return &API{API: api}
}

func apiHostFromLogin(host string) string {
	integration, _ := regexp.MatchString("^cobalt", host)
	staging, _ := regexp.MatchString("^moo", host)

	prefix := ""

	switch {
	case integration:
		prefix = "ca1-analytics-499"
	case staging:
		prefix = "moo-analytics"
	default:
		prefix = "analytics"
	}

	urlElems := strings.Split(host, ".")
	urlElems[0] = prefix
	return strings.Join(urlElems, ".")
}

// Initialize GenMetadata from each CA API generated metadata
func setupMetadata() (result map[string]*metadata.Resource) {
	result = make(map[string]*metadata.Resource)
	for n, r := range cac.GenMetadata {
		result[n] = r
	}
	return
}
