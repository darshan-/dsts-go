package dsts

import (
	"bytes"
	"html/template"
)

type page struct {
	Title    string
	Encoding string
	Styles   []string
	Scripts  []string

	content  bytes.Buffer
	templ    *template.Template
}

type Html5Page struct {
	page
}

var html5TemplStr =
`<!DOCTYPE html>
<html>
<head>
<title>{{.Title}}</title>
{{range .Styles}}` +
`<link rel="stylesheet" href="{{.}}" type="text/css" />
{{end}}` +
`{{range .Scripts}}` +
`<script type="text/javascript" src="{{.}}"></script>
{{end}}` +
`<meta http-equiv="Content-Type" content="text/html; charset={{.Encoding}}" />
</head>
<body>
{{.Content}}
</body>
</html>
`

func NewHtml5Page() *Html5Page {
	p := new(Html5Page)

	p.Encoding = "utf-8"
	p.templ, _ = template.New("").Parse(html5TemplStr)

	return p
}

func (p *page) Add(s string) {
	p.content.WriteString(s)
}

func (p page) Content() string {
	return p.content.String()
}

func (p page) String() string {
	buf := new(bytes.Buffer)

	err := p.templ.Execute(buf, p)
	if err != nil { panic(err) }

	return buf.String()
}
