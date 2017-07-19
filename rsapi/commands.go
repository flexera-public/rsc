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
	"time"

	"github.com/rightscale/rsc/metadata"
)

// ParsedCommand is the result of parsing the command line (ParseCommand method).
// The parser infers the Uri and HTTP method of the request that need to be made and builds the
// maps of query string and payload parameters. The parameter values are all coerced to the type
// dictated by the API metadata.
type ParsedCommand struct {
	HTTPMethod    string
	URI           string
	QueryParams   APIParams
	PayloadParams APIParams
}

// ActionCommand is a data structure initialized during command line parsing and used by
// ParseCommand to create the final ParsedCommand struct.
// Each specific API client registers the action commands under the main command line tool command.
// kingpin then initializes the values from the command line.
type ActionCommand struct {
	Href     string   // Resource or collection href
	Params   []string // Action parameters
	ShowHelp string   // Whether to list flags supported by resource action
}

// ActionCommands indexes action command values by action name.
type ActionCommands map[string]*ActionCommand

// ParseCommand actually runs a command.
func (a *API) ParseCommand(cmd, hrefPrefix string, values ActionCommands) (*ParsedCommand, error) {
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

	// 2. Coerce and validate given flag values, use array for payload as order is meaningful
	//    when building the data structure (e.g. in the case of array of hashes)
	var queryParams, payloadParams []APIParams
	var seen []string
	for _, p := range params {
		param, value, err := a.findParamAndValue(action, p)
		if err != nil {
			return nil, err
		}
		if param == nil {
			name := strings.SplitN(p, "=", 2)[0]
			if len(action.CommandFlags) > 0 {
				supported := make([]string, len(action.CommandFlags))
				for i, p := range action.CommandFlags {
					supported[i] = p.Name
				}
				return nil, fmt.Errorf("Unknown %s.%s flag '%s'. Supported flags are: %s",
					resource.Name, action.Name, name, strings.Join(supported, ", "))
			}
			return nil, fmt.Errorf("Unknown %s.%s flag '%s'. %s.%s does not accept any flag",
				resource.Name, action.Name, name, resource.Name, action.Name)
		}
		name := param.Name
		seen = append(seen, name)
		if err := validateFlagValue(value, param); err != nil {
			return nil, err
		}
		coerced := &queryParams
		if param.Location == metadata.PayloadParam {
			coerced = &payloadParams
		}
		switch param.Type {
		case "string", "[]string", "interface{}":
			*coerced = append(*coerced, APIParams{name: value})
		case "int", "[]int":
			val, err := strconv.Atoi(value)
			if err != nil {
				return nil, fmt.Errorf("Value for '%s' must be an integer, value provided was '%s'",
					name, value)
			}
			*coerced = append(*coerced, APIParams{name: val})
		case "bool", "[]bool":
			val, err := strconv.ParseBool(value)
			if err != nil {
				return nil, fmt.Errorf("Value for '%s' must be a bool, value provided was '%s'",
					name, value)
			}
			*coerced = append(*coerced, APIParams{name: val})
		case "*time.Time":
			// Try to parse the input as int64 and if it was successful, treat is as a Unix
			// time otherwise parse as RFC3339.
			if val, err := strconv.ParseInt(value, 10, 64); err == nil {
				*coerced = append(*coerced, APIParams{name: time.Unix(val, 0)})
			} else {
				val, err := time.Parse(time.RFC3339, value)
				if err != nil {
					return nil, fmt.Errorf("Invalid time '%s' for %s: %s", value, name, err)
				}
				*coerced = append(*coerced, APIParams{name: val})
			}
		case "map":
			velems := strings.SplitN(value, "=", 2)
			if len(velems) != 2 {
				return nil, fmt.Errorf("Value for '%s' must be of the form NAME=VALUE, got %s", name, value)
			}
			*coerced = append(*coerced, APIParams{fmt.Sprintf("%s[%s]", name, velems[0]): velems[1]})
		case "sourcefile":
			file, err := os.Open(value)
			if err != nil {
				return nil, fmt.Errorf("Invalid file upload path '%s' for %s: %s", value, name, err)
			}
			*coerced = append(*coerced, APIParams{name: &SourceUpload{Reader: file}})
		case "file":
			file, err := os.Open(value)
			if err != nil {
				return nil, fmt.Errorf("Invalid file upload path '%s' for %s: %s", value, name, err)
			}
			*coerced = append(*coerced, APIParams{name: &FileUpload{Name: name, Filename: value, Reader: file}})
		}
	}
	for _, p := range action.CommandFlags {
		var ok bool
		for _, s := range seen {
			if s == p.Name {
				ok = true
				break
			}
		}
		if p.Mandatory && !ok {
			return nil, fmt.Errorf("Missing required flag '%s'", p.Name)
		}
	}

	// Reconstruct data structure from flat values
	qParams, err := buildQuery(queryParams)
	if err != nil {
		return nil, err
	}
	pParams, err := buildPayload(payloadParams)
	if err != nil {
		return nil, err
	}

	return &ParsedCommand{path.HTTPMethod, path.Path, qParams, pParams}, nil
}

