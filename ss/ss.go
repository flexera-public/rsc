//go:generate praxisgen -metadata=ssd/restful_doc -output=ssd -pkg=ssd -target=1.0 -client=Api
//go:generate praxisgen -metadata=ssc/restful_doc -output=ssc -pkg=ssc -target=1.0 -client=Api
//go:generate praxisgen -metadata=ssm/restful_doc -output=ssm -pkg=ssm -target=1.0 -client=Api
package ss

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/rightscale/rsc/cmd"
	"github.com/rightscale/rsc/metadata"
	"github.com/rightscale/rsc/rsapi"
	"github.com/rightscale/rsc/ss/ssc"
	"github.com/rightscale/rsc/ss/ssd"
	"github.com/rightscale/rsc/ss/ssm"
)

// Self-Service 1.0 common client to all self-service APIs
type Api struct {
	*rsapi.Api     // not used to make actual API calls, just to parse command line, TBD: refactor so this is not needed
	designerClient *ssd.Api
	catalogClient  *ssc.Api
	managerClient  *ssm.Api
}

// Api dispatch function type
// Used to call proper "Dispatch" method depending on service (ssm vs. ssc vs. ssd)
type DispatchFunc func(method, actionUrl string, params, payload rsapi.ApiParams) (*http.Response, error)

// New returns a client that uses User oauth authentication.
// logger and client are optional.
// host may be blank in which case client attempts to resolve it using auth.
// If no HTTP client is specified then the default client is used.
func New(accountId int, refreshToken string, host string, logger *log.Logger,
	client rsapi.HttpClient) (*Api, error) {
	a, err := rsapi.New(accountId, refreshToken, host, logger, client)
	if err != nil {
		return nil, err
	}
	d, err := ssd.New(accountId, refreshToken, host, logger, client)
	if err != nil {
		return nil, err
	}
	c, err := ssc.New(accountId, refreshToken, host, logger, client)
	if err != nil {
		return nil, err
	}
	m, err := ssm.New(accountId, refreshToken, host, logger, client)
	if err != nil {
		return nil, err
	}
	api := Api{a, d, c, m}
	setupMetadata(a, d, c, m)

	return &api, nil
}

// Build client from command line
func FromCommandLine(cmdLine *cmd.CommandLine) (*Api, error) {
	// Hackyity hack
	if !strings.HasPrefix(cmdLine.Host, "selfservice") {
		cmdLine.Host = "selfservice-3.rightscale.com"
	}
	if cmdLine.RL10 {
		return nil, fmt.Errorf("RightLink 10 proxy not supported for Self-Service APIs")
	}
	a, err := rsapi.FromCommandLine(cmdLine)
	if err != nil {
		return nil, err
	}
	d, err := ssd.FromCommandLine(cmdLine)
	if err != nil {
		return nil, err
	}
	c, err := ssc.FromCommandLine(cmdLine)
	if err != nil {
		return nil, err
	}
	m, err := ssm.FromCommandLine(cmdLine)
	if err != nil {
		return nil, err
	}
	api := Api{a, d, c, m}
	setupMetadata(a, d, c, m)

	return &api, nil
}

// Merge all metadata so that Api object has access to all commands and actions for command line
// parsing.
func setupMetadata(a *rsapi.Api, d *ssd.Api, c *ssc.Api, m *ssm.Api) {
	md := map[string]*metadata.Resource{}
	for n, r := range d.Metadata {
		md[n] = r
	}
	for n, r := range c.Metadata {
		md[n] = r
	}
	for n, r := range m.Metadata {
		md[n] = r
	}
	a.Metadata = md
}
