package main // import "gopkg.in/rightscale/rsc.v6/gen/goav2gen"

import (
	"strings"
)

// Define the top level swagger defintion structs here.
// These definitions are good enough for parsing goa generated swaggers but definately
// don't reflect the complete swagger spec as of yet.

// Doc holds the swagger data structure
type Doc struct {
	SwaggerVersion      string                         `json:"swagger"`
	Info                Info                           `json:"info"`
	Host                string                         `json:"host"`
	BasePath            string                         `json:"basePath"`
	Paths               map[string]EndpointMap         `json:"paths"`
	Definitions         map[string]*Definition         `json:"definitions"`
	SecurityDefinitions map[string]*SecurityDefinition `json:"securityDefinitions"`
}

// Info holds additional info from the swagger document
type Info struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Version     string `json:"version"`
}

// EndpointMap is a map of http verbs ("get") -> Endpoint
type EndpointMap map[string]*Endpoint

// Endpoint defines an API endpoint
type Endpoint struct {
	Tags        []string          `json:"tags"`
	Summary     string            `json:"summary"`
	Description string            `json:"description"`
	OperationID string            `json:"operationId"`
	Parameters  []*Parameter      `json:"parameters"`
	Responses   map[int]*Response `json:"responses"`
	Schemes     []string          `json:"schemes"`
	Security    SecurityRefs      `json:"security"`
}

// Definition defines a result type
type Definition struct {
	Title       string               `json:"title"`
	Type        string               `json:"type"`
	Description string               `json:"description"`
	Example     interface{}          `json:"example"`
	Properties  map[string]*Property `json:"properties"`
	Required    []string             `json:"required"`
	Items       Ref                  `json:"items"`
}

// Property is an attribute of a result type
type Property struct {
	Fault       bool          `json:"fault"`
	Items       Ref           `json:"items"`
	Description string        `json:"description"`
	Example     interface{}   `json:"example"`
	Type        string        `json:"type"`
	Enum        []interface{} `json:"enum"`
	Schema      Ref           `json:"schema"`
	Pattern     string        `json:"pattern"`
	Format      string        `json:"format"`
	Default     interface{}   `json:"default"`
	Minimum     interface{}   `json:"minimum"`
	Maximum     interface{}   `json:"maximum"`
	Ref         string        `json:"$ref"`
}

// Parameter is an attribute for a request type
type Parameter struct {
	Description string        `json:"description"`
	Name        string        `json:"name"`
	In          string        `json:"in"`
	Required    bool          `json:"required"`
	Type        string        `json:"type"`
	Enum        []interface{} `json:"enum"`
	Schema      Ref           `json:"schema"`
	Pattern     string        `json:"pattern"`
	Format      string        `json:"format"`
}

// Ref references a definition
type Ref map[string]interface{}

// Response defines attributes about an HTTP response
type Response struct {
	Description string               `json:"description"`
	Schema      Ref                  `json:"schema"`
	Headers     map[string]Parameter `json:"headers"`
}

// SecurityDefinition defines the authentication scheme
type SecurityDefinition struct {
	Type        string `json:"type"`
	Description string `json:"description"`
	Name        string `json:"string"`
	In          string `json:"in"`
}

// SecurityRefs a series of pointers to security definitions
type SecurityRefs []SecurityRef

// SecurityRef is a single pointer to a security definition
type SecurityRef map[string]interface{}

// Ref gets a definition for a Schema reference, if it exists
func (d *Doc) Ref(r Ref) *Definition {
	if refIF, ok := r["$ref"]; ok {
		refKey := strings.TrimPrefix(refIF.(string), "#/definitions/")
		return d.Definitions[refKey]
	}
	return nil
}

// Type gets a type for a Schema reference, if it exists
func (r Ref) Type() string {
	if _, ok := r["$ref"]; ok {
		return "object"
	}
	if refIF, ok := r["type"]; ok {
		return refIF.(string)
	}
	return ""
}

// Required gets a type for a Schema reference, if it exists
func (r Ref) Required() []string {
	if refIF, ok := r["required"]; ok {
		return refIF.([]string)
	}
	return []string{}
}

// ID of the reference
func (r Ref) ID() string {
	if refIF, ok := r["$ref"]; ok {
		return strings.TrimPrefix(refIF.(string), "#/definitions/")
	}
	return ""
}

// Service returns the goa.v2 service
func (ep *Endpoint) Service() string {
	if len(ep.Tags) > 0 {
		return ep.Tags[0]
	}
	if len(ep.OperationID) > 0 {
		return strings.Split(ep.OperationID, "#")[0]
	}
	return ""
}

// Methods returns the goa.v2 method
func (ep *Endpoint) Method() string {
	if strings.Contains(ep.OperationID, "#") {
		return strings.Split(ep.OperationID, "#")[1]
	}
	return ""
}

// IsRef is whether this property is a reference to another object.
func (p *Property) IsRef() bool {
	return len(p.Ref) > 0
}

// GetRef will get the object pointed to by this property
func (p *Property) GetRef() Ref {
	return Ref{"$ref": p.Ref}
}
