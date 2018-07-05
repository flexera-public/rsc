package main

import (
	"regexp"
	"strings"

	"bitbucket.org/pkg/inflect"
	"github.com/kr/pretty"
	"github.com/rightscale/rsc/gen"
)

// AnalyzeEndpoint creates an API descriptor from goa v2 generated swagger definition
func (a *APIAnalyzer) AnalyzeEndpoint(verb string, path string, ep *Endpoint) error {
	path = joinPath(a.Doc.BasePath, path)
	dbg("\n------\nDEBUG AnalyzeEndpoint %s %s %+v\n", verb, path, ep)
	pattern := toPattern(verb, path)
	dbg("DEBUG AnalyzeEndpoint pattern %v\n", pattern)
	svc := ep.Service()
	svcName := svc + "Service"

	// Get Resource -- base it on the service name for now
	res := a.api.Resources[svcName]
	if res == nil {
		res = &gen.Resource{
			Name:       svcName,
			ClientName: a.ClientName,
			Actions:    []*gen.Action{},
		}
		a.api.Resources[svcName] = res
	}
	action := &gen.Action{
		Name:         ep.Method(),
		MethodName:   toMethodName(ep.Method()),
		Description:  cleanDescription(ep.Description),
		ResourceName: svcName,
		PathPatterns: []*gen.PathPattern{pattern},
	}
	res.Actions = append(res.Actions, action)

	var returnDT gen.DataType
	var hasLocation bool
	for code, response := range ep.Responses {
		if code >= 300 {
			continue
		}

		if response.Headers != nil {
			if _, ok := response.Headers["Location"]; ok {
				hasLocation = true
			}
		}

		if response.Schema == nil {
			dbg("DEBUG MISSING SCHEMA SKIP!\n")
			break
		}
		dbg("DEBUG AnalyzeEndpoint %d RESPONSE %#v\n", code, response)

		returnDef := a.Doc.Ref(response.Schema)
		if returnDef != nil {
			ec := EvalCtx{IsResult: true, Trail: nil, Svc: res, Method: action}
			if mediaType(returnDef.Title) == "" {
				dbg("DEBUG AnalyzeEndpoint UNKNOWN mediatype! %#v\n", returnDef)
				continue
			}
			//returnTypeName = a.guessType(svc, returnDef)
			returnDT = a.AnalyzeDefinition(ec, returnDef)
			if returnObj, ok := returnDT.(*gen.ObjectDataType); ok {
				isResourceType := verb == "get" && returnObj.TypeName == svc
				dbg("DEBUG AnalyzeEndpoint Path %s Verb %s returnTypeName %s svc %s\n", path, verb, returnObj.TypeName, svc)

				if isResourceType {
					res.Description = cleanDescription(ep.Description)
					res.Identifier = mediaType(returnDef.Title)
					//copyFieldsToResource(res, returnObj)
				}
				a.addType(ec, returnObj, response.Schema)
			}
		} else {
			//dbg("HERE")
			returnDT = basicType(response.Schema.Type())
			//returnTypeName = toGoTypeName(response.Schema.Type())
			//			returnType = basicType(returnTypeName)
		}

		break
	}

	for _, p := range ep.Parameters {
		ec := EvalCtx{IsResult: false, Trail: nil, Svc: res, Method: action}
		ap := a.AnalyzeParam(ec, p)

		switch p.In {
		case "header":
			action.HeaderParamNames = append(action.HeaderParamNames, p.Name)
			action.Params = append(action.Params, ap)
		case "query":
			action.QueryParamNames = append(action.QueryParamNames, p.Name)
			action.Params = append(action.Params, ap)
		case "path":
			action.PathParamNames = append(action.PathParamNames, p.Name)
			action.Params = append(action.Params, ap)
		case "body":
			def := a.Doc.Ref(p.Schema)
			if def != nil {
				if def.Type == "array" {
					fail("ERROR ANALYZE PARAMETERS UNHANDLED ARRAY!")
				} else if def.Type == "object" {
					// Flatten the first level of object
					dt := a.AnalyzeDefinition(ec, def)
					if obj, ok := dt.(*gen.ObjectDataType); ok {
						a.addType(ec, obj, p.Schema)
						action.Payload = obj
						for _, f := range obj.Fields {
							action.PayloadParamNames = append(action.PayloadParamNames, f.Name)
							action.Params = append(action.Params, f)
						}
					}

				}
			}
		}
	}
	if returnDT != nil {
		action.Return = signature(returnDT)

	}
	action.ReturnLocation = hasLocation
	dbg("DEBUG ACTION % #v", pretty.Formatter(action))
	return nil
}

func copyFieldsToResource(res *gen.Resource, returnObj *gen.ObjectDataType) {
	for _, f := range returnObj.Fields {
		switch f.Type.(type) {
		case *gen.ObjectDataType:
			attr := &gen.Attribute{
				Name:      f.Name,
				FieldName: toTypeName(f.Name),
				FieldType: f.Signature(),
			}
			dbg("COPYFIELDS %v\n", attr)
			//res.Attributes = append(res.Attributes, attr)
		case *gen.ArrayDataType:
			attr := &gen.Attribute{
				Name:      f.Name,
				FieldName: toTypeName(f.Name),
				FieldType: f.Signature(),
			}
			dbg("COPYFIELDS %v\n", attr)
		default:
			attr := &gen.Attribute{
				Name:      f.Name,
				FieldName: toTypeName(f.Name),
				FieldType: f.Signature(),
			}
			dbg("COPYFIELDS %v\n", attr)

			//res.Attributes = append(res.Attributes, attr)
		}

	}
}

// guessType tries to guess the resource name based on the definition and service.
// This info is not stored in the swagger. TBD manual overrides if needed
func (a *APIAnalyzer) guessType(ec EvalCtx, d *Definition) string {
	if mt := mediaType(d.Title); mt != "" {
		bits := strings.Split(mt, ".")
		if len(bits) > 1 {
			name := bits[len(bits)-1]
			attrs := mediaTypeAttrs(d.Title)
			if attrs["type"] != "" {
				name += "_" + attrs["type"]
			}
			return toTypeName(name)
		}
	}
	return normTitle(d.Title)
}

func normResponseName(s string) string {
	s = strings.TrimSuffix(s, "RequestBody")
	s = strings.TrimSuffix(s, "ResponseBody")
	return inflect.Underscore(s)
}

// locatorFunc returns the source for the function returning the resource locator built from its
// href field.
func locatorFunc(resource string) string {
	return "return api." + resource + "Locator(r.Href)"
}

// Regular expression that captures variables in a path
var pathVarsRe = regexp.MustCompile(`/{([^}]+)}`)
var pathVarsReQuoted = regexp.MustCompile(`/\\{([^}]+)\\}`)

// Create path pattern from HTTP verb and request path
func toPattern(verb, path string) *gen.PathPattern {
	pattern := gen.PathPattern{
		HTTPMethod: verb,
		Path:       path,
		Pattern:    pathVarsRe.ReplaceAllLiteralString(path, "/%s"),
		Regexp:     pathVarsReQuoted.ReplaceAllLiteralString(regexp.QuoteMeta(path), `/([^/]+)`),
	}
	matches := pathVarsRe.FindAllStringSubmatch(path, -1)
	if len(matches) > 0 {
		pattern.Variables = make([]string, len(matches))
		for i, m := range matches {
			pattern.Variables[i] = m[1]
		}
	}
	return &pattern
}

func joinPath(parts ...string) string {
	path := strings.Join(parts, "/")
	return strings.Replace(path, "//", "/", -1)
}
