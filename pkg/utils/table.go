package utils

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
)

func Table(h table.Row) *table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	if nil != h {
		t.AppendHeader(h)
	}

	t.Style().Options.DrawBorder = false
	t.Style().Options.SeparateColumns = false
	t.Style().Options.SeparateFooter = false
	t.Style().Options.SeparateHeader = false
	t.Style().Options.SeparateRows = false
	return &t
}
