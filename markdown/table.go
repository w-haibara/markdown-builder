package markdown

import (
	"bytes"

	"github.com/olekukonko/tablewriter"
)

type table struct {
	*tablewriter.Table
}

func Table(header []string, data [][]string) string {
	var buf bytes.Buffer
	table := tablewriter.NewWriter(&buf)
	table.SetHeader(header)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data)
	table.Render()
	return buf.String()
}
