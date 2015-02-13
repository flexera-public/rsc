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
		"joinParamNames":  joinParamNames,
		"paramsAsPayload": paramsAsPayload,
		"isPointer":       isPointer,
		"isArray":         isArray,
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
func (c *CodeWriter) WriteType(t *ObjectDataType, w io.Writer) {
	fmt.Fprintf(w, "%s\n\n", t.Declaration())
}

// Write code for a resource
func (c *CodeWriter) WriteResource(resource *ResourceData, w io.Writer) error {
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
func parameters(a *ResourceAction) string {
	params := []string{}
	hasOptional := false
	for _, name := range a.ParamNames {
		param := a.AllParams[name]
		if param.Mandatory {
			params = append(params, fmt.Sprintf("%s %s", name, param.Signature()))
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
func joinParamNames(p []*ActionParam) string {
	paramNames := make([]string, len(p))
	for i, param := range p {
		paramNames[i] = param.Name
	}
	return strings.Join(paramNames, ", ")
}

// Create map out of parameter names
func paramsAsPayload(p []*ActionParam) string {
	fields := []string{}
	hasOptional := false
	for _, param := range p {
		if param.Mandatory {
			fields = append(fields, fmt.Sprintf("\"%s\": %s,", param.NativeName, param.Name))
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
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Href
type Href string

// Convenience type
type ApiParams map[string]interface{}

// Helper function that merges optional parameters into payload
func mergeOptionals(params, options ApiParams) ApiParams {
	for name, value := range options {
		params[name] = value
	}
	return params
}
`

const resourceTmpl = `{{define "actionBody"}}` + actionBodyTmpl + `{{end}}
{{comment .Description}}
type {{.Name}} struct { {{range .Attributes}}
{{.Name}} {{.Signature}} ` + "`" + `json:"{{.JsonName}},omitempty"` + "`" + `{{end}}
}
{{range .Actions}}
// {{.HttpMethod}} {{.Path}}
{{comment .Description}}
func (c *Client) {{.Name}}({{parameters .}}){{if .Return}} ({{.Return}},{{end}} error{{if .Return}}){{end}} {
	{{template "actionBody" . }}
}
{{end}}
`

const actionBodyTmpl = `{{$action := .}}{{if .Return}}var res {{.Return}}
	{{end}}{{range .PathParams}}if {{.Name}} == "" {
		return {{if $action.Return}}res, {{end}}fmt.Errorf("{{.Name}} cannot be blank")
	}
	{{end}}{{if .PayloadParams}}{{range .PayloadParams}}{{if and .Mandatory (.Type.BlankConditionExp .Name)}}{{.Type.BlankConditionExp .Name}}
		return {{if $action.Return}}res, {{end}}fmt.Errorf("{{.Name}} is required")
	}
	{{end}}{{end}}payload := {{paramsAsPayload .PayloadParams}}
	b, err := json.Marshal(payload)
	if err != nil {
		{{if .Return}}return res, err{{else}}return err{{end}}
	}
	{{else}}b := []byte{}{{end}}
	body := bytes.NewReader(b)
	req, err := http.NewRequest("{{.HttpMethod}}", c.endpoint+{{.UrlExp}}, body)
	if err != nil {
		return {{if .Return}}res, {{end}}err
	}
	{{if .QueryParams}}{{range .QueryParams}}{{if isArray .Signature}}if temp, ok := options["{{.Name}}"]; ok {
		for _, v := range temp.([]string) {
			req.URL.Query().Add("{{.NativeName}}", v)
		}
	}
	{{else}}if temp, ok := options["{{.Name}}"]; ok {
		req.URL.Query().Set("{{.NativeName}}", temp.(string))
	}
	{{end}}{{end}}{{end}}{{if .PayloadParams}}req.Header.Set("Content-Type", "application/json")
	{{end}}ctx := c.beforeRequest(req)
	resp, err := c.client.Do(req)
	c.afterRequest(resp, ctx)
	if err != nil {
		return {{if .Return}}res, {{end}}err
	}
	{{if eq .Return "Href"}}loc := resp.Header.Get("Location")
	if len(loc) == 0 {
		return res, fmt.Errorf("Missing location header in response")
	} else {
		return Href(loc), nil
	}{{else if .Return}}defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}
	err = json.Unmarshal(respBody, {{if not (isPointer .Return)}}&{{end}}res)
	return res, err{{else}}return nil{{end}}`
