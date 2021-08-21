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

var maxUploadSize int64

//go:embed index.html
var htmlIndex string

func main() {
	flag.StringVar(&directory, "d", ".", "Destination folder to store the uploaded files")
	flag.Int64Var(&maxUploadSize, "m", 20<<20, "Maximum upload size for each file in bytes")
	flag.Parse()

	srv := middleware.New()
	srv.GET("/", index)
	srv.POST("/upload", upload)
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

func upload(w http.ResponseWriter, r *http.Request) {
	// Store up to 200 MiB of data in RAM.
	if err := r.ParseMultipartForm(200 << 20); err != nil {
		panic(err)
	}

	files, ok := r.MultipartForm.File["files"]

	if !ok {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	for _, fileHeader := range files {
		if fileHeader.Size > maxUploadSize {
			http.Error(w, "maximum upload size exceeded by "+fileHeader.Filename, http.StatusBadRequest)
			return
		}

		if err := uploadThisFile(fileHeader); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	http.Redirect(w, r, "/?success=true", http.StatusFound)
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
