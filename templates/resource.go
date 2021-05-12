package templates

var resourceStructure = `resource "{{type}}" "{{name}}" {
}`

type Resource struct {
	template *Template
}

func (r *Resource) Render(args map[string]interface{}) string {
	r.template = &Template{
		structure: resourceStructure,
	}

	return r.template.Render(args)
}
