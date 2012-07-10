package dsts

import (
	"bytes"
)

type page struct {
	content bytes.Buffer
}

type Html5Page struct {
	page
}

func (p *page) Add(s string) {
	p.content.WriteString(s)
}

func (p *page) String() string {
	return p.content.String()
}
