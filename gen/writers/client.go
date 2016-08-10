package writers

import (
	"fmt"
	"io"
	"strings"
	"text/template"

	"github.com/rightscale/rsc/gen"
)

// ClientWriter struct exposes methods to generate the go API client code
type ClientWriter struct {
	headerTmpl   *template.Template
	resourceTmpl *template.Template
}

// NewClientWriter is the  client writer factory.
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

// WriteHeader writes the header text.
func (c *ClientWriter) WriteHeader(pkg, version string, needTime, needJSON bool, w io.Writer) error {
	ctx := map[string]interface{}{
		"Pkg":        pkg,
		"APIVersion": version,
		"NeedTime":   needTime,
		"NeedJSON":   needJSON,
	}
	return c.headerTmpl.Execute(w, ctx)
}

// WriteResourceHeader writes the resource header.
func (c *ClientWriter) WriteResourceHeader(name string, w io.Writer) {
	fmt.Fprintf(w, "/******  %s ******/\n\n", name)
}

// WriteTypeSectionHeader writes the separator between resources and data types.
func (c *ClientWriter) WriteTypeSectionHeader(w io.Writer) {
	fmt.Fprintf(w, "\n/****** Parameter Data Types ******/\n\n\n")
}

// WriteType writest the type declaration for a resource action arguments.
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

// WriteResource writest the code for a resource.
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
	{{if .NeedJSON}}"encoding/json"
	{{end}}"fmt"
	"io/ioutil"
	{{if .NeedTime}}"time"
	{{end}}
	"github.com/rightscale/rsc/metadata"
	"github.com/rightscale/rsc/rsapi"
)

// API Version
const APIVersion = "{{.APIVersion}}"

// An Href contains the relative path to a resource or resource collection,
// e.g. "/api/servers/123" or "/api/servers".
type Href string

// ActionPath computes the path to the given resource action. For example given the href
// "/api/servers/123" calling ActionPath with resource "servers" and action "clone" returns the path
// "/api/servers/123/clone" and verb POST.
// The algorithm consists of extracting the variables from the href by looking up a matching
// pattern from the resource metadata. The variables are then substituted in the action path. 
// If there are more than one pattern that match the href then the algorithm picks the one that can
// substitute the most variables.
func (r *Href) ActionPath(rName, aName string) (*metadata.ActionPath, error) {
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
	return action.URL(vars)
}
`

const resourceTmpl = `{{$resource := .}}{{define "ActionBody"}}` + actionBodyTmpl + `{{end}}
{{comment .Description}}
type {{.Name}} struct { {{range .Attributes}}
{{.FieldName}} {{.FieldType}} ` + "`" + `json:"{{.Name}},omitempty"` + "`" + `{{end}}
}
{{if .LocatorFunc}}
// Locator returns a locator for the given resource
func (r *{{.Name}}) Locator(api *API) *{{.Name}}Locator {
	{{.LocatorFunc}}
}
{{end}}
{{if .Actions}}
//===== Locator

// {{.Name}}Locator exposes the {{.Name}} resource actions.
type {{.Name}}Locator struct {
	Href
	api *{{.ClientName}}
}

// {{.Name}}Locator builds a locator from the given href.
func (api *{{.ClientName}}) {{.Name}}Locator(href string) *{{.Name}}Locator {
	return &{{.Name}}Locator{Href(href), api}
}

//===== Actions
{{end}}{{range .Actions}}{{range .PathPatterns}}
// {{.HTTPMethod}} {{.Path}}{{end}}
//
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
	{{end}}{{end}}{{/* end range .Params */}}var params rsapi.APIParams{{paramsInitializer . 1 "params"}}
	var p rsapi.APIParams{{paramsInitializer . 2 "p"}}
	uri, err := loc.ActionPath("{{$action.ResourceName}}", "{{$action.Name}}")
	if err != nil {
		return {{if $action.Return}}res, {{end}}err
	}
	req, err := loc.api.BuildHTTPRequest(uri.HTTPMethod, uri.Path, APIVersion, params, p)
	if err != nil {
		return {{if $action.Return}}res, {{end}}err
	}
	resp, err := loc.api.PerformRequest(req)
	if err != nil {
		return {{if $action.Return}}res, {{end}}err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		sr := string(respBody)
		if sr != "" {
			sr = ": " + sr
		}
		return {{if $action.Return}}res, {{end}}fmt.Errorf("invalid response %s%s", resp.Status, sr)
	}
	{{if .ReturnLocation}}location := resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &{{stripStar .Return}}{Href(location), loc.api}, nil
	}{{else if .Return}}defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	{{if eq .Return "string"}}res = string(respBody)
	{{else}}err = json.Unmarshal(respBody, &res)
	{{end}}return res, err{{else}}return nil{{end}}`
