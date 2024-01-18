package main

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/a-h/templ"
)

//go:embed dist
var staticFiles embed.FS

func main() {

	page := contentPage()

	http.Handle("/", templ.Handler(layoutComponent(page)))

	distFS, _ := fs.Sub(staticFiles, "dist")
	http.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.FS(distFS))))

	http.ListenAndServe(":8081", nil)
}
