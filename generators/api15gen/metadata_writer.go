package main

import (
	"fmt"
	"io"
	"strings"
	"text/template"
	"time"
)

// MetadataWriter struct exposes methods to generate the go API client command line tool
type MetadataWriter struct {
	headerTmpl       *template.Template
	resourceTmpl     *template.Template
	resActionMapTmpl *template.Template
}

// Commands writer factory
func NewMetadataWriter() (*MetadataWriter, error) {
	funcMap := template.FuncMap{
		"comment":       comment,
		"now":           time.Now,
		"join":          strings.Join,
		"commandLine":   commandLine,
		"toStringArray": toStringArray,
		"toHelp":        toHelp,
		"flagType":      flagType,
	}
	headerT, err := template.New("header-metadata").Funcs(funcMap).Parse(headerMetadataTmpl)
	if err != nil {
		return nil, err
	}
	resourceT, err := template.New("resource-metadata").Funcs(funcMap).Parse(resourceMetadataTmpl)
	if err != nil {
		return nil, err
	}
	resActionMapT, err := template.New("resActionMap-metadata").Funcs(funcMap).Parse(resActionMapTmpl)
	if err != nil {
		return nil, err
	}
	return &MetadataWriter{
		headerTmpl:       headerT,
		resourceTmpl:     resourceT,
		resActionMapTmpl: resActionMapT,
	}, nil
}

// Write header text
func (c *MetadataWriter) WriteHeader(w io.Writer) error {
	return c.headerTmpl.Execute(w, nil)
}

// Write commands
func (c *MetadataWriter) WriteMetadata(d *ApiDescriptor, w io.Writer) error {
	var resources = make([]*Resource, len(d.ResourceNames))
	for i, n := range d.ResourceNames {
		resources[i] = d.Resources[n]
	}
	return c.resourceTmpl.Execute(w, resources)
}

// Write resource to actions map
func (c *MetadataWriter) WriteResourceActionMap(d *ApiDescriptor, w io.Writer) error {
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
	} else if _, ok := param.Type.(*EnumerableDataType); ok {
		return "map"
	}
	var b, ok = param.Type.(*BasicDataType)
	if !ok {
		panic("Wooaat? a object leaf???")
	}
	return string(*b)
}

//
const headerMetadataTmpl = `//************************************************************************//
//                     rsc - RightScale API 1.5 command line tool
//
{{comment "Generated " (now.Format "Jan 2, 2006 at 3:04pm (PST)")}}
// Command:
{{comment commandLine}}
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package rsapi15

import (
	"regexp"

	"github.com/rightscale/rsc/metadata"
)

`

const resourceMetadataTmpl = `{{define "action"}}` + actionMetadataTmpl + `{{end}}// API 1.5 resource commands
// API metadata, consists of a map of resource name to resource metadata.
var api_metadata = map[string]*metadata.Resource{ {{range .}}
	"{{.Name}}": &metadata.Resource{
		Name: "{{.Name}}",
		Description: ` + "`" + `{{toHelp .Description}}` + "`" + `,
		Actions: []*metadata.Action{ {{range .Actions}}
		{{template "action" .}}{{end}}
		},
	},{{end}}
}
`

const actionMetadataTmpl = `&metadata.Action {
				Name: "{{.Name}}",
				Description: ` + "`" + `{{toHelp .Description}}` + "`" + `,
				Params: []*metadata.ActionParam{ {{range .LeafParams}}
					&metadata.ActionParam{
						Name: "{{.QueryName}}",
						Description: ` + "`" + `{{toHelp .Description}}` + "`" + `,
						Type: "{{flagType .}}",
						Mandatory: {{.Mandatory}},
						NonBlank: {{.NonBlank}},{{if .Regexp}}
						Regexp: regexp.MustCompile("{{.Regexp}}"),{{end}}{{if .ValidValues}}
						ValidValues: []string{"{{join (toStringArray .ValidValues) "\", \""}}"},{{end}}
					},{{end}}
				},
				PathPatterns: []*metadata.PathPattern{ {{range .PathPatterns}}
					&metadata.PathPattern{
						Pattern: "{{.Pattern}}",
						Variables: []string{"{{join .Variables "\", \""}}"},
						Regexp: regexp.MustCompile(` + "`" + `{{.Regexp}}` + "`" + `),
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
		{{range .Actions}}ActionInfo{"{{.Name}}", ` + "`" + `{{toHelp .Description}}` + "`" + `},
		{{end}}},
	{{end}} }
`
