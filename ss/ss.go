package ss

import (
	"path"
	"strings"

	"github.com/rightscale/rsc/cmd"
	"github.com/rightscale/rsc/metadata"
	"github.com/rightscale/rsc/rsapi"
	"github.com/rightscale/rsc/ss/ssc"
	"github.com/rightscale/rsc/ss/ssd"
	"github.com/rightscale/rsc/ss/ssm"
)

// Metadata synthetized from all SS APIs metadata
var GenMetadata map[string]*metadata.Resource

// API is the Self-Service 1.0 common client to all self-service APIs.
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
	api.Metadata = GenMetadata
	if api.Auth != nil {
		api.Auth.SetHost(api.Host)
	}
	return &API{api}, nil
}

// New returns a Self-Service API client.
func New(h string, a rsapi.Authenticator) *API {
	api := rsapi.New(h, a)
	setupMetadata()
	api.Metadata = GenMetadata
	return &API{API: api}
}

// HostFromLogin returns the Self-service endpoint from its login endpoint.
// The following isn't great but seems better than having to enter by hand.
func HostFromLogin(host string) string {
	urlElems := strings.Split(host, ".")
	hostPrefix := urlElems[0]
	elems := strings.Split(hostPrefix, "-")
	if len(elems) < 2 {
		return host
	}
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
