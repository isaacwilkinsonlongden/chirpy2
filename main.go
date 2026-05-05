package main

import (
	"log"
	"net/http"
)

func main() {
	const filepathroot = "."
	const port = "8080"

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(filepathroot)))

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathroot, port)
	log.Fatal(srv.ListenAndServe())
}
