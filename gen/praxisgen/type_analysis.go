package main

import (
	"fmt"
	"regexp"
	"strings"

	"bitbucket.org/pkg/inflect"

	"github.com/rightscale/rsc/gen"
)

// Regular expression used to capture brackets in query name
var bracketRegexp = regexp.MustCompile(`(\[|\])+`)

// AnalyzeAttribute analyzes an attribute creating a corresponding ActionParam.
func (a *APIAnalyzer) AnalyzeAttribute(name, query string, attr map[string]interface{}) (*gen.ActionParam, error) {
	param := gen.ActionParam{Name: name, QueryName: query, VarName: toVarName(name)}
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
	if values, ok := attr["values"]; ok {
		param.ValidValues = values.([]interface{})
	}
	t := attr["type"].(map[string]interface{})
	dataType, err := a.AnalyzeType(t, query)
	if err != nil {
		return nil, err
	}
	param.Type = dataType
	switch dataType.(type) {
	case *gen.ArrayDataType:
		param.QueryName += "[]"
	}

	return &param, nil
}

// AnalyzeType analyzes a type given its JSON definition.
func (a *APIAnalyzer) AnalyzeType(typeDef map[string]interface{}, query string) (gen.DataType, error) {
	n, ok := typeDef["name"].(string)
	if !ok {
		n = "Struct" // Assume inline struct (e.g. payload types)
	}
	if strings.HasSuffix(n, "FileUpload") {
		// A little bit hacky but this is to make upload work with resticle
		// The idea is that a type named "FileUpload" is assumed to define a multipart
		// request with a file part.
		// The type must define a "name" and  "filename" string fields.
		t, ok := a.RawTypes[n]
		if !ok {
			return nil, fmt.Errorf("Unknown type %s for %s", n, prettify(typeDef))
		}
		attrs, ok := t["attributes"]
		if !ok {
			return nil, fmt.Errorf("Invalid file upload type %s for %s: no attributes", n, prettify(typeDef))
		}
		mattrs, ok := attrs.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("Invalid file upload type %s for %s: basic type", n, prettify(typeDef))
		}
		_, ok = mattrs["name"]
		if !ok {
			return nil, fmt.Errorf("Invalid file upload type %s for %s: no name", n, prettify(typeDef))
		}
		_, ok = mattrs["filename"]
		if !ok {
			return nil, fmt.Errorf("Invalid file upload type %s for %s: no filename", n, prettify(typeDef))
		}
		return &gen.UploadDataType{TypeName: n}, nil
	}
	if isBuiltInType(n) {
		n = "String"
	}
	var dataType gen.DataType
	switch n {
	case "Integer":
		i := gen.BasicDataType("int")
		dataType = &i
	case "Float":
		f := gen.BasicDataType("float64")
		dataType = &f
	case "String":
		s := gen.BasicDataType("string")
		dataType = &s
	case "Boolean":
		b := gen.BasicDataType("bool")
		dataType = &b
	case "Object":
		o := gen.BasicDataType("interface{}")
		dataType = &o
	case "DateTime":
		t := gen.BasicDataType("*time.Time") // Need pointer for case where value is null
		a.descriptor.NeedTime = true
		dataType = &t
	case "Collection", "Ids":
		member, ok := typeDef["member_attribute"].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("Missing \"member_attribute\" for %s", prettify(typeDef))
		}
		elemType, err := a.AnalyzeAttribute(n+"Member", query+"[]", member)
		if err != nil {
			return nil, fmt.Errorf("Failed to compute type of \"member_attribute\": %s", err)
		}
		dataType = &gen.ArrayDataType{elemType}
	case "Struct":
		attrs, ok := typeDef["attributes"].(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("Failed to retrieve attributes of struct for %s", prettify(typeDef))
		}
		obj, err := a.CreateType(query, attrs)
		if err != nil {
			return nil, err
		}
		dataType = obj
	case "Hash":
		keys, ok := typeDef["keys"].(map[string]interface{})
		if !ok {
			dataType = new(gen.EnumerableDataType)
		} else {
			obj, err := a.CreateType(query, keys)
			if err != nil {
				return nil, err
			}
			dataType = obj
		}
	default:
		// First check if we already analyzed that type
		if t := a.Registry.GetNamedType(n); t != nil {
			return t, nil
		}

		// No then analyze it
		t, ok := a.RawTypes[n]
		if !ok {
			return nil, fmt.Errorf("Unknown type %s for %s", n, prettify(typeDef))
		}
		attrs, ok := t["attributes"]
		if !ok {
			// No attribute, it's a string
			s := gen.BasicDataType("string")
			dataType = &s
		} else {
			att := attrs.(map[string]interface{})
			obj := a.Registry.CreateNamedType(n)
			obj.Fields = make([]*gen.ActionParam, len(att))

			for idx, an := range sortedKeys(att) {
				at := att[an]
				aq := fmt.Sprintf("%s[%s]", query, an)
				ap, err := a.AnalyzeAttribute(an, aq, at.(map[string]interface{}))
				if err != nil {
					return nil, err
				}
				obj.Fields[idx] = ap
			}

			// We're done
			dataType = obj
		}
	}

	return dataType, nil
}

// CreateType is a helper method that creates or retrieve a object data type given its attributes.
func (a *APIAnalyzer) CreateType(query string, attributes map[string]interface{}) (*gen.ObjectDataType, error) {
	name := inflect.Camelize(bracketRegexp.ReplaceAllLiteralString(query, "_") + "_struct")
	obj := a.Registry.CreateInlineType(name)
	obj.Fields = make([]*gen.ActionParam, len(attributes))
	for idx, an := range sortedKeys(attributes) {
		at := attributes[an]
		var childQ string
		if query == "payload" {
			childQ = an
		} else {
			childQ = fmt.Sprintf("%s[%s]", query, an)
		}
		att, err := a.AnalyzeAttribute(an, childQ, at.(map[string]interface{}))
		if err != nil {
			return nil, fmt.Errorf("Failed to compute type of attribute %s: %s", an, err)
		}
		obj.Fields[idx] = att
	}
	return obj, nil
}
