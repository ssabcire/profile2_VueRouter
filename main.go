package main

func main() {
	mux := http.NewServeMux()

	server := http.Server{
		Addr:	"127.0.0.1",
	}
	mux.ListenAndServe()
}
