package main // import "gopkg.in/rightscale/rsc.v8/gen/goav2gen"

import (
	"fmt"
	"sort"

	"gopkg.in/rightscale/rsc.v8/gen"
)

// APIAnalyzer is holds the analysis results.
type APIAnalyzer struct {
	// Swagger document
	Doc *Doc
	// Version being analyzed
	Version string
	// Name of golang struct to generate for API client
	ClientName string
	// AttrOverrides holds attribute overrides
	AttrOverrides map[string]string
	// TypeOverrides holds type overrides. We have to guess the goa type for result types from
	// methods and TypeOverrides is for the cases where we can't guess or the swagger is wrong
	// because of a bug
	TypeOverrides map[string]string

	// Temporary data structures used by analysis
	refByType map[string]string

	// Descriptor being built
	api *gen.APIDescriptor
}

// NewAPIAnalyzer is the factory method for APIAnalyzer.
func NewAPIAnalyzer(version, clientName string, doc *Doc, attrOverrides, typeOverrides map[string]string) *APIAnalyzer {
	return &APIAnalyzer{
		Doc:           doc,
		Version:       version,
		ClientName:    clientName,
		AttrOverrides: attrOverrides,
		TypeOverrides: typeOverrides,

		refByType: map[string]string{},
	}
}

// Analyze creates an API descriptor from raw resources and types.
func (a *APIAnalyzer) Analyze() (*gen.APIDescriptor, error) {
	api := gen.APIDescriptor{
		Version:   a.Version,
		Resources: make(map[string]*gen.Resource),
		Types:     make(map[string]*gen.ObjectDataType),
	}
	a.api = &api

	sortedVerbs := []string{"get", "post", "put", "delete"}
	paths := []string{}
	for path := range a.Doc.Paths {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths {
		endpointMap := a.Doc.Paths[path]
		for _, verb := range sortedVerbs {
			if ep, ok := endpointMap[verb]; ok {
				a.AnalyzeEndpoint(verb, path, ep)
			}
		}
	}

	a.finalize()

	return &api, nil
}

func (a *APIAnalyzer) finalize() {
	// Create resource fields
	for _, res := range a.api.Resources {
		if t, ok := a.api.Types[res.Name]; ok {
			convertToResource(res, t)
			delete(a.api.Types, res.Name)
		}
	}

	// Generate leaf params for command line
	for _, res := range a.api.Resources {
		for _, ax := range res.Actions {
			cmdLineParams := []*gen.ActionParam{}
			methodParams := []*gen.ActionParam{}
			for _, p := range ax.Params {
				switch p.Location {
				case gen.PathParam:
					//methodParams = append(methodParams, p)
				case gen.QueryParam:
					methodParams = append(methodParams, p)
					cmdLineParams = append(cmdLineParams, p)
				case gen.PayloadParam:
					methodParams = append(methodParams, p)
					extracted := extractCmdLineParams(p, p.Name, make(map[string]*[]*gen.ActionParam), !p.Mandatory)
					for _, e := range extracted {
						e.Location = gen.PayloadParam
					}
					cmdLineParams = append(cmdLineParams, extracted...)
				}
			}
			ax.LeafParams = cmdLineParams
			ax.Params = methodParams
		}
	}

	for n := range a.api.Resources {
		a.api.ResourceNames = append(a.api.ResourceNames, n)
	}
	sort.Strings(a.api.ResourceNames)

	for n := range a.api.Types {
		a.api.TypeNames = append(a.api.TypeNames, n)
	}
	sort.Strings(a.api.TypeNames)

}

// extractCmdLineParams generates flags for the command line
func extractCmdLineParams(a *gen.ActionParam, root string, seen map[string]*[]*gen.ActionParam, parentNotMandatory bool) []*gen.ActionParam {
	switch t := a.Type.(type) {
	case *gen.BasicDataType, *gen.EnumerableDataType, *gen.UploadDataType:
		dup := gen.ActionParam{
			Name:        a.Name,
			QueryName:   root,
			Description: a.Description,
			VarName:     a.VarName,
			Location:    a.Location,
			Type:        a.Type,
			Mandatory:   a.Mandatory && !parentNotMandatory, // yay for double negations!
			NonBlank:    a.NonBlank,
			Regexp:      a.Regexp,
			ValidValues: a.ValidValues,
			Min:         a.Min,
			Max:         a.Max,
		}
		return []*gen.ActionParam{&dup}
	case *gen.ArrayDataType:
		p := t.ElemType
		eq, ok := seen[p.Name]
		if !ok {
			eq = &[]*gen.ActionParam{}
			seen[p.Name] = eq
			*eq = extractCmdLineParams(p, root+"[]", seen, parentNotMandatory || !a.Mandatory)
		}
		return *eq
	case *gen.ObjectDataType:
		params := []*gen.ActionParam{}
		for _, f := range t.Fields {
			eq, ok := seen[f.Name]
			if !ok {
				eq = &[]*gen.ActionParam{}
				seen[f.Name] = eq
				*eq = extractCmdLineParams(f, fmt.Sprintf("%s[%s]", root, f.Name), seen, parentNotMandatory || !a.Mandatory)
			}
			params = append(params, *eq...)
		}
		return params
	}
	return nil
}

func convertToResource(res *gen.Resource, returnObj *gen.ObjectDataType) {
	for _, f := range returnObj.Fields {
		switch f.Type.(type) {
		case *gen.ObjectDataType:
			attr := &gen.Attribute{
				Name:      f.Name,
				FieldName: toTypeName(f.Name),
				FieldType: f.Signature(),
			}
			dbg("COPYFIELDS %v\n", attr)
			res.Attributes = append(res.Attributes, attr)
		case *gen.ArrayDataType:
			attr := &gen.Attribute{
				Name:      f.Name,
				FieldName: toTypeName(f.Name),
				FieldType: f.Signature(),
			}
			dbg("COPYFIELDS %v\n", attr)
			res.Attributes = append(res.Attributes, attr)
		default:
			attr := &gen.Attribute{
				Name:      f.Name,
				FieldName: toTypeName(f.Name),
				FieldType: f.Signature(),
			}
			dbg("COPYFIELDS %v\n", attr)
			res.Attributes = append(res.Attributes, attr)
		}

	}
}
