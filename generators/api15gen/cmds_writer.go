package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"text/template"
	"time"

	"bitbucket.org/pkg/inflect"
)

// CmdsWriter struct exposes methods to generate the go command line tool code
type CmdsWriter struct {
	HeaderTmpl   *template.Template
	ResourceTmpl *template.Template
	Resources    map[string]*ResourceCmd
}

// Resource command
type ResourceCmd struct {
	Name        string       // Name of resource, e.g. ServerArray
	Description string       // Description of resource if any
	ActionCmds  []*ActionCmd // Resource subcommands
}

type ActionCmd struct {
	Name              string // e.g. "createServerArray"
	Description       string
	RunnerFields      []*RunnerField // Runner struct fields
	ParamFields       []*ParamField  // Details on action parameters ordered by action method parameter order
	HasOptionalParams bool           // Whether action takes optional parameters
	Flags             []*ActionFlag
}

type ActionFlag struct {
	Name        string       // Flag name
	Type        string       // Flag kind i.e. "Flag or "FlagPattern"
	Description string       // Flag description if any
	IsRequired  bool         // Whether flag is required
	RunnerField *RunnerField // Field used to save flag value
	TypeVar     string       // "StringVar" or "StringsVar" or "IntVar" or "IntsVar"
}

// Runner struct field info
type RunnerField struct {
	Name        string // e.g. "serverArrayJson"
	Type        string // e.g. "[]string"
	CaptureName string // Name of field used to store flag capture if any (only applies to arrays and enums)
	IsRequired  bool   // Whether parameter is required
}

type ParamKind int

const (
	kBasic  ParamKind = iota // Param field is int or string
	kArray                   // Param field is array
	kObject                  // Param field is a struct
	kEnum                    // Param field is a enumerable (map[string]string)
)

// Action parameter info
type ParamField struct {
	Name          string        // e.g. "serverArray"
	NativeName    string        // e.g. "server_array"
	Initializer   string        // Code that initializes child structures recursively, e.g. "serverArray.Instance = new(serverArrayInstanceParam)"
	Type          string        // e.g. "ServerArrayParam"
	JsonFieldName string        // e.g. Name of runner field containing data structure JSON "serverArrayJson"
	BasicFields   []*BasicField // Non array and non enum fields with path ordered by Name
	ArrayFields   []*ArrayField // Array or array child fields
	EnumFields    []*BasicField // Enum fields
	Kind          ParamKind     // Param kind
	IsRequired    bool          // Whether parameter is required
}

func (p *ParamField) IsBasic() bool  { return p.Kind == kBasic }
func (p *ParamField) IsArray() bool  { return p.Kind == kArray }
func (p *ParamField) IsObject() bool { return p.Kind == kObject }
func (p *ParamField) IsEnum() bool   { return p.Kind == kEnum }

type BasicField struct {
	VarName    string // e.g. "serverArrayInstanceCloudHref" - shouldn't have "Names" or "Values" extensions for enum fields
	Path       string // e.g. "server.Instance.CloudHref"
	IsRequired bool   // Whether param field is required
}

type ArrayField struct {
	Name            string        // Name of array field e.g. "Schedules"
	VarName         string        // Name of array field variable e.g. "serverArraySchedules"
	Path            string        // Path to array field e.g. "serverArray.Schedules"
	ElemType        string        // Type of array elements, e.g. "*Schedule", "string"
	ElemInitializer string        // Go code used to initialize i-th array element, e.g. "&Schedule{Name: serverArrayScheduleNames[i]}"
	Leaves          []*BasicField // Leaf fields
}

