package main

import "fmt"

// Data structure used to describe API resources.
type Resource struct {
	Name        string       // Resource name, e.g. "ServerArray"
	Description string       // Resource description
	Actions     []*Action    // Resource actions, e.g. "index", "show", "update" ...
	Attributes  []*Attribute // Resource attributes
	HrefRegexp  string       // Regular expression that matches resource hrefs
}

// Data structure used to describe collection and resource actions
type Action struct {
	Name           string         // Action name, e.g. "create", "multi_terminate"
	MethodName     string         // Go method name, e.g. "Create", "MultiTerminate"
	Description    string         // Action description
	HttpMethod     string         // Action HTTP method, e.g. "GET", "POST"
	Paths          []string       // Action paths
	Suffix         string         // Action url suffix
	Params         []*ActionParam // Action method parameters
	LeafParams     []*ActionParam // Action parameter leaves (for command line)
	Return         string         // Type of method results, e.g. "*ServerArray"
	ReturnLocation bool           // Whether API returns a location header. True for all "Create" except OAuth2's ugh.
}

// Whether action requests takes a payload
func (a *Action) HasPayload() bool {
	return a.HttpMethod == "POST" || a.HttpMethod == "PUT" || a.HttpMethod == "GET"
}

// Whether action has a response body or header
func (a *Action) HasResponse() bool {
	return a.HttpMethod == "POST" || a.HttpMethod == "GET"
}

// Resource attributes
type Attribute struct {
	Name      string // Attribute name, e.g. "elasticity_params"
	FieldName string // Corresponding go struct field name, e.g. "ElasticityParams"
	FieldType string // Corresponding go struct type, e.g. "*ElasticityParams"
}

// Data structure used to render method params
type ActionParam struct {
	Name        string        // Name of parameter
	QueryName   string        // Query string style parameter name
	Description string        // Description of parameter
	VarName     string        // go variable name
	Type        DataType      // Type of parameter
	Mandatory   bool          // Whether parameter is mandatory
	NonBlank    bool          // Whether parameter must not be blank if provided (string or array)
	Regexp      string        // Regular expression string parameter must match
	ValidValues []interface{} // Allowed values (if not empty)
}

// Data type interface
type DataType interface {
	IsEquivalent(other DataType) bool // true if datatyoe and other represent the same data structure
}

// A basic data type only has a name, i.e. "int" or "string"
type BasicDataType string

// An array data type defines the type of its elements
type ArrayDataType struct {
	ElemType *ActionParam
}

// An object data type has a name and fields
type ObjectDataType struct {
	Name   string
	Fields []*ActionParam
}

// An enumerable is just a map
type EnumerableDataType int

// true if both b and other represent the same type
func (b *BasicDataType) IsEquivalent(other DataType) bool {
	t, ok := other.(*BasicDataType)
	if !ok {
		return false
	}
	if *t != *b {
		return false
	}
	return true
}

// true if other is also a array data type and element types of both arrays are equivalent
func (a *ArrayDataType) IsEquivalent(other DataType) bool {
	t, ok := other.(*ArrayDataType)
	if !ok {
		return false
	}
	if !a.ElemType.IsEquivalent(t.ElemType) {
		return false
	}
	return true
}

// true if other is a object data type and each field is equivalent recursively
func (o *ObjectDataType) IsEquivalent(other DataType) bool {
	a, ok := other.(*ObjectDataType)
	if !ok {
		return false
	}
	if o.Name != a.Name {
		return false
	}
	if len(o.Fields) != len(a.Fields) {
		return false
	}
	for i, f := range o.Fields {
		if !f.IsEquivalent(a.Fields[i]) {
			return false
		}
	}
	return true
}

// true if other is also an enumerable data type
func (e *EnumerableDataType) IsEquivalent(other DataType) bool {
	_, ok := other.(*EnumerableDataType)
	return ok
}

// Generate signature used e.g. when specifying param in function signatures
func (p *ActionParam) Signature() (sig string) {
	switch t := p.Type.(type) {
	case *BasicDataType:
		if p.Mandatory {
			sig = string(*t)
		} else {
			sig = "*" + string(*t)
		}
	case *ArrayDataType:
		cs := t.ElemType.Signature()
		if _, ok := t.ElemType.Type.(*BasicDataType); ok {
			if cs[0] == '*' {
				cs = cs[1:]
			}
		}
		sig = fmt.Sprintf("[]%s", cs)
	case *ObjectDataType:
		sig = fmt.Sprintf("*%s", t.Name)
	case *EnumerableDataType:
		sig = "map[string]string"
	}
	return
}

// true if action params have the same name and type
func (p *ActionParam) IsEquivalent(other *ActionParam) bool {
	return p.Name == other.Name && p.Type.IsEquivalent(other.Type)
}

// Make it possible to sort action parameters by name
type ByName []*ActionParam

func (b ByName) Len() int           { return len(b) }
func (b ByName) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b ByName) Less(i, j int) bool { return b[i].Name < b[j].Name }
