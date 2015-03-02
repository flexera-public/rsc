package rsapi

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/rightscale/rsc/metadata"
)

// Result of parsing the command line (ParseCommand method).
// The parser infers the Uri and HTTP method of the request that need to be made and builds the
// maps of query string and payload parameters. The parameter values are all coerced to the type
// dictated by the API metadata.
type ParsedCommand struct {
	HttpMethod    string
	Uri           string
	QueryParams   ApiParams
	PayloadParams ApiParams
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
func (a *Api) ParseCommand(cmd, hrefPrefix string, values ActionCommands) (*ParsedCommand, error) {
	// 0. Show help if requested
	val, ok := values[cmd]
	if !ok {
		return nil, fmt.Errorf("Unknown command '%s'", cmd)
	}
	if val.ShowHelp != "" {
		a.ShowHelp(cmd, hrefPrefix, values)
		return nil, nil
	}

	// 1. Initialize / find href as well as resource and action command definitions
	target, params, err := a.ParseCommandAndFlags(cmd, hrefPrefix, values)
	if err != nil {
		return nil, err
	}
	resource, action, path := target.Resource, target.Action, target.Path

	// 2. Coerce and validate given flag values
	queryParams := ApiParams{}
	payloadParams := ApiParams{}
	for _, p := range params {
		elems := strings.SplitN(p, "=", 2)
		if len(elems) != 2 {
			return nil, fmt.Errorf("Arguments must be of the form NAME=VALUE, value provided was '%s'", p)
		}
		name := elems[0]
		value := elems[1]
		var param *metadata.ActionParam
		for _, ap := range action.CommandFlags {
			if ap.Name == name {
				param = ap
				break
			}
		}
		if param == nil {
			supported := make([]string, len(action.CommandFlags))
			for i, p := range action.CommandFlags {
				supported[i] = p.Name
			}
			return nil, fmt.Errorf("Unknown %s.%s flag '%s'. Supported flags are: %s",
				resource.Name, action.Name, name, strings.Join(supported, ", "))
		}
		if err := validateFlagValue(value, param); err != nil {
			return nil, err
		}
		coerced := queryParams
		if param.Location == metadata.PayloadParam {
			coerced = payloadParams
		}
		switch param.Type {
		case "string":
			if _, ok := coerced[name]; ok {
				return nil, fmt.Errorf("Multiple values specified for '%s'", param.Name)
			}
			coerced[param.Name] = value
		case "[]string":
			if v, ok := coerced[param.Name]; ok {
				v = append(v.([]string), value)
			} else {
				coerced[param.Name] = []string{value}
			}
		case "int":
			if _, ok := coerced[name]; ok {
				return nil, fmt.Errorf("Multiple values specified for '%s'", param.Name)
			}
			val, err := strconv.Atoi(value)
			if err != nil {
				return nil, fmt.Errorf("Value for '%s' must be an integer, value provided was '%s'",
					param.Name, value)
			}
			coerced[param.Name] = val
		case "[]int":
			val, err := strconv.Atoi(value)
			if err != nil {
				return nil, fmt.Errorf("Value for '%s' must be an integer, value provided was '%s'",
					param.Name, value)
			}
			if v, ok := coerced[param.Name]; ok {
				v = append(v.([]int), val)
			} else {
				coerced[param.Name] = []int{val}
			}
		case "bool":
			if _, ok := coerced[name]; ok {
				return nil, fmt.Errorf("Multiple values specified for '%s'", param.Name)
			}
			val, err := strconv.ParseBool(value)
			if err != nil {
				return nil, fmt.Errorf("Value for '%s' must be a bool, value provided was '%s'",
					param.Name, value)
			}
			coerced[param.Name] = val
		case "[]bool":
			val, err := strconv.ParseBool(value)
			if err != nil {
				return nil, fmt.Errorf("Value for '%s' must be a bool, value provided was '%s'",
					param.Name, value)
			}
			if v, ok := coerced[param.Name]; ok {
				v = append(v.([]bool), val)
			} else {
				coerced[param.Name] = []bool{val}
			}
		}
	}
	for _, p := range action.CommandFlags {
		_, ok := queryParams[p.Name]
		if !ok {
			_, ok = payloadParams[p.Name]
		}
		if p.Mandatory && !ok {
			return nil, fmt.Errorf("Missing required flag '%s'", p.Name)
		}
	}

	// Reconstruct data structure from flat values
	if payloadParams, err = buildPayload(payloadParams); err != nil {
		return nil, err
	}

	return &ParsedCommand{path.HttpMethod, path.Path, queryParams, payloadParams}, nil
}

// Show help for given command and flags
func (a *Api) ShowHelp(cmd, hrefPrefix string, values ActionCommands) error {
	target, _, err := a.ParseCommandAndFlags(cmd, hrefPrefix, values)
	resource, action, href := target.Resource, target.Action, target.Href
	if err != nil {
		return err
	}
	if len(action.CommandFlags) == 0 {
		fmt.Printf("usage: rsc [<flags>] %s %s %s\n", strings.Split(cmd, " ")[0],
			action.Name, href)
		return nil
	}
	flagHelp := make([]string, len(action.CommandFlags))
	for i, f := range action.CommandFlags {
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

// Target of command: resource type and href as well as action details
type CommandTarget struct {
	Resource *metadata.Resource   // Resource command applies to
	Action   *metadata.Action     // Action command corresponds to
	Path     *metadata.ActionPath // Action path
	Href     string               // Resource href
}

// Parse command and flags and infer resource, action, href and params
func (a *Api) ParseCommandAndFlags(cmd, hrefPrefix string, commandValues ActionCommands) (*CommandTarget, []string, error) {
	flags := commandValues[cmd]
	if flags == nil {
		return nil, nil, fmt.Errorf("Invalid command line, try --help.")
	}
	elems := strings.Split(cmd, " ")
	actionName := elems[len(elems)-1]
	href := flags.Href
	if hrefPrefix != "" && !strings.HasPrefix(href, hrefPrefix) {
		href = path.Join(hrefPrefix, href)
	}

	var vars []*metadata.PathVariable
	var resource *metadata.Resource
	for _, res := range a.Metadata {
		var err error
		if vars, err = res.ExtractVariables(href); err == nil {
			resource = res
			break
		}
	}
	if resource == nil {
		return nil, nil, fmt.Errorf("Invalid href '%s'. Try --hrefs.", href)
	}
	var action *metadata.Action
	for _, a := range resource.Actions {
		if a.Name == actionName {
			action = a
			break
		}
	}
	if action == nil {
		supported := make([]string, len(resource.Actions))
		for i, a := range resource.Actions {
			supported[i] = a.Name
		}
		return nil, nil, fmt.Errorf("Unknown %s action '%s'. Supported actions are: %s",
			resource.Name, actionName, strings.Join(supported, ", "))
	}

	path, err := action.Url(vars)
	if err != nil {
		return nil, nil, err
	}

	return &CommandTarget{resource, action, path, href}, flags.Params, nil
}

// Print all known href patterns for selected resource.
// Show all hrefs if no selected resource.
func (a *Api) ShowHrefs(cmd, hrefPrefix string, values ActionCommands) error {
	target, _, err := a.ParseCommandAndFlags(cmd, hrefPrefix, values)
	resource := ""
	if err == nil {
		resource = target.Resource.Name
	}
	knownHrefs := make(map[string]interface{})
	for resName, res := range a.Metadata {
		if resource != "" && resName != resource {
			continue
		}
		for _, action := range res.Actions {
			for _, pattern := range action.PathPatterns {
				vars := pattern.Variables
				ivars := make([]interface{}, len(vars))
				for i, v := range vars {
					ivars[i] = interface{}(":" + v)
				}
				pat := fmt.Sprintf(pattern.Pattern, ivars...)
				knownHrefs[pat] = fmt.Sprintf("%s.%s", resName, action.Name)
			}
		}
	}
	keys := make([]string, len(knownHrefs))
	i := 0
	for k, _ := range knownHrefs {
		keys[i] = k
		i += 1
	}
	sort.Strings(keys)
	lines := make([]string, len(keys))
	for i, pat := range keys {
		names := knownHrefs[pat]
		lines[i] = fmt.Sprintf("%s\t%s", pat, names)
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 4, 1, ' ', 0)
	w.Write([]byte("Href pattern\t<resource>.<action>\n"))
	w.Write([]byte(strings.Join(lines, "\n")))
	w.Write([]byte("\n"))
	return w.Flush()
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
		found := false
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

// Reconstruct payload map from flatten values
func buildPayload(values ApiParams) (ApiParams, error) {
	payload := map[string]interface{}{}
	for name, value := range values {
		if _, err := normalize(payload, name, value); err != nil {
			return nil, err
		}
	}
	return payload, nil
}

// Regular expressions used to capture parts of query string encoded param name by algorithm below
var (
	nameRegex    = regexp.MustCompile(`\A[\[\]]*([^\[\]]+)\]*(.*)`)
	childRegexp  = regexp.MustCompile(`^\[\]\[([^\[\]]+)\]$`)
	childRegexp2 = regexp.MustCompile(`^\[\](.+)$`)
)

// Recursively travers a query string encoded name and build up the corresponding structure.
// "a[b]" produces a map[string]interface{} with key "b"
// "a[]" produces a []interface{}
func normalize(payload ApiParams, name string, value interface{}) (ApiParams, error) {
	matches := nameRegex.FindStringSubmatch(name)
	if len(matches) == 0 {
		return nil, nil
	}
	k := matches[1]
	if len(k) == 0 {
		return nil, nil
	}
	after := matches[2]
	if after == "" {
		payload[k] = value
	} else if after == "[" {
		payload[name] = value
	} else if after == "[]" {
		if _, ok := payload[k]; !ok {
			payload[k] = []interface{}{}
		}
		if _, ok := payload[k].([]interface{}); !ok {
			return nil, fmt.Errorf("expected array for param '%s'", k)
		}
		payload[k] = append(payload[k].([]interface{}), value)
	} else {
		matches = childRegexp.FindStringSubmatch(after)
		if len(matches) == 0 {
			matches = childRegexp2.FindStringSubmatch(after)
		}
		if len(matches) > 0 {
			childKey := matches[1]
			if _, ok := payload[k]; !ok {
				payload[k] = []interface{}{}
			}
			if _, ok := payload[k].([]interface{}); !ok {
				return nil, fmt.Errorf("expected array for param '%s'", k)
			}
			array := payload[k].([]interface{})
			if len(array) == 0 {
				p, err := normalize(ApiParams{}, childKey, value)
				if err != nil {
					return nil, err
				}
				payload[k] = append(array, p)
			} else {
				last, ok := array[len(array)-1].(ApiParams)
				if ok {
					_, ok = last[childKey]
				}
				if ok {
					normalize(last, childKey, value)
				} else {
					p, err := normalize(ApiParams{}, childKey, value)
					if err != nil {
						return nil, err
					}
					payload[k] = append(array, p)
				}
			}
		} else {
			if payload[k] == nil {
				payload[k] = ApiParams{}
			}
			if _, ok := payload[k].(ApiParams); !ok {
				return nil, fmt.Errorf("expected map for param '%s'", k)
			}
			p, err := normalize(payload[k].(ApiParams), after, value)
			if err != nil {
				return nil, err
			}
			payload[k] = p
		}
	}
	return payload, nil
}
