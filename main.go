package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/Xavier-Hsiao/Chirpy/docs"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/Xavier-Hsiao/Chirpy/internal/api/handlers"
	"github.com/Xavier-Hsiao/Chirpy/internal/api/middleware"
	"github.com/Xavier-Hsiao/Chirpy/internal/config"
	"github.com/Xavier-Hsiao/Chirpy/internal/database"
)

// @title			Chirpy API
// @version		1.0
// @description	This is the API server for Chirpy application
// @host			localhost:8080
// @BasePath		/
func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	const port = "8080"
	mux := http.NewServeMux()
	server := &http.Server{
		Handler: mux,
		Addr:    ":" + port,
	}

	dbURL := os.Getenv("DB_URL")
	fmt.Println("Database URL:", dbURL)
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Failed to connect to database!\n")
	}

	dbQueries := database.New(db)
	cfg := config.ApiConfig{
		DBQueries: dbQueries,
	}

	mux.Handle("/app/", middleware.MiddlewareMetricsInc(&cfg, http.StripPrefix("/app", http.FileServer(http.Dir(".")))))

	// API routes with separate handlers
	mux.HandleFunc("GET /admin/metrics", handlers.HandlerMetrics(&cfg))
	mux.HandleFunc("POST /admin/reset", handlers.HandlerReset(&cfg))
	mux.HandleFunc("GET /api/healthz", handlers.HandlerReadiness)
	mux.HandleFunc("POST /api/validate_chirp", handlers.HandlerValidateLength)
	mux.HandleFunc("POST /api/users", handlers.HandlerCreateUser(&cfg))

	// Swagger UI endpoint
	mux.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), // The URL pointing to your Swagger docs
	))

	log.Printf("Serving on port: %s", port)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to open the server!\n")
	}
}
