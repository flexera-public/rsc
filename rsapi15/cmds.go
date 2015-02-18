package rsapi15

import (
	"fmt"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"gopkg.in/alecthomas/kingpin.v1"
)

// Resource command
type ResourceCmd struct {
	Name        string
	Description string
	HrefRegexp  *regexp.Regexp // Regular expression that matches resource hrefs
	Actions     []*ActionCmd
}

// Resource action subcommand
type ActionCmd struct {
	Name        string
	Description string
	Flags       []*ActionFlag
}

// Resource action subcommand flag
type ActionFlag struct {
	Name        string
	Description string
	Type        string
	IsPattern   bool
	Mandatory   bool
	NonBlank    bool
	Regexp      *regexp.Regexp
	ValidValues []string
}

// Flag values indexed by flag name indexed by command name
type FlagValues map[string]*ActionParams

// Action params consist of href and map of query parameters
type ActionParams struct {
	Href     string       // Resource or collection href
	Params   *ParamsValue // Action parameters
	ShowHelp string       // Whether to list flags supported by resource action
}

// Map that holds flag values resulting from command line parsing
var flagValues FlagValues

// Command used to make API 1.5 requests
const api15Command = "api15"

// Register all commands with kinpin application
func RegisterCommands(app *kingpin.Application) {
	flagValues = make(FlagValues)
	var api15Cmd = app.Command(api15Command, "RightScale API 1.5 client")
	var actionNames = make([]string, len(actionMap))
	var i = 0
	for actionName, _ := range actionMap {
		actionNames[i] = actionName
		i += 1
	}
	sort.Strings(actionNames)
	actionNames = append([]string{"index", "show", "create", "update", "delete"}, actionNames...)
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
			for name, actions := range resourceActions {
				for _, a := range actions {
					if a.Name == action {
						actionDescription = a.Description
						resources = append(resources, name)
					}
				}
			}
			if len(resources) == 1 {
				description = actionDescription
			} else {
				description = "Action of resources " + strings.Join(resources[:len(resources)], ", ") + " and " + resources[len(resources)-1]
			}
		}
		var actionCmd = api15Cmd.Command(action, description)
		var actionParams = ActionParams{}
		actionCmd.Arg("href", "API Resource or resource collection href on which to act, e.g. '/api/servers'").Required().StringVar(&actionParams.Href)
		actionParams.Params = Params(actionCmd.Flag("params", "Action parameters in the form QUERY=VALUE, e.g. '-P server[name]=server42'").Short('P').PlaceHolder("QUERY=VALUE"))
		var key = fmt.Sprintf("%s %s", api15Command, action)
		flagValues[key] = &actionParams
	}
}

// Show help for given command and flags
func (a *Api15) ShowHelp(cmd string) error {
	var resource, action, href, _, err = parseCommandAndFlags(cmd)
	if err != nil {
		return err
	}
	var flagHelp = make([]string, len(action.Flags))
	for i, f := range action.Flags {
		var attrs = f.Type
		if f.Mandatory {
			attrs += ", required"
		} else {
			attrs += ", optional"
		}
		flagHelp[i] = fmt.Sprintf("--%s\n    <%s> %s", f.Name, attrs, f.Description)
	}
	fmt.Printf("usage: rsc [<flags>] api15 %s [<%s.%s flags>] %s\n", action.Name,
		resource.Name, action.Name, href)
	fmt.Printf("<%s.%s flags>:\n%s\n", resource.Name, action.Name, strings.Join(flagHelp, "\n\n"))
	return nil
}

