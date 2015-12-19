package templiceEchoRenderer

import (
	"io"

	"go.iondynamics.net/templice"
)

type Renderer struct {
	Template *templice.Template
}

func New(tpl *templice.Template) *Renderer {
	return &Renderer{Template: tpl}
}

func (r *Renderer) Render(w io.Writer, name string, data interface{}) error {
	return r.Template.ExecuteTemplate(w, name, data)
}