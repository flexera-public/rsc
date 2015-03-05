package ss

import (
	"fmt"
	"net/http"
	"path"
	"strings"

	"github.com/rightscale/rsc/cmd"
	"github.com/rightscale/rsc/metadata"
	"github.com/rightscale/rsc/rsapi"
	"github.com/rightscale/rsc/ss/ssc"
	"github.com/rightscale/rsc/ss/ssd"
	"github.com/rightscale/rsc/ss/ssm"
)

// Data structure that holds parsed command line values
var commandValues rsapi.ActionCommands

// Register all commands with kinpin application
func RegisterCommands(ssCmd cmd.CommandProvider) {
	commandValues = rsapi.ActionCommands{}
	var actionNames []string
	allResources := make(map[string]*metadata.Resource, len(ssc.GenMetadata)+len(ssd.GenMetadata)+len(ssm.GenMetadata))
	for n, r := range ssc.GenMetadata {
		allResources[n] = r
	}
	for n, r := range ssd.GenMetadata {
		allResources[n] = r
	}
	for n, r := range ssm.GenMetadata {
		allResources[n] = r
	}
	for _, r := range allResources {
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
			resources := []string{}
			var actionDescription string
			for name, resource := range allResources {
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
		actionCmd := ssCmd.Command(action, description)
		actionCmdValue := rsapi.ActionCommand{}
		hrefMsg := "API Resource or resource collection href on which to act, e.g. '/projects/1/executions/2'"
		paramsMsg := "Action parameters in the form QUERY=VALUE, e.g. 'execution[name]=my\\ execution'"
		actionCmd.Arg("href", hrefMsg).Required().StringVar(&actionCmdValue.Href)
		actionCmd.Arg("params", paramsMsg).StringsVar(&actionCmdValue.Params)
		commandValues[actionCmd.FullCommand()] = &actionCmdValue
	}
}

// Parse and run command
func (a *Api) RunCommand(cmd string) (*http.Response, error) {
	parsed, err := a.ParseCommand(cmd, "", commandValues)
	if err != nil {
		return nil, err
	}
	target, _, _ := a.ParseCommandAndFlags(cmd, "", commandValues)
	service, err := getResourceService(target.Resource)
	if err != nil {
		return nil, err
	}
	var dispatch DispatchFunc
	href := parsed.Uri
	switch service {
	case "manager":
		dispatch = a.managerClient.Dispatch
	case "catalog":
		dispatch = a.catalogClient.Dispatch
	case "designer":
		dispatch = a.designerClient.Dispatch
	}
	href = path.Join("/api/", service, href)
	return dispatch(parsed.HttpMethod, href, parsed.QueryParams, parsed.PayloadParams)
}

// Show command help
func (a *Api) ShowCommandHelp(cmd string) error {
	target, _, err := a.ParseCommandAndFlags(cmd, "", commandValues)
	if err != nil {
		return err
	}
	service, err := getResourceService(target.Resource)
	if err != nil {
		return err
	}
	var showHelp func(string, string, rsapi.ActionCommands) error
	switch service {
	case "manager":
		showHelp = a.managerClient.ShowHelp
	case "catalog":
		showHelp = a.catalogClient.ShowHelp
	case "designer":
		showHelp = a.designerClient.ShowHelp
	}
	return showHelp(cmd, "", commandValues)
}

// Show command hrefs
func (a *Api) ShowCommandHrefs(cmd string) error {
	target, _, err := a.ParseCommandAndFlags(cmd, "", commandValues)
	var service string
	if err == nil {
		service, err = getResourceService(target.Resource)
	}
	if err == nil {
		var showHrefs func(string, string, rsapi.ActionCommands) error
		switch service {
		case "manager":
			showHrefs = a.managerClient.ShowHrefs
		case "catalog":
			showHrefs = a.catalogClient.ShowHrefs
		case "designer":
			showHrefs = a.designerClient.ShowHrefs
		}
		return showHrefs(cmd, "", commandValues)
	} else {
		fmt.Println("DESIGNER")
		a.designerClient.ShowHrefs(cmd, "", commandValues)
		fmt.Println("\nCATALOG")
		a.catalogClient.ShowHrefs(cmd, "", commandValues)
		fmt.Println("\nMANAGER")
		a.managerClient.ShowHrefs(cmd, "", commandValues)
		return nil
	}
}

// Retrieve name of service containing given resource or error if none contain it
// name is one of "designer", "catalog" or "manager"
func getResourceService(resource *metadata.Resource) (string, error) {
	resourceName := resource.Name
	var service string
	switch {
	case containsResource(resourceName, ssc.GenMetadata):
		service = "catalog"
	case containsResource(resourceName, ssd.GenMetadata):
		service = "designer"
	case containsResource(resourceName, ssm.GenMetadata):
		service = "manager"
	default:
		return "", fmt.Errorf("Unknown resource %s", resourceName)
	}
	return service, nil
}

// Check whether resource map contains given resource name
func containsResource(name string, resources map[string]*metadata.Resource) bool {
	for n, _ := range resources {
		if n == name {
			return true
		}
	}
	return false

}
