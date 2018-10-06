package main

import "net/http"

func index(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, "layout", "header", "index")

	//t := template.Must(template.ParseFiles("template/MyIntroduction.html"))
	//err := t.ExecuteTemplate(w, "MyIntroduction.html", nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
}
