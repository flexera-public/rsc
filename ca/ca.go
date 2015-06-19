package ca

import (
	"github.com/rightscale/rsc/ca/cac"
	"github.com/rightscale/rsc/cmd"
	"github.com/rightscale/rsc/metadata"
	"github.com/rightscale/rsc/rsapi"
	"regexp"
	"time"
)

// Metadata synthetized from all CA APIs metadata
var GenMetadata map[string]*metadata.Resource

// CA 1.0 common client to all self-service APIs
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
	api.Auth = &rsapi.CAAuthenticator{Username: cmdLine.Username,
		Host:      cmdLine.Host,
		Password:  cmdLine.Password,
		RefreshAt: time.Now().Add(-time.Duration(5) * time.Minute),
		Client:    api.Client}
	return &Api{api}, nil
}

func apiHostFromLogin(host string) string {
	integration, _ := regexp.MatchString("^cobalt", host)
	staging, _ := regexp.MatchString("^moo", host)

	switch {
	case integration:
		return "ca1-analytics-499.test.rightscale.com"
	case staging:
		return "moo-analytics.test.rightscale.com"
	default:
		return "analytics.rightscale.com"
	}
}

// Initialize GenMetadata from each SS API generated metadata
func setupMetadata() {
	GenMetadata = map[string]*metadata.Resource{}
	for n, r := range cac.GenMetadata {
		GenMetadata[n] = r
	}
}
