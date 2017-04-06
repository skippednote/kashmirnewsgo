package main

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func cli(n news) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Source", "Title"})
	table.SetRowLine(true)
	table.SetColWidth(70)
	table.SetCenterSeparator("*")
	table.SetColumnSeparator("‡")
	table.SetAlignment(tablewriter.ALIGN_LEFT)

	for _, news := range n {
		var v []string
		v = append(v, news.Source)
		v = append(v, news.Title)
		table.Append(v)
	}

	table.Render()
}
