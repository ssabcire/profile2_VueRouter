package main

import (
	"html/template"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("template/MyIntroduction.html"))
	err := t.ExecuteTemplate(w, "MyIntroduction.html", nil)
	if err != nil {
		log.Fatal(err)
	}
}
