package main

import (
	"fmt"
	"sort"
	"strings"

	"bitbucket.org/pkg/inflect"
)

var (
	// Resources that don't have a media type
	NoMediaTypeResources = map[string]bool{
		"HealthCheck":          true,
		"Oauth2":               true,
		"Tag":                  true,
		"UserDatas":            true,
		"ChildAccounts":        true,
		"MonitoringMetricData": true,
		"ImportPreview":        true,
		"Changes":              true,
		"CookbookResolution":   true,
		"ResourceTag":          true,
	}
)

// The analyzer struct holds the analysis results
type ApiAnalyzer struct {
	// Raw resources as defined in API json metadata
	rawResources map[string]interface{}

	// Attribute type mappings defined in attributes.json
	attributeTypes map[string]string

	/* Analysis results */

	// Resources indexed by name
	Resources map[string]*ResourceData

	// Types used by resource actions indexed by name
	Types map[string]*ObjectDataType
}

// Factory method for API analyzer
func NewApiAnalyzer(resources map[string]interface{}, attributes map[string]string) *ApiAnalyzer {
	return &ApiAnalyzer{
		rawResources:   resources,
		attributeTypes: attributes,
		Resources:      make(map[string]*ResourceData),
		Types:          make(map[string]*ObjectDataType),
	}
}

// Analyze iterate through all resources and initializes the Resources and ParamTypes fields of
// the ApiAnalyzer struct accordingly.
func (a *ApiAnalyzer) Analyze() {
	for name, resource := range a.rawResources {
		a.AnalyzeResource(name, resource)
	}
}

// AnalyzeResource analyzes the given resource and updates the Resources and ParamTypes analyzer
// fields accordingly
func (a *ApiAnalyzer) AnalyzeResource(name string, resource interface{}) {
	res := resource.(map[string]interface{})

	// Compute description
	var description string
	if d, ok := res["description"].(string); ok {
		description = d
	}

	// Compute attributes
	var attributes []*ResourceAttribute
	var atts map[string]interface{}
	if m, ok := res["media_type"].(map[string]interface{}); ok {
		atts = m["attributes"].(map[string]interface{})
		attributes = make([]*ResourceAttribute, len(atts))
		for idx, n := range sortedKeys(atts) {
			attributes[idx] = &ResourceAttribute{attributeName(n), n, a.attributeTypes[n]}
		}
	} else {
		attributes = []*ResourceAttribute{}
	}

	// Compute actions
	methods := res["methods"].(map[string]interface{})
	actionNames := sortedKeys(methods)
	actions := make([]*ResourceAction, len(methods))
	idx := 0
	for _, actionName := range actionNames {
		m := methods[actionName]
		meth := m.(map[string]interface{})
		var params map[string]interface{}
		if p, ok := meth["parameters"]; ok {
			params = p.(map[string]interface{})
		}
		description := "<no description>"
		if d, _ := meth["description"]; d != nil {
			description = d.(string)
		}
		httpMethod, path := parseRoute(fmt.Sprintf("%s#%s", name, actionName),
			meth["route"].(string))
		var contentType string
		if c, ok := meth["content_type"].(string); ok {
			contentType = c
		}
		paramAnalyzer := NewAnalyzer(path, params)
		paramAnalyzer.Analyze()

		// Record new parameter types
		for name, pType := range paramAnalyzer.ParamTypes {
			if existing, ok := a.Types[name]; ok {
				if !existing.Compare(pType) {
					panic(fmt.Sprintf("SHOOT \n%+v\nvs.\n%+v\n", existing, pType))
				}
			} else {
				a.Types[name] = pType
			}
		}

		// Record action
		actions[idx] = &ResourceAction{
			Name:          methodName(actionName, name),
			Description:   description,
			HttpMethod:    httpMethod,
			Path:          path,
			NativeParams:  params,
			AllParams:     paramAnalyzer.Params,
			PathParams:    paramAnalyzer.PathParams,
			QueryParams:   paramAnalyzer.QueryParams,
			PayloadParams: paramAnalyzer.PayloadParams,
			UrlExp:        paramAnalyzer.UrlExp,
			Return:        parseReturn(actionName, name, contentType),
		}
		idx += 1
	}

	// We're done!
	name = inflect.Singularize(name)
	a.Resources[name] = &ResourceData{
		Name:        name,
		Description: description,
		Actions:     actions,
		Attributes:  attributes,
	}
}

