package main

import (
	"log"
	"net/http"
)

func main() {
	const filepathRoot = "."
	const port = "8080"

	mux := http.NewServeMux()
	mux.Handle("/app/",http.StripPrefix("/app/", http.FileServer(http.Dir(filepathRoot))))

	mux.HandleFunc("/healthz", func (res http.ResponseWriter,  request *http.Request) {
		res.Header().Add("Content-Type", "text/plain; charset=utf-8")
		res.WriteHeader(200)
		res.Write([]byte("OK"))
	})
	srv := &http.Server{
		Handler: mux, 
		Addr: ":" + port,
	}
	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}
