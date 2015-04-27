package writers

import (
	// "fmt"
	"io"	
	"strings"
	"text/template"
  "encoding/json"
  "fmt"

	"github.com/rightscale/rsc/gen"
	"github.com/ryanoleary/go-restful/swagger"
)

// SwaggerWriter struct exposes methods to generate the angular JS services code
type SwaggerWriter struct {
	swaggerTmpl *template.Template
	swaggerHeaderTmpl *template.Template
	swaggerFooterTmpl *template.Template
}

// Swagger writer factory
func NewSwaggerWriter() (*SwaggerWriter, error) {
	funcMap := template.FuncMap{
		"comment":     comment,
		"commandLine": commandLine,
		"inValue":     inValue,
		"jsonEncode":  jsonEncode,
		"toLowerVerb": toLowerVerb,
		"swaggerPath": swaggerPath,
	}
	resourceT, err := template.New("resource-client").Funcs(funcMap).Parse(swaggerTmpl)
	if err != nil {
		return nil, err
	}
	resourceTH, err := template.New("resource-client").Funcs(funcMap).Parse(swaggerHeaderTmpl)
	if err != nil {
		return nil, err
	}
	resourceTF, err := template.New("resource-client").Funcs(funcMap).Parse(swaggerFooterTmpl)
	if err != nil {
		return nil, err
	}
	return &SwaggerWriter{
		swaggerTmpl: resourceT,
		swaggerHeaderTmpl: resourceTH,
		swaggerFooterTmpl: resourceTF,
	}, nil
}


// A path pattern represents a possible path for a given action.
type SwaggerPath struct {
	SwaggerPattern 		*gen.PathPattern  
	SwaggerResource 	*gen.Resource  
	SwaggerActions   	[]*gen.Action  
}

type ApiDescriptorPlus struct {
	Api 		*gen.ApiDescriptor
	ApiHost	string
	ApiRoot string
}

// Write code for a resource
func (c *SwaggerWriter) WriteApi(api *gen.ApiDescriptor, w io.Writer, apiHost string, apiRoot string) error {

	td := ApiDescriptorPlus{api, apiHost, apiRoot}

	err := c.swaggerHeaderTmpl.Execute(w, td)
	if err != nil {
		return err
	}
	paths := make(map[string]*SwaggerPath)

	for _, resource := range api.ResourceNames {
		for _, action := range api.Resources[resource].Actions {

			if _, ok := paths[action.PathPatterns[0].Path]; ok {
				paths[action.PathPatterns[0].Path].SwaggerActions = append(paths[action.PathPatterns[0].Path].SwaggerActions, action)
			} else {
				aArray := []*gen.Action{ action }
				p := &SwaggerPath{ action.PathPatterns[0], api.Resources[resource], aArray }
				paths[action.PathPatterns[0].Path] = p
			}
		}
	}


	// for i, path := range paths {
		err = c.swaggerTmpl.Execute(w, paths )
		if err != nil {
			return err
		}		
	// }
	err = c.swaggerFooterTmpl.Execute(w, api )
	if err != nil {
		return err
	}

	return nil

}

// Required. The location of the parameter. Possible values are "query", "header", "path", "formData" or "body".
// 	Location    int           // 0 for path param, 1 for query param, 2 for payload param

func inValue(l int) string {
	in := "query" // default to query, cause why not

	switch l {
	case 0:
		in = "path"
	case 1:
		in = "query"
// TODO Swagger can only have one body param, but we want to test this so we'll hack it
	case 2:
		in = "body"
		in = "query"
	}

	return in
}

func jsonEncode(s string) string {
	e, err := json.Marshal(s)
	if err != nil {
		//return err
	}
	return string(e[:])
}

// GET => get
func toLowerVerb(s string) string {
	res := strings.ToLower(s)
	return res
}

func swaggerPath(p gen.PathPattern) string {
  vars := make([]interface{}, len(p.Variables))
  for i, v := range p.Variables { vars[i] = "{" + v + "}"}

  return fmt.Sprintf(p.Pattern, vars...)
}

// Inline templates

const swaggerHeaderTmpl = `{
  "swagger": "2.0",
  "info": {
    "description": "This is a TEST API",
    "version": "{{.Api.Version}}",
    "title": "RightScale TEST GO",
    "contact": {
      "email": "support@rightscale.com"
    }
  },
  "host": "{{.ApiHost}}",
  "basePath": "{{.ApiRoot}}",
  "schemes":[ "https" ], 
  "tags":[ {{ range $ri, $r := .Api.Resources }}{{if $ri}},{{end}}
  	{
    	"name": "{{$r.Name}}",
    	"description": "",
    	"ri": {{$ri}}
  	}{{end}}
  ],
  "paths":{
`
    	// "description": {{ jsonEncode $r.Description}}
// {{ range  $i, $e := . }}{{ if $i }}, {{ end }}{{ $e }}{{ end }}
const swaggerFooterTmpl = `
  }
}
`

