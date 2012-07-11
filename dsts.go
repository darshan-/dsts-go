package dsts

import (
	"bytes"
)

type Page struct {
	Title     string
	content   bytes.Buffer
	openPage  func() string
	closePage func() string
	openBody  func() string
	closeBody func() string
}

func NewHtml5Page() (*Page) {
	p := new(Page)

	p.openPage = func () string {
		return `<!DOCTYPE html>
<html>
<head>
<title>` + stripTags(p.Title) + `</title>
<meta http-equiv="Content-Type" content="text/html; charset=#{encoding}" />
`
}

	p.closePage = func() string {
		return "</html>\n"
	}

	p.openBody = func() string {
		return "</head>\n<body>\n"
	}

	p.closeBody = func() string {
		return "\n</body>\n"
	}

	return p
}

func (p *Page) Add(s string) {
	p.content.WriteString(s)
}

func (p *Page) String() string {
	return p.openPage() + p.openBody() + p.content.String() + p.closeBody() + p.closePage()
}

func stripTags(s string) string {
//    s.gsub(/(<.*?>)/, "")
	return s
}
