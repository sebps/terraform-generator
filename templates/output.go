package templates

var outputStructure = `output "{{name}}" {
	value = {{value}}	
}`

type Output struct {
	template *Template
}

func (o *Output) Parse(args map[string]string) string {
	o.template = &Template{
		structure: outputStructure,
	}

	return o.template.Parse(args)
}