const swaggerTmpl = `		{{range $k, $sp := .}}{{if $k}},{{end}}"{{ swaggerPath $sp.SwaggerPattern}}": { {{$rName := $sp.SwaggerResource.Name}}{{range $i, $a := $sp.SwaggerActions}}{{if $i}},{{end}}
  		"{{ toLowerVerb (index $a.PathPatterns 0).HttpMethod}}": {
  			"tags": ["{{$rName}}"],
  			"summary": "{{$a.Name}}",
  			"operationId": "{{$a.Name}}{{$rName}}",
  			"consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
  			"description" : {{ jsonEncode $a.Description}},
  			"parameters": [ {{range $j, $p := $a.Params}}{{if $j}},{{end}}
  			  {
  			  	"name": "{{$p.Name}}",
  			  	"in": "{{ inValue $p.Location }}",
  			  	"description": {{ jsonEncode $p.Description}},
  			  	"required": {{$p.Mandatory}},{{if ne (inValue $p.Location) "body"}}
  			  	"type": "string"{{end}}{{if eq (inValue $p.Location) "body"}}
  			  	"schema": {}{{end}}
  			  }{{end}}{{if (len $a.Params) gt 0}},{{end}}{{ $a0 := (index $sp.SwaggerActions 0)}}{{ range $m, $p := (index $a0.PathPatterns 0).Variables }}{{if $m}},{{end}}
	  			{
	  				"name": "{{.}}",
	  				"in": "path",
	  				"required": true,
	  				"type": "string"
	  			}{{end}}
  			],
  			"responses": {"default": {"description": "FAKE"}}
  		}{{end}},
  		"parameters": [{{ $a0 := (index $sp.SwaggerActions 0)}}{{ range $m, $p := (index $a0.PathPatterns 0).Variables }}{{if $m}},{{end}}
  			{
  				"name": "{{.}}",
  				"in": "path",
  				"required": true,
  				"type": "string"
  			}
  		{{end}}]
    }{{end}}
`
// 	Name              string         // Action name, e.g. "create", "multi_terminate"
// 	MethodName        string         // Go method name, e.g. "Create", "MultiTerminate"
// 	Description       string         // Action description
// 	ResourceName      string         // Name of resource that contains this action
// 	PathPatterns      []*PathPattern // Action path patterns
// 	Payload           DataType       // Payload type if any
// 	Params            []*ActionParam // Action method parameters
// 	LeafParams        []*ActionParam // Action parameter leaves (for command line)
// 	Return            string         // Type of method results, e.g. "*ServerArray"
// 	ReturnLocation    bool           // Whether API returns a location header
// 	PathParamNames    []string       // Name of path parameters if any (e.g. :id in /clouds/:id)
// 	QueryParamNames   []string       // Name of query string parameters if any
// 	PayloadParamNames []string 

// 	type ActionParam struct {
// 	Name        string        // Name of parameter
// 	QueryName   string        // Query string style parameter name
// 	Description string        // Description of parameter
// 	VarName     string        // go variable name
// 	Location    int           // 0 for path param, 1 for query param, 2 for payload param
// 	Type        DataType      // Type of parameter
// 	Mandatory   bool          // Whether parameter is mandatory
// 	NonBlank    bool          // Whether parameter must not be blank if provided (string or array)
// 	Regexp      string        // Regular expression string parameter must match
// 	ValidValues []interface{} // Allowed values (if not empty)
// 	Min         int           // Minimum value (int)
// 	Max         int           // Maximum value (int)
// }

//const angularTmpl = `
// //************************************************************************//
// //             RightScale {{.Name}} API Resource Client
// //
// // Generated with:
// {{comment commandLine}}
// //
// // The content of this file is auto-generated, DO NOT MODIFY
// //************************************************************************//

// app.service('{{.Name}}Client', function(ApiFactory) {
//   return ApiFactory.buildClient({
//     actions: { {{range .Actions}}{{$action := .}}
//       {{.Name}}: {
//         method: '{{(index .PathPatterns 0).HttpMethod}}',
//         path: '{{path .}}'{{if .QueryParamNames}},
//         queryParams: {
//         {{range .QueryParamNames}}  {{.}}: { {{if mandatory $action .}}required: true {{end}}},
//         {{end}}}{{end}}{{if .PathParamNames}},
//         pathParams: {
//         {{range .PathParamNames}}  {{.}}: {},
// 	{{end}}}{{end}}{{if .PayloadParamNames}},
//         payload: {
//         {{range .PayloadParamNames}}  {{.}}: { {{if mandatory $action .}}required: true {{end}}},
//         {{end}}}{{end}}
//       },{{end}}
//     },
//     model: {
//       name: '{{.Name}}',
//       fields: [{{range .Attributes}}'{{.Name}}',{{end}}]
//     }
//   });
// });
// `
