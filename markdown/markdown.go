package markdown

import (
	"bytes"
	"html/template"
	"log"
	"os"

	"github.com/yuin/goldmark"
)

func Parse(file string) template.HTML {
	f, err := os.ReadFile(file)
	if err != nil {
		log.Fatalf("error reading markdown file: %v", err)
	}
	var buf bytes.Buffer
	err = goldmark.Convert(f, &buf)
	if err != nil {
		log.Fatalf("error parsing file: %v", err)
	}

	return template.HTML(buf.String())
}
