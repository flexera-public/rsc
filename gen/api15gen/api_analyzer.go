package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/rightscale/rsc/gen"

	"bitbucket.org/pkg/inflect"
)

// aliases maps the alias resource action to their original actions. This is used to
// build the metadata with correct parameters, etc.
var aliases = map[string]string{
	"Deployments#servers":               "Servers#index",
	"ServerArrays#current_instances":    "Instances#index",
	"ServerArrays#launch":               "Instances#launch",
	"ServerArrays#multi_run_executable": "Instances#multi_run_executable",
	"ServerArrays#multi_terminate":      "Instances#multi_terminate",
	"Servers#launch":                    "Instances#launch",
	"Servers#terminate":                 "Instances#terminate",
}

// APIAnalyzer holds the analysis results.
type APIAnalyzer struct {
	// Raw resources as defined in API json metadata
	rawResources map[string]interface{}
	// Attribute type mappings defined in attributes.json
	attributeTypes map[string]string
	// Temporary analysis construct
	// Holds all types indexed by name, multiple actions could generate types with the same
	// name. Keep them all here then make names unique as needed once we gathered all of them.
	rawTypes map[string][]*gen.ObjectDataType
}

// NewAPIAnalyzer is the factory method for the API analyzer
func NewAPIAnalyzer(resources map[string]interface{}, attributeTypes map[string]string) *APIAnalyzer {
	return &APIAnalyzer{
		rawResources:   resources,
		attributeTypes: attributeTypes,
		rawTypes:       make(map[string][]*gen.ObjectDataType),
	}
}

// Analyze iterate through all resources and initializes the Resources and ParamTypes fields of
// the APIAnalyzer struct accordingly.
func (a *APIAnalyzer) Analyze() *gen.APIDescriptor {
	a.AnalyzeAliases()
	var descriptor = &gen.APIDescriptor{
		Resources: make(map[string]*gen.Resource),
		Types:     make(map[string]*gen.ObjectDataType),
	}
	var rawResourceNames = make([]string, len(a.rawResources))
	var idx = 0
	for n := range a.rawResources {
		rawResourceNames[idx] = n
		idx++
	}
	sort.Strings(rawResourceNames)
	for _, name := range rawResourceNames {
		var resource = a.rawResources[name]
		a.AnalyzeResource(name, resource, descriptor)
	}
	descriptor.FinalizeTypeNames(a.rawTypes)
	return descriptor
}

// AnalyzeAliases goes through the aliases and copies the details from original actions to the
// aliased actions. It skips the route field since we have the routes hard-coded in the
// ParseRoute function.
func (a *APIAnalyzer) AnalyzeAliases() {
	for from, to := range aliases {
		splits := strings.SplitN(from, "#", 2)
		fromResName := splits[0]
		fromActionName := splits[1]
		splits = strings.SplitN(to, "#", 2)
		toResName := splits[0]
		toActionName := splits[1]

		fromRes := a.rawResources[fromResName]
		fromAct := fromRes.(map[string]interface{})["methods"].(map[string]interface{})[fromActionName].(map[string]interface{})

		toRes := a.rawResources[toResName]
		toAct := toRes.(map[string]interface{})["methods"].(map[string]interface{})[toActionName].(map[string]interface{})
		fromAct["parameters"] = toAct["parameters"]
		fromAct["status_code"] = toAct["status_code"]
		fromAct["access_rules"] = toAct["access_rules"]
	}
}

