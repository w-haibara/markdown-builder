package markdown

func Bold(str string) string {
	return "**" + str + "**"
}

func Italic(str string) string {
	return "*" + str + "*"
}

func BoldAndItalic(str string) string {
	return "***" + str + "***"
}

func Strikethrough(str string) string {
	return "~~" + str + "~~"
}

func Br() string {
	return "\n\n"
}
