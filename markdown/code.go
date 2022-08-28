package markdown

func Code(str string) string {
	return "`" + str + "`"
}

func CodeBlock(str, lang string) string {
	return "```" + lang + "\n" + str + "\n```\n"
}
