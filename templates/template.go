package templates

import "strings"

type Template struct {
	structure string
}

func (t *Template) Parse(args map[string]string) string {
	result := t.structure

	for k, v := range args {
		result = strings.ReplaceAll(result, "{{"+k+"}}", v)
	}

	return result
}
