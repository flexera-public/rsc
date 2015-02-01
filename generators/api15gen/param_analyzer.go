package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"bitbucket.org/pkg/inflect"
)

// Capture root of path
var rootRegexp = regexp.MustCompile("([^\\[]+)\\[")

// Parent path regular expression
var parentPathRegexp = regexp.MustCompile(`^(.*)\[.+\]$`)

// Child path regular expression
var childPathRegexp = regexp.MustCompile(`^.*\[(.+)\]$`)

// Analyzer method parameters
func analyzeParams(path string, params map[string]interface{}, pathParamNames []string) ([]*ActionParam, []*ActionParam, []*ActionParam, []*ActionParam) {
	actionParams := parseParams(path, params)
	queryParams := []*ActionParam{}
	payloadParams := []*ActionParam{}
	for n, p := range actionParams {
		isPathParam := false
		for _, pp := range pathParamNames {
			if pp == n {
				isPathParam = true
				break
			}
		}
		if isPathParam {
			continue
		}
		if isQueryParam(n) {
			queryParams = append(queryParams, p)
		} else {
			payloadParams = append(payloadParams, p)
		}
	}
	pathParams := make([]*ActionParam, len(pathParamNames))
	for i, n := range pathParamNames {
		pathParams[i] = actionParams[n]
	}
	allParams := make([]*ActionParam, len(pathParams)+len(queryParams)+len(payloadParams))
	for i, p := range pathParams {
		allParams[i] = p
	}
	offset := len(pathParams)
	for i, p := range payloadParams {
		allParams[offset+i] = p
	}
	offset += len(payloadParams)
	for i, p := range queryParams {
		allParams[offset+i] = p
	}
	return allParams, pathParams, queryParams, payloadParams
}

// First extract parameters from action path (i.e. ":server_id" in "/servers/:server_id") then
// build a set of parameters from the list of all field names.
// Example parameters declaration:
//  "parameters": {
//      "security_group_rule[description]": {
//          "class": "String",
//          "non_blank": true
//      },
//      "security_group_rule": {
//          "class": "Hash",
//          "non_blank": true,
//          "mandatory": true
//      }
//  }
// Example of tricky cases:
// inputs: { "class": "Enumerable" },
// inputs[*]: { "class": "String" },
// inputs[][name]: { "class": "String" },
// inputs[][value]: { "class": "String" }
// server_array[elasticity_params][schedule][][max_count]: { "class": "String" },
// server_array[elasticity_params][schedule][][min_count]: { "class": "String" },
// server_array[elasticity_params][schedule][][day]: { "class": "String" },
// server_array[elasticity_params][schedule]: { "class": "Array" },
func parseParams(path string, params map[string]interface{}) map[string]*ActionParam {
	// Add path parameters
	elems := strings.Split(path, "/")
	for _, e := range elems {
		if strings.HasPrefix(e, ":") {
			if strings.HasSuffix(e, "(.:format)?") {
				e = e[:len(e)-11]
			}
			params[e[1:]] = map[string]interface{}{"class": "String", "mandatory": true, "non_blank": true}
		}
	}

	// Order keys using their length so "foo[bar]" is analyzed before "foo"
	paths := make([]string, len(params))
	i := 0
	for n, _ := range params {
		paths[i] = n
		i += 1
	}
	sort.Sort(ByReverseLength(paths))

	//
	parsed := map[string]*ActionParam{}
	top := map[string]*ActionParam{}
	for _, path := range paths {
		var child *ActionParam
		origPath := path
		// fmt.Printf("origPath: %s\n", origPath)
		origParam := params[path].(map[string]interface{})
		// fmt.Printf("origParam: %v\n", origParam)
		matches := parentPathRegexp.FindStringSubmatch(path)
		// fmt.Printf("matches: %v\n", matches)
		isTop := (matches == nil)
		for matches != nil {
			// fmt.Printf("\nPROCESSING %s - CHILD %v\n", path, child)
			param := params[path].(map[string]interface{})
			parentPath := matches[1]
			var isArrayChild bool
			if strings.HasSuffix(parentPath, "[]") {
				isArrayChild = true
			}
			// fmt.Printf("isArrayChild: %v\n", isArrayChild)
			if parent, ok := parsed[parentPath]; ok {
				// fmt.Printf("parent: %+v:%v\n", parent, parent.Type.Inspect())
				a, ok := parent.Type.(*ArrayDataType)
				if ok {
					parent = a.ElemType
				}
				dType := parseDataType(params, child, parsed, path)
				child = parseParam(path, param, dType)
				if _, ok = parent.Type.(*EnumerableDataType); !ok {
					o := parent.Type.(*ObjectDataType)
					o.Fields = append(o.Fields, child)
					parsed[path] = child
					// fmt.Printf("parsed[%s] = %+v: %v\n", path, child, child.Type.Inspect())
					break // No need to keep going back, we already have a parent
				}
			} else {
				dType := parseDataType(params, child, parsed, path)
				child = parseParam(path, param, dType)
				// fmt.Printf("no parent - parsed[%s] = %+v: %v\n", path, child, child.Type.Inspect())
				parsed[path] = child
				if isArrayChild {
					typeName := parseParamName(matches[1] + "[item]")
					parent = parseParam(matches[1]+"[item]",
						map[string]interface{}{},
						&ObjectDataType{typeName, []*ActionParam{child}})
					parsed[parentPath] = parent
					// fmt.Printf("parsed[%s] = %+v:%v\n", parentPath, parent, parent.Type.Inspect())
					child = parent
				}
			}
			matches = parentPathRegexp.FindStringSubmatch(parentPath)
		}
		// fmt.Printf("\nDONE PROCESSING CHILDREN CHILD: %v\n", child)
		if isTop {
			if _, ok := parsed[path]; !ok {
				dType := parseDataType(params, nil, parsed, path)
				actionParam := parseParam(path, origParam, dType)
				parsed[path] = actionParam
				// fmt.Printf("parsed[%s] = %+v:%v\n", path, actionParam, actionParam.Type.Inspect())
			}
			top[path] = parsed[path]
		} else {
			// fmt.Printf("PARSED BEFORE FINALIZE: %+v\n", parsed)
			matches := rootRegexp.FindStringSubmatch(origPath)
			rootPath := matches[1]
			// fmt.Printf("\nrootPath: %s\n", rootPath)
			if _, ok := parsed[rootPath]; !ok {
				parsed[rootPath] = parseParam(rootPath, params[rootPath].(map[string]interface{}),
					parseDataType(params, child, parsed, rootPath))
			}
		}
	}

	return top
}