// AnalyzeResource analyzes the given resource and updates the Resources and ParamTypes analyzer
// fields accordingly
func (a *APIAnalyzer) AnalyzeResource(name string, resource interface{}, descriptor *gen.APIDescriptor) {
	var res = resource.(map[string]interface{})

	// Compute description
	var description string
	if d, ok := res["description"].(string); ok {
		description = d
	}

	// Compute attributes, mediatype identifier, and links
	var attributes []*gen.Attribute
	var identifier string
	var atts map[string]interface{}
	var links map[string]string
	if m, ok := res["media_type"].(map[string]interface{}); ok {
		// Compute attributes
		atts = m["attributes"].(map[string]interface{})
		attributes = make([]*gen.Attribute, len(atts))
		for idx, n := range sortedKeys(atts) {
			at, ok := a.attributeTypes[n+"#"+name]
			if !ok {
				at = a.attributeTypes[n]
			}
			attributes[idx] = &gen.Attribute{n, inflect.Camelize(n), at}
		}

		// Compute mediatype identifier
		contentTypes := m["content_types"].([]interface{})
		if len(contentTypes) > 0 {
			identifier = contentTypes[0].(string)
		}

		// Compute links
		if l, ok := m["links"].(map[string]interface{}); ok {
			links = make(map[string]string, len(l))
			for n, d := range l {
				links[n] = d.(string)
			}
		}
	} else {
		attributes = []*gen.Attribute{}
	}

	// Compute actions
	var methods = res["methods"].(map[string]interface{})
	var actionNames = sortedKeys(methods)
	var actions = []*gen.Action{}
	for _, actionName := range actionNames {
		var m = methods[actionName]
		var meth = m.(map[string]interface{})
		var params map[string]interface{}
		if p, ok := meth["parameters"]; ok {
			params = p.(map[string]interface{})
		}
		var description = "No description provided for " + actionName + "."
		if d, _ := meth["description"]; d != nil {
			description = d.(string)
		}
		var routes = meth["routes"].([]interface{})
		var strRoutes = make([]string, len(routes))
		for i, r := range routes {
			strRoutes[i] = r.(string)
		}
		var pathPatterns = ParseRoute(fmt.Sprintf("%s#%s", name, actionName), strRoutes)
		if len(pathPatterns) == 0 {
			// Custom action
			continue
		}
		// HACK: /api/right_scripts/:id/update_source is a unique snowflake. It has no
		// params, it accepts the script body as a text/plain type type!. We inject a
		// fake param here that acts as a multipart file upload does.
		if name == "RightScripts" && actionName == "update_source" {
			params["filename"] = map[string]interface{}{
				"class":       "SourceUpload",
				"mandatory":   true,
				"non_blank":   true,
				"description": "The file name to update the RightScript source with.",
			}
		}
		var allParamNames = make([]string, len(params))
		var i = 0
		for n := range params {
			allParamNames[i] = n
			i++
		}
		sort.Strings(allParamNames)

		var contentType string
		if c, ok := meth["content_type"].(string); ok {
			contentType = c
		}
		var paramAnalyzer = NewAnalyzer(params)
		paramAnalyzer.Analyze()

		// Record new parameter types
		var paramTypeNames = make([]string, len(paramAnalyzer.ParamTypes))
		var idx = 0
		for n := range paramAnalyzer.ParamTypes {
			paramTypeNames[idx] = n
			idx++
		}
		sort.Strings(paramTypeNames)
		for _, name := range paramTypeNames {
			var pType = paramAnalyzer.ParamTypes[name]
			if _, ok := a.rawTypes[name]; ok {
				a.rawTypes[name] = append(a.rawTypes[name], pType)
			} else {
				a.rawTypes[name] = []*gen.ObjectDataType{pType}
			}
		}

		// Update description with parameter descriptions
		var mandatory []string
		var optional []string
		for _, p := range paramAnalyzer.Params {
			if p.Mandatory {
				desc := p.Name
				if p.Description != "" {
					desc += ": " + strings.TrimSpace(p.Description)
				}
				mandatory = append(mandatory, desc)
			} else {
				desc := p.Name
				if p.Description != "" {
					desc += ": " + strings.TrimSpace(p.Description)
				}
				optional = append(optional, desc)
			}
		}
		if len(mandatory) > 0 {
			sort.Strings(mandatory)
			if !strings.HasSuffix(description, "\n") {
				description += "\n"
			}
			description += "Required parameters:\n\t" + strings.Join(mandatory, "\n\t")
		}
		if len(optional) > 0 {
			sort.Strings(optional)
			if !strings.HasSuffix(description, "\n") {
				description += "\n"
			}
			description += "Optional parameters:\n\t" + strings.Join(optional, "\n\t")
		}

		// Sort parameters by location
		actionParams := paramAnalyzer.Params
		leafParams := paramAnalyzer.LeafParams
		var pathParamNames []string
		var queryParamNames []string
		var payloadParamNames []string
		for _, p := range leafParams {
			n := p.Name
			if isQueryParam(n) {
				queryParamNames = append(queryParamNames, n)
				p.Location = gen.QueryParam
			} else if isPathParam(n, pathPatterns) {
				pathParamNames = append(pathParamNames, n)
				p.Location = gen.PathParam
			} else {
				payloadParamNames = append(payloadParamNames, n)
				p.Location = gen.PayloadParam
			}
		}
		for _, p := range actionParams {
			done := false
			for _, ap := range leafParams {
				if ap == p {
					done = true
					break
				}
			}
			if done {
				continue
			}
			n := p.Name
			if isQueryParam(n) {
				p.Location = gen.QueryParam
			} else if isPathParam(n, pathPatterns) {
				p.Location = gen.PathParam
			} else {
				p.Location = gen.PayloadParam
			}
		}

		// Mix in filters information
		if filters, ok := meth["filters"]; ok {
			var filterParam *gen.ActionParam
			for _, p := range actionParams {
				if p.Name == "filter" {
					filterParam = p
					break
				}
			}
			if filterParam != nil {
				values := sortedKeys(filters.(map[string]interface{}))
				ivalues := make([]interface{}, len(values))
				for i, v := range values {
					ivalues[i] = v
				}
				filterParam.ValidValues = ivalues
			}
		}

		// Record action
		action := gen.Action{
			Name:              actionName,
			MethodName:        inflect.Camelize(actionName),
			Description:       removeBlankLines(description),
			ResourceName:      inflect.Singularize(name),
			PathPatterns:      pathPatterns,
			Params:            actionParams,
			LeafParams:        paramAnalyzer.LeafParams,
			Return:            parseReturn(actionName, name, contentType),
			ReturnLocation:    actionName == "create" && name != "Oauth2",
			PathParamNames:    pathParamNames,
			QueryParamNames:   queryParamNames,
			PayloadParamNames: payloadParamNames,
		}
		actions = append(actions, &action)
	}

	// We're done!
	name = inflect.Singularize(name)
	descriptor.Resources[name] = &gen.Resource{
		Name:        name,
		ClientName:  "API",
		Description: removeBlankLines(description),
		Actions:     actions,
		Attributes:  attributes,
		LocatorFunc: LocatorFunc(attributes, name),
		Identifier:  identifier,
		Links:       links,
	}
}

