package templates

var providerStructure = `provider "{{name}}" {
}`

type Provider struct {
	template *Template
}

func (p *Provider) Render(args map[string]interface{}) string {
	p.template = &Template{
		structure: providerStructure,
	}

	return p.template.Render(args)
}
