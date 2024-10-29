package main

import (
	"log"
	"net/http"

	"github.com/Xavier-Hsiao/Chirpy/internal/api/handlers"
	"github.com/Xavier-Hsiao/Chirpy/internal/api/middleware"
	"github.com/Xavier-Hsiao/Chirpy/internal/config"
)

func main() {
	const port = "8080"
	mux := http.NewServeMux()
	server := &http.Server{
		Handler: mux,
		Addr:    ":" + port,
	}

	cfg := config.ApiConfig{}

	mux.Handle("/app/", middleware.MiddlewareMetricsInc(&cfg, http.StripPrefix("/app", http.FileServer(http.Dir(".")))))
	mux.HandleFunc("/healthz", handlers.HandlerReadiness)
	mux.HandleFunc("/metrics", handlers.HandlerMetrics(&cfg))
	mux.HandleFunc("/reset", handlers.HandlerReset(&cfg))

	log.Printf("Serving on port: %s", port)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to open the server!\n")
	}
}
