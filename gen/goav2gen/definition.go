package main

import (
	"sort"

	"github.com/rightscale/rsc/gen"
)

// AnalyzeDefinition analyzes a definition which is an object. Definitions are referenced by
// both inputs to requests (parameters) and outputs from requests (return types). EvalCtx
// will store which we're analyzing currently.
func (a *APIAnalyzer) AnalyzeDefinition(ec EvalCtx, def *Definition) gen.DataType {
	switch def.Type {
	case "array":
		elemT := a.typeForRef(ec, def.Items)
		apElem := &gen.ActionParam{
			Location: gen.PayloadParam,
			Type:     elemT,
		}
		return &gen.ArrayDataType{
			ElemType: apElem,
		}
	default:
		// object assumed
		name := a.guessType(ec, def)
		isRequired := map[string]bool{}
		for _, r := range def.Required {
			isRequired[r] = true
		}

		obj := &gen.ObjectDataType{TypeName: name}
		keys := []string{}
		for propName := range def.Properties {
			keys = append(keys, propName)
		}
		sort.Strings(keys)
		for _, propName := range keys {
			prop := def.Properties[propName]
			attr := a.AnalyzeProperty(ec.WithTrail(name), propName, isRequired[propName], prop)
			obj.Fields = append(obj.Fields, attr)
		}
		return obj
	}

}

func (a *APIAnalyzer) AnalyzeProperty(ec EvalCtx, name string, required bool, p *Property) *gen.ActionParam {
	queryName := name
	if len(ec.Trail) > 0 {
		queryName = ec.Trail[0]
		for i := 1; i < len(ec.Trail); i++ {
			queryName += "[" + ec.Trail[i] + "]"
		}
		queryName += "[" + name + "]"
	}
	if p.Type == "Array" {
		queryName += "[]"
	}
	ap := &gen.ActionParam{
		Name:        name,
		QueryName:   queryName,
		Description: cleanDescription(p.Description),
		VarName:     toVarName(name),
		Location:    gen.PayloadParam,
		Mandatory:   required,
		NonBlank:    (required || p.Pattern != "") && (p.Default != nil),
		Regexp:      p.Pattern,
		ValidValues: p.Enum,
	}
	if p.IsRef() {
		ap.Type = a.typeForRef(ec, p.GetRef())
		return ap
	}
	switch p.Type {
	case "array":
		elemT := a.typeForRef(ec.WithTrail(name+"[]"), p.Items)
		apElem := &gen.ActionParam{
			Location: gen.PayloadParam,
			Type:     elemT,
		}
		ap.Type = &gen.ArrayDataType{
			ElemType: apElem,
		}
	case "object":
		if a.Doc.Ref(p.Schema) != nil {
			ap.Type = a.typeForRef(ec.WithTrail(name), p.Schema)
		} else {
			dt := gen.EnumerableDataType(0)
			ap.Type = &dt
		}
	default:
		ap.Type = basicType(p.Type)
	}

	return ap
}
