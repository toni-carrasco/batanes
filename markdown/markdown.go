package markdown

import (
	"bytes"
	"html/template"
	"log"
	"os"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/renderer/html"
)

var mdParser = goldmark.New(goldmark.WithRendererOptions(html.WithUnsafe()))

func Parse(file string) template.HTML {
	f, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("error reading markdown file: %v", err)
	}
	var buf bytes.Buffer
	err = mdParser.Convert(f, &buf)
	if err != nil {
		log.Fatalf("error parsing file: %v", err)
	}

	return template.HTML(buf.String())
}
