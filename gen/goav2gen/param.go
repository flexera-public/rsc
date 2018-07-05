package main

import (
	"strings"

	"github.com/kr/pretty"
	"github.com/rightscale/rsc/gen"
)

var loc = map[string]int{
	"path":   gen.PathParam,
	"query":  gen.QueryParam,
	"body":   gen.PayloadParam,
	"header": gen.HeaderParam,
}

type EvalCtx struct {
	IsResult bool     // true means this is a ResultType. false means input to function
	Trail    []string // trail of previous types we're a child of
	Svc      *gen.Resource
	Method   *gen.Action
}

func (ec EvalCtx) WithTrail(t string) EvalCtx {
	newEC := ec
	trailCopy := make([]string, 0, len(ec.Trail)+1)
	for _, val := range ec.Trail {
		trailCopy = append(trailCopy, val)
	}
	newEC.Trail = append(trailCopy, t)
	return newEC
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
	// if _, ok := a.api.Types[dt.TypeName]; ok {
	// 	dbg("DEBUG: Not creating new type %s, exists as resource", dt.TypeName)
	// 	return
	// }
	dbg("DEBUG NEW TYPE % #v\n", pretty.Formatter(dt))
	a.api.NeedJSON = true
	a.refByType[dt.TypeName] = r.ID()
	a.api.Types[dt.TypeName] = dt
}

func normTitle(s string) string {
	s = strings.TrimSuffix(s, "RequestBody")
	s = strings.TrimSuffix(s, "ResponseBody")
	return s
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
