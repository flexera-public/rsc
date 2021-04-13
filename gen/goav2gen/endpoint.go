package main // import "gopkg.in/rightscale/rsc.v9/gen/goav2gen"

import (
	"regexp"
	"strings"

	"gopkg.in/rightscale/rsc.v9/gen"
)

// EvalCtx stores what is currently under evaluation
type EvalCtx struct {
	// IsResult true means this is a ResultType. false means input to function
	IsResult bool
	// Trail of where we're evaluating is we're evaluating down a chain of objects
	Trail []string
	// Svc links to the goa v2 service
	Svc *gen.Resource
	// Method links to the goa v2 method/action
	Method *gen.Action
}

// WithTrail creates a new context with trail appended to
func (ec EvalCtx) WithTrail(t string) EvalCtx {
	newEC := ec
	trailCopy := make([]string, 0, len(ec.Trail)+1)
	for _, val := range ec.Trail {
		trailCopy = append(trailCopy, val)
	}
	newEC.Trail = append(trailCopy, t)
	return newEC
}

// AnalyzeEndpoint creates an API descriptor from goa v2 generated swagger definition
func (a *APIAnalyzer) AnalyzeEndpoint(verb string, path string, ep *Endpoint) error {
	path = joinPath(a.Doc.BasePath, path)
	dbg("\n------\nDEBUG AnalyzeEndpoint %s %s %+v\n", verb, path, ep)
	pattern := toPattern(verb, path)
	dbg("DEBUG AnalyzeEndpoint pattern %v\n", pattern)
	svc := ep.Service()

	// Get Resource -- base it on the service name for now
	res := a.api.Resources[svc]
	if res == nil {
		res = &gen.Resource{
			Name:       svc,
			ClientName: a.ClientName,
			Actions:    []*gen.Action{},
		}
		a.api.Resources[svc] = res
	}
	action := &gen.Action{
		Name:         ep.Method(),
		MethodName:   toMethodName(ep.Method()),
		Description:  cleanDescription(ep.Description),
		ResourceName: svc,
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
				warn("Warning: AnalyzeEndpoint: MediaType not set for %s, will be hard to guess the result type\n", response.Schema.ID())
				continue
			}
			returnDT = a.AnalyzeDefinition(ec, returnDef, response.Schema.ID())
			if returnObj, ok := returnDT.(*gen.ObjectDataType); ok {
				isResourceType := verb == "get" && returnObj.TypeName == svc
				dbg("DEBUG AnalyzeEndpoint Path %s Verb %s returnTypeName %s svc %s\n", path, verb, returnObj.TypeName, svc)

				if isResourceType {
					res.Description = cleanDescription(ep.Description)
					res.Identifier = mediaType(returnDef.Title)
				}
				a.addType(ec, returnObj, response.Schema)
			}
		} else {
			returnDT = basicType(response.Schema.Type())
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
					fail("Array type for body not implemented yet")
				} else if def.Type == "object" {
					// Flatten the first level of object
					dt := a.AnalyzeDefinition(ec, def, p.Schema.ID())
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
	dbg("DEBUG ACTION %s", prettify(action))
	return nil
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
		HTTPMethod: strings.ToUpper(verb),
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
