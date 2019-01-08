package main // import "gopkg.in/rightscale/rsc.v6/gen/goav2gen"

import (
	"gopkg.in/rightscale/rsc.v6/gen"
)

var loc = map[string]int{
	"path":   gen.PathParam,
	"query":  gen.QueryParam,
	"body":   gen.PayloadParam,
	"header": gen.HeaderParam,
}

// AnalyzeParam analyzes input parameters to methods
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
