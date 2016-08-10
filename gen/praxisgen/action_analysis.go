package main

import (
	"fmt"
	"regexp"

	"bitbucket.org/pkg/inflect"

	"github.com/rightscale/rsc/gen"
)

// Regular expression that captures variables in a path
var pathVariablesRegexp = regexp.MustCompile(`/:([^/]+)`)

// AnalyzeActions parses out a resource actions.
// Resource actions in the JSON consist of an array of map. The map keys are:
// "description", "name", "metadata", "urls", "headers", "params", "payload" and "responses".
func (a *APIAnalyzer) AnalyzeActions(resourceName string, resource map[string]interface{}) ([]*gen.Action, error) {
	methods := resource["actions"].([]interface{})
	actions := make([]*gen.Action, len(methods))
	for i, m := range methods {
		meth := m.(map[string]interface{})
		actionName := meth["name"].(string)
		description := fmt.Sprintf("No description provided for %s.", actionName)
		if d, _ := meth["description"]; d != nil {
			description = d.(string)
		}
		pathPatterns, err := a.ParseUrls(meth["urls"])
		if err != nil {
			return nil, err
		}
		params := []*gen.ActionParam{} // Query, path and payload params
		pathParamNames := []string{}
		queryParamNames := []string{}
		payloadParamNames := []string{}

		// Query and path params analysis
		if p, ok := meth["params"]; ok {
			param := p.(map[string]interface{})
			t, ok := param["type"]
			if !ok {
				return nil, fmt.Errorf("Missing type declaration in %s", prettify(p))
			}
			attrs := t.(map[string]interface{})["attributes"].(map[string]interface{})
			attrNames := sortedKeys(attrs)
			for _, pn := range attrNames {
				pt := attrs[pn]
				att, err := a.AnalyzeAttribute(pn, pn, pt.(map[string]interface{}))
				if err != nil {
					return nil, fmt.Errorf("Failed to compute type of param %s: %s", pn, err.Error())
				}
				isPathParam := false
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
					params = append(params, att)
				}
			}
		}

		// Initialize leaf params, all path and query params are leaf params
		leafParams := make([]*gen.ActionParam, len(queryParamNames))
		idx := 0
		for _, p := range params {
			if p.Location == gen.QueryParam {
				leafParams[idx] = p
				idx++
			}
		}
		paramNames := make([]string, len(queryParamNames))
		for i, n := range queryParamNames {
			paramNames[i] = n
		}

		// Payload params analysis
		var payload gen.DataType
		if p, ok := meth["payload"]; ok {
			as, ok := p.(map[string]interface{})["type"]
			if ok {
				pd, err := a.AnalyzeType(as.(map[string]interface{}), "payload")
				if err != nil {
					return nil, err
				}
				if po, ok := pd.(*gen.ObjectDataType); ok {

					// Remove the type since we are "flattening" the attributes
					// as top level params.
					// This is a bit hacky and should be refactored
					// (it should be possible to get the type without having
					// it be registered). Works for now(TM).
					delete(a.Registry.InlineTypes, po.TypeName)

					for _, att := range po.Fields {
						payloadParamNames = append(payloadParamNames, att.Name)
						att.Location = gen.PayloadParam
						params = append(params, att)
						extracted := extractLeafParams(att, att.Name, make(map[string]*[]*gen.ActionParam), !att.Mandatory)
						for _, e := range extracted {
							e.Location = gen.PayloadParam
						}
						leafParams = append(leafParams, extracted...)
					}
				} else {
					// Raw payload (no attributes)
					payload = pd
					param := rawPayload(pd, p)
					params = append(params, param)
					leafParams = append(leafParams, param)
				}
			}
		}

		// Heuristic to check whether response returns a location header
		// Also extract response type from success response media type
		// TBD: support actions returning multiple success responses with media types?
		hasLocation := false
		var returnTypeName string
		responses, ok := meth["responses"]
		if ok {
			resps := responses.(map[string]interface{})
			respNames := sortedKeys(resps)
			for _, rName := range respNames {
				r := resps[rName]
				resp := r.(map[string]interface{})
				status := resp["status"]
				s := int(status.(float64))
				if s < 200 || s > 299 {
					continue // Skip error responses
				}
				if s == 201 && actionName == "create" {
					hasLocation = true
				} else if headers, ok := resp["headers"]; ok {
					if hname, ok := headers.(string); ok {
						// TBD is there a better way?
						hasLocation = hname == "Location" && actionName == "create"
					} else {
						head := headers.(map[string]interface{})
						keys, ok := head["keys"]
						if ok {
							headerKeys := keys.(map[string]interface{})
							for _, k := range headerKeys {
								// TBD is there a better way?
								if k == "Location" && actionName == "create" {
									hasLocation = true
								}
								break
							}
						}
					}
				}
				if returnTypeName == "" {
					if media, ok := resp["media_type"]; ok {
						m := media.(map[string]interface{})
						if name, ok := m["name"]; ok {
							returnTypeName = toGoReturnTypeName(name.(string), actionName == "index")
							a.descriptor.NeedJSON = true
							// Analyze return type to make sure it gets recorded
							_, err := a.AnalyzeType(a.RawTypes[name.(string)], "return")
							if err != nil {
								return nil, err
							}
						} else {
							// Default to string
							returnTypeName = "string"
						}
					} else if mime, ok := resp["mime_type"]; ok {
						// Resticle compat
						for n, r := range a.RawResources {
							if mt, ok := r["mime_type"]; ok {
								if mt == mime {
									if actionName == "index" {
										returnTypeName = "[]*" + n
									} else {
										returnTypeName = "*" + n
									}
									a.descriptor.NeedJSON = true
									break
								}
							}
						}
					}
				}
			}
		}
		if hasLocation {
			returnTypeName = fmt.Sprintf("*%sLocator", resourceName)
		}

		// Record action
		action := gen.Action{
			Name:              actionName,
			MethodName:        inflect.Camelize(actionName),
			Description:       removeBlankLines(description),
			ResourceName:      resourceName,
			PathPatterns:      pathPatterns,
			Payload:           payload,
			Params:            params,
			LeafParams:        leafParams,
			Return:            returnTypeName,
			ReturnLocation:    hasLocation,
			QueryParamNames:   queryParamNames,
			PayloadParamNames: payloadParamNames,
			PathParamNames:    pathParamNames,
		}
		actions[i] = &action
	}
	return actions, nil
}

