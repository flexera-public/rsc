package writers

import (
	"fmt"
	"io"
	"text/template"

	"github.com/rightscale/rsc/gen"
)

// AngularWriter struct exposes methods to generate the angular JS services code
type AngularWriter struct {
	angularTmpl *template.Template
}

// NewAngularWriter creates a new code writer that generates angular.js types.
func NewAngularWriter() (*AngularWriter, error) {
	funcMap := template.FuncMap{
		"comment":     comment,
		"commandLine": commandLine,
		"path":        path,
		"mandatory":   mandatory,
	}
	resourceT, err := template.New("resource-client").Funcs(funcMap).Parse(angularTmpl)
	if err != nil {
		return nil, err
	}
	return &AngularWriter{
		angularTmpl: resourceT,
	}, nil
}

// WriteResource writes the code for a resource.
func (c *AngularWriter) WriteResource(resource *gen.Resource, w io.Writer) error {
	return c.angularTmpl.Execute(w, resource)
}

// Path for given action, for now simplify and take first path in PathPatterns.
// In the future we may want to create one "action" in generated JS per path.
func path(a *gen.Action) string {
	pattern := a.PathPatterns[0]
	vars := pattern.Variables
	ivars := make([]interface{}, len(vars))
	for i, v := range vars {
		ivars[i] = interface{}(":" + v)
	}
	return fmt.Sprintf(pattern.Pattern, ivars...)
}

// Returns true if parameter with given name is mandatory
func mandatory(a gen.Action, param string) bool {
	for _, p := range a.Params {
		if p.Name == param {
			return p.Mandatory
		}
	}
	panic("praxisgen bug: Unknown param " + param + " for action " + a.Name)
}

// Inline templates

const angularTmpl = `
//************************************************************************//
//             RightScale {{.Name}} API Resource Client
//
// Generated with:
{{comment commandLine}}
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

app.service('{{.Name}}Client', function(ApiFactory) {
  return ApiFactory.buildClient({
    actions: { {{range .Actions}}{{$action := .}}
      {{.Name}}: {
        method: '{{(index .PathPatterns 0).HTTPMethod}}',
        path: '{{path .}}'{{if .QueryParamNames}},
        queryParams: {
        {{range .QueryParamNames}}  {{.}}: { {{if mandatory $action .}}required: true {{end}}},
        {{end}}}{{end}}{{if .PathParamNames}},
        pathParams: {
        {{range .PathParamNames}}  {{.}}: {},
	{{end}}}{{end}}{{if .PayloadParamNames}},
        payload: {
        {{range .PayloadParamNames}}  {{.}}: { {{if mandatory $action .}}required: true {{end}}},
        {{end}}}{{end}}
      },{{end}}
    },
    model: {
      name: '{{.Name}}',
      fields: [{{range .Attributes}}'{{.Name}}',{{end}}]
    }
  });
});
`
