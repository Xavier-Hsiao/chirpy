package main

import (
	"log"
	"net/http"

	"github.com/Xavier-Hsiao/Chirpy/internal/api/handlers"
)

func main() {
	const port = "8080"
	mux := http.NewServeMux()
	server := &http.Server{
		Handler: mux,
		Addr:    ":" + port,
	}

	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir("."))))
	mux.HandleFunc("/healthz", handlers.HandlerReadiness)

	log.Printf("Serving on port: %s", port)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to open the server!\n")
	}
}
