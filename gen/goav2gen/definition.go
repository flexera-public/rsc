package main

import (
	"sort"
	"strings"

	"github.com/rightscale/rsc/gen"
)

// AnalyzeDefinition analyzes a definition which is an object. Definitions are referenced by
// both inputs to requests (parameters) and outputs from requests (return types). EvalCtx
// will store which we're analyzing currently.
func (a *APIAnalyzer) AnalyzeDefinition(ec EvalCtx, def *Definition, refID string) gen.DataType {
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
		name := a.guessType(ec, def, refID)
		if name == "string" || name == "boolean" || name == "integer" {
			return basicType("string")
		}
		// else: object is assumed
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

// guessType tries to guess the resource name based on the definition and service.
// This info is not stored in the swagger. TBD manual overrides if needed.
func (a *APIAnalyzer) guessType(ec EvalCtx, d *Definition, refID string) string {
	// First get the type name and and view from the swagger reference definition
	// name -- are a few cases where that's the only place that has the view
	if t, ok := a.TypeOverrides[refID]; ok {
		return t
	}
	var name, view string
	if strings.Contains(refID, "RequestBody") {
		bits := strings.Split(refID, "RequestBody")
		name = bits[0]
		if len(bits) > 1 {
			view = strings.ToLower(bits[1])
		}
	} else if strings.Contains(refID, "ResponseBody") {
		bits := strings.Split(refID, "ResponseBody")
		name = bits[0]
		if len(bits) > 1 {
			view = strings.ToLower(bits[1])
		}
	} else {
		name = refID
	}

	// Now try and get it from the media type -- this is preferred if its set.
	if mt := mediaType(d.Title); mt != "" {
		if strings.Contains(mt, "application") {
			bits := strings.Split(mt, ".")
			name := bits[len(bits)-1]
			attrs := mediaTypeAttrs(d.Title)
			if attrs["type"] != "" {
				name += "_" + attrs["type"]
			}
			if attrs["view"] != "" && attrs["view"] != "default" {
				name += "_" + attrs["view"]
			} else if view != "" {
				name += "_" + view
			}
			dbg("DEBUG media type refID:%#v title:%#v name:%#v view:%#v -> type:%#v\n", refID, d.Title, name, view, name)
			return toTypeName(name)
		} else if strings.Contains(mt, "text/") {
			return "string"
		} else {
			fail("Don't know how to handle media type %s", mt)
		}
	}
	if view != "" {
		return name + "_" + view
	}

	return name
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
	if min, ok := p.Minimum.(float64); ok {
		ap.Min = int(min)
	}
	if max, ok := p.Maximum.(float64); ok {
		ap.Max = int(max)
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
		if p.Format == "date-time" {
			a.api.NeedTime = true
			t := gen.BasicDataType("*time.Time")
			ap.Type = &t
		} else if t := a.overrideAttr(ec, name); t != nil {
			ap.Type = t
		} else {
			ap.Type = basicType(p.Type)
		}
	}

	return ap
}

func (a *APIAnalyzer) overrideAttr(ec EvalCtx, name string) gen.DataType {
	if override, ok := a.AttrOverrides[name]; ok {
		switch override {
		case "sourcefile":
			return new(gen.UploadDataType)
		}
		return basicType(override)
	}
	return nil
}

// typeForRef converts a "$ref" reference in the swagger definition to a concrete type, usually an object.
func (a *APIAnalyzer) typeForRef(ec EvalCtx, r Ref) gen.DataType {
	switch r.Type() {
	case "array":
		fail("References to arrays not implemented yet")
	case "object":
		def := a.Doc.Ref(r)
		if def == nil {
			fail("No ref for %+v", r)
		}
		dt := a.AnalyzeDefinition(ec, def, r.ID())
		if objDT, ok := dt.(*gen.ObjectDataType); ok {
			a.addType(ec, objDT, r)
		}
		return dt
	}
	return basicType(r.Type())

}

// addType conditionally adds a new type, trying its best to avoid type collisions
func (a *APIAnalyzer) addType(ec EvalCtx, dt *gen.ObjectDataType, r Ref) {
	a.api.NeedJSON = true
	if a.refByType[dt.TypeName] == r.ID() {
		return
	}

	if other, ok := a.api.Types[dt.TypeName]; ok {
		if !ec.IsResult {
			// If its an input parameter, fix the collision by amending this types name
			dt.TypeName += "Param"
			if a.refByType[dt.TypeName] == r.ID() {
				return
			}
		}
		oldFields := []string{}
		newFields := []string{}
		for _, f := range other.Fields {
			oldFields = append(oldFields, f.Name)
		}
		for _, f := range dt.Fields {
			newFields = append(newFields, f.Name)
		}
		use := "Old"
		if len(newFields) > len(oldFields) {
			use = "New"
		}
		if strings.Join(oldFields, ",") != strings.Join(newFields, ",") {
			warn("Warning: Type collision when adding new type %s!\n  New: id %s fields %v\n  Old: id %s fields %v\n  Using %s, which has more fields\n",
				dt.TypeName, r.ID(), newFields, a.refByType[dt.TypeName], oldFields, use)
		}
		if use == "Old" {
			return
		}
	}
	dbg("DEBUG NEW TYPE %s\n", prettify(dt))
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
