package main  // import "gopkg.in/rightscale/rsc.v1-unstable/gen/praxisgen"

import (
	"fmt"

	"bitbucket.org/pkg/inflect"
	"gopkg.in/rightscale/rsc.v1-unstable/gen" // import "gopkg.in/rightscale/rsc.v1-unstable/gen"
)

// Create API descriptor from raw resources and types
func (a *ApiAnalyzer) AnalyzeResource(name string, res map[string]interface{}, desc *gen.ApiDescriptor) error {
	name = inflect.Singularize(name)
	resource := gen.Resource{Name: name, ClientName: a.ClientName}

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
	attributes := []*gen.Attribute{}
	if ok {
		attrs, ok := t["attributes"].(map[string]interface{})
		if ok {
			attributes = make([]*gen.Attribute, len(attrs))
			for idx, n := range sortedKeys(attrs) {
				param, err := a.AnalyzeAttribute(n, n, attrs[n].(map[string]interface{}))
				if err != nil {
					return err
				}
				attributes[idx] = &gen.Attribute{n, inflect.Camelize(n), param.Signature()}
			}
		}
	}
	resource.Attributes = attributes

	// Actions
	actions, err := a.AnalyzeActions(name, res)
	if err != nil {
		return err
	}
	resource.Actions = actions

	// Name and done
	resName := toGoTypeName(name, false)
	desc.Resources[resName] = &resource
	desc.ResourceNames = append(desc.ResourceNames, resName)

	return nil
}
