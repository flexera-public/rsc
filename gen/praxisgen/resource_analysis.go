package main

import (
	"fmt"

	"bitbucket.org/pkg/inflect"
	"github.com/rightscale/rsc/gen"
)

// Create API descriptor from raw resources and types
func (a *ApiAnalyzer) AnalyzeResource(name string, res map[string]interface{}, desc *gen.ApiDescriptor) error {
	var resource = gen.Resource{Name: name}

	// Description
	if d, ok := res["description"]; ok {
		resource.Description = removeBlankLines(d.(string))
	}

	// Attributes
	m, ok := res["media_type"].(string)
	if !ok {
		return fmt.Errorf("Resource %s has no media type (missing key \"media_type\")", name)
	}
	t, ok := a.RawTypes[m]
	if !ok {
		return fmt.Errorf("Missing definition for Resource %s media type %s", name, prettify(m))
	}
	attrs, ok := t["attributes"].(map[string]interface{})
	if !ok {
		return fmt.Errorf("Missing attributes for media type %s", prettify(m))
	}
	var attributes = make([]*gen.Attribute, len(attrs))
	for idx, n := range sortedKeys(attrs) {
		var param, err = a.AnalyzeAttribute(n, n, attrs[n].(map[string]interface{}))
		if err != nil {
			return err
		}
		attributes[idx] = &gen.Attribute{n, inflect.Camelize(n), param.SignatureIgnoreMandatory()}
	}
	resource.Attributes = attributes

	// Actions
	actions, err := a.AnalyzeActions(name, res)
	if err != nil {
		return err
	}
	resource.Actions = actions

	// Name and done
	var resName = toGoTypeName(name, false)
	desc.Resources[resName] = &resource
	desc.ResourceNames = append(desc.ResourceNames, resName)

	return nil
}
