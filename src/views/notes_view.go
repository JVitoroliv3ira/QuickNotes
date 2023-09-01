package views

import (
	"github.com/olekukonko/tablewriter"
	"os"
	"quick_notes/src/types"
)

func Render(notes types.Notes) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Id", "Conteúdo"})
	table.SetAlignment(1)

	for _, note := range notes.Data {
		table.Append([]string{
			note.Id,
			note.Text,
		})
	}

	table.Render()
}
