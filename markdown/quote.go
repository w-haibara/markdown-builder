package markdown

import (
	"strings"
)

func Quote(str string) string {
	lines := strings.Split(str, "\n")
	for i, line := range lines {
		lines[i] = "> " + line
	}

	return strings.Join(lines, "\n")
}
