package generator

import (
	"github.com/fatrbaby/poke-hugo-ass/config"
	"github.com/fatrbaby/poke-hugo-ass/parser"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var temp = loadStub()

type Generator struct {
	contents map[string]*parser.MarkdownContent
}

func New(contents map[string]*parser.MarkdownContent) *Generator {
	g := new(Generator)
	g.contents = contents

	return g
}

func (g *Generator) Generate() {
	for s, content := range g.contents {
		c := compile(content)
		if write(s, c) {
			log.Printf("%s generated\n", s)
		} else {
			log.Fatalf("%s generate failed\n", s)
		}
	}
}

func loadStub() string {
	bytes, err := ioutil.ReadFile("generator/markdown.stub")

	if err != nil {
		log.Fatal(err)
	}

	return string(bytes)
}

func write(filename, content string) bool {
	err := ioutil.WriteFile(config.Conf.Dir.Target+"/"+filename, []byte(content), os.ModePerm)

	return err == nil
}

func compile(md *parser.MarkdownContent) string {
	var t string
	t = strings.Replace(temp, "{{title}}", md.Title, -1)
	t = strings.Replace(t, "{{date}}", md.Date, -1)
	t = strings.Replace(t, "{{draft}}", config.Conf.Document.Draft, -1)
	t = strings.Replace(t, "{{content}}", md.Content, -1)
	t = strings.Replace(t, "{{tags}}", md.Tags, -1)

	return t
}
