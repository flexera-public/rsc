package rsapi

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rightscale/rsc/metadata"
)

// Result of parsing the command line (ParseCommand method).
// The parser infers the Uri and HTTP method of the request that need to be made and builds the map
// of parameters coerced to the right type (as dictated by the API metadata).
type ParsedCommand struct {
	HttpMethod string
	Uri        string
	Params     map[string]interface{}
}

// Data structure initialized during command line parsing and used by ParseCommand to create the
// final ParsedCommand struct. Each specific API client registers the action commands under the
// main command line tool command. kingpin then initializes the values from the command line.
type ActionCommand struct {
	Href     string   // Resource or collection href
	Params   []string // Action parameters
	ShowHelp string   // Whether to list flags supported by resource action
}

// Action command values indexed by action name
type ActionCommands map[string]*ActionCommand

// Actually run command
func (a *Api) ParseCommand(cmd string, values ActionCommands) (*ParsedCommand, error) {
	// 1. Initialize / find href as well as resource and action command definitions
	var resource, action, path, _, params, err = a.parseCommandAndFlags(cmd, values)
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
			var val, err = strconv.Atoi(value)
			if err != nil {
				return nil, fmt.Errorf("Value for '%s' must be an integer, value provided was '%s'",
					flag.Name, value)
			}
			coerced[flag.Name] = val
		case "[]int":
			var val, err = strconv.Atoi(value)
			if err != nil {
				return nil, fmt.Errorf("Value for '%s' must be an integer, value provided was '%s'",
					flag.Name, value)
			}
			if v, ok := coerced[flag.Name]; ok {
				v = append(v.([]int), val)
			} else {
				coerced[flag.Name] = []int{val}
			}
		case "bool":
			if _, ok := coerced[name]; ok {
				return nil, fmt.Errorf("Multiple values specified for '%s'", flag.Name)
			}
			var val, err = strconv.ParseBool(value)
			if err != nil {
				return nil, fmt.Errorf("Value for '%s' must be a bool, value provided was '%s'",
					flag.Name, value)
			}
			coerced[flag.Name] = val
		case "[]bool":
			var val, err = strconv.ParseBool(value)
			if err != nil {
				return nil, fmt.Errorf("Value for '%s' must be a bool, value provided was '%s'",
					flag.Name, value)
			}
			if v, ok := coerced[flag.Name]; ok {
				v = append(v.([]bool), val)
			} else {
				coerced[flag.Name] = []bool{val}
			}
		}
	}
	for _, p := range action.Params {
		_, ok := coerced[p.Name]
		if p.Mandatory && !ok {
			return nil, fmt.Errorf("Missing required flag '%s'", p.Name)
		}
	}

	return &ParsedCommand{path.HttpMethod, path.Path, coerced}, nil
}

// Show help for given command and flags
func (a *Api) ShowHelp(cmd string, values ActionCommands) error {
	var resource, action, _, href, _, err = a.parseCommandAndFlags(cmd, values)
	if err != nil {
		return err
	}
	if len(action.Params) == 0 {
		fmt.Printf("usage: rsc [<flags>] %s %s %s\n", strings.Split(cmd, " ")[0],
			action.Name, href)
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
	fmt.Printf("usage: rsc [<flags>] %s %s %s [<%s %s params>]\n\n", strings.Split(cmd, " ")[0],
		action.Name, href, strings.ToLower(resource.Name), action.Name)
	fmt.Printf("%s %s params:\n%s\n", resource.Name, action.Name, strings.Join(flagHelp, "\n\n"))
	return nil
}

// Parse command and flags and infer resource, action, href and params
func (a *Api) parseCommandAndFlags(cmd string, commandValues ActionCommands) (resource *metadata.Resource, action *metadata.Action,
	path *metadata.ActionPath, href string, params []string, err error) {

	var flags = commandValues[cmd]
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
	for _, res := range a.Metadata {
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
