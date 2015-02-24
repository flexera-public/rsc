package rsapi15

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/rightscale/rsc/cmd"
	"github.com/rightscale/rsc/metadata"
)

// Action params consist of href and map of query parameters
type ActionFlags struct {
	Href     string   // Resource or collection href
	Params   []string // Action parameters
	ShowHelp string   // Whether to list flags supported by resource action
}

// Flag values indexed by action name indexed by command name
type FlagValues map[string]*ActionFlags

// Map that holds flag values resulting from command line parsing
var flagValues = make(FlagValues)

// Register all commands with kinpin application
func RegisterCommands(api15Cmd cmd.CommandProvider) {
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
		var actionCmd = api15Cmd.Command(action, description)
		var actionFlags = ActionFlags{}
		actionCmd.Arg("href", "API Resource or resource collection href on which to act, e.g. '/api/servers'").Required().StringVar(&actionFlags.Href)
		actionCmd.Arg("params", "Action parameters in the form QUERY=VALUE, e.g. 'server[name]=server42'").StringsVar(&actionFlags.Params)
		flagValues[actionCmd.FullCommand()] = &actionFlags
	}
}

// Show help for given command and flags
func (a *Api15) ShowHelp(cmd string) error {
	var resource, action, _, href, _, err = parseCommandAndFlags(cmd)
	if err != nil {
		return err
	}
	if len(action.Params) == 0 {
		fmt.Printf("usage: rsc [<flags>] api15 %s %s\n", action.Name, href)
		return nil
	}
	var flagHelp = make([]string, len(action.Params))
	for i, f := range action.Params {
		var attrs string
		if f.Mandatory {
			attrs = "required"
		} else {
			attrs = "optional"
		}
		if len(f.ValidValues) > 0 {
			attrs += ", [" + strings.Join(f.ValidValues, "|") + "]"
		}
		if f.Regexp != nil {
			attrs += ", /" + f.Regexp.String() + "/"
		}
		flagHelp[i] = fmt.Sprintf("%s=%s\n    <%s> %s", f.Name, f.Type, attrs, f.Description)
	}
	fmt.Printf("usage: rsc [<flags>] api15 %s %s [<%s %s params>]\n\n", action.Name,
		href, strings.ToLower(resource.Name), action.Name)
	fmt.Printf("%s %s params:\n%s\n", resource.Name, action.Name, strings.Join(flagHelp, "\n\n"))
	return nil
}

// Actually run command
func (a *Api15) RunCommand(cmd string) (*http.Response, error) {
	// 1. Initialize / find href as well as resource and action command definitions
	var resource, action, path, _, params, err = parseCommandAndFlags(cmd)
	if err != nil {
		return nil, err
	}

	// 2. Coerce and validate given flag values
	var coerced = map[string]interface{}{}
	for _, p := range params {
		var elems = strings.SplitN(p, "=", 2)
		if len(elems) != 2 {
			return nil, fmt.Errorf("Arguments must be of the form NAME=VALUE, value provided was '%s'", p)
		}
		var name = elems[0]
		var value = elems[1]
		var flag *metadata.ActionParam
		for _, ap := range action.Params {
			if ap.Name == name {
				flag = ap
				break
			}
		}
		if flag == nil {
			var supported = make([]string, len(action.Params))
			for i, p := range action.Params {
				supported[i] = p.Name
			}
			return nil, fmt.Errorf("Unknown %s.%s flag '%s'. Supported flags are: %s",
				resource.Name, action.Name, name, strings.Join(supported, ", "))
		}
		if err := validateFlagValue(value, flag); err != nil {
			return nil, err
		}
		switch flag.Type {
		case "string":
			if _, ok := coerced[name]; ok {
				return nil, fmt.Errorf("Multiple values specified for '%s'", flag.Name)
			}
			coerced[flag.Name] = value
		case "[]string":
			if v, ok := coerced[flag.Name]; ok {
				v = append(v.([]string), value)
			} else {
				coerced[flag.Name] = []string{value}
			}
		case "int":
			if _, ok := coerced[name]; ok {
				return nil, fmt.Errorf("Multiple values specified for '%s'", flag.Name)
			}
			val, err := strconv.Atoi(value)
			if err != nil {
				return nil, fmt.Errorf("Value for '%s' must be an integer, value provided was '%s'",
					flag.Name, value)
			}
			coerced[flag.Name] = val
		}
	}
	for _, p := range action.Params {
		_, ok := coerced[p.Name]
		if p.Mandatory && !ok {
			return nil, fmt.Errorf("Missing required flag '%s'", p.Name)
		}
	}

	return a.Dispatch(path.HttpMethod, path.Path, coerced)
}

// Parse command and flags and infer resource, action, href and params
func parseCommandAndFlags(cmd string) (resource *metadata.Resource, action *metadata.Action,
	path *metadata.ActionPath, href string, params []string, err error) {

	var flags = flagValues[cmd]
	if flags == nil {
		err = fmt.Errorf("Invalid command line, try --help.")
		return
	}
	var elems = strings.Split(cmd, " ")
	var actionName = elems[len(elems)-1]
	href = flags.Href
	if !strings.HasPrefix(href, "/api") {
		if strings.HasPrefix(href, "/") {
			href = "/api" + href
		} else {
			href = "/api/" + href
		}
	}
	var vars []*metadata.PathVariable
	for _, res := range api_metadata {
		var err error
		if vars, err = res.ExtractVariables(href); err == nil {
			resource = res
			break
		}
	}
	if resource == nil {
		err = fmt.Errorf("Invalid href '%s' (does not match any known href)", href)
		return
	}
	for _, a := range resource.Actions {
		if a.Name == actionName {
			action = a
			break
		}
	}
	if action == nil {
		var supported = make([]string, len(resource.Actions))
		for i, a := range resource.Actions {
			supported[i] = a.Name
		}
		err = fmt.Errorf("Unknown %s action '%s'. Supported actions are: %s",
			resource.Name, actionName, strings.Join(supported, ", "))
		return
	}

	if path, err = action.Url(vars); err != nil {
		return
	}

	params = flags.Params
	return
}

// Validate flag value using validation criteria provided in metadata
func validateFlagValue(value string, param *metadata.ActionParam) error {
	if param.Regexp != nil {
		if !param.Regexp.MatchString(value) {
			return fmt.Errorf("Invalid value '%s' for '%s', value must validate /%s/",
				value, param.Name, param.Regexp.String())
		}
	}
	if param.NonBlank && value == "" {
		return fmt.Errorf("Invalid value for '%s', value must not be blank",
			param.Name)
	}
	if len(param.ValidValues) > 0 {
		var found = false
		for _, v := range param.ValidValues {
			if v == value {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("Invalid value for '%s', value must be one of %s, value provided was '%s'",
				param.Name, strings.Join(param.ValidValues, ", "), value)
		}
	}

	return nil
}
