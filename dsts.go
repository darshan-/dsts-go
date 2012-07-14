package dsts

/*******************************************************************************
 *
 * Client interface:
 *
 *  1. Instantiate a page creation object, likely specifiying the format it will generate
 *    e.g. p := NewHtml5Page()
 *
 *  2. Optionally set properties and add content
 *    e.g. p.Title = "Help"
 *         p.AddScript("default.css")
 *         p.Add("<p>Once upon a time...</p>\n")
 *
 *  3. Generate the page
 *    e.g. p.String()
 *
 *  4. And the tricky part: Make it extensible, such that a website named Egg Sample could easily
 *    leverage the library to make new formats based on existing ones, which themselves can form the
 *    basis of subformats:
 *      e.g. p  := NewEggSamplePage()
 *           p2 := NewEggSampleForumPage()
 *
 *
 *  Points where sub-formats can add content:
 *    * At the end of the html <head>
 *    * At the beginning of the html <body>
 *    * At the end of the html <body>
 *
 *  So HtmlPage struct has a string or bytes.Buffer for each of those,
 *    the constructor for each of those calls it's "super's" constructor,
 *    then adds to those
 *
 *  You get either the top of Xhtml or Html5, depending on which you chose, then
 *    The title and header according to your title, scripts, and styles
 *    Then the extra headers stuff, first from the super then decending down to the final subtype last
 *    Then the extra body-top stuff, first from the super, decending down to the final subtype last
 *    Then the extra body-bottom stuff, first from the super, decending down to the final subtype last
 *    The end of all html-type pages
 *
 *
 ******************************************************************************/




import (
	"bytes"
	"strings"
)

type page struct {
	Title    string

	content  bytes.Buffer
}

type HtmlPage struct {
	page

	Encoding   string
	HeaderMisc string

	styles   []string
	scripts  []string
}

type Html5Page struct {
	HtmlPage
}

type XhtmlPage struct {
	HtmlPage

	Doctype string
}

type pageStringer interface {
	openPage()   string
	openBody()   string
	contentStr() string
	closeBody()  string
	closePage()  string
}

type htmlStringer interface {
	pageStringer

	openHead()   string
	titleStr()   string
	stylesStr()  string
	scriptsStr() string
	contentTypeStr() string
}

func PageString(p pageStringer) string {
	return headerString(p) + p.contentStr() + footerString(p)
}

func headerString(p pageStringer) string {
	return p.openPage() + p.openBody()
}

func HtmlPageString(p htmlStringer) string {
	return htmlHeaderString(p) + p.contentStr() + footerString(p)
}

func htmlHeaderString(p htmlStringer) string {
	return p.openPage() +
		p.openHead() + p.titleStr() + p.stylesStr() + p.scriptsStr() + p.contentTypeStr() +
		p.openBody()
}

func footerString(p pageStringer) string {
	return p.closeBody() + p.closePage()
}

func (p *page) contentStr() string {
	return p.content.String()
}

func (p *HtmlPage) openPage() string {
	return "<html>\n"
}

func (p *Html5Page) openPage() string {
	return "<!DOCTYPE html>\n<html>\n"
}

func (p *XhtmlPage) openPage() string {
	return `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 ` +
		strings.ToUpper(p.Doctype) +
		`//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-` +
	strings.ToLower(p.Doctype) +
		`.dtd">
<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en" lang="en">`
}

func (p *HtmlPage) openBody() string {
	return p.HeaderMisc + "  </head>\n<body>\n"
}

func (p *HtmlPage) closeBody() string {
	return "</body>\n"
}

func (p *HtmlPage) closePage() string {
	return "</html>\n"
}

func (p *HtmlPage) openHead() string {
	return "  <head>\n"
}

func (p *XhtmlPage) openHead() string {
	return `  <head profile="http://www.w3.org/2005/10/profile">` + "\n"
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

func NewHtml5Page() *Html5Page {
	p := new(Html5Page)

	p.Encoding = "utf-8"

	return p
}

func NewXhtmlPage() *XhtmlPage {
	p := new(XhtmlPage)

	p.Encoding = "utf-8"
	p.Doctype  = "strict"

	return p
}

func (p *page) Add(s string) {
	p.content.WriteString(s)
}

func (p *HtmlPage) AddStyle(s string) {
	p.styles = append(p.styles, s)
}

func (p *HtmlPage) AddScript(s string) {
	p.scripts = append(p.scripts, s)
}