// Actually run command
func (a *Api15) RunCommand(cmd string) (*http.Response, error) {
	// 1. Initialize / find href as well as resource and action command definitions
	var resource, action, href, params, err = parseCommandAndFlags(cmd)
	if err != nil {
		return nil, err
	}

	// 2. Coerce and validate given flag values
	var coerced = map[string]interface{}{}
	for name, value := range params {
		var flag *ActionFlag
		for _, f := range action.Flags {
			if f.Name == name {
				flag = f
				break
			}
		}
		if flag == nil {
			var supported = make([]string, len(action.Flags))
			for i, f := range action.Flags {
				supported[i] = f.Name
			}
			return nil, fmt.Errorf("Unknown %s.%s flag '%s'. Supported flags are: %s",
				resource.Name, action.Name, name, strings.Join(supported, ", "))
		}
		switch flag.Type {
		case "string":
			if len(value) > 1 {
				return nil, fmt.Errorf("Multiple values specified for '%s'", flag.Name)
			}
			if err := validateFlagValue(value[0], flag); err != nil {
				return nil, err
			}
			coerced[flag.Name] = value[0]
		case "[]string":
			for _, val := range value {
				if err := validateFlagValue(val, flag); err != nil {
					return nil, err
				}
			}
			coerced[flag.Name] = value
		case "int":
			if len(value) > 1 {
				return nil, fmt.Errorf("Multiple values specified for '%s'", flag.Name)
			}
			if err := validateFlagValue(action.Name, flag); err != nil {
				return nil, err
			}
			val, err := strconv.Atoi(value[0])
			if err != nil {
				return nil, fmt.Errorf("Value for '%s' must be an integer, value provided was '%s'",
					flag.Name, value[0])
			}
			coerced[flag.Name] = val
		case "map":
			var val = make(map[string]string)
			for _, v := range value {
				elems := strings.Split(v, ":")
				if len(elems) != 2 {
					return nil, fmt.Errorf("Values of '%s' must be of the form NAME:VALUE, value provided was '%s'",
						flag.Name, v)
				}
				if err := validateFlagValue(elems[1], flag); err != nil {
					return nil, err
				}
				val[elems[0]] = elems[1]
			}
			coerced[flag.Name] = val
		}
	}
	for _, f := range action.Flags {
		_, ok := coerced[f.Name]
		if f.Mandatory && !ok {
			return nil, fmt.Errorf("Missing required flag '%s'", f.Name)
		}
	}

	return a.Do(href, action.Name, coerced)
}

// Parse command and flags and infer resource, action, href and params
func parseCommandAndFlags(cmd string) (resource *ResourceCmd, action *ActionCmd, href string, params ParamsValue, err error) {
	var elems = strings.Split(cmd, " ")
	if len(elems) != 2 {
		err = fmt.Errorf("Invalid command '%s'", cmd)
		return
	}
	var flags = flagValues[cmd]
	href = flags.Href
	if !strings.HasPrefix(href, "/api") {
		if strings.HasPrefix(href, "/") {
			href = "/api" + href
		} else {
			href = "/api/" + href
		}
	}
	var actionName = elems[1]
	for _, res := range commands {
		if res.HrefRegexp.MatchString(href) {
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
	params = *flags.Params
	return
}

// Validate flag value using validation criteria provided in metadata
func validateFlagValue(value string, flag *ActionFlag) error {
	if flag.Regexp != nil {
		if !flag.Regexp.MatchString(value) {
			return fmt.Errorf("Invalid value '%s' for '%s', value must validate /%s/",
				value, flag.Name, flag.Regexp.String())
		}
	}
	if flag.NonBlank && value == "" {
		return fmt.Errorf("Invalid value for '%s', value must not be blank",
			flag.Name)
	}
	if len(flag.ValidValues) > 0 {
		var found = false
		for _, v := range flag.ValidValues {
			if v == value {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("Invalid value for '%s', value must be one of %s, value provided was '%s'",
				flag.Name, strings.Join(flag.ValidValues, ", "), value)
		}
	}

	return nil
}

// Params command line argument parser

var stringMapRegex = regexp.MustCompile("[:=]")

type ParamsValue map[string][]string

func (p *ParamsValue) Set(value string) error {
	parts := stringMapRegex.Split(value, 2)
	if len(parts) != 2 {
		return fmt.Errorf("expected KEY=VALUE got '%s'", value)
	}
	(*p)[parts[0]] = append((*p)[parts[0]], parts[1])
	return nil
}

func (p *ParamsValue) String() string {
	return fmt.Sprintf("%s", map[string][]string(*p))
}

func (p *ParamsValue) IsCumulative() bool {
	return true
}

func Params(s kingpin.Settings) (target *ParamsValue) {
	target = &ParamsValue{}
	s.SetValue(target)
	return
}
