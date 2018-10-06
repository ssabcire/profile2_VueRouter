package main

import "net/http"

func skill(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, "layout", "header", "skill")
}

func sns(w http.ResponseWriter, r *http.Request) {
	generateHTML(w, "layout", "header", "sns")
}

func usingTech(w htto.ResponseWriter, r *http.Request) {
	generateHTML(w, "layout", "header", "usingTech")
}
