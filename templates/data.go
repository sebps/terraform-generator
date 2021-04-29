package templates

var dataStructure = `data "{{type}}" "{{name}}" {
}`

type Data struct {
	template *Template
}

func (d *Data) Parse(args map[string]string) string {
	d.template = &Template{
		structure: resourceStructure,
	}

	return d.template.Parse(args)
}
