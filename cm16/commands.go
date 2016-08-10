package cm16

import (
	"net/http"
	"path"
	"strings"

	"github.com/rightscale/rsc/rsapi"
)

const (
	// APIName is used by rsc to display the command line help.
	APIName = "RightScale CM API 1.6"
)

// Data structure that holds parsed command line values
var commandValues rsapi.ActionCommands

// RegisterCommands registers all the commands with the kinpin application.
func RegisterCommands(registrar *rsapi.Registrar) {
	commandValues = rsapi.ActionCommands{}
	registrar.RegisterActionCommands(APIName, GenMetadata, commandValues)
}

// RunCommand parses and runs the command with the given name.
func (a *API) RunCommand(cmd string) (*http.Response, error) {
	parsed, err := a.ParseCommand(cmd, "/api", commandValues)
	if err != nil {
		return nil, err
	}
	href := parsed.URI
	if !strings.HasPrefix(href, "/api") {
		href = path.Join("/api", href)
	}
	req, err := a.BuildHTTPRequest("GET", href, "1.6", parsed.QueryParams, nil)
	if err != nil {
		return nil, err
	}
	return a.PerformRequest(req)
}

// ShowCommandHelp displays help for the given command.
func (a *API) ShowCommandHelp(cmd string) error {
	return a.ShowHelp(cmd, "/api", commandValues)
}

// ShowAPIActions displays the command hrefs.
func (a *API) ShowAPIActions(cmd string) error {
	return a.ShowActions(cmd, "/api", commandValues)
}
