package main

import (
	"sort"

	"github.com/rightscale/rsc/gen"
)

// APIAnalyzer is holds the analysis results.
type APIAnalyzer struct {
	// Raw resources as defined in API json metadata
	RawResources map[string]map[string]interface{}
	// Raw types as defined in API json metadata
	RawTypes map[string]map[string]interface{}
	// Version being analyzed
	Version string
	// Name of golang struct to generate for API client
	ClientName string

	// Temporary data structures used by analysis

	// Descriptor being built
	descriptor *gen.APIDescriptor
	// Temporary data structure used to create types and keep track of names
	Registry *TypeRegistry
}

// NewAPIAnalyzer is the factory method for APIAnalyzer.
func NewAPIAnalyzer(version, clientName string, resources, types map[string]map[string]interface{}) *APIAnalyzer {
	return &APIAnalyzer{
		RawResources: resources,
		RawTypes:     types,
		Version:      version,
		ClientName:   clientName,
		Registry:     NewTypeRegistry(),
	}
}

// Analyze creates an API descriptor from raw resources and types.
func (a *APIAnalyzer) Analyze() (*gen.APIDescriptor, error) {
	descriptor := gen.APIDescriptor{
		Version:   a.Version,
		Resources: make(map[string]*gen.Resource),
		Types:     make(map[string]*gen.ObjectDataType),
	}
	a.descriptor = &descriptor

	// Sort resource names so iterations are always done in the same order
	rawResourceNames := make([]string, len(a.RawResources))
	idx := 0
	for name := range a.RawResources {
		rawResourceNames[idx] = name
		idx++
	}
	sort.Strings(rawResourceNames)

	// Analyze each resource
	for _, name := range rawResourceNames {
		err := a.AnalyzeResource(name, a.RawResources[name], &descriptor)
		if err != nil {
			return nil, err
		}
	}

	// We're done
	a.Registry.FinalizeTypeNames(&descriptor)
	return &descriptor, nil
}

// TypeRegistry keeps track of all created types.
// There are types that have a one to one mapping with types defined in the metadata (named types)
// and types that are created from inline structs and hashes (inline types). We need to
// differentiate because the names for the later are not guarenteed to be unique.
// Keep track of all types during analysis, FinalizeTypeNames should be called at the end to
// resolve all the type names and make sure that different data structures end up with different
// names.
type TypeRegistry struct {
	NamedTypes  map[string]*gen.ObjectDataType
	InlineTypes map[string][]*gen.ObjectDataType
}

// NewTypeRegistry creates a type registry.
func NewTypeRegistry() *TypeRegistry {
	return &TypeRegistry{
		NamedTypes:  make(map[string]*gen.ObjectDataType),
		InlineTypes: make(map[string][]*gen.ObjectDataType),
	}
}

// GetNamedType retrieves a type given its name.
func (reg *TypeRegistry) GetNamedType(name string) *gen.ObjectDataType {
	return reg.NamedTypes[toGoTypeName(name)]
}

// CreateNamedType returns a new type given a name, the name must be unique.
func (reg *TypeRegistry) CreateNamedType(name string) *gen.ObjectDataType {
	goName := toGoTypeName(name)
	obj := gen.ObjectDataType{TypeName: goName}
	if _, ok := reg.NamedTypes[goName]; ok {
		panic("BUG: Can't create two named types with same name....")
	}
	reg.NamedTypes[goName] = &obj
	return &obj
}

// CreateInlineType creates a new inline type.
func (reg *TypeRegistry) CreateInlineType(name string) *gen.ObjectDataType {
	goName := toGoTypeName(name)
	obj := gen.ObjectDataType{TypeName: goName}
	reg.InlineTypes[goName] = append(reg.InlineTypes[goName], &obj)
	return &obj
}

// FinalizeTypeNames makes sure type names are unique, it should be called after analysis
// has completed.
func (reg *TypeRegistry) FinalizeTypeNames(d *gen.APIDescriptor) {
	for n, named := range reg.NamedTypes {
		reg.InlineTypes[n] = append(reg.InlineTypes[n], named)
	}
	d.FinalizeTypeNames(reg.InlineTypes)
}
