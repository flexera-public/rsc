package ss

import (
	"path"
	"time"

	"github.com/rightscale/rsc/cmd"
	"github.com/rightscale/rsc/metadata"
	"github.com/rightscale/rsc/rsapi"
	"github.com/rightscale/rsc/ss/ssc"
	"github.com/rightscale/rsc/ss/ssd"
	"github.com/rightscale/rsc/ss/ssm"
)

// Metadata synthetized from all SS APIs metadata
var GenMetadata map[string]*metadata.Resource

// Self-Service 1.0 common client to all self-service APIs
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
	fiveMnAgo := time.Now().Add(-time.Duration(5) * time.Minute)
	api.Auth = &rsapi.SSAuthenticator{api.Auth, api.AccountId, fiveMnAgo, api.Client}
	return &Api{api}, nil
}

// Whether we've already adjusted the action path patterns in the SS APIs generated metadata
var pathFixupDone bool

// Initialize GenMetadata from each SS API generated metadata
func setupMetadata() {
	GenMetadata = map[string]*metadata.Resource{}
	for n, r := range ssd.GenMetadata {
		GenMetadata[n] = r
		if pathFixupDone {
			continue
		}
		for _, a := range r.Actions {
			for _, p := range a.PathPatterns {
				p.Pattern = path.Join("designer", p.Pattern)
			}
		}
	}
	for n, r := range ssc.GenMetadata {
		GenMetadata[n] = r
		if pathFixupDone {
			continue
		}
		for _, a := range r.Actions {
			for _, p := range a.PathPatterns {
				p.Pattern = path.Join("catalog", p.Pattern)
			}
		}
	}
	for n, r := range ssm.GenMetadata {
		GenMetadata[n] = r
		if pathFixupDone {
			continue
		}
		for _, a := range r.Actions {
			for _, p := range a.PathPatterns {
				p.Pattern = path.Join("manager", p.Pattern)
			}
		}
	}
	pathFixupDone = true
}
