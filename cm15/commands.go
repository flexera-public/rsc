package cm15 // import "gopkg.in/rightscale/rsc.v9/cm15"

import (
	"net/http"

	"gopkg.in/rightscale/rsc.v9/rsapi"
)

const (
	// APIName is used by rsc to display command line help.
	APIName = "RightScale CM API 1.5"
)

// Data structure that holds parsed command line values
var commandValues rsapi.ActionCommands

// RegisterCommands registers all commands with kinpin application.
func RegisterCommands(registrar rsapi.APICommandRegistrar) {
	commandValues = rsapi.ActionCommands{}
	registrar.RegisterActionCommands(APIName, GenMetadata, commandValues)
}

// RunCommand parses and runs a command given its name.
func (a *API) RunCommand(cmd string) (*http.Response, error) {
	c, err := a.ParseCommand(cmd, "/api", commandValues)
	if err != nil {
		return nil, err
	}
	req, err := a.BuildHTTPRequest(c.HTTPMethod, c.URI, "1.5", c.QueryParams, c.PayloadParams)
	if err != nil {
		return nil, err
	}
	return a.PerformRequest(req)
}

// ShowCommandHelp displays a command help given its name.
func (a *API) ShowCommandHelp(cmd string) error {
	return a.ShowHelp(cmd, "/api", commandValues)
}

// ShowAPIActions displays the command hrefs.
func (a *API) ShowAPIActions(cmd string) error {
	return a.ShowActions(cmd, "/api", commandValues)
}