// ShowHelp displays help for the given command and flags.
func (a *API) ShowHelp(cmd, hrefPrefix string, values ActionCommands) error {
	target, _, err := a.ParseCommandAndFlags(cmd, hrefPrefix, values)
	if err != nil {
		return err
	}
	_, action, href := target.Resource, target.Action, target.Href
	if len(action.CommandFlags) == 0 {
		fmt.Printf("usage: rsc [<FLAGS>] %s %s %s\n\n%s\n", strings.Split(cmd, " ")[0],
			action.Name, href, action.Description)
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
	fmt.Printf("usage: rsc [<FLAGS>] %s %s %s [<PARAMS>]\n\n%s\n\n", strings.Split(cmd, " ")[0],
		action.Name, href, action.Description)
	fmt.Printf("PARAMS:\n%s\n", strings.Join(flagHelp, "\n\n"))
	return nil
}

// CommandTarget is the target of a command: resource type and href as well as action details.
type CommandTarget struct {
	Resource *metadata.Resource   // Resource command applies to
	Action   *metadata.Action     // Action command corresponds to
	Path     *metadata.ActionPath // Action path
	Href     string               // Resource href
}

// ParseCommandAndFlags parses a command flag and infers the resource, action, href and params.
func (a *API) ParseCommandAndFlags(cmd, hrefPrefix string, values ActionCommands) (*CommandTarget, []string, error) {
	resource, vars, err := a.parseResource(cmd, hrefPrefix, values)
	if err != nil {
		return nil, nil, err
	}
	var action *metadata.Action
	elems := strings.Split(cmd, " ")
	actionName := elems[len(elems)-1]
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

	path, err := action.URL(vars)
	if err != nil {
		return nil, nil, err
	}
	flags := values[cmd]

	return &CommandTarget{resource, action, path, flags.Href}, flags.Params, nil
}

// ShowActions prints all known actions for an API or the selected resource if any.
func (a *API) ShowActions(cmd, hrefPrefix string, values ActionCommands) error {
	res, _, err := a.parseResource(cmd, hrefPrefix, values)
	resource := ""
	if err == nil {
		resource = res.Name
	}
	actions := make(map[string][][2]string)
	resNames := make([]string, len(a.Metadata))
	idx := 0
	for n := range a.Metadata {
		resNames[idx] = n
		idx++
	}
	sort.Strings(resNames)
	for _, resName := range resNames {
		res := a.Metadata[resName]
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
				subPattern := pattern.Pattern
				pat := fmt.Sprintf(subPattern, ivars...)
				var ok bool
				pat, ok = shortenPattern(res, pat, "/actions/"+action.Name)
				if !ok {
					pat, _ = shortenPattern(res, pat, "/"+action.Name)
				}
				actions[action.Name] = append(actions[action.Name], [2]string{pat, resName})
			}
		}
	}
	keys := make([]string, len(actions))
	i := 0
	for k := range actions {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	var lines []string
	actionTitle := "Action"
	hrefTitle := "Href"
	resourceTitle := "Resource"
	maxActionLen := len(actionTitle)
	maxHrefLen := len(hrefTitle)
	maxResourceLen := len(resourceTitle)
	for _, name := range keys {
		if len(name) > maxActionLen {
			maxActionLen = len(name)
		}
		as := actions[name]
		for _, action := range as {
			if len(action[0]) > maxHrefLen {
				maxHrefLen = len(action[0])
			}
			if len(action[1]) > maxResourceLen {
				maxResourceLen = len(action[1])
			}
		}
	}
	for idx, name := range keys {
		as := actions[name]
		for i, action := range as {
			title := ""
			if i == 0 {
				title = name
				if idx > 0 {
					lines = append(lines, fmt.Sprintf("%s\t%s\t%s",
						strings.Repeat("-", maxActionLen),
						strings.Repeat("-", maxHrefLen),
						strings.Repeat("-", maxResourceLen)))
				}
			}
			lines = append(lines, fmt.Sprintf("%s\t%s\t%s", title, action[0], action[1]))
		}
	}
	w := tabwriter.NewWriter(os.Stdout, 0, 4, 1, ' ', 0)
	w.Write([]byte(fmt.Sprintf("%s\t%s\t%s\n", actionTitle, hrefTitle, resourceTitle)))
	w.Write([]byte(fmt.Sprintf("%s\t%s\t%s\n", strings.Repeat("=", maxActionLen),
		strings.Repeat("=", maxHrefLen),
		strings.Repeat("=", maxResourceLen))))
	w.Write([]byte(strings.Join(lines, "\n")))
	w.Write([]byte("\n"))
	return w.Flush()
}

