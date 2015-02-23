package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/rightscale/rsc/generators/text"
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
		"toVerb":          toVerb,
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
func (c *ClientWriter) WriteHeader(w io.Writer) error {
	return c.headerTmpl.Execute(w, nil)
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
func (c *ClientWriter) WriteType(o *ObjectDataType, w io.Writer) {
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
func (c *ClientWriter) WriteResource(resource *Resource, w io.Writer) error {
	return c.resourceTmpl.Execute(w, resource)
}

/***** Format helpers *****/

// Produce line comments by concatenating given strings and producing 80 characters long lines
// starting with "//"
func comment(elems ...string) string {
	t := strings.Join(elems, "")
	return text.Indent(t, "// ")
}

// Serialize action parameters
func parameters(a *Action) string {
	var m = a.MandatoryParams()
	var hasOptional = a.HasOptionalParams()
	var countParams = len(m)
	if hasOptional {
		countParams += 1
	}
	var params = make([]string, countParams)
	for i, param := range m {
		params[i] = fmt.Sprintf("%s %s", param.VarName, param.Signature())
	}
	if hasOptional {
		params[countParams-1] = "options ApiParams"
	}

	return strings.Join(params, ", ")
}

// Create map out of parameter names
func paramsAsPayload(p []*ActionParam) string {
	if len(p) == 0 {
		return "map[string]interface{}{}"
	}
	fields := []string{}
	hasOptional := false
	for _, param := range p {
		if param.Mandatory {
			fields = append(fields, fmt.Sprintf("\"%s\": %s,", param.Name, param.VarName))
		} else {
			hasOptional = true
		}
	}
	mandatory := fmt.Sprintf("ApiParams{\n%s\n}", strings.Join(fields, "\n\t"))
	if !hasOptional {
		return mandatory
	}
	return fmt.Sprintf("mergeOptionals(%s, options)", mandatory)
}

// Return true if signature contains pointer, false otherwise
func isPointer(sig string) bool {
	return strings.HasPrefix(sig, "*")
}

// Command line used to run tool
func commandLine() string {
	return fmt.Sprintf("$ api15gen %s", strings.Join(os.Args[1:], " "))
}

// Code that checks whether variable with given name and type contains a blank value (empty string,
// empty array or empy map).
// Return empty string if type of variable cannot produce blank values
func blankCondition(name string, t DataType) (blank string) {
	switch actual := t.(type) {
	case *BasicDataType:
		if *actual == "string" {
			blank = fmt.Sprintf("if %s == \"\" {", name)
		}
	case *ArrayDataType:
		blank = fmt.Sprintf("if len(%s) == 0 {", name)
	case *ObjectDataType:
		blank = fmt.Sprintf("if %s == nil {", name)
	case *EnumerableDataType:
		blank = fmt.Sprintf("if len(%s) == 0 {", name)
	}
	return
}

// GET => Get
func toVerb(text string) (res string) {
	res = strings.ToUpper(string(text[0])) + strings.ToLower(text[1:])
	if text == "GET" || text == "POST" {
		res += "Raw"
	}
	return
}

// *ServerArrayLocator => ServerArrayLocator
func stripStar(text string) string {
	if text[0] == '*' {
		return text[1:]
	}
	return text
}

// Inline templates

const headerTmpl = `
//************************************************************************//
//                     RightScale API 1.5 go client
//
{{comment "Generated " (now.Format "Jan 2, 2006 at 3:04pm (PST)")}}
// Command:
{{comment commandLine}}
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package rsapi15

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

// Url resolver produces an action URL from its name and a given resource href.
// The algorithm consists of first extracting the variables from the href and then substituing them
// in the action path. If there are more than one action paths then the algorithm picks the one that
// can substitute the most variables.
type UrlResolver string

func (r *UrlResolver) Url(rName, aName string) (string, error) {
	var res, ok = api_metadata[rName]
	if !ok {
		return "", fmt.Errorf("No resource with name '%s'", rName)
	}
	var action *metadata.Action
	for _, a := range res.Actions {
		if a.Name == aName {
			action = a
			break
		}
	}
	if action == nil {
		return "", fmt.Errorf("No action with name '%s' on %s", aName, rName)
	}
	var vars, err = res.ExtractVariables(string(*r))
	if err != nil {
		return "", err
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
{{end}}{{range .Actions}}{{$httpMethod := .HttpMethod}}{{range .PathPatterns}}
// {{$httpMethod}} {{.Path}}{{end}}
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
	{{end}}{{end}}{{/* end range .Params */}}{{if not (eq .HttpMethod "DELETE")}}var params = {{paramsAsPayload .Params}}{{end}}
	var uri, err = loc.Url("{{$action.ResourceName}}", "{{$action.Name}}")
	if err != nil {
		return {{if $action.Return}}res, {{end}}err
	}
	var {{if .HasResponse}}{{if .Return}}resp, {{else}}_, {{end}}{{end}}err2 = loc.api.{{toVerb .HttpMethod}}(uri{{if not (eq .HttpMethod "DELETE")}}, params{{end}})
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