/** Helper methods for parsing raw JSON **/

var (
	// Regular expression used to extract routes from JSON
	routeRegexp = regexp.MustCompile(`\{[^\}]+\}`)

	// Regular expression that captures variables in a path
	routeVariablesRegexp = regexp.MustCompile(`/:([^/]+)`)
)

// LocatorFunc returns the source code for building a locator instance from a resource.
func LocatorFunc(attributes []*gen.Attribute, name string) string {
	hasLinks := false
	for _, a := range attributes {
		if a.FieldName == "Links" {
			hasLinks = true
			break
		}
	}
	if !hasLinks {
		return ""
	}
	return `for _, l := range r.Links {
			if l["rel"] == "self" {
				return api.` + name + `Locator(l["href"])
			}
		}
		return nil`
}

// ParseRoute parses a API 1.5 route and returns corresponding path patterns.
func ParseRoute(moniker string, routes []string) (pathPatterns []*gen.PathPattern) {
	// :(((( some routes are empty
	var paths []string
	var method string
	switch moniker {
	case "Deployments#servers":
		method, paths = "GET", []string{"/api/deployments/:id/servers"}
	case "ServerArrays#current_instances":
		method, paths = "GET", []string{"/api/server_arrays/:id/current_instances"}
	case "ServerArrays#launch":
		method, paths = "POST", []string{"/api/server_arrays/:id/launch"}
	case "ServerArrays#multi_run_executable":
		method, paths = "POST", []string{"/api/server_arrays/:id/multi_run_executable"}
	case "ServerArrays#multi_terminate":
		method, paths = "POST", []string{"/api/server_arrays/:id/multi_terminate"}
	case "Servers#launch":
		method, paths = "POST", []string{"/api/servers/:id/launch"}
	case "Servers#terminate":
		method, paths = "POST", []string{"/api/servers/:id/terminate"}
	default:
		for _, route := range routes {
			bound := routeRegexp.FindStringIndex(route)
			match := route[0:bound[0]]
			method = strings.TrimRight(match[0:7], " ")
			path := strings.TrimRight(match[7:], " ")
			path = strings.TrimSuffix(path, "(.:format)?")
			if isDeprecated(path) || isCustom(method, path) {
				continue
			}
			paths = append(paths, path)
		}
	}
	pathPatterns = make([]*gen.PathPattern, len(paths))
	for i, p := range paths {
		rx := routeVariablesRegexp.ReplaceAllLiteralString(regexp.QuoteMeta(p), `/([^/]+)`)
		rx = fmt.Sprintf("^%s$", rx)
		pattern := gen.PathPattern{
			HTTPMethod: method,
			Path:       p,
			Pattern:    routeVariablesRegexp.ReplaceAllLiteralString(p, "/%s"),
			Regexp:     rx,
		}
		matches := routeVariablesRegexp.FindAllStringSubmatch(p, -1)
		if len(matches) > 0 {
			pattern.Variables = make([]string, len(matches))
			for i, m := range matches {
				pattern.Variables[i] = m[1]
			}
		}
		pathPatterns[i] = &pattern
	}
	return
}

