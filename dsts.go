package dsts

import (
	"bytes"
	"strings"
)

type Page struct {
	preContent  string
	content     bytes.Buffer
	postContent string
}

type HtmlPage struct {
	Page

	Title      string
	Encoding   string
	HeadExtras string

	// Exported, but meant as "protected" fields, to be optionally modified by
	//  "subclass" outside of this package
	DocTop     string
	HeadTop    string
	HeadBottom string
	BodyTop    string
	BodyAttrs  string
	BodyBottom string
	DocBottom  string

	styles  []string
	scripts []string
}

type Html5Page struct {
	*HtmlPage
}

type XhtmlPage struct {
	*HtmlPage

	Doctype string
}

func (p *Page) String() string {
	return p.preContent + p.content.String() + p.postContent
}

func (p *HtmlPage) String() string {
	p.preContent = p.DocTop + p.HeadTop +
		p.titleStr() + p.stylesStr() + p.scriptsStr() + p.contentTypeStr() +
		p.HeadExtras + p.HeadBottom + p.BodyTop
	p.postContent = p.BodyBottom + p.DocBottom

	return p.Page.String()
}

func (p *XhtmlPage) String() string {
	p.DocTop = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 ` +
		strings.ToUpper(p.Doctype) +
		`//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-` +
		strings.ToLower(p.Doctype) +
		`.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en" lang="en">
`

	return p.HtmlPage.String()
}

func NewHtmlPage() *HtmlPage {
	p := new(HtmlPage)

	p.DocTop     = "<html>\n"
	p.HeadTop    = "  <head>\n"
	p.HeadBottom = "  </head>\n"
	p.BodyTop    = "<body>\n"
	p.BodyBottom = "</body>\n"
	p.DocBottom  = "</html>\n"

	p.Encoding = "utf-8"

	return p
}

func NewHtml5Page() *Html5Page {
	p := new(Html5Page)

	p.HtmlPage = NewHtmlPage()

	p.DocTop = "<!DOCTYPE html>\n" + p.DocTop

	return p
}

func NewXhtmlPage() *XhtmlPage {
	p := new(XhtmlPage)

	p.HtmlPage = NewHtmlPage()

	p.DocTop  = ""
	p.HeadTop = `  <head profile="http://www.w3.org/2005/10/profile">` + "\n"

	p.Doctype = "strict"

	return p
}

func (p *HtmlPage) titleStr() string {
	return "    <title>" + p.Title + "</title>\n"
}

func (p *HtmlPage) stylesStr() (s string) {
	for _, style := range p.styles {
		s += `    <link rel="stylesheet" href="` + style + `" type="text/css" />` + "\n"
	}

	return s
}

func (p *HtmlPage) scriptsStr() (s string) {
	for _, script := range p.scripts {
		s += `    <script type="text/javascript" src="` + script + `"></script>` + "\n"
	}

	return s
}

func (p *HtmlPage) contentTypeStr() string {
	return `    <meta http-equiv="Content-Type" content="text/html; charset=` +
		p.Encoding + `" />` + "\n"
}

func (p *Page) Add(s string) {
	p.content.WriteString(s)
}

func (p *HtmlPage) AddStyle(s string) {
	p.styles = append(p.styles, s)
}

func (p *HtmlPage) AddScript(s string) {
	p.scripts = append(p.scripts, s)
}
