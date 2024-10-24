package main

import (
	"log"
	"net/http"
)

func main() {
	const port = "8080"
	mux := http.NewServeMux()
	server := &http.Server{
		Handler: mux,
		Addr:    ":" + port,
	}

	mux.Handle("/", http.FileServer(http.Dir(".")))

	log.Printf("Serving on port: %s", port)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to open the server!\n")
	}
}
