package gen

import (
	"text/template"
)

type Template struct {
	input interface{}
	*template.Template
}

func NewTemplate(csvPath string) (*Template, error) {
	tmpl, err := template.ParseFiles("./cfg.template")

	if err != nil {
		return nil, err
	}

	return &Template{
		Template: tmpl,
	}, nil
}
