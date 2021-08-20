package main

import (
	_ "embed"

	"html/template"
	"log"
	"net/http"

	"github.com/cixtor/middleware"
)

//go:embed index.html
var htmlIndex string

func main() {
	srv := middleware.New()
	srv.GET("/", index)
	log.Fatal(srv.ListenAndServe(":3000"))
}

func index(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.New("index").Parse(htmlIndex)

	if err != nil {
		panic(err)
	}

	if err := tpl.Execute(w, nil); err != nil {
		panic(err)
	}
}
