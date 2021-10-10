package template

import (
	"fmt"
	"io"
	"text/template"
)

type WriteTo struct {
	Data     [][]string
	RowName  string
	RowsName string
	tmp      *template.Template
}

func (wt *WriteTo) WriteTo(w io.Writer) (n int64, err error) {
	type Input struct {
		Row    string
		Rows   string
		Header []string
		Data   [][]string
	}

	input := Input{Row: wt.RowName, Rows: wt.RowsName, Data: wt.Data}

	for i, record := range wt.Data[0] {
		input.Header = append(input.Header, fmt.Sprintf("%v %v", record, wt.Data[1][i]))
	}

	if err := wt.tmp.Execute(w, input); err != nil {
		return 0, err
	}

	return 0, nil
}

func NewWriteTo(rowName, rowsName string, t *template.Template, d [][]string) io.WriterTo {
	return &WriteTo{RowName: rowName, RowsName: rowsName, tmp: t, Data: d}
}
