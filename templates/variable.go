package templates

var variableStructure = `variable "{{name}}" {
	type = {{type}}	
}`

type Variable struct {
	template *Template
}

func (v *Variable) Parse(args map[string]string) string {
	v.template = &Template{
		structure: resourceStructure,
	}

	return v.template.Parse(args)
}
