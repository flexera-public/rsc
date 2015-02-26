package rsapi16

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
	var actionNames []string
	for _, r := range api_metadata {
		for _, a := range r.Actions {
			var name = a.Name
			var exists = false
			for _, e := range actionNames {
				if e == name {
					exists = true
					break
				}
			}
			if !exists {
				actionNames = append(actionNames, name)
			}
		}
	}
	for _, action := range actionNames {
		var description string
		switch action {
		case "show":
			description = "Show information about a single resource."
		case "index":
			description = "Lists all resources of given type in account."
		case "create":
			description = "Create new resource."
		case "update":
			description = "Update existing resource."
		case "delete":
			description = "Destroy a single resource."
		default:
			var resources = []string{}
			var actionDescription string
			for name, resource := range api_metadata {
				for _, a := range resource.Actions {
					if a.Name == action {
						actionDescription = a.Description
						resources = append(resources, name)
					}
				}
			}
			if len(resources) == 1 {
				description = actionDescription
			} else {
				description = "Action of resources " + strings.Join(resources[:len(resources)-1], ", ") + " and " + resources[len(resources)-1]
			}
		}
		var actionCmd = api16Cmd.Command(action, description)
		var actionCmdValue = rsapi.ActionCommand{}
		var hrefMsg = "API Resource or resource collection href on which to act, e.g. '/api/servers'"
		var paramsMsg = "Action parameters in the form QUERY=VALUE, e.g. 'server[name]=server42'"
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
	return a.Dispatch(parsed.HttpMethod, parsed.Uri, parsed.Params)
}
