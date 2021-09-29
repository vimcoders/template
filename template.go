package template

import (
	"encoding/csv"
	"fmt"
	"text/template"
)

type Input struct {
	Row    string
	Rows   string
	Header []string
	Data   [][]string
}

func NewTemplate(src *csv.Reader) (*template.Template, interface{}, error) {
	tmpl, err := template.ParseFiles("./cfg.template")

	if err != nil {
		return nil, nil, err
	}

	records, err := src.ReadAll()

	if err != nil {
		return nil, nil, err
	}

	input := Input{Row: "NetRow", Rows: "NetRows", Data: records[2:]}

	for i, record := range records[0] {
		input.Header = append(input.Header, fmt.Sprintf("%v %v", record, records[1][i]))
	}

	return tmpl, &input, nil
}
