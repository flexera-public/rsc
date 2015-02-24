package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/rightscale/rsc/gen"

	"bitbucket.org/pkg/inflect"
)

var (
	// Regexp used to replace e.g. ':cloud_id' from action URLs
	pathParamsRegex = regexp.MustCompile(`:[^/]+`)
)

// The analyzer struct holds the analysis results
type ApiAnalyzer struct {
	// Raw resources as defined in API json metadata
	rawResources map[string]map[string]interface{}
	// Raw types as defined in API json metadata
	rawTypes map[string]map[string]interface{}
	// Module name, e.g. "V1_6"
	moduleName string
}

// Factory method for API analyzer
func NewApiAnalyzer(resources, types map[string]map[string]interface{}) *ApiAnalyzer {
	var moduleName string
	if len(resources) > 0 {
		for n, _ := range resources {
			moduleName = strings.Split(n, "::")[0]
			break
		}
	}
	return &ApiAnalyzer{
		rawResources: resources,
		rawTypes:     types,
		moduleName:   moduleName,
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
	return descriptor
}

// Regular expression that catches href variables (e.g. ':id' in '/servers/:id')
var hrefVarRegexp = regexp.MustCompile(`/:[^/]+`)

// AnalyzeResource analyzes the given resource and updates the Resources and ParamTypes analyzer
// fields accordingly
func (a *ApiAnalyzer) AnalyzeResource(name string, resource map[string]interface{}, descriptor *gen.ApiDescriptor) error {
	// Compute description
	var description string
	if d, ok := resource["description"].(string); ok {
		description = d
	}

	// Compute attributes
	m, ok := resource["media_type"].(string)
	if !ok {
		return fmt.Errorf("Resource %s has no media type (missing key \"media_type\")", name)
	}
	t, ok := a.rawTypes[m]
	if !ok {
		return fmt.Errorf("Missing definition for Resource %s media type %s", name, m)
	}
	attrs, ok := t["attributes"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("Missing attributes for media type %s", m)
	}
	var attributes = make([]*gen.Attribute, len(attrs))
	for idx, n := range sortedKeys(attrs) {
		var attr = attrs[n].(map[string]interface{})
		at, ok := attr["type"].(map[string]interface{})
		if !ok {
			return fmt.Errorf("Missing attribute type for attribute %s of media type %s", n, m)
		}
		var tn, err = a.toType(n, at)
		if err != nil {
			return fmt.Errorf("Failed to compute type for attribute %s of media type %s: %s", n, m, err.Error())
		}
		attributes[idx] = &gen.Attribute{n, inflect.Camelize(n), tn}
	}

	// Compute actions
	var methods = resource["actions"].(map[string]interface{})
	var actionNames = sortedKeys(methods)
	var actions = []*gen.Action{}
	for _, actionName := range actionNames {
		var m = methods[actionName]
		var meth = m.(map[string]interface{})
		var params map[string]interface{}
		if p, ok := meth["params"]; ok {
			params = p.(map[string]interface{})
		}
		var description = "No description provided for " + actionName + "."
		if d, _ := meth["description"]; d != nil {
			description = d.(string)
		}
		httpMethod, pathPatterns := ParseRoute(meth["url"].(map[string]string))
		var allParamNames = make([]string, len(params))
		var i = 0
		for n, _ := range params {
			allParamNames[i] = n
			i += 1
		}
		sort.Strings(allParamNames)
		var queryParamNames = []string{}
		var payloadParamNames = []string{}
		var hasPayload = httpMethod == "POST" || httpMethod == "PUT"
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
			QueryParamNames:   queryParamNames,
			PayloadParamNames: payloadParamNames,
		}
		actions = append(actions, &action)
	}

	// We're done!
	name = inflect.Singularize(name)
	descriptor.Resources[name] = &gen.Resource{
		Name:        name,
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

func ParseRoute(moniker string, route string) (method string, pathPatterns []*gen.PathPattern) {
	// :(((( some routes are empty
	var paths []string
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
		var pattern = gen.PathPattern{
			Path:    p,
			Pattern: routeVariablesRegexp.ReplaceAllLiteralString(p, "/%s"),
			Regexp:  routeVariablesRegexp.ReplaceAllLiteralString(regexp.QuoteMeta(p), `/([^/]+)`),
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

// sameType returns true if both argument describe the same praxis type, false otherwise.
func sameType(praxisType, other map[string]interface{}) bool {
	var type1, ok = praxisType["type"].(map[string]interface{})
	if !ok {
		return false
	}
	name1, ok := type1["name"].(string)
	if !ok {
		return false
	}
	type2, ok := other["type"].(map[string]interface{})
	if !ok {
		return false
	}
	name2, ok := type2["name"]
	if !ok {
		return false
	}
	if name1 != name2 {
		return false
	}
	if name1 == "Struct" {
		attr1, ok := praxisType["attributes"].(map[string]interface{})
		if !ok {
			return false
		}
		attr2, ok := other["attributes"].(map[string]interface{})
		if !ok {
			return false
		}
		if len(attr1) != len(attr2) {
			return false
		}
		for an, av := range attr1 {
			avv, ok := av.(map[string]interface{})
			if !ok {
				return false
			}
			ov, ok := attr2[an]
			if !ok {
				return false
			}
			ovv, ok := ov.(map[string]interface{})
			if !ok {
				return false
			}
			if !sameType(avv, ovv) {
				return false
			}
		}
	}
	if name1 == "Collection" {
		mem1, ok := praxisType["member_attribute"].(map[string]interface{})
		if !ok {
			return false
		}
		mem2, ok := other["member_attribute"].(map[string]interface{})
		if !ok {
			return false
		}
		if !sameType(mem1, mem2) {
			return false
		}
	}
	return true
}

// Adds type to raw types, make sure name is unique
func (a *ApiAnalyzer) addType(name string, praxisType map[string]interface{}) (string, error) {
	var existing, ok = a.rawTypes[name]
	if ok {
		var base = name
		var idx = 2
		for ok {
			if sameType(praxisType, existing) {
				break
			}
			name = base + strconv.Itoa(idx)
			existing, ok = a.rawTypes[name]
		}
	}
	a.rawTypes[name] = praxisType
	return name, nil
}

// Convert praxis type to go type
func (a *ApiAnalyzer) toType(name string, praxisType map[string]interface{}) (string, error) {
	tn, ok := praxisType["name"].(string)
	if !ok {
		return "", fmt.Errorf("Missing attribute type name")
	}
	switch tn {
	case "String", "Symbol":
		return "string", nil
	case "Integer":
		return "int", nil
	case "Boolean":
		return "bool", nil
	case "Struct":
		return a.addType(name, praxisType)
	case "Collection":
		member, ok := praxisType["member_attribute"].(map[string]interface{})
		if !ok {
			return "", fmt.Errorf("Missing collection \"member_attribute\" key.")
		}
		at, ok := member["type"].(map[string]interface{})
		if !ok {
			return "", fmt.Errorf("Missing collection \"member_attribute\" type.")
		}
		var res, err = a.toType("[]"+name, at)
		if err != nil {
			return "", fmt.Errorf("Failed to compute type of \"member_attribute\": ", err.Error())
		}
		return "[]" + res, nil
	default:
		var elems = strings.Split(tn, "::")
		return strings.Join(elems[2:len(elems)-1], ""), nil
	}
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
