package main

import (
	"fmt"
	"io"
	"strings"
	"text/template"
	"time"
)

// CmdsWriter struct exposes methods to generate the go API client command line tool
type CmdsWriter struct {
	headerTmpl       *template.Template
	resourceTmpl     *template.Template
	resActionMapTmpl *template.Template
}

// Commands writer factory
func NewCmdsWriter() (*CmdsWriter, error) {
	funcMap := template.FuncMap{
		"comment":       comment,
		"now":           time.Now,
		"join":          strings.Join,
		"commandLine":   commandLine,
		"toStringArray": toStringArray,
		"toHelp":        toHelp,
		"flagType":      flagType,
	}
	headerT, err := template.New("header-cmds").Funcs(funcMap).Parse(headerCmdsTmpl)
	if err != nil {
		return nil, err
	}
	resourceT, err := template.New("resource-cmds").Funcs(funcMap).Parse(resourceCmdsTmpl)
	if err != nil {
		return nil, err
	}
	resActionMapT, err := template.New("resActionMap-code").Funcs(funcMap).Parse(resActionMapTmpl)
	if err != nil {
		return nil, err
	}
	return &CmdsWriter{
		headerTmpl:       headerT,
		resourceTmpl:     resourceT,
		resActionMapTmpl: resActionMapT,
	}, nil
}

// Write header text
func (c *CmdsWriter) WriteHeader(w io.Writer) error {
	return c.headerTmpl.Execute(w, nil)
}

// Write commands
func (c *CmdsWriter) WriteCommands(d *ApiDescriptor, w io.Writer) error {
	var resources = make([]*Resource, len(d.ResourceNames))
	for i, n := range d.ResourceNames {
		resources[i] = d.Resources[n]
	}
	return c.resourceTmpl.Execute(w, resources)
}

// Write resource to actions map
func (c *CmdsWriter) WriteResourceActionMap(d *ApiDescriptor, w io.Writer) error {
	return c.resActionMapTmpl.Execute(w, d)
}

/** Format Helpers **/

// Convert []interface{} into []string
func toStringArray(a []interface{}) []string {
	var res = make([]string, len(a))
	for i, v := range a {
		res[i] = fmt.Sprintf("%v", v)
	}
	return res
}

// Convert description into flag help message
func toHelp(long string) string {
	lines := strings.Split(long, "\n")
	if len(lines) < 2 {
		return long
	}
	if lines[0][len(lines[0])-1] == '.' {
		return lines[0]
	}
	sentences := strings.Split(long, ".")
	if strings.Count(sentences[0], "\n") > 0 {
		sentence := strings.Split(sentences[0], "\n")
		return strings.Join(sentence, " ") + "..."
	}
	return sentences[0]
}

// Type of flag, one of "string", "[]string", "int" or "map"
func flagType(param *ActionParam) string {
	var path = param.QueryName
	if strings.HasSuffix(path, "[*]") {
		return "map"
	}
	if strings.HasSuffix(path, "[]") {
		return "[]string"
	}
	if strings.Contains(path, "[]") {
		var _, ok = param.Type.(*BasicDataType)
		if !ok {
			panic("Wooaat? an array with a non basic leaf???")
		}
		return "[]string"
	}
	if _, ok := param.Type.(*ArrayDataType); ok {
		return "[]string"
	}
	var b, ok = param.Type.(*BasicDataType)
	if !ok {
		panic("Wooaat? a non basic or array leaf???")
	}
	return string(*b)
}

//
const headerCmdsTmpl = `//************************************************************************//
//                     rsc - RightScale API 1.5 command line tool
//
{{comment "Generated " (now.Format "Jan 2, 2006 at 3:04pm (PST)")}}
// Command:
{{comment commandLine}}
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package rsapi15

import "regexp"

`

const resourceCmdsTmpl = `{{define "action"}}` + actionCmdTmpl + `{{end}}// API 1.5 resource commands
// Each command contains sub-commands for all resource actions
var commands = map[string]*ResourceCmd{ {{range .}}
	"{{.Name}}": &ResourceCmd{
		Name: "{{.Name}}",
		Description: ` + "`" + `{{toHelp .Description}}` + "`" + `,
		HrefRegexp: regexp.MustCompile(` + "`" + `{{.HrefRegexp}}` + "`" + `),
		Actions: []*ActionCmd{ {{range .CollectionActions}}
		{{template "action" .}}{{end}}{{range .ResourceActions}}
		{{template "action" .}}{{end}}
		},
	},{{end}}
}
`

const actionCmdTmpl = `&ActionCmd {
				Name: "{{.Name}}",
				Description: ` + "`" + `{{toHelp .Description}}` + "`" + `,
				Flags: []*ActionFlag{ {{range .LeafParams}}
					&ActionFlag{
						Name: "{{.QueryName}}",
						Description: ` + "`" + `{{toHelp .Description}}` + "`" + `,
						Type: "{{flagType .}}",
						Mandatory: {{.Mandatory}},
						NonBlank: {{.NonBlank}},{{if .Regexp}}
						Regexp: regexp.MustCompile("{{.Regexp}}"),{{end}}{{if .ValidValues}}
						ValidValues: []string{"{{join (toStringArray .ValidValues) "\", \""}}"},{{end}}
					},{{end}}
				},
			},
`

// Resource to actions map
const resActionMapTmpl = `// Action info struct
type ActionInfo struct {
	Name string        // Action name
	Description string // Action description
}

// Map resource names to action info
var resourceActions = map[string][]ActionInfo{
	{{range $name, $resource := .Resources}}"{{$name}}": []ActionInfo{
		{{range .CollectionActions}}ActionInfo{"{{.Name}}", ` + "`" + `{{toHelp .Description}}` + "`" + `},
		{{end}}{{range .ResourceActions}}ActionInfo{"{{.Name}}", ` + "`" + `{{toHelp .Description}}` + "`" + `},
		{{end}} },
	{{end}} }
`
