package templates

import (
	"github.com/sebps/template-engine/rendering"
)

type Template struct {
	structure string
}

func (t *Template) Render(args map[string]interface{}) string {
	return rendering.Render(t.structure, args)
}
