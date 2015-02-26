package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/rightscale/rsc/gen"

	"bitbucket.org/pkg/inflect"
)

// The analyzer struct holds the analysis results
type ApiAnalyzer struct {
	// Raw resources as defined in API json metadata
	rawResources map[string]interface{}
	// Attribute type mappings defined in attributes.json
	attributeTypes map[string]string
	// Temporary analysis construct
	// Holds all types indexed by name, multiple actions could generate types with the same
	// name. Keep them all here then make names unique as needed once we gathered all of them.
	rawTypes map[string][]*gen.ObjectDataType
}

// Factory method for API analyzer
func NewApiAnalyzer(resources map[string]interface{}, attributeTypes map[string]string) *ApiAnalyzer {
	return &ApiAnalyzer{
		rawResources:   resources,
		attributeTypes: attributeTypes,
		rawTypes:       make(map[string][]*gen.ObjectDataType),
	}
}

// Analyze iterate through all resources and initializes the Resources and ParamTypes fields of
// the ApiAnalyzer struct accordingly.
func (a *ApiAnalyzer) Analyze() *gen.ApiDescriptor {
	var descriptor = &gen.ApiDescriptor{
		Resources: make(map[string]*gen.Resource),
		Types:     make(map[string]*gen.ObjectDataType),
	}
	var rawResourceNames = make([]string, len(a.rawResources))
	var idx = 0
	for n, _ := range a.rawResources {
		rawResourceNames[idx] = n
		idx += 1
	}
	sort.Strings(rawResourceNames)
	for _, name := range rawResourceNames {
		var resource = a.rawResources[name]
		a.AnalyzeResource(name, resource, descriptor)
	}
	descriptor.FinalizeTypeNames(a.rawTypes)
	return descriptor
}

