package main

import (
	"fmt"

	"bitbucket.org/pkg/inflect"

	"github.com/rightscale/rsc/gen"
)

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
		var typeName string
		for n, _ := range attrs {
			typeName = inflect.Camelize(n) + "Struct"
			break
		}
		var obj = gen.ObjectDataType{Name: typeName}
		obj.Fields = make([]*gen.ActionParam, len(attrs))

		// Record newly analyzed type to avoid infinite recursion
		a.typeNames[n] = typeName
		a.descriptor.Types[typeName] = &obj

		var idx = 0
		for an, at := range attrs {
			att, err := a.AnalyzeAttribute(an, fmt.Sprintf("%s[%s]", query, an), at.(map[string]interface{}))
			if err != nil {
				return nil, fmt.Errorf("Failed to compute type of attribute %s: %s", an, err.Error())
			}
			obj.Fields[idx] = att
			idx += 1
		}
		dataType = &obj
	case "Hash":
		keys, ok := typeDef["keys"].(map[string]map[string]interface{})
		if !ok {
			dataType = new(gen.EnumerableDataType)
		} else {
			var typeName = toGoTypeName(n, false)
			var obj = gen.ObjectDataType{Name: typeName}
			obj.Fields = make([]*gen.ActionParam, len(keys))

			// Record newly analyzed type to avoid infinite recursion
			a.typeNames[n] = typeName
			a.descriptor.Types[typeName] = &obj

			var idx = 0
			for kn, k := range keys {
				kt, err := a.AnalyzeAttribute(kn, fmt.Sprintf("%s[%s]", query, kn), k)
				if err != nil {
					return nil, fmt.Errorf("Failed to compute type of key %s of hash for %s",
						kn, prettify(typeDef))
				}
				obj.Fields[idx] = kt
				idx += 1
			}
			dataType = &obj
		}
	default:
		// First check if we already analyzed that type
		var typeName = toGoTypeName(n, false)
		if _, ok := a.typeNames[n]; ok {
			return a.descriptor.Types[typeName], nil
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
		var obj = gen.ObjectDataType{Name: typeName}
		obj.Fields = make([]*gen.ActionParam, len(att))

		// Record newly analyzed type to avoid infinite recursion
		a.typeNames[n] = typeName
		a.descriptor.Types[typeName] = &obj

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
		dataType = &obj
	}

	return dataType, nil
}