// Build CmdWrite and all its dependencies from the info exposed by data types
func NewCmdsWriter(d *ApiDescriptor) (*CmdsWriter, error) {
	funcMap := template.FuncMap{
		"comment":     comment,
		"commandLine": commandLine,
		"now":         time.Now,
		"join":        join,
		"title":       strings.Title,
	}
	headerT, err := template.New("header-cmds").Funcs(funcMap).Parse(headerCmdTmpl)
	if err != nil {
		return nil, err
	}
	resourceT, err := template.New("resource-cmds").Funcs(funcMap).Parse(resourceCmdTmpl)
	if err != nil {
		return nil, err
	}
	resources := make(map[string]*ResourceCmd, len(d.ResourceNames))
	for _, r := range d.ResourceNames {
		resource := d.Resources[r]
		actionCmds := make([]*ActionCmd, len(resource.Actions))
		for j, action := range resource.Actions {
			params := action.AllParams
			var actionFlags []*ActionFlag
			var runnerFields []*RunnerField
			paramFields := make([]*ParamField, len(params))
			for i, n := range action.ParamNames {
				param := params[n]
				tmpFlags, tmpFields := toActionRunner(param)
				actionFlags = append(actionFlags, tmpFlags...)
				runnerFields = append(runnerFields, tmpFields...)
				paramFields[i] = toParamField(param)
			}
			hasOptionalParams := false
			for _, p := range paramFields {
				if !p.IsRequired {
					hasOptionalParams = true
					break
				}
			}
			actionCmds[j] = &ActionCmd{
				Name:              action.Name,
				Description:       toShortDescription(action.Description),
				RunnerFields:      runnerFields,
				ParamFields:       paramFields,
				Flags:             actionFlags,
				HasOptionalParams: hasOptionalParams,
			}
		}
		resources[r] = &ResourceCmd{
			Name:        r,
			Description: toShortDescription(resource.Description),
			ActionCmds:  actionCmds,
		}
	}
	return &CmdsWriter{
		HeaderTmpl:   headerT,
		ResourceTmpl: resourceT,
		Resources:    resources,
	}, nil
}

// Write header text
func (c *CmdsWriter) WriteHeader(d *ApiDescriptor, w io.Writer) error {
	return c.HeaderTmpl.Execute(w, d)
}

// Write resource header
func (c *CmdsWriter) WriteResourceHeader(name string, w io.Writer) {
	fmt.Fprintf(w, "/****** %s ******/\n\n", name)
}

// Write code for a resource
func (c *CmdsWriter) WriteResource(resource *ResourceData, w io.Writer) error {
	/* fmt.Printf("GENERATING %s\n", resource.Name)*/
	//spew.Dump(c.Resources[resource.Name])
	//fmt.Printf("\nFROM:\n\n")
	//spew.Dump(resource)
	/*fmt.Print("\n\n\n")*/
	return c.ResourceTmpl.Execute(w, c.Resources[resource.Name])
}

/* Helper cmd builder methods */

