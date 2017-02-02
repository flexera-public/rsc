package ss

import (
	"regexp"
	"strings"

	"github.com/rightscale/rsc/cmd"
	"github.com/rightscale/rsc/metadata"
	"github.com/rightscale/rsc/rsapi"
	"github.com/rightscale/rsc/ss/ssc"
	"github.com/rightscale/rsc/ss/ssd"
	"github.com/rightscale/rsc/ss/ssm"
)

// Metadata synthetized from all SS APIs metadata; setup once
var GenMetadata = setupMetadata()

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
	api.Metadata = GenMetadata
	if api.Auth != nil {
		api.Auth.SetHost(api.Host)
	}
	return &API{api}, nil
}

// New returns a Self-Service API client.
func New(h string, a rsapi.Authenticator) *API {
	api := rsapi.New(h, a)
	api.Metadata = GenMetadata
	return &API{API: api}
}

// HostFromLogin returns the Self-service endpoint from its login endpoint.
// The following isn't great but seems better than having to enter by hand.
func HostFromLogin(host string) string {
	urlElems := strings.Split(host, ".")
	hostPrefix := urlElems[0]
	elems := strings.Split(hostPrefix, "-")

	if len(elems) == 1 && elems[0] == "cm" {
		// accommodates micromoo host inference, such as "cm.rightscale.local" => "selfservice.rightscale.local"
		elems[0] = "selfservice"
	} else if len(elems) < 2 {
		// don't know how to compute this ss host; use the cm host
		return host
	} else {
		elems[len(elems)-2] = "selfservice"
	}
	ssLoginHostPrefix := strings.Join(elems, "-")
	return strings.Join(append([]string{ssLoginHostPrefix}, urlElems[1:]...), ".")
}

func copyPathPattern(p *metadata.PathPattern) (newP *metadata.PathPattern) {
	newP = &metadata.PathPattern{HTTPMethod: p.HTTPMethod, Pattern: p.Pattern}
	copy(newP.Variables, p.Variables)
	newP.Regexp = &regexp.Regexp{}
	*newP.Regexp = *p.Regexp
	return
}

// Removes the specified number of prefixes from a regexp and returns a new regexp.
// This basically loosens validations on the regexp by making the specified number of
// prefixes optional. num must be greater than the number of prefixes. For example:
// r.String()                       // => "/api/catalog/collections/([^/]+)/templates/actions/dependencies"
// removePrefixes(&r, 2).String()   // => "/collections/([^/]+)/templates/actions/dependencies"
func removePrefixes(r *regexp.Regexp, num int) (result *regexp.Regexp) {
	path := strings.TrimLeft(r.String(), "/")
	paths := strings.Split(path, "/")
	result = regexp.MustCompile("/" + strings.Join(paths[num:], "/"))
	return
}

// Initialize GenMetadata from each SS API generated metadata
func setupMetadata() (result map[string]*metadata.Resource) {
	result = make(map[string]*metadata.Resource)
	for n, r := range ssd.GenMetadata {
		result[n] = r
		for _, a := range r.Actions {
			for _, p := range a.PathPatterns {
				// remove "/api/designer" prefix
				p.Regexp = removePrefixes(p.Regexp, 2)
			}
		}
	}
	for n, r := range ssc.GenMetadata {
		result[n] = r
		for _, a := range r.Actions {
			for _, p := range a.PathPatterns {
				// remove "/api/catalog" prefix
				p.Regexp = removePrefixes(p.Regexp, 2)
			}
		}
	}
	for n, r := range ssm.GenMetadata {
		result[n] = r
		for _, a := range r.Actions {
			for _, p := range a.PathPatterns {
				// remove "/api/manager" prefix
				p.Regexp = removePrefixes(p.Regexp, 2)
			}
		}
	}
	return
}