// Identify action resource if any
// Return error if no resource could be identified
//
// This is a "subtle" algorithm...
// We want to be able to:
//
// * Match a resource action on partial href, e.g. /clouds/1/instances/2 should be ok to invoke
//   terminate (whose path has /terminate at the end)
//
// * If two resources have actions with different names but same path pattern then return the right
//   one, e.g. both Deployment and Server have an action with path /deployments/1/servers/2 but it's
//   called "servers" in the case of Deployment and "index" in the case of servers.
//
// * If more than one actions on different resources have the same name and match the href then we
//   want to fail so we don't risk running the action on the wrong resource.
func (a *API) parseResource(cmd, hrefPrefix string, commandValues ActionCommands) (*metadata.Resource, []*metadata.PathVariable, error) {
	elems := strings.Split(cmd, " ")
	actionName := elems[len(elems)-1]
	flags := commandValues[cmd]
	if flags == nil {
		return nil, nil, fmt.Errorf("Invalid command line, try --help.")
	}
	href := flags.Href
	if hrefPrefix != "" && !strings.HasPrefix(href, hrefPrefix) {
		href = path.Join(hrefPrefix, href)
	}

	var vars []*metadata.PathVariable
	var candidates []*metadata.Resource
Metadata:
	for _, res := range a.Metadata {
		if v, err := res.ExtractVariables(href); err == nil {
			vars = v
			for _, a := range res.Actions {
				if a.Name == actionName && a.MatchHref(href) {
					// We found an exact match!
					candidates = []*metadata.Resource{res}
					break Metadata
				}
			}
			candidates = append(candidates, res)
		}
	}
	if len(candidates) == 0 {
		return nil, nil, fmt.Errorf("Invalid href '%s'. Try '%s %s actions'.", href,
			os.Args[0], strings.Split(cmd, " ")[0])
	}
	if len(candidates) > 1 {
		return nil, nil, fmt.Errorf("Invalid href '%s'. Multiple resources have an action '%s' that match this href. Try '%s %s actions' and specify a more complete href.",
			href, actionName, os.Args[0], strings.Split(cmd, " ")[0])
	}
	return candidates[0], vars, nil
}

// Capture enumerable input key value
var captureEnumRegex = regexp.MustCompile(`.*\[(.+)\]$`)

