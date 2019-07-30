package main

import (
	"html/template"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	//CSS系
	// files := http.FileServer(http.Dir("template/"))
	// mux.Handle("/template/", http.StripPrefix("/template/", files))

	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	//template.html1枚のみを提供する
	t := template.Must(template.ParseFiles("template/template.html"))
	err := t.ExecuteTemplate(w, "template.html", nil)
	if err != nil {
		os.Exit(1)
	}
}
