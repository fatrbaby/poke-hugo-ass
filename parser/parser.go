package parser

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/fatrbaby/poke-hugo-ass/config"
	"log"
	"os"
	"path"
)

type Parser struct {
	contents map[string]*MarkdownContent
}

func New() *Parser {
	p := new(Parser)
	p.contents = make(map[string]*MarkdownContent)

	return p
}

func (p *Parser) Parse() {
	var mdName string

	s := newScanner(config.Conf.Dir.Html)
	s.scan()

	for _, result := range s.results {
		mdName = toMarkdownName(result)
		document := readAsDocument(result)
		p.contents[mdName] = convert(document)
	}
}

func (p *Parser) Contents() map[string]*MarkdownContent {
	return p.contents
}

func readAsDocument(file string) *goquery.Document {
	f, err := os.Open(file)

	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(f)

	if err != nil {
		log.Fatal(err)
	}

	return doc
}

func toMarkdownName(name string) string {
	dir := path.Dir(name)
	base := path.Base(dir)

	return base + ".md"
}
