package markdown

import (
	"fmt"
	"strings"
)

func MustHeading(str string, i int) string {
	c, err := Heading(str, i)
	if err != nil {
		panic(err.Error())
	}

	return c
}

func Heading(str string, i int) (string, error) {
	if i < 1 || i > 6 {
		return string(""), fmt.Errorf("invalid index for heading: %d", i)
	}

	return strings.Repeat("#", i) + " " + str + "\n", nil
}
