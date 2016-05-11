package metadata

type Attribute struct {
	Name      string // Attribute name, e.g. "elasticity_params"
	FieldName string // Corresponding go struct field name, e.g. "ElasticityParams"
	FieldType string // Corresponding go struct type, e.g. "*ElasticityParams"
}