// AnalyzeResource analyzes the given resource and updates the Resources and ParamTypes analyzer
// fields accordingly
func (a *ApiAnalyzer) AnalyzeResource(name string, resource interface{}, descriptor *gen.ApiDescriptor) {
	var res = resource.(map[string]interface{})

	// Compute description
	var description string
	if d, ok := res["description"].(string); ok {
		description = d
	}

	// Compute attributes
	var attributes []*gen.Attribute
	var atts map[string]interface{}
	if m, ok := res["media_type"].(map[string]interface{}); ok {
		atts = m["attributes"].(map[string]interface{})
		attributes = make([]*gen.Attribute, len(atts))
		for idx, n := range sortedKeys(atts) {
			attributes[idx] = &gen.Attribute{n, inflect.Camelize(n), a.attributeTypes[n]}
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
		var pathPatterns = ParseRoute(fmt.Sprintf("%s#%s", name, actionName),
			meth["route"].(string))
		if len(pathPatterns) == 0 {
			// Custom action
			continue
		}
		var allParamNames = make([]string, len(params))
		var i = 0
		for n, _ := range params {
			allParamNames[i] = n
			i += 1
		}
		sort.Strings(allParamNames)
		var pathParamNames = []string{}
		for _, pattern := range pathPatterns {
			for _, v := range pattern.Variables {
				var existing = false
				for _, e := range pathParamNames {
					existing = e == v
					if existing {
						break
					}
				}
				if !existing {
					pathParamNames = append(pathParamNames, v)
				}
			}
		}
		var queryParamNames = []string{}
		var payloadParamNames = []string{}
		var hasPayload = pathPatterns[0].HttpMethod == "POST" || pathPatterns[0].HttpMethod == "PUT"
		for _, p := range allParamNames {
			if isQueryParam(p) {
				queryParamNames = append(queryParamNames, p)
			} else if hasPayload && !strings.Contains(p, "[") {
				payloadParamNames = append(payloadParamNames, p)
			}
		}
		var contentType string
		if c, ok := meth["content_type"].(string); ok {
			contentType = c
		}
		var paramAnalyzer = NewAnalyzer(params)
		paramAnalyzer.Analyze()

		// Record new parameter types
		var paramTypeNames = make([]string, len(paramAnalyzer.ParamTypes))
		var idx = 0
		for n, _ := range paramAnalyzer.ParamTypes {
			paramTypeNames[idx] = n
			idx += 1
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
		var mandatory = []string{}
		var optional = []string{}
		for _, p := range paramAnalyzer.Params {
			if p.Mandatory {
				if p.Description != "" {
					var desc = fmt.Sprintf("%s: %s", p.VarName, p.Description)
					mandatory = append(mandatory, desc)
				}
			} else {
				var desc = p.Name
				if p.Description != "" {
					desc += ": " + p.Description
				}
				optional = append(optional, desc)
			}
		}
		sort.Strings(mandatory)
		sort.Strings(optional)
		if len(mandatory) > 0 {
			description += "\n\t" + strings.Join(mandatory, "\n\t")
		}
		if len(optional) > 0 {
			description += "\n-- Optional parameters:\n\t" +
				strings.Join(optional, "\n\t")
		}

		// Record action
		var action = gen.Action{
			Name:              actionName,
			MethodName:        inflect.Camelize(actionName),
			Description:       removeBlankLines(description),
			ResourceName:      name,
			PathPatterns:      pathPatterns,
			Params:            paramAnalyzer.Params,
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
		ClientName:  "Api15",
		Description: removeBlankLines(description),
		Actions:     actions,
		Attributes:  attributes,
	}
}

/** Helper methods for parsing raw JSON **/

// Regular expression used to extract routes from JSON
var routeRegexp = regexp.MustCompile(`\{[^\}]+\}`)

// Regular expression that captures variables in a path
var routeVariablesRegexp = regexp.MustCompile(`/:([^/]+)`)

func ParseRoute(moniker string, route string) (pathPatterns []*gen.PathPattern) {
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
		method, paths = "POST", []string{"/api/servers/:id/teminate"}
	default:
		var bounds = routeRegexp.FindAllStringIndex(route, -1)
		var matches = make([]string, len(bounds))
		var prev = 0
		for i, bound := range bounds {
			matches[i] = route[prev:bound[0]]
			prev = bound[1]
		}
		method = strings.TrimRight(matches[0][0:7], " ")
		paths = make([]string, len(bounds))
		var j = 0
		for _, r := range matches {
			var path = strings.TrimRight(r[7:], " ")
			path = strings.TrimSuffix(path, "(.:format)?")
			if isDeprecated(path) || isCustom(method, path) {
				continue
			}
			paths[j] = path
			j += 1
		}
		paths = paths[:j]
	}
	pathPatterns = make([]*gen.PathPattern, len(paths))
	for i, p := range paths {
		var rx = routeVariablesRegexp.ReplaceAllLiteralString(regexp.QuoteMeta(p), `/([^/]+)`)
		rx = fmt.Sprintf("^%s$", rx)
		var pattern = gen.PathPattern{
			HttpMethod: method,
			Path:       p,
			Pattern:    routeVariablesRegexp.ReplaceAllLiteralString(p, "/%s"),
			Regexp:     rx,
		}
		var matches = routeVariablesRegexp.FindAllStringSubmatch(p, -1)
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
		return fmt.Sprintf("*%s", resourceType(resName))
	case "index":
		return fmt.Sprintf("[]*%s", resourceType(resName))
	case "create":
		if _, ok := noMediaTypeResources[resName]; ok {
			return "map[string]interface{}"
		} else {
			return "*" + inflect.Singularize(resName) + "Locator"
		}
	case "update", "destroy":
		return ""
	default:
		switch {
		case len(contentType) == 0:
			return ""
		case strings.Index(contentType, "application/vnd.rightscale.") == 0:
			if contentType == "application/vnd.rightscale.text" {
				return "string"
			}
			var elems = strings.SplitN(contentType[27:], ";", 2)
			var name = resourceType(inflect.Camelize(elems[0]))
			if len(elems) > 1 && elems[1] == "type=collection" {
				name = "[]*" + name
			}
			return name
		default: // Shouldn't be here
			panic("api15gen: Unknown content type " + contentType)
		}
	}

}

// Name of go type for resource with given name
// It should always be the same (camelized) but there are some resources that don't have a media
// type so for these we use a map.
func resourceType(resName string) string {
	if resName == "ChildAccounts" {
		return "Account"
	}
	if _, ok := noMediaTypeResources[resName]; ok {
		return "map[string]string"
	} else {
		return inflect.Singularize(resName)
	}
}

// Return keys of given maps sorted
func sortedKeys(m map[string]interface{}) []string {
	var keys = make([]string, len(m))
	var idx = 0
	for k, _ := range m {
		keys[idx] = k
		idx += 1
	}
	sort.Strings(keys)
	return keys
}
