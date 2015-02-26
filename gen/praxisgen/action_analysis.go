package main

import (
	"fmt"
	"regexp"

	"bitbucket.org/pkg/inflect"

	"github.com/rightscale/rsc/gen"
)

// Regular expression that captures variables in a path
var pathVariablesRegexp = regexp.MustCompile(`/:([^/]+)`)

// Parse out resource actions
// Resource actions in the JSON consist of an array of map. The map keys are:
// "description", "name", "metadata", "urls", "headers", "params", "payload" and "responses".
func (a *ApiAnalyzer) AnalyzeActions(resourceName string, resource map[string]interface{}) ([]*gen.Action, error) {
	var methods = resource["actions"].([]interface{})
	var actions = make([]*gen.Action, len(methods))
	for i, m := range methods {
		var meth = m.(map[string]interface{})
		var actionName = meth["name"].(string)
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
		if p, ok := meth["params"]; ok {
			var param = p.(map[string]interface{})
			t, ok := param["type"]
			if !ok {
				return nil, fmt.Errorf("Misins type declaration in %s", prettify(p))
			}
			var attrs = t.(map[string]interface{})["attributes"].(map[string]interface{})
			for pn, pt := range attrs {
				att, err := a.AnalyzeAttribute(pn, pn, pt.(map[string]interface{}))
				if err != nil {
					return nil, fmt.Errorf("Failed to compute type of param %s: %s", pn, err.Error())
				}
				var isPathParam = false
				for _, pat := range pathPatterns {
					for _, v := range pat.Variables {
						if v == pn {
							isPathParam = true
							break
						}
					}
					if isPathParam {
						break
					}
				}
				if isPathParam {
					pathParamNames = append(pathParamNames, pn)
				} else {
					queryParamNames = append(queryParamNames, pn)
				}
				params = append(params, att)
			}
		}

		// Initialize leaf params, all path and query params are leaf params
		var leafParams = make([]*gen.ActionParam, len(params))
		for i, p := range params {
			leafParams[i] = p
		}

		// Payload params analysis
		if p, ok := meth["payload"]; ok {
			var attrs = p.(map[string]interface{})["attributes"].(map[string]interface{})
			for pn, pt := range attrs {
				att, err := a.AnalyzeAttribute(pn, pn, pt.(map[string]interface{}))
				if err != nil {
					return nil, fmt.Errorf("Failed to compute type of param %s: %s", pn, err.Error())
				}
				payloadParamNames = append(payloadParamNames, pn)
				params = append(params, att)
				leafParams = append(leafParams, extractLeafParams(att)...)
			}
		}

		// Heuristic to check whether response returns a location header
		var hasLocation = false
		var returnTypeName string
		responses, ok := meth["responses"]
		if ok {
			var resps = responses.(map[string]interface{})
			for _, r := range resps {
				var resp = r.(map[string]interface{})
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
					var s = int(status.(float64))
					if s > 199 && s < 300 {
						if media, ok := resp["media_type"]; ok {
							var m = media.(map[string]interface{})
							if name, ok := m["name"]; ok {
								returnTypeName = toGoTypeName(name.(string), true)
							}
						} else if mime, ok := resp["mime_type"]; ok {
							if mime == "controller_defined" {
								returnTypeName = toGoTypeName(resource["media_type"].(string), true)
							}
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
		actions[i] = &action
	}
	return actions, nil
}

// Extract path patterns from action urls
// Urls consist of an array of map, each map has the following keys:
// "verb", "path", "version"
// Also make it support the resticle format: array of pairs of verb + path.
func (a *ApiAnalyzer) ParseUrls(urls interface{}) ([]*gen.PathPattern, error) {
	var patterns []*gen.PathPattern
	if urlElems, ok := urls.([]interface{}); ok {
		patterns = make([]*gen.PathPattern, len(urlElems))
		for i, elem := range urlElems {
			if pair, ok := elem.([]string); ok {
				if len(pair) != 2 {
					return nil, fmt.Errorf("Invalid URL pair %v, must be [verb, path]", pair)
				}
				patterns[i] = toPattern(pair[0], pair[1])
			} else if url, ok := elem.(map[string]interface{}); ok {
				var verb, ok = url["verb"].(string)
				if !ok {
					return nil, fmt.Errorf("Missing verb in url %v", url)
				}
				var path, ok2 = url["path"].(string)
				if !ok2 {
					return nil, fmt.Errorf("Missing path in url %v", url)
				}
				patterns[i] = toPattern(verb, path)
			} else {
				return nil, fmt.Errorf("Invalid url format %v", elem)
			}
		}
	} else {
		return nil, fmt.Errorf("Invalid \"urls\" format %v", prettify(urls))
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
