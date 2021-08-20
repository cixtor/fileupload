package main

import (
	_ "embed"

	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/cixtor/middleware"
)

var directory string

//go:embed index.html
var htmlIndex string

func main() {
	flag.StringVar(&directory, "d", ".", "Destination folder to store the uploaded files")
	flag.Parse()

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

func uploadThisFile(fileHeader *multipart.FileHeader) error {
	file, err := fileHeader.Open()

	if err != nil {
		return fmt.Errorf("uploadThisFile %w", err)
	}

	defer file.Close()

	dst, err := os.Create(directory + "/" + fileHeader.Filename)

	if err != nil {
		return fmt.Errorf("uploadThisFile %w", err)
	}

	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		return fmt.Errorf("uploadThisFile %w", err)
	}

	return nil
}
