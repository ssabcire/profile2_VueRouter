package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func generateHTML(w http.ResponseWriter, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("template/%s.html", file))
		templates := template.Must(template.ParseFiles(files...))
		templates.ExecuteTemplate(w, "layout", nil)
	}
}
