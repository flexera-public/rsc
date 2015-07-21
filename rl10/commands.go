package rl10

import (
	"net/http"

	"github.com/rightscale/rsc/rsapi"
)

const (
	// Used by rsc to display command line help
	ApiName = "RightScale RightLink10 API"
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
	c, err := a.ParseCommand(cmd, "/rll", commandValues)
	if err != nil {
		return nil, err
	}
	req, err := a.BuildHTTPRequest(c.HttpMethod, c.Uri, "", c.QueryParams, c.PayloadParams)
	if err != nil {
		return nil, err
	}
	return a.PerformRequest(req)
}

// Show command help
func (a *Api) ShowCommandHelp(cmd string) error {
	return a.ShowHelp(cmd, "/rll", commandValues)
}

// Show command hrefs
func (a *Api) ShowApiActions(cmd string) error {
	return a.ShowActions(cmd, "/rll", commandValues)
}
