package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/rightscale/rsclient/generators/text"
)

// CodeWriter struct exposes methods to generate the go API client code
type CodeWriter struct {
	headerTmpl   *template.Template
	resourceTmpl *template.Template
}

// Code writer factory
func NewCodeWriter() (*CodeWriter, error) {
	funcMap := template.FuncMap{
		"comment":         comment,
		"now":             time.Now,
		"join":            strings.Join,
		"commandLine":     commandLine,
		"parameters":      parameters,
		"joinParams":      joinParams,
		"paramsAsPayload": paramsAsPayload,
		"isPointer":       isPointer,
		"isArray":         isArray,
		"blankCondition":  blankCondition,
		"toVerb":          toVerb,
		"stripStar":       stripStar,
	}
	headerT, err := template.New("header-code").Funcs(funcMap).Parse(headerTmpl)
	if err != nil {
		return nil, err
	}
	resourceT, err := template.New("resource-code").Funcs(funcMap).Parse(resourceTmpl)
	if err != nil {
		return nil, err
	}
	return &CodeWriter{
		headerTmpl:   headerT,
		resourceTmpl: resourceT,
	}, nil
}

// Write header text
func (c *CodeWriter) WriteHeader(w io.Writer) error {
	return c.headerTmpl.Execute(w, nil)
}

// Write resource header
func (c *CodeWriter) WriteResourceHeader(name string, w io.Writer) {
	fmt.Fprintf(w, "/******  %s ******/\n\n", name)
}

// Write separator between resources and data types
func (c *CodeWriter) WriteTypeSectionHeader(w io.Writer) {
	fmt.Fprintln(w, "\n/****** Parameter Data Types ******/\n\n")
}

// Write type declaration for resource action arguments
func (c *CodeWriter) WriteType(o *ObjectDataType, w io.Writer) {
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
func (c *CodeWriter) WriteResource(resource *Resource, w io.Writer) error {
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
	params := []string{}
	hasOptional := false
	for _, param := range a.Params {
		if param.Mandatory {
			params = append(params, fmt.Sprintf("%s %s", param.VarName, param.Signature()))
		} else {
			hasOptional = true
		}
	}
	if hasOptional {
		params = append(params, "options ApiParams")
	}

	return strings.Join(params, ", ")
}

// Serialize action parameter names
func joinParams(p []*ActionParam) string {
	var params = make([]string, len(p))
	for i, param := range p {
		params[i] = fmt.Sprintf("%s %s", param.Name, param.Signature())
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

// Return true if signature contains an array, false otherwise
func isArray(sig string) bool {
	return strings.HasPrefix(sig, "[]")
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
	"net/url"
)

// Helper function that merges optional parameters into payload
func mergeOptionals(params, options ApiParams) ApiParams {
	for name, value := range options {
		params[name] = value
	}
	return params
}

`

const resourceTmpl = `{{$resource := .}}{{define "ActionBody"}}` + actionBodyTmpl + `{{end}}
{{comment .Description}}{{if .Attributes}}
type {{.Name}} struct { {{range .Attributes}}
{{.FieldName}} {{.FieldType}} ` + "`" + `json:"{{.Name}},omitempty"` + "`" + `{{end}}
}
{{end}}{{if .CollectionActions}}
// {{.Name}} collection locator, exposes collection actions.
type {{.CollectionName}}Locator struct {
	api *Api15
	Href string
}

// {{.Name}} collection locator factory
func (api *Api15) {{.CollectionName}}Locator(href string) *{{.CollectionName}}Locator {
	return &{{.CollectionName}}Locator{api, href}
}
{{end}}{{if .ResourceActions}}
// {{.Name}} resource locator, exposes resource actions.
type {{.Name}}Locator struct {
	api *Api15
	Href string
}

// {{.Name}} resource locator factory
func (api *Api15) {{.Name}}Locator(href string) *{{.Name}}Locator {
	return &{{.Name}}Locator{api, href}
}
{{end}}{{if .CollectionActions}}
//===== Collection actions
{{end}}{{range .CollectionActions}}{{$httpMethod := .HttpMethod}}{{range .Paths}}
// {{$httpMethod}} {{.}}{{end}}
{{comment .Description}}
func (loc *{{$resource.CollectionName}}Locator) {{.MethodName}}({{parameters .}}){{if .Return}} ({{.Return}},{{end}} error{{if .Return}}){{end}} {
	{{template "ActionBody" . }}
}
{{end}}{{if .ResourceActions}}
//===== Resource actions
{{end}}{{range .ResourceActions}}{{$httpMethod := .HttpMethod}}{{range .Paths}}
// {{$httpMethod}} {{.}}{{end}}
{{comment .Description}}
func (loc *{{$resource.Name}}Locator) {{.MethodName}}({{parameters .}}){{if .Return}} ({{.Return}},{{end}} error{{if .Return}}){{end}} {
	{{template "ActionBody" . }}
}
{{end}}
`

const actionBodyTmpl = `{{$action := .}}{{if .Return}}var res {{.Return}}
	{{end}}{{if .HasPayload}}{{range .Params}}{{if and .Mandatory (blankCondition .VarName .Type)}}{{blankCondition .VarName .Type}}
		return {{if $action.Return}}res, {{end}}fmt.Errorf("{{.VarName}} is required")
	}
	{{end}}{{end}}{{/* end range .Params */}}var payload = {{paramsAsPayload .Params}}
	{{end}}{{/* end if .HasPayload */}}var href = loc.Href{{with $suffix := .Suffix}}{{if $suffix}}+"{{$suffix}}"{{end}}{{end}}
	{{if not .HasPayload}}{{if .Params}}var u, err = url.Parse(href)
	if err != nil {
		return {{if .Return}}res, {{end}}fmt.Errorf("Invalid href '%s' - %s", href, err.Error())
	}
	{{range .Params}}{{if .Mandatory}}u.Query().Set("{{.Name}}", {{.VarName}})
	{{else}}{{if isArray .Signature}}if temp, ok := options["{{.Name}}"]; ok {
		for _, v := range temp.([]string) {
			u.Query().Add("{{.Name}}", v)
		}
	}
	{{else}}if temp, ok := options["{{.Name}}"]; ok {
		u.Query().Set("{{.Name}}", temp.(string))
	}
	{{end}}href = u.String()
	{{end}}{{end}}{{end}}{{end}}{{/* end not .HasPayload */}}var {{if .HasResponse}}{{if .Return}}resp, {{else}}_, {{end}}{{end}}err2 = loc.api.{{toVerb .HttpMethod}}(href{{if .HasPayload}}, payload{{end}})
	if err2 != nil {
		return {{if $action.Return}}res, {{end}}err2
	}
	{{if eq $action.Name "Create"}}var location = resp.Header.Get("Location")
	if len(location) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return &{{stripStar .Return}}{loc.api, location}, nil
	}{{else if .Return}}defer resp.Body.Close()
	var respBody, err3 = ioutil.ReadAll(resp.Body)
	if err3 != nil {
		return res, err3
	}
	var err4 = json.Unmarshal(respBody, {{if not (isPointer .Return)}}&{{end}}res)
	return res, err4{{else}}return nil{{end}}`
