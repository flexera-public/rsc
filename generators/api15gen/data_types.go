package main

import (
	"fmt"
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
	Url           string
	NativeParams  map[string]interface{} // Params as defined in JSON
	AllParams     []*ActionParam         // All parameters
	PathParams    []*ActionParam         // Params used to build URL
	PayloadParams []*ActionParam         // Params that should be sent in payload
	QueryParams   []*ActionParam         // Params that should be sent in query string
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
	Mandatory   bool
	NonBlank    bool
	Regexp      string
	ValidValues []interface{}
}

// All type structures implement the DataType interface
type DataType interface {
	Signature() string // Produce go compatible signature, e.g. to define function
	Inspect() string   // Produce pretty print, mainly for debug
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

// Object data type declaration
func (o *ObjectDataType) Declaration() string {
	fields := make([]string, len(o.Fields))
	for i, f := range o.Fields {
		fields[i] = fmt.Sprintf("%s %s", f.Name, f.Type.Signature())
	}
	return fmt.Sprintf("type %s struct {\n%s\n}", o.Name, strings.Join(fields, "\n\t"))
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
