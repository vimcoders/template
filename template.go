package template

import (
	"encoding/csv"
	"fmt"
	"text/template"
)

func NewTemplate(src *csv.Reader) (*template.Template, interface{}, error) {
	tmpl, err := template.ParseFiles("./cfg.template")

	if err != nil {
		return nil, nil, err
	}

	type Input struct {
		Name   string
		Header map[string]string
		Rows   [][]string
	}

	records, err := src.ReadAll()

	if err != nil {
		return nil, nil, err
	}

	fmt.Println(records)

	input := Input{"NetRow", make(map[string]string), records[2:]}

	for i, record := range records[0] {
		input.Header[record] = records[1][i]
	}

	return tmpl, &input, nil
}
