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

// Data structure used to render resource code
// Note: we don't want to use hashes in structs used to generate the code to guarantee that code
// is always generated in the same order (to allow for diff etc.)
type ResourceData struct {
	Name        string
	Description string
	Actions     []*ResourceAction
	Attributes  []*ResourceAttribute
}

// Data structure used to render resource method
type ResourceAction struct {
	Name        string
	Description string
	HttpMethod  string
	Path        string
	Params      []*ActionParam
	Return      string
}

// Resource attributes, 'Type' is string representation of go type, e.g. "*time.Time"
type ResourceAttribute struct {
	Name      string
	Signature string
}

// Data structure used to render method params
type ActionParam struct {
	Name      string
	Signature string
}

// Code writer factory
func NewCodeWriter() (*CodeWriter, error) {
	funcMap := template.FuncMap{
		"comment":     comment,
		"now":         time.Now,
		"join":        strings.Join,
		"commandLine": commandLine,
		"joinParams":  joinParams,
	}
	headerT, err := template.New("header").Funcs(funcMap).Parse(headerTmpl)
	if err != nil {
		return nil, err
	}
	resourceT, err := template.New("resource").Funcs(funcMap).Parse(resourceTmpl)
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

// Write code for a resource
func (c *CodeWriter) WriteResource(resource *ResourceData, w io.Writer) error {
	return c.resourceTmpl.Execute(w, resource)
}

// Produce line comments by concatenating given strings and producing 80 characters long lines
// wrapped in "/*" "*/"
func comment(elems ...string) string {
	t := strings.Join(elems, "")
	return text.Indent(text.Wrap(t, 77), "// ")
}

// Serialize action parameters
func joinParams(p []*ActionParam) string {
	params := make([]string, len(p))
	for i, param := range p {
		params[i] = fmt.Sprintf("%s %s", param.Name, param.Signature)
	}
	return strings.Join(params, ", ")
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
{{comment "Command: '" commandLine "'"}}
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package main

// Href type
type Href string
`

const resourceTmpl = `
// == {{.Name}} ==

{{ comment .Description}}
type {{.Name}} struct { {{range .Attributes}}
	{{.Name}} {{.Signature}}{{end}}
}
{{range .Actions}}
// {{.HttpMethod}} {{.Path}}
{{comment .Description}}
func (c *Client) {{.Name}}({{joinParams .Params}}){{if .Return}} ({{.Return}},{{end}} error{{if .Return}}){{end}} {

}
func (c *Client) {{.Name}}G(p *Params){{if .Return}} ({{.Return}},{{end}} error{{if .Return}}){{end}} {

}
{{end}}
`
