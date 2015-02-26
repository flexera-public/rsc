package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"regexp"

	"bitbucket.org/pkg/inflect"

	"github.com/rightscale/rsc/gen"
)

// Regular expression used to capture brackets in query name
var bracketRegexp = regexp.MustCompile(`(\[|\])+`)

// Analyze an attribute, create corresponding ActionParam
func (a *ApiAnalyzer) AnalyzeAttribute(name, query string, attr map[string]interface{}) (*gen.ActionParam, error) {
	var param = gen.ActionParam{Name: name, QueryName: query, VarName: toVarName(name)}
	if d, ok := attr["description"]; ok {
		param.Description = removeBlankLines(d.(string))
	}
	if r, ok := attr["required"]; ok {
		if r.(bool) {
			param.Mandatory = true
		}
	}
	if options, ok := attr["options"]; ok {
		opts, ok := options.(map[string]interface{})
		if ok {
			for n, o := range opts {
				switch n {
				case "max":
					param.Max = int(o.(float64))
				case "min":
					param.Min = int(o.(float64))
				case "regexp":
					param.Regexp = o.(string)
				}
			}
		}
	}
	var t = attr["type"].(map[string]interface{})
	var dataType, err = a.AnalyzeType(t, query)
	if err != nil {
		return nil, err
	}
	param.Type = dataType

	return &param, nil
}

// Analyze type given its json definition
func (a *ApiAnalyzer) AnalyzeType(typeDef map[string]interface{}, query string) (gen.DataType, error) {
	var n, ok = typeDef["name"].(string)
	if !ok {
		return nil, fmt.Errorf("Missing type name in %s", prettify(typeDef))
	}
	if isBuiltInType(n) {
		n = "String"
	}
	var dataType gen.DataType
	switch n {
	case "Integer":
		var i = gen.BasicDataType("int")
		dataType = &i
	case "String":
		var s = gen.BasicDataType("string")
		dataType = &s
	case "Boolean":
		var b = gen.BasicDataType("bool")
		dataType = &b
	case "DateTime":
		var t = gen.BasicDataType("time.Time")
		dataType = &t
	case "Collection", "Ids":
		member, ok := typeDef["member_attribute"].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("Missing \"member_attribute\" for %s", prettify(typeDef))
		}
		var elemType, err = a.AnalyzeAttribute(n+"Member", query+"[]", member)
		if err != nil {
			return nil, fmt.Errorf("Failed to compute type of \"member_attribute\": %s", err.Error())
		}
		dataType = &gen.ArrayDataType{elemType}
	case "Struct":
		attrs, ok := typeDef["attributes"].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("Failed to retrieve attributes of struct for %s", prettify(typeDef))
		}
		var obj, err = a.CreateOrGetType(query, attrs)
		if err != nil {
			return nil, err
		}
		dataType = obj
	case "Hash":
		keys, ok := typeDef["keys"].(map[string]interface{})
		if !ok {
			dataType = new(gen.EnumerableDataType)
		} else {
			var obj, err = a.CreateOrGetType(query, keys)
			if err != nil {
				return nil, err
			}
			dataType = obj
		}
	default:
		// First check if we already analyzed that type
		if t := a.GetType(n); t != nil {
			return t, nil
		}

		// No then analyze it
		t, ok := a.RawTypes[n]
		if !ok {
			return nil, fmt.Errorf("Unknown type %s for %s", n, prettify(typeDef))
		}
		var attrs, ok2 = t["attributes"]
		if !ok2 {
			return nil, fmt.Errorf("Type %s has no attributes: %s", n, prettify(typeDef))
		}
		var att = attrs.(map[string]interface{})
		var obj = a.CreateType(n)
		obj.Fields = make([]*gen.ActionParam, len(att))

		var idx = 0
		for an, at := range att {
			var aq = fmt.Sprintf("%s[%s]", query, an)
			var ap, err = a.AnalyzeAttribute(an, aq, at.(map[string]interface{}))
			if err != nil {
				return nil, err
			}
			obj.Fields[idx] = ap
			idx += 1
		}

		// We're done
		dataType = obj
	}

	return dataType, nil
}

// Helper method that creates or retrieve a object data type given its attributes.
func (a *ApiAnalyzer) CreateOrGetType(query string, attributes map[string]interface{}) (*gen.ObjectDataType, error) {
	var hasher = md5.New()
	hasher.Write([]byte(fmt.Sprintf("%v", attributes)))
	var md5str = hex.EncodeToString(hasher.Sum(nil))
	var name = inflect.Camelize(bracketRegexp.ReplaceAllLiteralString(query, "_") + "_struct")
	var obj = a.GetOrCreate(md5str, name)
	obj.Fields = make([]*gen.ActionParam, len(attributes))
	var idx = 0
	for an, at := range attributes {
		att, err := a.AnalyzeAttribute(an, fmt.Sprintf("%s[%s]", query, an), at.(map[string]interface{}))
		if err != nil {
			return nil, fmt.Errorf("Failed to compute type of attribute %s: %s", an, err.Error())
		}
		obj.Fields[idx] = att
		idx += 1
	}
	return obj, nil
}
