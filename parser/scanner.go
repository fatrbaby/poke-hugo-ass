package parser

import (
	"github.com/fatrbaby/poke-hugo-ass/config"
	"os"
	"path/filepath"
	"strings"
)

type scanner struct {
	dir     string
	results []string
}

func newScanner(dir string) *scanner {
	s := new(scanner)
	s.dir = dir

	return s
}

func (s *scanner) scan() {
	indexPage := config.Conf.Dir.Html + "/index.html"

	filepath.Walk(s.dir, func(path string, info os.FileInfo, err error) error {
		if path != indexPage && strings.HasSuffix(path, ".html") {
			s.results = append(s.results, path)
		}

		return nil
	})
}
