package main

import (
	"fmt"
	"regexp"
	"strings"
)

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
	Name          string
	Description   string
	HttpMethod    string
	Path          string
	UrlExp        string
	NativeParams  map[string]interface{}  // Params as defined in JSON
	AllParams     map[string]*ActionParam // All parameters
	ParamNames    []string                // Names of all parameters ordered alphabetically
	PathParams    []*ActionParam          // Params used to build URL
	PayloadParams []*ActionParam          // Params that should be sent in payload
	QueryParams   []*ActionParam          // Params that should be sent in query string
	Return        string
}

// Resource attributes, 'Type' is string representation of go type, e.g. "*time.Time"
type ResourceAttribute struct {
	Name      string
	JsonName  string
	Signature string
}

// Data structure used to render method params
type ActionParam struct {
	Name        string
	Description string
	NativeName  string
	Type        DataType
	Declaration string
	Mandatory   bool
	NonBlank    bool
	Regexp      string
	ValidValues []interface{}
}

// Same name and type => identical (validations and description are not taken into account)
func (p *ActionParam) Compare(other *ActionParam) bool {
	if p.Name != other.Name {
		return false
	}
	if p.NativeName != other.NativeName {
		return false
	}
	switch t := p.Type.(type) {
	case *BasicDataType:
		b, ok := other.Type.(*BasicDataType)
		if !ok {
			return false
		}
		if *t != *b {
			return false
		}
	case *EnumerableDataType:
		_, ok := other.Type.(*EnumerableDataType)
		if !ok {
			return false
		}
	case *ArrayDataType:
		a, ok := other.Type.(*ArrayDataType)
		if !ok {
			return false
		}
		if !a.ElemType.Compare(t.ElemType) {
			return false
		}
	case *ObjectDataType:
		a, ok := other.Type.(*ObjectDataType)
		if !ok {
			return false
		}
		if !a.Compare(t) {
			return false
		}
	}
	return true
}

// Flag definitions consist of a map of definition to unique name.
// Names can be used to prefix command line args.
type FlagDefs []*FlagDef

// A flag definition contains either a literal flag name or a flag pattern. It also defines the
// type of the value returned by the command line parser for that flag.
type FlagDef struct {
	Path      string // Used to build variable name
	Value     string // Actual flag name
	Type      string // Flag type, "int", "string", "ints" or "strings"
	IsPattern bool   // Whether value is a regexp
}

// Generate flag definitions needed to create a command line parser that will gather all the
// data needed to build an instance of the parameter value.
func (p *ActionParam) Flags() FlagDefs {
	children := p.childFlags()
	var res FlagDefs
	if len(children) > 0 {
		res = make(FlagDefs, len(children))
		for i, c := range children {
			value := c.Value
			path := p.Name
			if len(value) == 0 {
				value = path
			} else {
				if c.IsPattern {
					value = regexp.QuoteMeta(path+".") + c.Value
				} else {
					value = path + "." + c.Value
				}
			}
			if len(c.Path) > 0 {
				path += "." + c.Path
			}
			res[i] = &FlagDef{path, value, c.Type, c.IsPattern}
		}
	} else {
		res = FlagDefs{&FlagDef{p.Name, p.Name, string(*p.Type.(*BasicDataType)), false}}
	}
	return res
}

// Child flags, used for recursion
func (p *ActionParam) childFlags() []*FlagDef {
	var res []*FlagDef
	switch t := p.Type.(type) {
	case *BasicDataType:
		res = []*FlagDef{}
	case *ArrayDataType:
		flags := t.ElemType.childFlags()
		if len(flags) > 0 {
			res = make([]*FlagDef, len(flags))
			for i, f := range flags {
				fType := f.Type
				if !strings.HasSuffix(fType, "s") {
					fType += "s"
				}
				value := f.Value
				if !f.IsPattern {
					value = regexp.QuoteMeta(value)
				}
				res[i] = &FlagDef{f.Path, fmt.Sprintf(`(\d+)\.%s`, value), fType, true}
			}
		} else {
			res = []*FlagDef{&FlagDef{"", `(\d+)`, "strings", true}}
		}
	case *ObjectDataType:
		res := []*FlagDef{}
		for _, field := range t.Fields {
			flags := field.childFlags()
			temp := make([]*FlagDef, len(flags))
			if len(flags) > 0 {
				for i, f := range flags {
					temp[i] = &FlagDef{field.Name + "." + f.Path,
						fmt.Sprintf("%s.%s", field.Name, f.Value),
						f.Type, f.IsPattern}
				}
				res = append(res, temp...)
			} else {
				// No child means field's type is a basic type
				fType := string(*field.Type.(*BasicDataType))
				res = append(res, &FlagDef{field.Name, field.Name, fType, false})
			}
		}
	case *EnumerableDataType:
		res = []*FlagDef{&FlagDef{"", `(.+)`, "strings", true}}
	}
	return res
}

