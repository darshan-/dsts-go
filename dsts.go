package dsts

import (
	"bytes"
	"html/template"
	"fmt"
)

type page struct {
	Title     string
	Encoding  string

	content   bytes.Buffer

	openPage  string
	closePage string
	openBody  string
	closeBody string
}

type Html5Page struct {
	page
}

func NewHtml5Page() *Html5Page {
	p := new(Html5Page)

	p.Encoding = "utf-8"

	p.openPage = `<!DOCTYPE html>
<html>
<head>
<title>{{.Title}}</title>
<meta http-equiv="Content-Type" content="text/html; charset={{.Encoding}}" />
`
	p.closePage = "</html>\n"
	p.openBody  = "</head>\n<body>\n"
	p.closeBody = "\n</body>\n"

	return p
}

func (p *page) Add(s string) {
	p.content.WriteString(s)
}

func (p page) String() string {
	buf := new(bytes.Buffer)

	tmpl, err := template.New("").Parse(p.openPage)
	if err != nil { panic(err) }
	err = tmpl.Execute(buf, p)
	if err != nil { panic(err) }

	fmt.Fprint(buf, p.openBody, p.content.String(), p.closeBody, p.closePage)

	return buf.String()
}
