package main

import (
	"fmt"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/charmbracelet/glamour"
	md "github.com/w-haibara/markdown-builder/markdown"
)

func main() {
	m := md.NewMarkdown()

	m.Write(md.MustHeading("Heading1", 1))
	m.Write(md.MustHeading("Heading2", 2))
	m.Write(md.MustHeading("Heading3", 3))
	m.Write(md.MustHeading("Heading4", 4))
	m.Write(md.MustHeading("Heading5", 5))
	m.Write(md.MustHeading("Heading6", 6))

	m.Write("text text text\n")

	m.Write(md.Br())

	m.Write(heredoc.Doc(`
	long text long text long text long text
	long text long text long text long text
	long text long text long text long text
	`))

	m.Write(md.Br())

	m.Write(md.RootList(
		md.Bold("Bold"),
		md.Italic("Italic"),
		md.Strikethrough("Strikethrough"),
		md.BoldAndItalic("Bold and Italic"),
	))

	m.Write(md.Br())

	m.Write(md.RootNumList(
		"Item1",
		"Item2",
		md.NumList(
			"Item2-1",
			"Item2-2",
			md.NumList(
				"Item2-2-1",
				"Item2-2-2",
			),
			md.List(
				"Item2-2-3",
				"Item2-2-4",
			),
			"Item2-3",
		),
		"Item3",
	))

	m.Write(md.Br())

	m.Write(md.RootList(
		[]string{
			"Item1",
			"Item2",
		},
	))

	m.Write(md.Br())

	m.Write(md.RootTaskList(
		md.Task(true, "Task1"),
		md.Task(false, "Task2"),
		md.TaskList(
			md.Task(true, "Task2-1"),
			md.Task(false, "Task2-2"),
		),
	))

	m.Write(md.Br())

	m.Write(md.Quote(heredoc.Doc(`
		quote text
		quote text
	`)))

	m.Write(md.Br())

	m.Write(md.Code("code quote"), md.Br())
	m.Write(md.CodeBlock(heredoc.Doc(`
		# code block (ruby)
		def foo(x)
			return 3
		end
	`), "ruby"))

	m.Write(md.Br())

	m.Write(md.Link("link title", "https://example.com"), md.Br())

	m.Write(md.Br())

	header := []string{"AAA", "BBB", "CCC"}
	data := [][]string{
		{"111", "222", "333"},
		{"First", "Second", "Third"},
		{"Alpha", "Beta", "Gamma"},
	}
	m.Write(md.Table(header, data))

	out, err := glamour.Render(m.Render(), "dark")
	if err != nil {
		panic(err.Error)
	}
	fmt.Println(out)
}
