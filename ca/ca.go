package ca

import (
	"github.com/rightscale/rsc/ca/cac"
	"github.com/rightscale/rsc/cmd"
	"github.com/rightscale/rsc/metadata"
	"github.com/rightscale/rsc/rsapi"
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
	api.Metadata = GenMetadata
	return &Api{api}, nil
}

// Initialize GenMetadata from each SS API generated metadata
func setupMetadata() {
	GenMetadata = map[string]*metadata.Resource{}
	for n, r := range cac.GenMetadata {
		GenMetadata[n] = r
	}
}
