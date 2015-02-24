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

// The analyzer struct holds the analysis results
type ApiAnalyzer struct {
	// Raw resources as defined in API json metadata
	rawResources map[string]map[string]interface{}
	// Raw types as defined in API json metadata
	rawTypes map[string]map[string]interface{}
}

// Factory method for API analyzer
func NewApiAnalyzer(resources, types map[string]map[string]interface{}) *ApiAnalyzer {
	return &ApiAnalyzer{
		rawResources: resources,
		rawTypes:     types,
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

// AnalyzeResource analyzes the given resource and updates the Resources and ParamTypes analyzer
// fields accordingly
func (a *ApiAnalyzer) AnalyzeResource(name string, resource map[string]interface{},
	descriptor *gen.ApiDescriptor) error {

	// Compute description
	var description string
	if d, ok := resource["description"].(string); ok {
		description = d
	}

	// Compute attributes
	var attributes, err = a.AnalyzeAttributes(name, resource)
	if err != nil {
		return err
	}

	// Compute actions
	var actions, err2 = a.AnalyzeActions(name, resource)
	if err2 != nil {
		return err2
	}

	// We're done!
	name = inflect.Singularize(name)
	descriptor.Resources[name] = &gen.Resource{
		Name:        name,
		Description: removeBlankLines(description),
		Actions:     actions,
		Attributes:  attributes,
	}
	return nil
}

// Parse out resource attributes
func (a *ApiAnalyzer) AnalyzeAttributes(resourceName string, resource map[string]interface{}) ([]*gen.Attribute, error) {
	m, ok := resource["media_type"].(string)
	if !ok {
		return nil, fmt.Errorf("Resource %s has no media type (missing key \"media_type\")", resourceName)
	}
	t, ok := a.rawTypes[m]
	if !ok {
		return nil, fmt.Errorf("Missing definition for Resource %s media type %s", resourceName, m)
	}
	attrs, ok := t["attributes"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("Missing attributes for media type %s", m)
	}
	var attributes = make([]*gen.Attribute, len(attrs))
	for idx, n := range sortedKeys(attrs) {
		var attr = attrs[n].(map[string]interface{})
		at, ok := attr["type"].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("Missing attribute type for attribute %s of media type %s", n, m)
		}
		var tn, err = a.toTypeName(n, at)
		if err != nil {
			return nil, fmt.Errorf("Failed to compute type for attribute %s of media type %s: %s", n, m, err.Error())
		}
		attributes[idx] = &gen.Attribute{n, inflect.Camelize(n), tn}
	}
	return attributes, nil
}

// Parse out resource actions
// Resource actions in the JSON consist of an array of map. The map keys are:
// "description", "name", "metadata", "urls", "headers", "params", "payload" and "responses".
func (a *ApiAnalyzer) AnalyzeActions(resourceName string, resource map[string]interface{}) ([]*gen.Action, error) {
	var methods = resource["actions"].(map[string]interface{})
	var actionNames = sortedKeys(methods)
	var actions = []*gen.Action{}
	for _, actionName := range actionNames {
		var m = methods[actionName]
		var meth = m.(map[string]interface{})
		var description = fmt.Sprintf("No description provided for %s.", actionName)
		if d, _ := meth["description"]; d != nil {
			description = d.(string)
		}
		var pathPatterns, err = a.ParseUrls(meth["urls"])
		if err != nil {
			return nil, err
		}
		var params = []*gen.ActionParam{} // Query, path and payload params
		var pathParamNames = []string{}
		var queryParamNames = []string{}
		var payloadParamNames = []string{}

		// Query and path params analysis
		if p, ok := meth["params"].(map[string]interface{}); ok {
			// The "params" hash is the same as a object type description, leverage
			// that to generate an action param so we can pluck out the type fields.
			var obj, err = a.toParam("params", "params", p)
			if err != nil {
				return nil, err
			}
			o, ok := obj.Type.(*gen.ObjectDataType)
			if !ok {
				return nil, fmt.Errorf("Invalid params type %v", obj)
			}
			for _, f := range o.Fields {
				var fn = f.Name
				var isPathParam = false
				for _, pat := range pathPatterns {
					for _, v := range pat.Variables {
						if v == fn {
							isPathParam = true
							break
						}
					}
					if isPathParam {
						break
					}
				}
				if isPathParam {
					pathParamNames = append(pathParamNames, fn)
				} else {
					queryParamNames = append(queryParamNames, fn)
				}
				params = append(params, f)
			}
		}

		// Initialize leaf params, all path and query params are leaf params
		var leafParams = make([]*gen.ActionParam, len(params))
		for i, p := range params {
			leafParams[i] = p
		}

		// Payload params analysis
		if p, ok := meth["payload"].(map[string]interface{}); ok {
			var obj, err = a.toParam("payload", "payload", p)
			if err != nil {
				return nil, err
			}
			o, ok := obj.Type.(*gen.ObjectDataType)
			if !ok {
				payloadParamNames = []string{"payload"}
				params = append(params, obj)
				leafParams = append(leafParams, extractLeafParams(obj)...)
			} else {
				for _, f := range o.Fields {
					payloadParamNames = append(payloadParamNames, f.Name)
					params = append(params, f)
					leafParams = append(leafParams, extractLeafParams(f)...)
				}
			}
		}

		// Heuristic to check whether response returns a location header
		var hasLocation = false
		var returnTypeName string
		responses, ok := meth["responses"]
		if ok {
			var resp = responses.(map[string]interface{})
			if headers, ok := resp["headers"]; ok {
				var head = headers.(map[string]interface{})
				keys, ok := head["keys"]
				if ok {
					var headerKeys = keys.(map[string]interface{})
					for _, k := range headerKeys {
						if k == "Location" {
							hasLocation = true
						}
						break
					}
				}
				if hasLocation {
					break
				}
			}
			if status, ok := resp["status"]; ok {
				var s = status.(int)
				if s > 199 && s < 300 {
					if media, ok := resp["media_type"]; ok {
						var m = media.(map[string]interface{})
						if name, ok := m["name"]; ok {
							returnTypeName = toGoTypeName(name.(string))
						}
					} else if mime, ok := resp["mime_type"]; ok {
						if mime == "controller_defined" {
							returnTypeName = toGoTypeName(resource["media_type"].(string))
						}
					}
				}
			}
		}

		// Record action
		var action = gen.Action{
			Name:              actionName,
			MethodName:        inflect.Camelize(actionName),
			Description:       removeBlankLines(description),
			ResourceName:      resourceName,
			PathPatterns:      pathPatterns,
			Params:            params,
			LeafParams:        leafParams,
			Return:            returnTypeName,
			ReturnLocation:    hasLocation,
			QueryParamNames:   queryParamNames,
			PayloadParamNames: payloadParamNames,
		}
		actions = append(actions, &action)
	}
	return actions, nil
}

// Regular expression that captures variables in a path
var pathVariablesRegexp = regexp.MustCompile(`/:([^/]+)`)

// Extract path patterns from action urls
// Urls consist of an array of map, each map has the following keys:
// "verb", "path", "version"
// Also make it support the resticle format: array of pairs of verb + path.
func (a *ApiAnalyzer) ParseUrls(urls interface{}) ([]*gen.PathPattern, error) {
	var patterns []*gen.PathPattern
	if urlPairs, ok := urls.([][]string); ok {
		patterns = make([]*gen.PathPattern, len(urlPairs))
		for i, pair := range urlPairs {
			if len(pair) != 2 {
				return nil, fmt.Errorf("Invalid URL pair %v, must be [verb, path]", pair)
			}
			patterns[i] = toPattern(pair[0], pair[1])
		}
	} else if urlMaps, ok := urls.([]map[string]string); ok {
		patterns = make([]*gen.PathPattern, len(urlMaps))
		for i, url := range urlMaps {
			var verb, ok = url["verb"]
			if !ok {
				return nil, fmt.Errorf("Missing verb in url %v", url)
			}
			var path, ok2 = url["path"]
			if !ok2 {
				return nil, fmt.Errorf("Missing path in url %v", url)
			}
			patterns[i] = toPattern(verb, path)
		}
	} else {
		return nil, fmt.Errorf("Invalid \"urls\" format %v", urls)
	}
	return patterns, nil
}

// Create path pattern from HTTP verb and request path
func toPattern(verb, path string) *gen.PathPattern {
	var pattern = gen.PathPattern{
		HttpMethod: verb,
		Path:       path,
		Pattern:    pathVariablesRegexp.ReplaceAllLiteralString(path, "/%s"),
		Regexp: pathVariablesRegexp.ReplaceAllLiteralString(regexp.QuoteMeta(path),
			`/([^/]+)`),
	}
	var matches = pathVariablesRegexp.FindAllStringSubmatch(path, -1)
	if len(matches) > 0 {
		pattern.Variables = make([]string, len(matches))
		for i, m := range matches {
			pattern.Variables[i] = m[1]
		}
	}
	return &pattern
}

// Convert raw (metadata) type definition into ActionParam
func (a *ApiAnalyzer) toParam(name, query string, def map[string]interface{}) (*gen.ActionParam, error) {
	var t, ok = def["type"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("Missing type definition in %v", def)
	}
	var description = fmt.Sprintf("No description provided for %s", name)
	if d, ok := def["description"]; ok {
		description = d.(string)
	}
	var param = gen.ActionParam{
		Name:        name,
		QueryName:   query,
		Description: description,
		VarName:     toVarName(name),
	}
	if r, ok := def["required"]; ok {
		if r.(bool) {
			param.Mandatory = true
		}
	}
	if options, ok := def["options"]; ok {
		opts, ok := options.(map[string]interface{})
		if ok {
			for n, o := range opts {
				switch n {
				case "max":
					param.Max = o.(int)
				case "min":
					param.Min = o.(int)
				case "regexp":
					param.Regexp = o.(string)
				}
			}
		}
	}
	var n, ok2 = t["name"].(string)
	if !ok2 {
		return nil, fmt.Errorf("Missing type name in %v", def)
	}
	switch n {
	case "Integer":
		var i = gen.BasicDataType("int")
		param.Type = &i
	case "String":
		var s = gen.BasicDataType("string")
		param.Type = &s
	case "Boolean":
		var b = gen.BasicDataType("bool")
		param.Type = &b
	case "Collection", "Ids":
		member, ok := t["member_attribute"].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("Missing \"member_attribute\" for %v", t)
		}
		var elemType, err = a.toParam(name, "[]"+query, member)
		if err != nil {
			return nil, fmt.Errorf("Failed to compute type of \"member_attribute\": %s", err.Error())
		}
		param.Type = &gen.ArrayDataType{elemType}
	case "Struct":
		attrs, ok := t["attributes"].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("Failed to retrieve attributes of struct for %v", def)
		}
		var obj = gen.ObjectDataType{Name: n}
		for an, at := range attrs {
			a, err := a.toParam(an, fmt.Sprintf("%s[%s]", query, an), at.(map[string]interface{}))
			if err != nil {
				return nil, fmt.Errorf("Failed to compute type of attribute %s: %s", an, err.Error())
			}
			obj.Fields = append(obj.Fields, a)
		}
		param.Type = &obj
	case "Hash":
		keys, ok := t["keys"].(map[string]map[string]interface{})
		if !ok {
			param.Type = new(gen.EnumerableDataType)
		} else {
			var obj = gen.ObjectDataType{Name: n}
			for kn, k := range keys {
				kt, err := a.toParam(kn, fmt.Sprintf("%s[%s]", query, kn), k)
				if err != nil {
					return nil, fmt.Errorf("Failed to compute type of key %s of hash for %v", kn, def)
				}
				obj.Fields = append(obj.Fields, kt)
			}
			param.Type = &obj
		}
	default:
		t, ok := a.rawTypes[n]
		if !ok {
			return nil, fmt.Errorf("Unknown type %s for %v", n, def)
		}
		return a.toParam(name, query, t)
	}
	return &param, nil
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
func (a *ApiAnalyzer) toTypeName(name string, praxisType map[string]interface{}) (string, error) {
	tn, ok := praxisType["name"].(string)
	if !ok {
		return "", fmt.Errorf("Missing attribute type name")
	}
	switch tn {
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
		var res, err = a.toTypeName("[]"+name, at)
		if err != nil {
			return "", fmt.Errorf("Failed to compute type of \"member_attribute\": %s", err.Error())
		}
		return "[]" + res, nil
	default:
		return toGoTypeName(tn), nil
	}
}

