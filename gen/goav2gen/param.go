package main

import (
	"github.com/rightscale/rsc/gen"
)

var loc = map[string]int{
	"path":   gen.PathParam,
	"query":  gen.QueryParam,
	"body":   gen.PayloadParam,
	"header": gen.HeaderParam,
}

func (a *APIAnalyzer) AnalyzeParam(ec EvalCtx, p *Parameter) *gen.ActionParam {
	location, ok := loc[p.In]
	if !ok {
		location = -1
	}
	ap := &gen.ActionParam{
		Name:        p.Name,
		QueryName:   p.Name,
		Description: cleanDescription(p.Description),
		VarName:     toVarName(p.Name),
		Location:    location,
		Mandatory:   p.Required,
		NonBlank:    p.Required || p.Pattern != "",
		Regexp:      p.Pattern,
		ValidValues: p.Enum,
	}
	if p.Schema != nil {
		ap.Type = a.typeForRef(ec, p.Schema)
	} else {
		ap.Type = basicType(p.Type)
	}

	return ap
}

func (a *APIAnalyzer) typeForRef(ec EvalCtx, r Ref) gen.DataType {
	switch r.Type() {
	case "array":
		panic("here")
	case "object":
		def := a.Doc.Ref(r)
		if def == nil {
			fail("No ref for %+v", r)
		}
		//defType := a.guessType("", def)
		dt := a.AnalyzeDefinition(ec, def)
		if objDT, ok := dt.(*gen.ObjectDataType); ok {
			a.addType(ec, objDT, r)
		}
		return dt
	}
	return basicType(r.Type())

}

func (a *APIAnalyzer) addType(ec EvalCtx, dt *gen.ObjectDataType, r Ref) {
	if other, ok := a.api.Types[dt.TypeName]; ok {
		if a.refByType[dt.TypeName] == r.ID() {
			return
		}
		if !ec.IsResult {
			dt.TypeName += "Param"
			if a.refByType[dt.TypeName] == r.ID() {
				return
			}
		} else {
			dbg("DEBUG: Type collision %s!\nNEW: %s\n%#v\nOLD: %s\n%#v\n",
				dt.TypeName, r.ID(), dt, a.refByType[dt.TypeName], other)
		}
	}
	dbg("DEBUG NEW TYPE %s\n", prettify(dt))
	a.api.NeedJSON = true
	a.refByType[dt.TypeName] = r.ID()
	a.api.Types[dt.TypeName] = dt
}

func basicType(typ string) gen.DataType {
	var goTyp = typ
	switch typ {
	case "string":
		goTyp = "string"
	case "integer":
		goTyp = "int"
	case "boolean":
		goTyp = "bool"
	default:
		goTyp = "interface{}"
	}
	bt := gen.BasicDataType(goTyp)
	return &bt
}