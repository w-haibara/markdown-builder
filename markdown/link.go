package markdown

func Link(title, url string) string {
	return "[" + title + "]" + "(" + url + ")"
}