// Produce go type name from given ruby type name
func toGoTypeName(name string) string {
	switch name {
	case "String", "Symbol":
		return "string"
	case "Integer":
		return "int"
	case "Boolean":
		return "bool"
	case "Struct", "Collection":
		panic("Uh oh, trying to infer a go type name for a unnamed struct or collection (" + name + ")")
	default:
		var elems = strings.Split(name, "::")
		return "*" + strings.Join(elems[2:len(elems)-1], "")
	}

}

// Capture all alphanumerical parts to build go identifier from raw param name
var partsRegexp = regexp.MustCompile("[^[:alnum:]]+")

// Parse native names into go parameter names
func toVarName(name string) string {
	if name == "type" {
		return "type_"
	}
	p := partsRegexp.ReplaceAllString(name, "_")
	return inflect.CamelizeDownFirst(p)
}

// Extract leaf parameters from given action param
func extractLeafParams(a *gen.ActionParam) []*gen.ActionParam {
	switch t := a.Type.(type) {
	case *gen.BasicDataType, *gen.EnumerableDataType:
		return []*gen.ActionParam{a}
	case *gen.ArrayDataType:
		return extractLeafParams(t.ElemType)
	case *gen.ObjectDataType:
		var params = []*gen.ActionParam{}
		for _, f := range t.Fields {
			params = append(params, extractLeafParams(f)...)
		}
		return params
	}
	return nil // not reached
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

// Check whether string only contains blank characters
var blankRegexp = regexp.MustCompile(`^\s*$`)

// Helper method that removes blank lines from strings
func removeBlankLines(doc string) string {
	lines := strings.Split(doc, "\n")
	fullLines := make([]string, len(lines))
	i := 0
	for _, line := range lines {
		if len(line) > 0 && !blankRegexp.MatchString(line) {
			fullLines[i] = line
			i += 1
		}
	}
	return strings.Join(fullLines[:i], "\n")
}
