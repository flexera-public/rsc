package main // import "gopkg.in/rightscale/rsc.v2/gen/praxisgen"

import (
	"bitbucket.org/pkg/inflect"
	"gopkg.in/rightscale/rsc.v2/gen"
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
	attributes := []*gen.Attribute{}
	m, ok := res["media_type"].(string)
	if ok {
		t, ok := a.RawTypes[m]
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
