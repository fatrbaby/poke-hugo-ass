package parser

import (
	"bytes"
	h2m "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/fatrbaby/poke-hugo-ass/config"
	"log"
	"strings"
)

var converter = h2m.NewConverter("", true, nil)

type MarkdownContent struct {
	Title   string
	Date    string
	Tags    string
	Content string
}

func convert(doc *goquery.Document) *MarkdownContent {
	md := new(MarkdownContent)

	if len(config.Conf.Pattern.Title) > 0 {
		md.Title = doc.Find(config.Conf.Pattern.Title).Text()
	}

	if len(config.Conf.Pattern.Date) > 0 {
		t := doc.Find(config.Conf.Pattern.Date).Text()
		md.Date = strings.Replace(t, ".", "-", -1)
	}

	if len(config.Conf.Pattern.Tags) > 0 {
		attr, exists := doc.Find(config.Conf.Pattern.Tags).Attr("content")

		if exists {
			md.Tags = formatTags(attr)
		}
	}

	if len(config.Conf.Pattern.Content) > 0 {
		sections := doc.Find(config.Conf.Pattern.Content)
		html, err := goquery.OuterHtml(sections)

		if err != nil {
			log.Fatal(err)
		}

		markdown, err := converter.ConvertString(html)
		if err != nil {
			log.Fatal(err)
		}

		md.Content = markdown
	}

	return md
}

func formatTags(s string) string {
	tags := strings.Split(s, ",")
	var buffer bytes.Buffer
	buffer.WriteString("[")

	for _, tag := range tags {
		buffer.WriteString(`"`)
		buffer.WriteString(strings.TrimSpace(tag))
		buffer.WriteString(`",`)
	}

	return strings.Trim(buffer.String(), ",") + "]"
}