// Extract parameter from metadata that correspond to given command line flag
// The flag is of the form "NAME=VALUE", the one complication is for the "enumerable" case where
// the flag is of the form "NAME[KEY]=VALUE" (e.g. inputs).
// First check if there is a parameter with the given flag name, then check for the enumerable case
// if not found.
func (a *API) findParamAndValue(action *metadata.Action, flag string) (*metadata.ActionParam, string, error) {
	elems := strings.SplitN(flag, "=", 2)
	if len(elems) != 2 {
		return nil, "", fmt.Errorf("Arguments must be of the form NAME=VALUE, value provided was '%s'", flag)
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
	if param == nil && strings.Contains(name, "[") {
		// Handle enumerable case
		name = name[:strings.LastIndex(name, "[")]
		for _, ap := range action.CommandFlags {
			if ap.Name == name {
				param = ap
				es := captureEnumRegex.FindStringSubmatch(elems[0])
				if es == nil {
					return nil, "", fmt.Errorf("Key/value arguments must be of the form NAME[KEY]=VALUE, value provided was '%s'", flag)
				}
				value = es[1] + "=" + value
				break
			}
		}
	}
	return param, value, nil
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
	if len(param.ValidValues) > 0 && param.Name != "filter[]" { // filter[] is special: it has values just so --help can list them
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

// Reconstruct query from flatten values
func buildQuery(values []APIParams) (APIParams, error) {
	query := APIParams{}
	for _, value := range values {
		// Only one iteration below, flatten params only have one element each
		for name, param := range value {
			if q, ok := query[name]; ok {
				if a, ok := q.([]interface{}); ok {
					query[name] = append(a, param)
				} else {
					query[name] = []interface{}{q, param}
				}
			} else {
				query[name] = param
			}
		}
	}
	return query, nil
}

// Reconstruct payload map from flatten values
func buildPayload(values []APIParams) (APIParams, error) {
	payload := APIParams{}
	for _, value := range values {
		// Only one iteration below, flatten params only have one element each
		for name, param := range value {
			if _, err := Normalize(payload, name, param); err != nil {
				return nil, err
			}
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

// Normalize recursively traverses a query string encoded name and builds up the corresponding
// structure.
// "a[b]" produces a map[string]interface{} with key "b"
// "a[]" produces a []interface{}
func Normalize(payload APIParams, name string, v interface{}) (APIParams, error) {
	var matches = nameRegex.FindStringSubmatch(name)
	if len(matches) == 0 {
		return nil, nil
	}
	var k = matches[1]
	if len(k) == 0 {
		return nil, nil
	}
	var after = matches[2]
	if after == "" {
		payload[k] = v
	} else if after == "[" {
		payload[name] = v
	} else if after == "[]" {
		if _, ok := payload[k]; !ok {
			payload[k] = []interface{}{}
		}
		a, ok := payload[k].([]interface{})
		if !ok {
			return nil, fmt.Errorf("expected array for param '%s'", k)
		}
		payload[k] = append(a, v)
	} else {
		matches = childRegexp.FindStringSubmatch(after)
		if len(matches) == 0 {
			matches = childRegexp2.FindStringSubmatch(after)
		}
		if len(matches) > 0 {
			var childKey = matches[1]
			if _, ok := payload[k]; !ok {
				payload[k] = []interface{}{}
			}
			var array []interface{}
			if a, ok := payload[k].([]interface{}); ok {
				array = a
			} else {
				return nil, fmt.Errorf("expected array for param '%s'", k)
			}
			var handled bool
			if len(array) > 0 {
				if last, ok := array[len(array)-1].(APIParams); ok {
					if _, ok := last[childKey]; !ok {
						handled = true
						Normalize(last, childKey, v)
					}
				}
			}
			if !handled {
				var p, err = Normalize(APIParams{}, childKey, v)
				if err != nil {
					return nil, err
				}
				payload[k] = append(array, p)
			}
		} else {
			if payload[k] == nil {
				payload[k] = APIParams{}
			}
			if _, ok := payload[k].(APIParams); !ok {
				return nil, fmt.Errorf("expected map for param '%s'", k)
			}
			var p, err = Normalize(payload[k].(APIParams), after, v)
			if err != nil {
				return nil, err
			}
			payload[k] = p
		}
	}
	return payload, nil
}

// Attempt to shorten action pattern for display by looking at other action hrefs
// and picking one that doesn't have the suffix if there is one.
func shortenPattern(res *metadata.Resource, pattern, suffix string) (string, bool) {
	if strings.HasSuffix(pattern, suffix) {
		pat := strings.TrimSuffix(pattern, suffix)
		for _, action := range res.Actions {
			for _, pattern2 := range action.PathPatterns {
				vars := pattern2.Variables
				ivars := make([]interface{}, len(vars))
				for i, v := range vars {
					ivars[i] = interface{}(":" + v)
				}
				subPattern := pattern2.Pattern
				pat2 := fmt.Sprintf(subPattern, ivars...)
				if pat == pat2 {
					return pat, true
				}
			}
		}
	}
	return pattern, false
}