// Make it possible to sort action parameters by name
type ByName []*ActionParam

func (b ByName) Len() int           { return len(b) }
func (b ByName) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByName) Less(i, j int) bool { return b[i].Name < b[j].Name }

// All type structures implement the DataType interface
type DataType interface {
	Signature() string                    // Produce go compatible signature, e.g. to define function
	Inspect() string                      // Produce pretty print, mainly for debug
	BlankConditionExp(name string) string // Go condition to test whether type value is blank. Empty string if type values can't be blank.
}

// A basic data type only has a name, e.g. "int" or "string"
type BasicDataType string

// Implement DataType
func (b *BasicDataType) Signature() string {
	return string(*b)
}

func (b *BasicDataType) Inspect() string {
	return string(*b)
}

func (b *BasicDataType) BlankConditionExp(name string) string {
	if *b == "string" {
		return fmt.Sprintf("if %s == \"\" {", name)
	} else {
		return ""
	}
}

// An array data type defines the type of its elements
type ArrayDataType struct {
	ElemType *ActionParam
}

// Implement DataType
func (a *ArrayDataType) Signature() string {
	return fmt.Sprintf("[]%s", a.ElemType.Type.Signature())
}

func (a *ArrayDataType) Inspect() string {
	return fmt.Sprintf("Array of %s", a.ElemType.Type.Inspect())
}

func (a *ArrayDataType) BlankConditionExp(name string) string {
	return fmt.Sprintf("if len(%s) == 0 {", name)
}

// An object data type has a name and fields
type ObjectDataType struct {
	Name   string
	Fields []*ActionParam
}

// Implement DataType
func (o *ObjectDataType) Signature() string {
	return fmt.Sprintf("*%s", o.Name)
}

func (o *ObjectDataType) Inspect() string {
	var fields []string
	for _, f := range o.Fields {
		fields = append(fields, fmt.Sprintf("%s: %s", f.Name, f.Type.Inspect()))
	}
	return fmt.Sprintf("%s:{\n\t%s\n}", o.Name, strings.Join(fields, "\n\t"))
}

func (o *ObjectDataType) BlankConditionExp(name string) string {
	return fmt.Sprintf("if %s == nil {", name)
}

// Object data type declaration
func (o *ObjectDataType) Declaration() string {
	fields := make([]string, len(o.Fields))
	for i, f := range o.Fields {
		fields[i] = fmt.Sprintf("%s %s `json:\"%s,omitempty\"`", strings.Title(f.Name),
			f.Type.Signature(), f.NativeName)
	}
	return fmt.Sprintf("type %s struct {\n%s\n}", o.Name,
		strings.Join(fields, "\n\t"))
}

// Object data type comparison
func (o *ObjectDataType) Compare(other *ObjectDataType) bool {
	if o.Name != other.Name {
		return false
	}
	if len(o.Fields) != len(other.Fields) {
		return false
	}
	for i, f := range o.Fields {
		if !f.Compare(other.Fields[i]) {
			return false
		}
	}
	return true
}

// An enumerable is just a map
type EnumerableDataType int

// Implement DataType
func (e *EnumerableDataType) Signature() string {
	return "map[string]string"
}

func (e *EnumerableDataType) Inspect() string {
	return fmt.Sprintf("Enumeration")
}

func (e *EnumerableDataType) BlankConditionExp(name string) string {
	return fmt.Sprintf("if len(%s) == 0 {", name)
}
