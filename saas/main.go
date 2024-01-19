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
	col := []string{"", "Name", "Job", "Favourite Color"}
	data := [][]string{
		{"1", "Cy Genderton", "Quality Control Specialist", "Blue"},
		{"2", "Hart Hagerty", "Desktop Support Technician", "Indigo"},
		{"3", "Brice Swyer", "Tax Accountant", "Red"},
	}
	tableData := TableData{
		Column: col,
		Data:   data,
	}
	page := contentPage(tableData)

	http.Handle("/", templ.Handler(layoutComponent(page)))

	distFS, _ := fs.Sub(staticFiles, "dist")
	http.Handle("/dist/", http.StripPrefix("/dist/", http.FileServer(http.FS(distFS))))

	http.ListenAndServe(":8081", nil)
}