// Parse data type in context
func parseDataType(params map[string]interface{}, child *ActionParam, parsed map[string]*ActionParam,
	path string) DataType {
	// fmt.Printf("PATH: %s\n", path)
	param := params[path].(map[string]interface{})
	class := "String"
	if c, ok := param["class"].(string); ok {
		class = c
	}
	var res DataType
	switch class {
	case "Integer":
		i := BasicDataType("int")
		res = &i
	case "String":
		s := BasicDataType("string")
		res = &s
	case "Array":
		if child != nil {
			res = &ArrayDataType{child}
		} else {
			s := BasicDataType("string")
			p := parseParam(fmt.Sprintf("%s[item]", path),
				map[string]interface{}{}, &s)
			res = &ArrayDataType{p}
		}
	case "Enumerable":
		res = new(EnumerableDataType)
	case "Hash":
		if current, ok := parsed[path]; ok {
			res = current.Type
			o := res.(*ObjectDataType)
			o.Fields = append(o.Fields, child)
		} else {
			res = &ObjectDataType{parseParamName(path), []*ActionParam{child}}
		}
	}
	// fmt.Printf("PARSED DATA TYPE %v\n", res.Inspect())
	return res
}

func parseParam(path string, param map[string]interface{}, dType DataType) *ActionParam {
	var description, regexp string
	var mandatory, nonBlank bool
	var validValues []interface{}
	if d, ok := param["description"]; ok {
		description = d.(string)
	}
	if m, ok := param["mandatory"]; ok {
		mandatory = m.(bool)
	}
	if n, ok := param["non_blank"]; ok {
		nonBlank = n.(bool)
	}
	if r, ok := param["regexp"]; ok {
		regexp = r.(string)
	}
	if v, ok := param["valid_values"]; ok {
		validValues = v.([]interface{})
	}
	native := path
	matches := childPathRegexp.FindStringSubmatch(path)
	if matches != nil {
		native = matches[1]
	}
	return &ActionParam{
		Name:        parseParamName(native),
		Description: description,
		Type:        dType,
		NativeName:  native,
		Mandatory:   mandatory,
		NonBlank:    nonBlank,
		Regexp:      regexp,
		ValidValues: validValues,
	}
}

func parseParamName(name string) string {
	r := regexp.MustCompile("[^[:alnum:]]+")
	p := r.ReplaceAllString(name, "_")
	if p == "r_s_version" {
		return "rsVersion"
	}
	return inflect.CamelizeDownFirst(p)
}
