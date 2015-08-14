package main // import "gopkg.in/rightscale/rsc.v3/gen/praxisgen"

import (
	"sort"

	"gopkg.in/rightscale/rsc.v3/gen"
)

// The analyzer struct holds the analysis results
type ApiAnalyzer struct {
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
	descriptor *gen.ApiDescriptor
	// Temporary data structure used to create types and keep track of names
	Registry *TypeRegistry
}

// Factory method for API analyzer
func NewApiAnalyzer(version, clientName string, resources, types map[string]map[string]interface{}) *ApiAnalyzer {
	return &ApiAnalyzer{
		RawResources: resources,
		RawTypes:     types,
		Version:      version,
		ClientName:   clientName,
		Registry:     NewTypeRegistry(),
	}
}

// Create API descriptor from raw resources and types
func (a *ApiAnalyzer) Analyze() (*gen.ApiDescriptor, error) {
	descriptor := gen.ApiDescriptor{
		Version:   a.Version,
		Resources: make(map[string]*gen.Resource),
		Types:     make(map[string]*gen.ObjectDataType),
	}
	a.descriptor = &descriptor

	// Sort resource names so iterations are always done in the same order
	rawResourceNames := make([]string, len(a.RawResources))
	idx := 0
	for name, _ := range a.RawResources {
		rawResourceNames[idx] = name
		idx += 1
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

// Registry of all created types
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

// Type registry factory
func NewTypeRegistry() *TypeRegistry {
	return &TypeRegistry{
		NamedTypes:  make(map[string]*gen.ObjectDataType),
		InlineTypes: make(map[string][]*gen.ObjectDataType),
	}
}

// Retrieve named type
func (reg *TypeRegistry) GetNamedType(name string) *gen.ObjectDataType {
	return reg.NamedTypes[toGoTypeName(name, false)]
}

// Create named type
func (reg *TypeRegistry) CreateNamedType(name string) *gen.ObjectDataType {
	goName := toGoTypeName(name, false)
	obj := gen.ObjectDataType{TypeName: goName}
	if _, ok := reg.NamedTypes[goName]; ok {
		panic("BUG: Can't create two named types with same name....")
	}
	reg.NamedTypes[goName] = &obj
	return &obj
}

// Create inline type
func (reg *TypeRegistry) CreateInlineType(name string) *gen.ObjectDataType {
	goName := toGoTypeName(name, false)
	obj := gen.ObjectDataType{TypeName: goName}
	reg.InlineTypes[goName] = append(reg.InlineTypes[goName], &obj)
	return &obj
}

// Finalize all type names, should be called once after analysis.
func (reg *TypeRegistry) FinalizeTypeNames(d *gen.ApiDescriptor) {
	for n, named := range reg.NamedTypes {
		reg.InlineTypes[n] = append(reg.InlineTypes[n], named)
	}
	d.FinalizeTypeNames(reg.InlineTypes)
}
