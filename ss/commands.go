package ss

import (
	"fmt"
	"net/http"
	"path"

	"github.com/rightscale/rsc/metadata"
	"github.com/rightscale/rsc/rsapi"
	"github.com/rightscale/rsc/ss/ssc"
	"github.com/rightscale/rsc/ss/ssd"
	"github.com/rightscale/rsc/ss/ssm"
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
	registrar.RegisterActionCommands(ApiName, ssd.GenMetadata, commandValues)
	registrar.RegisterActionCommands(ApiName, ssc.GenMetadata, commandValues)
	registrar.RegisterActionCommands(ApiName, ssm.GenMetadata, commandValues)
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

// Show command actions
func (a *Api) ShowApiActions(cmd string) error {
	target, _, err := a.ParseCommandAndFlags(cmd, "", commandValues)
	var service string
	if err == nil {
		service, err = getResourceService(target.Resource)
	}
	if err == nil {
		var showHrefs func(string, string, rsapi.ActionCommands) error
		switch service {
		case "manager":
			showHrefs = a.managerClient.ShowActions
		case "catalog":
			showHrefs = a.catalogClient.ShowActions
		case "designer":
			showHrefs = a.designerClient.ShowActions
		}
		return showHrefs(cmd, "", commandValues)
	} else {
		fmt.Println("DESIGNER")
		a.designerClient.ShowActions(cmd, "", commandValues)
		fmt.Println("\nCATALOG")
		a.catalogClient.ShowActions(cmd, "", commandValues)
		fmt.Println("\nMANAGER")
		a.managerClient.ShowActions(cmd, "", commandValues)
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
