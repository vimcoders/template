package gen

import (
	"encoding/csv"
	"text/template"
)

type Template struct {
	input interface{}
	*template.Template
}

func NewTemplate(r *csv.Reader) (*Template, error) {
	tmpl, err := template.ParseFiles("./cfg.template")

	if err != nil {
		return nil, err
	}

	type Input struct {
		Name   string
		Header map[string]string
		Rows   [][]string
	}

	records, err := r.ReadAll()

	if err != nil {
		return nil, err
	}

	input := Input{"test", make(map[string]string), records[2:]}

	for i, record := range records[0] {
		input.Header[record] = records[1][i]
	}

	return &Template{
		Template: tmpl,
		input:    &input,
	}, nil
}
