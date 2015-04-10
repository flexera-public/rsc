package ss

import (
	"net/http"

	"gopkg.in/rightscale/rsc.v1-unstable/rsapi" // import "gopkg.in/rightscale/rsc.v1-unstable/rsapi"
)

const (
	// Used by rsc to display command line help
	ApiName = "RightScale Self-Service APIs"
)

// Data structure that holds parsed command line values
var commandValues rsapi.ActionCommands

// Register all commands with kinpin application
func RegisterCommands(registrar rsapi.ApiCommandRegistrar) {
	commandValues = rsapi.ActionCommands{}
	setupMetadata()
	registrar.RegisterActionCommands(ApiName, GenMetadata, commandValues)
}

// Parse and run command
func (a *Api) RunCommand(cmd string) (*http.Response, error) {
	parsed, err := a.ParseCommand(cmd, "", commandValues)
	if err != nil {
		return nil, err
	}
	return a.Dispatch(parsed.HttpMethod, parsed.Uri, parsed.QueryParams, parsed.PayloadParams)
}

// Show command help
func (a *Api) ShowCommandHelp(cmd string) error {
	return a.ShowHelp(cmd, "", commandValues)
}

// Show command actions
func (a *Api) ShowApiActions(cmd string) error {
	return a.ShowActions(cmd, "", commandValues)
}
