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
			var attrNames = sortedKeys(attrs)
			for _, pn := range attrNames {
				var pt = attrs[pn]
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
					att.Location = gen.PathParam
				} else {
					queryParamNames = append(queryParamNames, pn)
					att.Location = gen.QueryParam
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
		var paramNames = make([]string, len(pathParamNames)+len(queryParamNames))
		for i, n := range pathParamNames {
			paramNames[i] = n
		}
		for i, n := range queryParamNames {
			paramNames[len(pathParamNames)+i] = n
		}
		if p, ok := meth["payload"]; ok {
			as, ok := p.(map[string]interface{})["attributes"]
			if ok {
				var attrs = as.(map[string]interface{})
				var attrNames = sortedKeys(attrs)
				for _, pn := range attrNames {
					var pt = attrs[pn]
					var queryName = gen.MakeUniq(pn, paramNames)
					att, err := a.AnalyzeAttribute(pn, queryName, pt.(map[string]interface{}))
					if err != nil {
						return nil, fmt.Errorf("Failed to compute type of param %s: %s", pn, err.Error())
					}
					payloadParamNames = append(payloadParamNames, pn)
					att.Location = gen.PayloadParam
					params = append(params, att)
					var extracted = extractLeafParams(att)
					for _, e := range extracted {
						e.Location = gen.PayloadParam
					}
					leafParams = append(leafParams, extracted...)
				}
			}
		}

		// Heuristic to check whether response returns a location header
		// Also extract response type from success response media type
		// TBD: support actions returning multiple success responses with media types?
		var hasLocation = false
		var returnTypeName string
		responses, ok := meth["responses"]
		if ok {
			var resps = responses.(map[string]interface{})
			var respNames = sortedKeys(resps)
			for _, rName := range respNames {
				var r = resps[rName]
				var resp = r.(map[string]interface{})
				status := resp["status"]
				var s = int(status.(float64))
				if s < 200 || s > 299 {
					continue // Skip error responses
				}
				if headers, ok := resp["headers"]; ok {
					if hname, ok := headers.(string); ok {
						// TBD is there a better way?
						hasLocation = hname == "Location" && actionName == "create"
					} else {
						var head = headers.(map[string]interface{})
						keys, ok := head["keys"]
						if ok {
							var headerKeys = keys.(map[string]interface{})
							for _, k := range headerKeys {
								// TBD is there a better way?
								if k == "Location" && actionName == "create" {
									hasLocation = true
								}
								break
							}
						}
					}
					if hasLocation {
						returnTypeName = fmt.Sprintf("*%sLocator", resourceName)
					}
				}
				if returnTypeName == "" {
					if media, ok := resp["media_type"]; ok {
						var m = media.(map[string]interface{})
						if name, ok := m["name"]; ok {
							returnTypeName = toGoTypeName(name.(string), true)
						}
					} else if mime, ok := resp["mime_type"]; ok {
						// Resticle compat
						if mime == "controller_defined" {
							returnTypeName = toGoTypeName(resource["media_type"].(string), true)
						} else {
							for n, r := range a.RawResources {
								if mt, ok := r["mime_type"]; ok {
									if mt == mime {
										if actionName == "index" {
											returnTypeName = "[]" + n
										} else {
											returnTypeName = n
										}
										break
									}
								}
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
			if pair, ok := elem.([]interface{}); ok {
				if len(pair) != 2 {
					return nil, fmt.Errorf("Invalid URL pair %v, must be [verb, path]", pair)
				}
				patterns[i] = toPattern(pair[0].(string), pair[1].(string))
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
				return nil, fmt.Errorf("Invalid url format %#v", elem)
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
