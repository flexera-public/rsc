package main // import "gopkg.in/rightscale/rsc.v4/gen/praxisgen"

import (
	"bitbucket.org/pkg/inflect"
	"gopkg.in/rightscale/rsc.v4/gen"
)

// AnalyzeResource creates an API descriptor from raw resources and types.
func (a *APIAnalyzer) AnalyzeResource(name string, res map[string]interface{}, desc *gen.APIDescriptor) error {
	name = inflect.Singularize(name)
	resource := gen.Resource{Name: name, ClientName: a.ClientName}

	// Description
	if d, ok := res["description"]; ok {
		resource.Description = removeBlankLines(d.(string))
	}

	// Attributes
	hasHref := false
	attributes := []*gen.Attribute{}
	m, ok := res["media_type"].(string)
	if ok {
		t, ok := a.RawTypes[m]
		if ok {
			attrs, ok := t["attributes"].(map[string]interface{})
			if ok {
				attributes = make([]*gen.Attribute, len(attrs))
				for idx, n := range sortedKeys(attrs) {
					if n == "href" {
						hasHref = true
					}
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
	if hasHref {
		resource.LocatorFunc = locatorFunc(name)
	}

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

// locatorFunc returns the source for the function returning the resource locator built from its
// href field.
func locatorFunc(resource string) string {
	return "return api." + resource + "Locator(r.Href)"
}
