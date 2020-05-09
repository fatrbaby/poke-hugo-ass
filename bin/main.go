package main

import (
	"github.com/fatrbaby/poke-hugo-ass/config"
	"github.com/fatrbaby/poke-hugo-ass/generator"
	"github.com/fatrbaby/poke-hugo-ass/parser"
)

func main() {
	config.Load("config.yml")
	p := parser.New()
	p.Parse()
	contents := p.Contents()
	g := generator.New(contents)
	g.Generate()
}
