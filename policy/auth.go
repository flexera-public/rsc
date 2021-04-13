package policy // import "gopkg.in/rightscale/rsc.v9/policy"

import (
	"strings"

	"gopkg.in/rightscale/rsc.v9/cmd"
	"gopkg.in/rightscale/rsc.v9/rsapi"
)

// API 1.0 client
// Just a vanilla RightScale API client.
type API struct {
	*rsapi.API
}

// New returns a API 1.0 client.
// It makes a test request to API 1.0 and returns an error if authentication fails.
// host may be blank in which case client attempts to resolve it using auth.
func New(host string, auth rsapi.Authenticator) *API {
	return fromAPI(rsapi.New(host, auth))
}

// FromCommandLine builds a client from the command line.
func FromCommandLine(cmdLine *cmd.CommandLine) (*API, error) {
	raw, err := rsapi.FromCommandLine(cmdLine)
	cmdLine.Host = HostFromLogin(cmdLine.Host)

	if err != nil {
		return nil, err
	}
	return fromAPI(raw), nil
}

// Wrap generic client into API 1.0 client
func fromAPI(api *rsapi.API) *API {
	api.FileEncoding = rsapi.FileEncodingJSON
	api.Host = HostFromLogin(api.Host)
	api.Metadata = GenMetadata
	api.VersionHeader = "Api-Version"
	return &API{api}
}

// HostFromLogin returns the policy endpoint from its login endpoint.
// The following isn't great but seems better than having to enter by hand.
func HostFromLogin(host string) string {
	urlElems := strings.Split(host, ".")
	hostPrefix := urlElems[0]
	elems := strings.Split(hostPrefix, "-")

	if len(elems) == 1 && elems[0] == "cm" {
		// accommodates micromoo host inference, such as "cm.rightscale.local" => "selfservice.rightscale.local"
		elems[0] = "governance"
	} else if len(elems) < 2 {
		// don't know how to compute this policy host; use the cm host
		return host
	} else {
		elems[len(elems)-2] = "governance"
	}
	policyHostPrefix := strings.Join(elems, "-")
	return strings.Join(append([]string{policyHostPrefix}, urlElems[1:]...), ".")
}
