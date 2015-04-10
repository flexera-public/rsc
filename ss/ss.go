package ss  // import "gopkg.in/rightscale/rsc.v1-unstable/ss"

import (
	"path"
	"strings"
	"time"

	"gopkg.in/rightscale/rsc.v1-unstable/cmd"      // import "gopkg.in/rightscale/rsc.v1-unstable/cmd"
	"gopkg.in/rightscale/rsc.v1-unstable/metadata" // import "gopkg.in/rightscale/rsc.v1-unstable/metadata"
	"gopkg.in/rightscale/rsc.v1-unstable/rsapi"    // import "gopkg.in/rightscale/rsc.v1-unstable/rsapi"
	"gopkg.in/rightscale/rsc.v1-unstable/ss/ssc"   // import "gopkg.in/rightscale/rsc.v1-unstable/ss/ssc"
	"gopkg.in/rightscale/rsc.v1-unstable/ss/ssd"   // import "gopkg.in/rightscale/rsc.v1-unstable/ss/ssd"
	"gopkg.in/rightscale/rsc.v1-unstable/ss/ssm"   // import "gopkg.in/rightscale/rsc.v1-unstable/ss/ssm"
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

// Return Self-service endpoint from login endpoint
// The following isn't great but seems better than having to enter by hand
func HostFromLogin(host string) string {
	urlElems := strings.Split(host, ".")
	hostPrefix := urlElems[0]
	elems := strings.Split(hostPrefix, "-")
	elems[len(elems)-2] = "selfservice"
	ssLoginHostPrefix := strings.Join(elems, "-")
	return strings.Join(append([]string{ssLoginHostPrefix}, urlElems[1:]...), ".")
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
