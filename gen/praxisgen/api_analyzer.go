package main

import (
	"fmt"
	"sort"

	"github.com/rightscale/rsc/gen"
)

// The analyzer struct holds the analysis results
type ApiAnalyzer struct {
	// Raw resources as defined in API json metadata
	RawResources map[string]map[string]interface{}
	// Raw types as defined in API json metadata
	RawTypes map[string]map[string]interface{}
	// Version being analyzed
	Version string

	// Temporary data structures used by analysis

	// Descriptor being built
	descriptor *gen.ApiDescriptor
	// Map of ruby type name to go type name
	// Also used to record types that have been processed.
	typeNames map[string]string
}

// Factory method for API analyzer
func NewApiAnalyzer(version string, resources, types map[string]map[string]interface{}) *ApiAnalyzer {
	return &ApiAnalyzer{
		RawResources: resources,
		RawTypes:     types,
		Version:      version,
		typeNames:    make(map[string]string),
	}
}

// Create API descriptor from raw resources and types
func (a *ApiAnalyzer) Analyze() (*gen.ApiDescriptor, error) {
	var descriptor = gen.ApiDescriptor{
		Version:   a.Version,
		Resources: make(map[string]*gen.Resource),
		Types:     make(map[string]*gen.ObjectDataType),
	}
	a.descriptor = &descriptor

	// 1. Do a first pass to collect all types
	var rawResourceNames = make([]string, len(a.RawResources))
	var idx = 0
	for name, _ := range a.RawResources {
		rawResourceNames[idx] = name
		idx += 1
	}
	sort.Strings(rawResourceNames)
	for _, name := range rawResourceNames {
		var resource = a.RawResources[name]
		var err = a.AnalyzeMediaType(resource["media_type"].(string))
		if err != nil {
			return nil, err
		}
		actions, ok := resource["actions"]
		if !ok {
			continue
		}
		for _, ac := range actions.([]interface{}) {
			var action = ac.(map[string]interface{})
			responses, ok := action["responses"]
			if ok {
				for _, r := range responses.(map[string]interface{}) {
					mediaType, ok := r.(map[string]interface{})["media_type"]
					if ok {
						var m = mediaType.(map[string]interface{})
						err := a.AnalyzeMediaType(m["name"].(string))
						if err != nil {
							return nil, err
						}
					}
				}
			}
		}
	}
	var typeNames = make([]string, len(a.typeNames))
	idx = 0
	for _, name := range a.typeNames {
		typeNames[idx] = name
		idx += 1
	}
	sort.Strings(typeNames)
	descriptor.TypeNames = typeNames

	// 2. Now do a second pass to analyze each resource
	for _, name := range rawResourceNames {
		var err = a.AnalyzeResource(name, a.RawResources[name], &descriptor)
		if err != nil {
			return nil, err
		}
	}

	// We're done
	return &descriptor, nil
}

// Analyze media type, recurse through all types and create corresponding ObjectDataTypes and
// ActionParams.
func (a *ApiAnalyzer) AnalyzeMediaType(name string) error {
	var mtype, ok = a.RawTypes[name]
	if !ok {
		return fmt.Errorf("Unknown media type %s", name)
	}
	var typeName = toGoTypeName(name, false)
	if _, ok := a.typeNames[typeName]; ok {
		return nil // Already analyzed
	}
	var obj = gen.ObjectDataType{Name: typeName}
	var attributes = mtype["attributes"].(map[string]interface{})
	var fields = make([]*gen.ActionParam, len(attributes))
	var idx = 0
	for _, attName := range sortedKeys(attributes) {
		var att = attributes[attName]
		var param, err = a.AnalyzeAttribute(attName, attName, att.(map[string]interface{}))
		if err != nil {
			return err
		}
		fields[idx] = param
		idx += 1
	}
	obj.Fields = fields
	a.typeNames[name] = typeName
	a.descriptor.Types[typeName] = &obj
	return nil
}
