package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/Xavier-Hsiao/Chirpy/internal/api/handlers"
	"github.com/Xavier-Hsiao/Chirpy/internal/api/middleware"
	"github.com/Xavier-Hsiao/Chirpy/internal/config"
	"github.com/Xavier-Hsiao/Chirpy/internal/database"
)

func main() {
	const port = "8080"
	mux := http.NewServeMux()
	server := &http.Server{
		Handler: mux,
		Addr:    ":" + port,
	}

	dbURL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Failed to connect to database!\n")
	}

	dbQueries := database.New(db)
	cfg := config.ApiConfig{
		DBQueries: dbQueries,
	}

	mux.Handle("/app/", middleware.MiddlewareMetricsInc(&cfg, http.StripPrefix("/app", http.FileServer(http.Dir(".")))))
	mux.HandleFunc("GET /admin/metrics", handlers.HandlerMetrics(&cfg))
	mux.HandleFunc("POST /admin/reset", handlers.HandlerReset(&cfg))
	mux.HandleFunc("GET /api/healthz", handlers.HandlerReadiness)
	mux.HandleFunc("POST /api/validate_chirp", handlers.HandlerValidateLength)

	log.Printf("Serving on port: %s", port)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to open the server!\n")
	}
}
