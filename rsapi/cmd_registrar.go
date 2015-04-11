package rsapi  // import "gopkg.in/rightscale/rsc.v1-unstable/rsapi"

import (
	"fmt"
	"sort"
	"strings"

	"gopkg.in/alecthomas/kingpin.v1"
	"gopkg.in/rightscale/rsc.v1-unstable/metadata"
)

// Interface implemented by registrar used by each API client to register its subcommands
type ApiCommandRegistrar interface {
	// Register subcommands for all resource actions
	// Store subcommand parse results into given command line value recipient
	RegisterActionCommands(apiName string, metadata map[string]*metadata.Resource,
		cmdValues ActionCommands)
}

// Implements ApiCommandRegistrar
// Create one of these per API client and initialize it with specific API subcommand
type Registrar struct {
	// Kingpin command under which API subcommands should be registered
	ApiCmd *kingpin.CmdClause
}

// ApiCommandRegistrar implementation
func (r *Registrar) RegisterActionCommands(apiName string, res map[string]*metadata.Resource,
	cmds ActionCommands) {

	var actionNames []string
	for _, r := range res {
		for _, a := range r.Actions {
			name := a.Name
			exists := false
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
	sort.Strings(actionNames)
	for _, action := range actionNames {
		if _, ok := cmds[fmt.Sprintf("%s %s", r.ApiCmd.FullCommand(), action)]; ok {
			continue // Already registered - can happen with SS where multiple APIs are registered under the same command
		}
		resources := []string{}
		var description string
		for name, resource := range res {
			for _, a := range resource.Actions {
				if a.Name == action {
					if description == "" {
						description = a.Description
					}
					resources = append(resources, name)
				}
			}
		}
		if len(resources) > 1 || description == "" {
			switch action {
			case "show":
				description = "Show information about a single resource."
			case "index":
				description = "Lists all resources of given type in account."
			case "create":
				description = "Create new resource."
			case "update":
				description = "Update existing resource."
			case "destroy":
				description = "Destroy a single resource."
			default:
				if len(resources) > 1 {
					description = "Action of resources " + strings.Join(resources[:len(resources)-1], ", ") + " and " + resources[len(resources)-1]
				} else {
					description = "<no description provided>"
				}
			}

		}
		actionCmd := r.ApiCmd.Command(action, description)
		actionCmdValue := ActionCommand{}
		hrefMsg := "API Resource or resource collection href on which to act, e.g. '/api/servers'"
		paramsMsg := "Action parameters in the form QUERY=VALUE, e.g. 'server[name]=server42'"
		actionCmd.Arg("href", hrefMsg).Required().StringVar(&actionCmdValue.Href)
		actionCmd.Arg("params", paramsMsg).StringsVar(&actionCmdValue.Params)
		cmds[actionCmd.FullCommand()] = &actionCmdValue
	}
	// Add special "actions" action
	_, ok := cmds[fmt.Sprintf("%s %s", r.ApiCmd.FullCommand(), "actions")]
	if !ok {
		pattern := "List all %s actions and associated hrefs. If a resource href is provided only list actions for that resource."
		description := fmt.Sprintf(pattern, apiName)
		actionsCmd := r.ApiCmd.Command("actions", description)
		actionsCmdValue := ActionCommand{}
		hrefMsg := "Href of resource to show actions for."
		actionsCmd.Arg("href", hrefMsg).StringVar(&actionsCmdValue.Href)
		cmds[actionsCmd.FullCommand()] = &actionsCmdValue
	}
}
