package writers

import (
	"io"
	"strings"
	"text/template"

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
		"join":          strings.Join,
		"commandLine":   commandLine,
		"toStringArray": toStringArray,
		"flagType":      flagType,
		"location":      location,
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
	resources := make([]*gen.Resource, len(d.ResourceNames))
	for i, n := range d.ResourceNames {
		resources[i] = d.Resources[n]
	}
	return c.resourceTmpl.Execute(w, resources)
}

// Return code corresponding to param location
func location(p *gen.ActionParam) string {
	switch p.Location {
	case gen.PathParam:
		return "metadata.PathParam"
	case gen.QueryParam:
		return "metadata.QueryParam"
	case gen.PayloadParam:
		return "metadata.PayloadParam"
	default:
		return ""
	}
}

const headerMetadataTmpl = `//************************************************************************//
//                     rsc - RightScale API command line tool
//
// Generated with:
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
var GenMetadata = map[string]*metadata.Resource{ {{range .}}
	"{{.Name}}": &metadata.Resource{
		Name: "{{.Name}}",
		Description: ` + "`" + `{{.Description}}` + "`" + `,
		Actions: []*metadata.Action{ {{range .Actions}}
		{{template "action" .}}{{end}}
		},
	},{{end}}
}
`

const actionMetadataTmpl = `&metadata.Action {
				Name: "{{.Name}}",
				Description: ` + "`" + `{{.Description}}` + "`" + `,
				PathPatterns: []*metadata.PathPattern{ {{range .PathPatterns}}
					&metadata.PathPattern{
						HttpMethod: "{{.HttpMethod}}",
						Pattern: "{{.Pattern}}",
						Variables: []string{ {{if .Variables}}"{{join .Variables "\", \""}}"{{end}}},
						Regexp: regexp.MustCompile(` + "`" + `{{.Regexp}}` + "`" + `),
					},{{end}}
				},
				CommandFlags: []*metadata.ActionParam{ {{range .LeafParams}}
					&metadata.ActionParam{
						Name: "{{.QueryName}}",
						Description: ` + "`" + `{{.Description}}` + "`" + `,
						Type: "{{flagType .}}",
						Location: {{location .}},
						Mandatory: {{.Mandatory}},
						NonBlank: {{.NonBlank}},{{if .Regexp}}
						Regexp: regexp.MustCompile("{{.Regexp}}"),{{end}}{{if .ValidValues}}
						ValidValues: []string{"{{join (toStringArray .ValidValues) "\", \""}}"},{{end}}
					},{{end}}
				},{{if .Payload}}
				Payload: "{{.Payload.Name}}",{{end}}
				ApiParams: []*metadata.ActionParam{ {{range .Params}}
					&metadata.ActionParam{
						Name: "{{.QueryName}}",
						Description: ` + "`" + `{{.Description}}` + "`" + `,
						Type: "{{.Signature}}",
						Location: {{location .}},
						Mandatory: {{.Mandatory}},
						NonBlank: {{.NonBlank}},{{if .Regexp}}
						Regexp: regexp.MustCompile("{{.Regexp}}"),{{end}}{{if .ValidValues}}
						ValidValues: []string{"{{join (toStringArray .ValidValues) "\", \""}}"},{{end}}
					},{{end}}
				},
			},
`