// rawPayload is a helper function that creates a ActionParam for a raw (non object) payload
func rawPayload(typ gen.DataType, p interface{}) *gen.ActionParam {
	var required bool
	if req, ok := p.(map[string]interface{})["required"]; ok {
		required = req.(bool)
	}
	return &gen.ActionParam{
		Name:      "payload",
		QueryName: "payload",
		VarName:   "payload",
		Type:      typ,
		Location:  gen.PayloadParam,
		Mandatory: required,
	}
}

// ParseUrls extracts the path patterns from an action URLs.
// Urls consist of an array of map, each map has the following keys:
// "verb", "path", "version"
// Also make it support the resticle format: array of pairs of verb + path.
func (a *APIAnalyzer) ParseUrls(urls interface{}) ([]*gen.PathPattern, error) {
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
				verb, ok := url["verb"].(string)
				if !ok {
					return nil, fmt.Errorf("Missing verb in url %v", url)
				}
				path, ok := url["path"].(string)
				if !ok {
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
	pattern := gen.PathPattern{
		HTTPMethod: verb,
		Path:       path,
		Pattern:    pathVariablesRegexp.ReplaceAllLiteralString(path, "/%s"),
		Regexp: pathVariablesRegexp.ReplaceAllLiteralString(regexp.QuoteMeta(path),
			`/([^/]+)`),
	}
	matches := pathVariablesRegexp.FindAllStringSubmatch(path, -1)
	if len(matches) > 0 {
		pattern.Variables = make([]string, len(matches))
		for i, m := range matches {
			pattern.Variables[i] = m[1]
		}
	}
	return &pattern
}

// Extract leaf parameters from given action param
func extractLeafParams(a *gen.ActionParam, root string, seen map[string]*[]*gen.ActionParam, parentNotMandatory bool) []*gen.ActionParam {
	switch t := a.Type.(type) {
	case *gen.BasicDataType, *gen.EnumerableDataType, *gen.UploadDataType:
		dup := gen.ActionParam{
			Name:        a.Name,
			QueryName:   root,
			Description: a.Description,
			VarName:     a.VarName,
			Location:    a.Location,
			Type:        a.Type,
			Mandatory:   a.Mandatory && !parentNotMandatory, // yay for double negations!
			NonBlank:    a.NonBlank,
			Regexp:      a.Regexp,
			ValidValues: a.ValidValues,
			Min:         a.Min,
			Max:         a.Max,
		}
		return []*gen.ActionParam{&dup}
	case *gen.ArrayDataType:
		p := t.ElemType
		eq, ok := seen[p.Name]
		if !ok {
			eq = &[]*gen.ActionParam{}
			seen[p.Name] = eq
			*eq = extractLeafParams(p, root+"[]", seen, parentNotMandatory || !a.Mandatory)
		}
		return *eq
	case *gen.ObjectDataType:
		params := []*gen.ActionParam{}
		for _, f := range t.Fields {
			eq, ok := seen[f.Name]
			if !ok {
				eq = &[]*gen.ActionParam{}
				seen[f.Name] = eq
				*eq = extractLeafParams(f, fmt.Sprintf("%s[%s]", root, f.Name), seen, parentNotMandatory || !a.Mandatory)
			}
			params = append(params, *eq...)
		}
		return params
	}
	return nil
}
