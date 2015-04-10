package writers

import (
	"fmt"
	"io"
	"strings"
	"text/template"

	"gopkg.in/rightscale/rsc.v1-unstable/gen" // import "gopkg.in/rightscale/rsc.v1-unstable/gen"
)

// ClientWriter struct exposes methods to generate the go API client code
type ClientWriter struct {
	headerTmpl   *template.Template
	resourceTmpl *template.Template
}

// Client writer factory
func NewClientWriter() (*ClientWriter, error) {
	funcMap := template.FuncMap{
		"comment":           comment,
		"commandLine":       commandLine,
		"parameters":        parameters,
		"paramsInitializer": paramsInitializer,
		"blankCondition":    blankCondition,
		"stripStar":         stripStar,
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
func (c *ClientWriter) WriteHeader(pkg string, needTime, needJson bool, w io.Writer) error {
	ctx := map[string]interface{}{"Pkg": pkg, "NeedTime": needTime, "NeedJson": needJson}
	return c.headerTmpl.Execute(w, ctx)
}

// Write resource header
func (c *ClientWriter) WriteResourceHeader(name string, w io.Writer) {
	fmt.Fprintf(w, "/******  %s ******/\n\n", name)
}

// Write separator between resources and data types
func (c *ClientWriter) WriteTypeSectionHeader(w io.Writer) {
	fmt.Fprintf(w, "\n/****** Parameter Data Types ******/\n\n\n")
}

// Write type declaration for resource action arguments
func (c *ClientWriter) WriteType(o *gen.ObjectDataType, w io.Writer) {
	fields := make([]string, len(o.Fields))
	for i, f := range o.Fields {
		fields[i] = fmt.Sprintf("%s %s `json:\"%s,omitempty\"`", strings.Title(f.VarName),
			f.Signature(), f.Name)
	}
	decl := fmt.Sprintf("type %s struct {\n%s\n}", o.TypeName,
		strings.Join(fields, "\n\t"))
	fmt.Fprintf(w, "%s\n\n", decl)
}

// Write code for a resource
func (c *ClientWriter) WriteResource(resource *gen.Resource, w io.Writer) error {
	return c.resourceTmpl.Execute(w, resource)
}

// Inline templates

const headerTmpl = `
//************************************************************************//
//                     RightScale API client
//
// Generated with:
{{comment commandLine}}
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package {{.Pkg}}

import (
	{{if .NeedJson}}"encoding/json"
	{{end}}"fmt"
	"io/ioutil"
	{{if .NeedTime}}"time"
	{{end}}
	"github.com/rightscale/rsc/metadata"
	"github.com/rightscale/rsc/rsapi"
)

// Url resolver produces an action URL and HTTP method from its name and a given resource href.
// The algorithm consists of first extracting the variables from the href and then substituing them
// in the action path. If there are more than one action paths then the algorithm picks the one that
// can substitute the most variables.
type UrlResolver string

func (r *UrlResolver) Url(rName, aName string) (*metadata.ActionPath, error) {
	res, ok := GenMetadata[rName]
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
	vars, err := res.ExtractVariables(string(*r))
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
	api *{{.ClientName}}
}

// {{.Name}} resource locator factory
func (api *{{.ClientName}}) {{.Name}}Locator(href string) *{{.Name}}Locator {
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
	{{end}}{{end}}{{/* end range .Params */}}var queryParams rsapi.ApiParams{{paramsInitializer . 1 "queryParams"}}
	var payloadParams rsapi.ApiParams{{paramsInitializer . 2 "payloadParams"}}
	uri, err := loc.Url("{{$action.ResourceName}}", "{{$action.Name}}")
	if err != nil {
		return {{if $action.Return}}res, {{end}}err
	}
	{{if .Return}}resp, {{else}}_, {{end}}err {{if .Return}}:{{end}}= loc.api.Dispatch(uri.HttpMethod, uri.Path, queryParams, payloadParams)
	if err != nil {
		return {{if $action.Return}}res, {{end}}err
	}
	{{if .ReturnLocation}}location := resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &{{stripStar .Return}}{UrlResolver(location), loc.api}, nil
	}{{else if .Return}}defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	{{if eq .Return "string"}}res = string(respBody)
	{{else}}err = json.Unmarshal(respBody, &res)
	{{end}}return res, err{{else}}return nil{{end}}`
