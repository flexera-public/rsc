package rsapi16

import (
	"net/http"

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
		var actionCmd = api16Cmd.Command(action, description)
		var actionCmdValue = rsapi.ActionCommand{}
		var hrefMsg = "API Resource or resource collection href on which to act, e.g. '/api/deployments'"
		var paramsMsg = "Action parameters in the form QUERY=VALUE, e.g. 'view=default'"
		actionCmd.Arg("href", hrefMsg).Required().StringVar(&actionCmdValue.Href)
		actionCmd.Arg("params", paramsMsg).StringsVar(&actionCmdValue.Params)
		commandValues[actionCmd.FullCommand()] = &actionCmdValue
	}
}

// Parse and run command
func (a *Api16) RunCommand(cmd string) (*http.Response, error) {
	var parsed, err = a.ParseCommand(cmd, commandValues)
	if err != nil {
		return nil, err
	}
	return a.GetRaw(parsed.Uri, parsed.QueryParams)
}

// Show command help
func (a *Api16) ShowCommandHelp(cmd string) error {
	return a.ShowHelp(cmd, commandValues)
}
