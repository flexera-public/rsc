package cm15 // import "gopkg.in/rightscale/rsc.v3/cm15"

import (
	"net/http"

	"gopkg.in/rightscale/rsc.v3/rsapi"
)

const (
	// Used by rsc to display command line help
	ApiName = "RightScale CM API 1.5"
)

// Data structure that holds parsed command line values
var commandValues rsapi.ActionCommands

// Register all commands with kinpin application
func RegisterCommands(registrar rsapi.ApiCommandRegistrar) {
	commandValues = rsapi.ActionCommands{}
	registrar.RegisterActionCommands(ApiName, GenMetadata, commandValues)
}

// Parse and run command
func (a *Api) RunCommand(cmd string) (*http.Response, error) {
	c, err := a.ParseCommand(cmd, "/api", commandValues)
	if err != nil {
		return nil, err
	}
	req, err := a.BuildHTTPRequest(c.HttpMethod, c.Uri, "1.5", c.QueryParams, c.PayloadParams)
	if err != nil {
		return nil, err
	}
	return a.PerformRequest(req)
}

// Show command help
func (a *Api) ShowCommandHelp(cmd string) error {
	return a.ShowHelp(cmd, "/api", commandValues)
}

// Show command hrefs
func (a *Api) ShowApiActions(cmd string) error {
	return a.ShowActions(cmd, "/api", commandValues)
}
