package templates

var outputStructure = `output "{{name}}" {
	value = {{value}}	
}`

type Output struct {
	template *Template
}

func (o *Output) Render(args map[string]interface{}) string {
	o.template = &Template{
		structure: outputStructure,
	}

	return o.template.Render(args)
}