func toShortDescription(long string) string {
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

func toActionFlag(f *FlagDef) *ActionFlag {
	flagType := "Flag"
	if f.IsArray || f.IsEnum {
		flagType = "FlagPattern"
	}
	typeVar := "String"
	if f.Type == "int" {
		typeVar = "Int"
	}
	if f.IsArray {
		typeVar += "s"
	}
	typeVar += "Var"
	return &ActionFlag{
		Name:        f.Value,
		Type:        flagType,
		Description: toShortDescription(f.Description),
		IsRequired:  f.IsRequired,
		TypeVar:     typeVar,
	}
}

// Analyze action flags and produce data structures needed for code generation
func toActionRunner(param *ActionParam) ([]*ActionFlag, []*RunnerField) {
	flagDefs := param.Flags()
	actionFlags := make([]*ActionFlag, len(flagDefs))
	runnerFields := make([]*RunnerField, len(flagDefs))
	for i, flagDef := range flagDefs {
		var field RunnerField
		parts := strings.Split(flagDef.Path, ".")
		nameParts := make([]string, len(parts))
		for i, p := range parts {
			nameParts[i] = inflect.Camelize(p)
		}
		field.Name = strings.Join(nameParts, "")
		field.Name = strings.ToLower(string(field.Name[0])) + field.Name[1:]
		field.Type = flagDef.Type
		if flagDef.IsArray {
			field.Type = "[]" + field.Type
			field.CaptureName = field.Name + "Pos"
		} else if flagDef.IsEnum {
			field.Type = "[]string"
			field.CaptureName = field.Name + "Names"
			field.Name += "Values"
		} else if !flagDef.IsRequired {
			field.Type = "*" + field.Type
		}
		field.IsRequired = flagDef.IsRequired
		runnerFields[i] = &field
		flag := toActionFlag(flagDef)
		flag.RunnerField = &field
		actionFlags[i] = flag
	}
	return actionFlags, runnerFields
}

// Generate param field from action param
func toParamField(param *ActionParam) *ParamField {
	var paramField ParamField
	paramField.Name = param.Name
	paramField.NativeName = param.NativeName
	paramField.Type = param.Signature()
	switch t := param.Type.(type) {
	case *BasicDataType:
		paramField.Kind = kBasic
	case *ArrayDataType:
		paramField.Kind = kArray
	case *ObjectDataType:
		paramField.Type = paramField.Type[1:] // Not great I know, we don't want the "*"
		paramField.JsonFieldName = param.Name + "Json"
		for _, f := range t.Fields {
			// TBD: There's some duplication between the code called below and
			// the code used to compute the FlagDefs in data_types.go - might
			// be good to refactor one day and move the Flags() function out of
			// there and do everything here.
			tmpBasicFields, tmpArrayFields, tmpEnumFields := parseFields(f, []string{param.Name})
			paramField.BasicFields = append(paramField.BasicFields, tmpBasicFields...)
			paramField.ArrayFields = append(paramField.ArrayFields, tmpArrayFields...)
			paramField.EnumFields = append(paramField.EnumFields, tmpEnumFields...)
		}
		paramField.Kind = kObject
	case *EnumerableDataType:
		paramField.Kind = kEnum
	}
	paramField.IsRequired = param.Mandatory

	return &paramField
}

func arrayElemInitializer(elemType *ActionParam) string {

	return "//TBD"
}

// Parse param fields recursively and categorize them for code generation
// We simplify by assumming an array does not have child fields that are arrays as well
// (true for API 1.5).
func parseFields(f *ActionParam, path []string) ([]*BasicField, []*ArrayField, []*BasicField) {
	if f.Name != "item" {
		path = append(path, f.Name)
	}
	varName := path[0]
	if len(path) > 1 {
		parts := make([]string, len(path)-1)
		for i := 1; i < len(path); i++ {
			parts[i-1] = strings.Title(path[i])
		}
		varName += strings.Join(parts, "")
	}
	var basicFields []*BasicField
	var arrayFields []*ArrayField
	var enumFields []*BasicField
	switch t := f.Type.(type) {
	case *BasicDataType:
		basicFields = []*BasicField{&BasicField{varName, toPath(path), f.Mandatory}}
	case *ArrayDataType:
		tmpBasicFields, tmpArrayFields, tmpEnumFields := parseFields(t.ElemType, path)
		if len(tmpArrayFields) > 0 {
			panic(fmt.Sprintf("Generator does not support arrays in arrays (parameter %s has such a recursive array)", f.Name))
		}
		elemType := t.ElemType.Signature()
		var init string
		if _, ok := t.ElemType.Type.(*BasicDataType); ok {
			init = fmt.Sprintf("r.%s[r.%sPos[i]]", varName, varName)
		} else {
			init = elemType + "{\n"
			init += arrayElemInitializer(t.ElemType)
			init += "\n}"
		}
		arrayField := ArrayField{
			Name:            f.Name,
			VarName:         varName,
			Path:            toPath(path),
			ElemType:        elemType,
			ElemInitializer: init,
			Leaves:          append(tmpBasicFields, tmpEnumFields...),
		}
		arrayFields = append(arrayFields, &arrayField)
	case *ObjectDataType:
		for _, cf := range t.Fields {
			childBasicFields, childArrayFields, childEnumFields := parseFields(cf, path)
			basicFields = append(basicFields, childBasicFields...)
			arrayFields = append(arrayFields, childArrayFields...)
			enumFields = append(enumFields, childEnumFields...)
		}
	case *EnumerableDataType:
		enumFields = []*BasicField{&BasicField{varName, strings.Join(path, "."), f.Mandatory}}
	}

	return basicFields, arrayFields, enumFields
}

// Convert array of strings into path into runner struct, ["foo", "bar", "baz"] => "foo.Bar.Baz"
func toPath(parts []string) string {
	var elems
	if len(parts) > 1 {
		elems2 := make([]string, len(parts)-1)
		for i, p := range parts[1:] {
			elems2[i] = strings.ToUpper(string(p[0])) + p[1:]
		}
	}
	return strings.Join(elems, ".")
}

/* Helper format methods */

// Generate go expression to call action
func (a *ActionCmd) CallExp() string {
	// TBD: Deal with optional parameters
	var b bytes.Buffer
	b.WriteString(a.Name)
	b.WriteString("(")
	ln := len(a.ParamFields) - 1
	hasOptional := false
	for i, p := range a.ParamFields {
		if !p.IsRequired {
			hasOptional = true
			continue
		}
		if p.Kind == kObject {
			b.WriteString("&")
		} else if p.Kind == kBasic {
			b.WriteString("r.")
		}
		b.WriteString(p.Name)
		if i < ln || hasOptional {
			b.WriteString(", ")
		}
	}
	if hasOptional {
		b.WriteString("options")
	}
	b.WriteString(")")
	return b.String()
}

func join(s ...string) string {
	return strings.Join(s, "")
}

// Inline templates

const headerCmdTmpl = `
//************************************************************************//
//                RightScale API 1.5 command line client
//
{{comment "Generated " (now.Format "Jan 2, 2006 at 3:04pm (PST)")}}
// Command:
{{comment commandLine}}
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package rsapi15

import (
	"fmt"

	"gopkg.in/alecthomas/kingpin.v1"
)

// Registry of all sub-commands
var registry map[string]CommandRunner = make(map[string]CommandRunner)

// Register all sub-commands
func RegisterCommands(app *kingpin.Application) {
{{range .ResourceNames}}register{{.}}Commands(app)
{{end}} }
`
const resourceCmdTmpl = `{{$resource := .}}{{range .ActionCmds}}{{$runner := join .Name "Runner"}}
type {{$runner}} struct {
	{{range .RunnerFields}}{{.Name}} {{.Type}}{{if .CaptureName}}
	{{.CaptureName}} []string{{end}}
{{end}} }

func (r *{{$runner}}) Run(c *Client) (interface{}, error) { {{range .ParamFields}}{{if .IsObject }}

	/** Handle {{.Name}} parameter **/
	var {{.Name}} {{.Type}}{{if .Initializer}}
	{{.Initializer}}
	{{end}}
	// Load JSON if provided
	if len(r.{{.JsonFieldName}}) > 0 {
		if err := Json.Unmarshal(r.{{.JsonFieldName}}, &{{.Name}}); err != nil {
			return nil, fmt.Errorf("Could not load {{.Name}} JSON: %s", err.Error())
		}
	}

	{{if .BasicFields}}// Override with any provided basic field{{range .BasicFields}}{{if .IsRequired}}
	if len(r.{{.VarName}}) > 0 { {{else}}
	if r.{{.VarName}} != nil { {{end}}
		{{.Path}} = r.{{.VarName}}
	}
	{{end}}{{/* end range .BasicFields */}}
	{{end}}{{/* end if .BasicFields */}}{{if .ArrayFields}}// Create array fields from flags
	var seenPos map[int]bool{{range .ArrayFields}}{{$array := .}}{{range .Leaves}}
	max{{title .VarName}}Pos := -1
	seenPos = make(map[int]bool)
	for _, p := range r.{{.VarName}}Pos {
		pos, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for {{.Path}} field of {{$array.Name}} array", p)
		}
		if _, ok := seenPos[pos]; ok {
			return nil, fmt.Errorf("Value of {{.Path}} field of {{$array.Name}} array defined multiple times for position %s", p)
		}
		seenPos[pos] = true
		if pos > max{{title .VarName}}Pos {
			max{{title .VarName}}Pos = pos
		}
	}
	if len(r.{{.VarName}}) != max{{title .VarName}}Pos {
		return nil, fmt.Errof("Invalid flags for {{.Path}} field of {{$array.Name}} array, %s were defined but max position is %s",
			len(r.{{.VarName}}), max{{title .VarName}}Pos)
	}
	{{end}}{{/* end range Leaves */}}{{$first := (index .Leaves 0).VarName}}if max{{title $first}}Pos > -1 {
		{{.VarName}} := make([]{{.ElemType}}, max{{title $first}}Pos+1)
		for i := 0; i < max{{title $first}}Pos+1; i++ {
			{{.VarName}}[i] = {{.ElemInitializer}}
		}
		{{.Path}} = {{.VarName}}
	}
	{{end}}{{/* end range .ArrayFields */}}
	{{end}}{{/* end if .ArrayFields */}}{{if .EnumFields}}{{range .EnumFields}}// Create enumarable fields from flags
	{{.VarName}} := make(map[string]string, len(r.{{.VarName}}Names))
	for i, n := range r.{{.VarName}}Names {
		{{.VarName}}[n] = r.{{.VarName}}Values[i]
	}
	{{.Path}} = {{.VarName}}
	{{end}}{{end}}{{/* end if .EnumFields */}}
	{{end}}{{if .IsEnum}}

	/** Handle {{.Name}} parameter **/
	var {{.Name}} {{.Type}}

	for i, n := range r.{{.Name}}Names {
		{{.Name}}[n] = r.{{.Name}}Values[i]
	}
	{{end}}{{if .IsArray}}

	/** Handle {{.Name}} parameter **/
	var {{.Name}} {{.Type}}

	for i, v := range r.{{.Name}} {
		pos, err := strconv.Atoi(r.{{.Name}}Pos[i])
		if err != nil {
			return nil, fmt.Errorf("Invalid position value %s for {{.Name}} array", r.{{.Name}}Pos[i])
		}
		{{.Name}}[pos] = v
	}
	{{end}}{{end}}{{/* end range .ParamFields*/}}{{if .HasOptionalParams}}

	/** Handle optional parameters **/
	options := make(ApiParams)
	{{range .ParamFields}}{{if not .IsRequired}}{{if .IsBasic}}if r.{{.Name}} != nil {
		{{end}}options["{{.NativeName}}"] = {{if .IsBasic}}*r.{{end}}{{.Name}}
	{{if .IsBasic}} }
	{{end}}{{end}}{{end}}{{end}}{{/* end if .HasOptionalParams */}}
	return c.{{.CallExp}}
}
{{end}}{{/* end range .ActionCmds */}}
// Register all {{.Name}} actions
func register{{.Name}}Cmds(app *kingpin.Application) {
	resCmd := app.Cmd("{{.Name}}", ` + "`" + `{{.Description}}` + "`" + `)
	{{range .ActionCmds}}{{$runner := join .Name $resource.Name "Runner"}}{{$runnerVar := join .Name "Runner"}}
	{{$runnerVar}} := new({{$runner}}){{$cmdVar := join .Name "Cmd"}}
	{{$cmdVar}} := resCmd.Command("{{.Name}}", ` + "`" + `{{.Description}}` + "`" + `)
	{{range .Flags}}{{$runnerVar}}.{{.Type}}(` + "`" + `{{.Name}}` + "`" + `, ` + "`" + `{{.Description}}` + "`" + `).{{/*
	*/}}{{if .IsRequired}}Required().{{end}}{{if .RunnerField.CaptureName}}Capture(&{{$runnerVar}}.{{/*
	*/}}{{.RunnerField.CaptureName}}).{{end}}{{.TypeVar}}({{if .IsRequired}}&{{end}}{{$runnerVar}}.{{.RunnerField.Name}})
	{{end}}{{/* end range .Flags */}}registry[{{$cmdVar}}.FullCommand()] = {{$runnerVar}}
	{{end}}{{/* end range .ActionCmds */}} }
`