// true if path is for a deprecated API
func isDeprecated(path string) bool {
	return strings.Contains(path, "/api/session") && !strings.Contains(path, "/api/sessions")
}

// Is action code not generated?
func isCustom(method, path string) bool {
	return method == "POST" && (path == "/api/sessions" || path == "/api/sessions/instance")
}

// Heuristic to determine whether given param is a query string param
// For now only consider view and filter...
func isQueryParam(n string) bool {
	return n == "view" || n == "filter"
}

// Look in given path patterns to chek whether given parameter name corresponds to a variable
// name.
func isPathParam(p string, pathPatterns []*gen.PathPattern) bool {
	for _, pattern := range pathPatterns {
		for _, v := range pattern.Variables {
			if p == v {
				return true
			}
		}
	}
	return false
}

// Resources that don't have a media type
var noMediaTypeResources = map[string]bool{
	"HealthCheck":          true,
	"Oauth2":               true,
	"Tag":                  true,
	"UserDatas":            true,
	"MonitoringMetricData": true,
	"ImportPreview":        true,
	"Changes":              true,
	"CookbookResolution":   true,
	"ResourceTag":          true,
}

func parseReturn(kind, resName, contentType string) string {
	switch kind {
	case "show":
		return refType(resName)
	case "index":
		return fmt.Sprintf("[]%s", refType(resName))
	case "create":
		if _, ok := noMediaTypeResources[resName]; ok {
			return "map[string]interface{}"
		}
		return "*" + inflect.Singularize(resName) + "Locator"
	case "update", "destroy":
		return ""
	case "current_instances":
		return "[]*Instance"
	default:
		switch {
		case len(contentType) == 0:
			return ""
		case strings.Index(contentType, "application/vnd.rightscale.") == 0:
			if contentType == "application/vnd.rightscale.text" {
				return "string"
			}
			elems := strings.SplitN(contentType[27:], ";", 2)
			name := refType(inflect.Camelize(elems[0]))
			if len(elems) > 1 && elems[1] == "type=collection" {
				name = "[]" + refType(inflect.Camelize(elems[0]))
			}
			return name
		default: // Shouldn't be here
			panic("api15gen: Unknown content type " + contentType)
		}
	}

}

// refType returns the Go code required to declare a variable of the given resource type.
// i.e. var foo *ResourceName
func refType(resName string) string {
	res := resourceType(resName)
	if res != "map[string]interface{}" {
		res = "*" + res
	}
	return res
}

// Name of go type for resource with given name
// It should always be the same (camelized) but there are some resources that don't have a media
// type so for these we use a map.
func resourceType(resName string) string {
	if resName == "ChildAccounts" {
		return "Account"
	}
	if _, ok := noMediaTypeResources[resName]; ok {
		return "map[string]interface{}"
	}
	return inflect.Singularize(resName)
}

// Return keys of given maps sorted
func sortedKeys(m map[string]interface{}) []string {
	keys := make([]string, len(m))
	idx := 0
	for k := range m {
		keys[idx] = k
		idx++
	}
	sort.Strings(keys)
	return keys
}
