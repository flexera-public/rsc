package ss

import (
	"net/http"

	"github.com/rightscale/rsc/rsapi"
)

const (
	// APIName is used by rsc to display command line help.
	APIName = "RightScale Self-Service APIs"
)

// Data structure that holds parsed command line values
var commandValues rsapi.ActionCommands

// RegisterCommands registers all commands with the kinpin application.
func RegisterCommands(registrar rsapi.APICommandRegistrar) {
	commandValues = rsapi.ActionCommands{}
	registrar.RegisterActionCommands(APIName, GenMetadata, commandValues)
}

// RunCommand parses and runs a command given its name.
func (a *API) RunCommand(cmd string) (*http.Response, error) {
	c, err := a.ParseCommand(cmd, "", commandValues)
	if err != nil {
		return nil, err
	}
	req, err := a.BuildHTTPRequest(c.HTTPMethod, c.URI, "1.0", c.QueryParams, c.PayloadParams)
	if err != nil {
		return nil, err
	}
	return a.PerformRequest(req)
}

// ShowCommandHelp displays a command help.
func (a *API) ShowCommandHelp(cmd string) error {
	return a.ShowHelp(cmd, "", commandValues)
}

// ShowAPIActions displays a command actions.
func (a *API) ShowAPIActions(cmd string) error {
	return a.ShowActions(cmd, "", commandValues)
}
