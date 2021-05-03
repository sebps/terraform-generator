package types

type AttributeSchema struct {
	Name            string
	Type            interface{}
	DescriptionKind string
	Required        bool
	Computed        bool
	Optional        bool
}

type ResourceSchema struct {
	Name       string
	Attributes []*AttributeSchema
}

type ProviderSchema struct {
	Name        string
	Registry    string
	Resources   []*ResourceSchema
	DataSources []*ResourceSchema
}
