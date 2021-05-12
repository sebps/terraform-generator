package templates

var variableStructure = `variable "{{name}}" {
	type = {{type}}	
}`

type Variable struct {
	template *Template
}

func (v *Variable) Render(args map[string]interface{}) string {
	v.template = &Template{
		structure: variableStructure,
	}

	return v.template.Render(args)
}
