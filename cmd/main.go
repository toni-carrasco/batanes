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
	lp, err := litepage.New("alcalde.senseparaula.cat")
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
	t := template.Must(tmpl.ParseFiles("./view/base.html", "./view/index.html"))
	data := struct {
		Content template.HTML
	}{
		Content: markdown.Parse("content/index.md"),
	}

	return func(w io.Writer) {
		err := t.ExecuteTemplate(w, "base", data)
		if err != nil {
			log.Fatalf("error while rendering template: %v", err)
		}
	}

}
