package dsts

import (
	"bytes"
	"html/template"
	"fmt"
)

type page struct {
	Title    string
	Encoding string

	content  bytes.Buffer

	preContTempl  *template.Template
	postContTempl *template.Template
}

type Html5Page struct {
	page
}

func NewHtml5Page() *Html5Page {
	p := new(Html5Page)

	p.Encoding = "utf-8"

	pre := `<!DOCTYPE html>
<html>
<head>
<title>{{.Title}}</title>
<meta http-equiv="Content-Type" content="text/html; charset={{.Encoding}}" />
</head>
<body>
`
	post := `</body>
</html>
`

	var err error
	p.preContTempl, err = template.New("").Parse(pre)
	if err != nil { panic(err) }

	p.postContTempl, err = template.New("").Parse(post)
	if err != nil { panic(err) }

	return p
}

func (p *page) Add(s string) {
	p.content.WriteString(s)
}

func (p page) String() string {
	buf := new(bytes.Buffer)

	err := p.preContTempl.Execute(buf, p)
	if err != nil { panic(err) }

	fmt.Fprintln(buf, p.content.String())

	err = p.postContTempl.Execute(buf, p)
	if err != nil { panic(err) }

	return buf.String()
}
