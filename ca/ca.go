package ca // import "gopkg.in/rightscale/rsc.v3/ca"

import (
	"gopkg.in/rightscale/rsc.v3/ca/cac"
	"gopkg.in/rightscale/rsc.v3/cmd"
	"gopkg.in/rightscale/rsc.v3/metadata"
	"gopkg.in/rightscale/rsc.v3/rsapi"
	"regexp"
	"strings"
)

// Metadata synthetized from all CA APIs metadata
var GenMetadata map[string]*metadata.Resource

// CA 1.0 common client to all cloud analytics APIs
type Api struct {
	*rsapi.Api
}

// Build client from command line
func FromCommandLine(cmdLine *cmd.CommandLine) (*Api, error) {
	api, err := rsapi.FromCommandLine(cmdLine)
	if err != nil {
		return nil, err
	}
	setupMetadata()
	api.Host = apiHostFromLogin(cmdLine.Host)
	api.Metadata = GenMetadata
	return &Api{api}, nil
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
func setupMetadata() {
	GenMetadata = map[string]*metadata.Resource{}
	for n, r := range cac.GenMetadata {
		GenMetadata[n] = r
	}
}
