package dsts

import (
	"bytes"
)

type page struct {
	Title     string
	content   bytes.Buffer
}

type Page interface {
	Add(string)
	getContent() string

	openPage()  string
	closePage() string
	openBody()  string
	closeBody() string
}

type Html5Page struct {
	page
}

func (p Html5Page) openPage() string {
		return `<!DOCTYPE html>
<html>
<head>
<title>` + stripTags(p.Title) + `</title>
<meta http-equiv="Content-Type" content="text/html; charset=#{encoding}" />
`
}

func (p Html5Page) closePage() string {
	return "</html>\n"
}

func (p Html5Page) openBody() string {
	return "</head>\n<body>\n"
}

func (p Html5Page) closeBody() string {
	return "\n</body>\n"
}

func (p *page) Add(s string) {
	p.content.WriteString(s)
}

func (p page) getContent() string {
	return p.content.String()
}

func ToString(p Page) string {
	return p.openPage() + p.openBody() + p.getContent() + p.closeBody() + p.closePage()
}

func stripTags(s string) string {
//    s.gsub(/(<.*?>)/, "")
	return s
}
