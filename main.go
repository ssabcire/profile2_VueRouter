package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	files := http.FileServer(http.Dir("template/"))
	mux.Handle("/template/", http.StripPrefix("/template/", files))

	mux.HandleFunc("/", index)

	mux.HandleFunc("/skill", skill)
	mux.HandleFunc("/sns", sns)
	mux.HandleFunc("/using", usingTech)

	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
