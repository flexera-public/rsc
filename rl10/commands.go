package rl10

import (
	"net/http"

	"github.com/rightscale/rsc/rsapi"
)

const (
	// APIName is used by rsc to display the command line help.
	APIName = "RightScale RightLink10 API"
)

// Data structure that holds parsed command line values
var commandValues rsapi.ActionCommands

// RegisterCommands registers all the commands with the kinpin application.
func RegisterCommands(registrar rsapi.APICommandRegistrar) {
	commandValues = rsapi.ActionCommands{}
	registrar.RegisterActionCommands(APIName, GenMetadata, commandValues)
}

// RunCommand parses and runs the command with the given name.
func (a *API) RunCommand(cmd string) (*http.Response, error) {
	c, err := a.ParseCommand(cmd, "/rll", commandValues)
	if err != nil {
		return nil, err
	}
	req, err := a.BuildHTTPRequest(c.HTTPMethod, c.URI, "", c.QueryParams, c.PayloadParams)
	if err != nil {
		return nil, err
	}
	return a.PerformRequest(req)
}

// ShowCommandHelp displays the command help.
func (a *API) ShowCommandHelp(cmd string) error {
	return a.ShowHelp(cmd, "/rll", commandValues)
}

// ShowAPIActions displays the command hrefs.
func (a *API) ShowAPIActions(cmd string) error {
	return a.ShowActions(cmd, "/rll", commandValues)
}
