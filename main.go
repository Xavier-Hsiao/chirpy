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
	mux.HandleFunc("GET /admin/metrics", handlers.HandlerMetrics(&cfg))
	mux.HandleFunc("POST /admin/reset", handlers.HandlerReset(&cfg))
	mux.HandleFunc("GET /api/healthz", handlers.HandlerReadiness)
	mux.HandleFunc("POST /api/validate_chirp", handlers.HandlerValidateLength)

	log.Printf("Serving on port: %s", port)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to open the server!\n")
	}
}
