package template

import (
	"encoding/csv"
	"fmt"
	"os"
	"testing"
	"text/template"
)

func TestNewTeamplate(t *testing.T) {
	file, err := os.OpenFile("./test.csv", os.O_RDONLY, os.ModeAppend)

	if err == nil {
		t.Log(err)
	}

	d, err := csv.NewReader(file).ReadAll()

	if err != nil {
		t.Log(err)
	}

	tmp, err := template.ParseFiles("./cfg.template")

	if err != nil {
		t.Log(err)
	}

	rowName := fmt.Sprintf("%vRow", file.Name())
	rowsName := fmt.Sprintf("%vRows", file.Name())

	w := NewWriteTo(rowName, rowsName, tmp, d)

	if _, err := w.WriteTo(os.Stdout); err != nil {
		t.Log(err)
	}
}
