package writers

import (
	"fmt"
	"io"
	"strings"
	"text/template"
	"time"

	"github.com/rightscale/rsc/gen"
)

// ClientWriter struct exposes methods to generate the go API client code
type ClientWriter struct {
	headerTmpl   *template.Template
	resourceTmpl *template.Template
}

// Client writer factory
func NewClientWriter() (*ClientWriter, error) {
	funcMap := template.FuncMap{
		"comment":         comment,
		"now":             time.Now,
		"commandLine":     commandLine,
		"parameters":      parameters,
		"paramsAsPayload": paramsAsPayload,
		"isPointer":       isPointer,
		"blankCondition":  blankCondition,
		"stripStar":       stripStar,
	}
	headerT, err := template.New("header-client").Funcs(funcMap).Parse(headerTmpl)
	if err != nil {
		return nil, err
	}
	resourceT, err := template.New("resource-client").Funcs(funcMap).Parse(resourceTmpl)
	if err != nil {
		return nil, err
	}
	return &ClientWriter{
		headerTmpl:   headerT,
		resourceTmpl: resourceT,
	}, nil
}

// Write header text
func (c *ClientWriter) WriteHeader(pkg string, w io.Writer) error {
	return c.headerTmpl.Execute(w, pkg)
}

// Write resource header
func (c *ClientWriter) WriteResourceHeader(name string, w io.Writer) {
	fmt.Fprintf(w, "/******  %s ******/\n\n", name)
}

// Write separator between resources and data types
func (c *ClientWriter) WriteTypeSectionHeader(w io.Writer) {
	fmt.Fprintln(w, "\n/****** Parameter Data Types ******/\n\n")
}

// Write type declaration for resource action arguments
func (c *ClientWriter) WriteType(o *gen.ObjectDataType, w io.Writer) {
	var fields = make([]string, len(o.Fields))
	for i, f := range o.Fields {
		fields[i] = fmt.Sprintf("%s %s `json:\"%s,omitempty\"`", strings.Title(f.VarName),
			f.Signature(), f.Name)
	}
	decl := fmt.Sprintf("type %s struct {\n%s\n}", o.Name,
		strings.Join(fields, "\n\t"))
	fmt.Fprintf(w, "%s\n\n", decl)
}

// Write code for a resource
func (c *ClientWriter) WriteResource(resource *gen.Resource, w io.Writer) error {
	return c.resourceTmpl.Execute(w, resource)
}

/***** Format helpers *****/

// Inline templates

const headerTmpl = `
//************************************************************************//
//                     RightScale API client
//
{{comment "Generated " (now.Format "Jan 2, 2006 at 3:04pm (PST)")}}
// Command:
{{comment commandLine}}
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package {{.}}

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/rightscale/rsc/metadata"
)

// Helper function that merges optional parameters into payload
func mergeOptionals(params, options ApiParams) ApiParams {
	for name, value := range options {
		params[name] = value
	}
	return params
}

// Url resolver produces an action URL and HTTP method from its name and a given resource href.
// The algorithm consists of first extracting the variables from the href and then substituing them
// in the action path. If there are more than one action paths then the algorithm picks the one that
// can substitute the most variables.
type UrlResolver string

func (r *UrlResolver) Url(rName, aName string) (*metadata.ActionPath, error) {
	var res, ok = api_metadata[rName]
	if !ok {
		return nil, fmt.Errorf("No resource with name '%s'", rName)
	}
	var action *metadata.Action
	for _, a := range res.Actions {
		if a.Name == aName {
			action = a
			break
		}
	}
	if action == nil {
		return nil, fmt.Errorf("No action with name '%s' on %s", aName, rName)
	}
	var vars, err = res.ExtractVariables(string(*r))
	if err != nil {
		return nil, err
	}
	return action.Url(vars)
}

`

const resourceTmpl = `{{$resource := .}}{{define "ActionBody"}}` + actionBodyTmpl + `{{end}}
{{comment .Description}}{{if .Attributes}}
type {{.Name}} struct { {{range .Attributes}}
{{.FieldName}} {{.FieldType}} ` + "`" + `json:"{{.Name}},omitempty"` + "`" + `{{end}}
}
{{end}}
{{if .Actions}}
//===== Locator

// {{.Name}} resource locator, exposes resource actions.
type {{.Name}}Locator struct {
	UrlResolver
	api *Api15
}

// {{.Name}} resource locator factory
func (api *Api15) {{.Name}}Locator(href string) *{{.Name}}Locator {
	return &{{.Name}}Locator{UrlResolver(href), api}
}
//===== Actions
{{end}}{{range .Actions}}{{range .PathPatterns}}
// {{.HttpMethod}} {{.Path}}{{end}}
{{comment .Description}}
func (loc *{{$resource.Name}}Locator) {{.MethodName}}({{parameters .}}){{if .Return}} ({{.Return}},{{end}} error{{if .Return}}){{end}} {
	{{template "ActionBody" . }}
}
{{end}}
`

const actionBodyTmpl = `{{$action := .}}{{if .Return}}var res {{.Return}}
	{{end}}{{range .Params}}{{if and .Mandatory (blankCondition .VarName .Type)}}{{blankCondition .VarName .Type}}
		return {{if $action.Return}}res, {{end}}fmt.Errorf("{{.VarName}} is required")
	}
	{{end}}{{end}}{{/* end range .Params */}}var params = {{paramsAsPayload .Params}}
	var uri, err = loc.Url("{{$action.ResourceName}}", "{{$action.Name}}")
	if err != nil {
		return {{if $action.Return}}res, {{end}}err
	}
	var {{if .Return}}resp, {{else}}_, {{end}}err2 = loc.api.Dispatch(uri.HttpMethod, uri.Path, params)
	if err2 != nil {
		return {{if $action.Return}}res, {{end}}err2
	}
	{{if .ReturnLocation}}var location = resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &{{stripStar .Return}}{UrlResolver(location), loc.api}, nil
	}{{else if .Return}}defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, {{if not (isPointer .Return)}}&{{end}}res)
	return res, err4{{else}}return nil{{end}}`
