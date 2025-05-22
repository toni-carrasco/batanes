package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"time"

	"github.com/man-on-box/litepage"
	"github.com/toni-carrasco/batanes/markdown"
)

func main() {
	lp, err := litepage.New("toni-carrasco.github.io", litepage.WithBasePath("/batanes"))
	if err != nil {
		log.Fatalf("Could not create app: %v", err)
	}

	lp.Page("/index.html", handleHomepage())

	err = lp.BuildOrServe()
	if err != nil {
		log.Fatalf("Could not run app: %v", err)
	}
}

var tmpl = template.New("").Funcs(template.FuncMap{
	"version": func() string {
		return fmt.Sprintf("%d", time.Now().Unix())
	},
})

func handleHomepage() func(w io.Writer) {
	content := markdown.Parse("content/index.md")
	t := template.Must(tmpl.ParseFiles("./view/base.html", "./view/index.html"))
	data := struct {
		Content template.HTML
	}{
		Content: content,
	}

	return func(w io.Writer) {
		err := t.ExecuteTemplate(w, "base", data)
		if err != nil {
			log.Printf("error while rendering template: %v", err)
		}
	}

}
