package templates

var resourceStructure = `resource "{{type}}" "{{name}}" {
}`

type Resource struct {
	template *Template
}

func (r *Resource) Parse(args map[string]string) string {
	r.template = &Template{
		structure: resourceStructure,
	}

	return r.template.Parse(args)
}