// Categorize parameters between path, query and payload.
func categorizeParams(params map[string]*ActionParam, pathParamNames []string) ([]*ActionParam, []*ActionParam, []*ActionParam, []*ActionParam) {
	allParams := make([]*ActionParam, len(params))
	idx := 0
	for _, param := range params {
		allParams[idx] = param
		idx += 1
	}
	sort.Sort(ByName(allParams))
	pathParams := []*ActionParam{}
	queryParams := []*ActionParam{}
	payloadParams := []*ActionParam{}
	for _, param := range allParams {
		pname := param.Name
		if isQueryParam(pname) {
			queryParams = append(queryParams, param)
		} else {
			isPathParam := false
			for _, p := range pathParamNames {
				if p == pname {
					isPathParam = true
					break
				}
			}
			if isPathParam {
				pathParams = append(pathParams, param)
			} else {
				payloadParams = append(payloadParams, param)
			}
		}
	}

	return allParams, pathParams, queryParams, payloadParams
}

// Heuristic to determine whether given param is a query string param
// For now only consider view and filter...
func isQueryParam(n string) bool {
	return n == "view" || n == "filter"
}

func parseRoute(moniker string, route string) (string, string) {
	// :(((( some routes are empty
	switch moniker {
	case "Deployments#servers":
		return "GET", "api/deployments/:id/servers"
	case "ServerArrays#current_instances":
		return "GET", "api/server_arrays/:id/current_instances"
	case "ServerArrays#launch":
		return "POST", "api/server_arrays/:id/launch"
	case "ServerArrays#multi_run_executable":
		return "POST", "api/server_arrays/:id/multi_run_executable"
	case "ServerArrays#multi_terminate":
		return "POST", "api/server_arrays/:id/multi_terminate"
	case "Servers#launch":
		return "POST", "api/servers/:id/launch"
	case "Servers#terminate":
		return "POST", "api/servers/:id/teminate"

	}
	method := strings.TrimRight(route[0:7], " ")
	path := strings.TrimRight(strings.Split(route[8:], "{")[0], " ")

	return method, path
}

func parseUrl(path string) (string, []string) {
	if strings.HasSuffix(path, "(.:format)?") {
		path = path[:len(path)-11]
	}
	elems := strings.Split(path, "/")
	urlElems := make([]string, len(elems))
	params := []string{}
	j := 0
	acc := ""
	for i, e := range elems {
		if strings.HasPrefix(e, ":") {
			prefix := acc
			acc = ""
			if len(prefix) > 0 {
				prefix = fmt.Sprintf("\"/%s/\"", prefix)
			}
			if i > 0 || len(prefix) > 0 {
				prefix += "+"
			}
			suffix := ""
			if i < len(elems)-1 {
				suffix = "+"
			}
			p := parseParamName(e[1:])
			params = append(params, e[1:])
			urlElems[j] = fmt.Sprintf("%s%s%s", prefix, p, suffix)
			j += 1
		} else {
			if len(acc) > 0 {
				acc = fmt.Sprintf("%s%s%s", acc, "/", e)
			} else {
				acc = e
			}
		}
	}
	if len(acc) > 0 {
		urlElems[j] = fmt.Sprintf("\"/%s\"", acc)
		j += 1
	}
	urlElems = urlElems[:j]
	return strings.Join(urlElems, ""), params
}

func parseReturn(kind, resName, contentType string) string {
	switch kind {
	case "show":
		return fmt.Sprintf("*%s", resourceType(resName))
	case "index":
		return fmt.Sprintf("[]%s", resourceType(resName))
	case "create":
		if _, ok := NoMediaTypeResources[resName]; ok {
			return "map[string]interface{}"
		} else {
			return "Href"
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
			elems := strings.SplitN(contentType[27:], ";", 2)
			name := resourceType(inflect.Camelize(elems[0]))
			if len(elems) > 1 && elems[1] == "type=collection" {
				name = "[]" + name
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
	if _, ok := NoMediaTypeResources[resName]; ok {
		return "map[string]string"
	} else {
		return inflect.Singularize(resName)
	}
}

// Heuristic to create method name from action and resource names
func methodName(action, resource string) string {
	if action != "index" && strings.Index(action, "multi") != 0 {
		resource = inflect.Singularize(resource)
	}
	return inflect.Camelize(action) + inflect.Camelize(resource)
}

// Escape attribute names using go keywords
func attributeName(name string) string {
	return inflect.Camelize(name)
}

// Return keys of given maps sorted
func sortedKeys(m map[string]interface{}) []string {
	keys := make([]string, len(m))
	idx := 0
	for k, _ := range m {
		keys[idx] = k
		idx += 1
	}
	sort.Strings(keys)
	return keys
}

// Sort array of string by length
type ByReverseLength []string

func (s ByReverseLength) Len() int {
	return len(s)
}
func (s ByReverseLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByReverseLength) Less(i, j int) bool {
	return len(s[i]) > len(s[j])
}
