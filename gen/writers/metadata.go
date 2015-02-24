package writers

import (
	"io"
	"strings"
	"text/template"
	"time"

	"github.com/rightscale/rsc/gen"
)

// MetadataWriter struct exposes methods to generate the go API client command line tool
type MetadataWriter struct {
	headerTmpl   *template.Template
	resourceTmpl *template.Template
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
	return &MetadataWriter{
		headerTmpl:   headerT,
		resourceTmpl: resourceT,
	}, nil
}

// Write header text
func (c *MetadataWriter) WriteHeader(pkg string, w io.Writer) error {
	return c.headerTmpl.Execute(w, pkg)
}

// Write metadata
func (c *MetadataWriter) WriteMetadata(d *gen.ApiDescriptor, w io.Writer) error {
	var resources = make([]*gen.Resource, len(d.ResourceNames))
	for i, n := range d.ResourceNames {
		resources[i] = d.Resources[n]
	}
	return c.resourceTmpl.Execute(w, resources)
}

const headerMetadataTmpl = `//************************************************************************//
//                     rsc - RightScale API command line tool
//
{{comment "Generated " (now.Format "Jan 2, 2006 at 3:04pm (PST)")}}
// Command:
{{comment commandLine}}
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package {{.}}

import (
	"regexp"

	"github.com/rightscale/rsc/metadata"
)

`

const resourceMetadataTmpl = `{{define "action"}}` + actionMetadataTmpl + `{{end}}// Consists of a map of resource name to resource metadata.
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
				PathPatterns: []*metadata.PathPattern{ {{range .PathPatterns}}
					&metadata.PathPattern{
						HttpMethod: "{{.HttpMethod}}",
						Pattern: "{{.Path}}",
						Variables: []string{ {{if .Variables}}"{{join .Variables "\", \""}}"{{end}}},
						Regexp: regexp.MustCompile(` + "`" + `{{.Regexp}}` + "`" + `),
					},{{end}}
				},
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
				QueryParamNames: []string{ {{if .QueryParamNames}}"{{join .QueryParamNames ", "}}"{{end}}},
				PayloadParamNames: []string{ {{if .PayloadParamNames}}"{{join .PayloadParamNames "\", \""}}"{{end}}},
			},
`
