package template

import (
	"encoding/csv"
	"os"
	"testing"
)

func TestNewTeamplate(t *testing.T) {
	file, err := os.OpenFile("./test.csv", os.O_RDONLY, os.ModeAppend)

	if err == nil {
		t.Log(err)
	}

	temp, input, err := NewTemplate(csv.NewReader(file))

	if err == nil {
		t.Log(err)
	}

	if err := temp.Execute(os.Stdout, input); err != nil {
		t.Log(err)
	}
}
