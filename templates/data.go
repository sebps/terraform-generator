package templates

var dataStructure = `data "{{type}}" "{{name}}" {
}`

type Data struct {
	template *Template
}

func (d *Data) Render(args map[string]interface{}) string {
	d.template = &Template{
		structure: resourceStructure,
	}

	return d.template.Render(args)
}
