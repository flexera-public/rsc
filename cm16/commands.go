package cm16

import (
	"net/http"
	"strings"

	"github.com/rightscale/rsc/cmd"
	"github.com/rightscale/rsc/rsapi"
)

// Data structure that holds parsed command line values
var commandValues rsapi.ActionCommands

// Register all commands with kinpin application
func RegisterCommands(api16Cmd cmd.CommandProvider) {
	commandValues = rsapi.ActionCommands{}
	for _, action := range []string{"index", "show"} {
		var description string
		switch action {
		case "index":
			description = "Lists all resources of given type in account."
		case "show":
			description = "Show information about a single resource."
		}
		actionCmd := api16Cmd.Command(action, description)
		actionCmdValue := rsapi.ActionCommand{}
		hrefMsg := "API Resource or resource collection href on which to act, e.g. '/api/deployments'"
		paramsMsg := "Action parameters in the form QUERY=VALUE, e.g. 'view=default'"
		actionCmd.Arg("href", hrefMsg).Required().StringVar(&actionCmdValue.Href)
		actionCmd.Arg("params", paramsMsg).StringsVar(&actionCmdValue.Params)
		commandValues[actionCmd.FullCommand()] = &actionCmdValue
	}
}

// Parse and run command
func (a *Api) RunCommand(cmd string) (*http.Response, error) {
	parsed, err := a.ParseCommand(cmd, "/api", commandValues)
	if err != nil {
		return nil, err
	}
	href := parsed.Uri
	if !strings.HasPrefix(href, "/api") {
		if strings.HasPrefix(href, "/") {
			href = "/api" + href
		} else {
			href = "/api/" + href
		}
	}
	return a.GetRaw(href, parsed.QueryParams)
}

// Show command help
func (a *Api) ShowCommandHelp(cmd string) error {
	return a.ShowHelp(cmd, "/api", commandValues)
}

// Show command hrefs
func (a *Api) ShowCommandHrefs(cmd string) error {
	return a.ShowHrefs(cmd, "/api", commandValues)
}
