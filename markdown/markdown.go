package markdown

import "strings"

type Markdown struct {
	builder strings.Builder
}

func NewMarkdown() Markdown {
	return Markdown{}
}

func (m *Markdown) Write(comps ...string) {
	for _, c := range comps {
		m.builder.WriteString(c)
	}
}

func (m Markdown) Render() string {
	return m.builder.String()
}
