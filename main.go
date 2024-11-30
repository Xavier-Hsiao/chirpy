package main

import (
	"database/sql"
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
	const port = "8080"

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mux := http.NewServeMux()

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL must be set in your .env file")
	}

	// Use it to protect /admin/reset endpoint
	platform := os.Getenv("PLATFORM")
	if platform == "" {
		log.Fatal("PLATFORM must be set")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if platform == "" {
		log.Fatal("JWT_SECRET must be set")
	}

	dbConn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Failed to connect to database!\n")
	}

	dbQueries := database.New(dbConn)
	cfg := config.ApiConfig{
		DBQueries: dbQueries,
		Platform:  platform,
		JWTSecret: jwtSecret,
	}

	mux.Handle("/app/", middleware.MiddlewareMetricsInc(&cfg, http.StripPrefix("/app", http.FileServer(http.Dir(".")))))

	// API routes with separate handlers
	mux.HandleFunc("GET /admin/metrics", handlers.HandlerMetrics(&cfg))
	mux.HandleFunc("POST /admin/reset", handlers.HandlerReset(&cfg))
	mux.HandleFunc("GET /api/healthz", handlers.HandlerReadiness)
	mux.HandleFunc("POST /api/validate_chirp", handlers.HandlerValidateLength)

	mux.HandleFunc("POST /api/users", handlers.HandlerCreateUser(&cfg))
	mux.HandleFunc("PUT /api/users", handlers.HandlerUpdateUsers(&cfg))
	mux.HandleFunc("POST /api/login", handlers.HandlerLogin(&cfg))

	mux.HandleFunc("POST /api/refresh", handlers.HandlerRefreshJWT(&cfg))
	mux.HandleFunc("POST /api/revoke", handlers.HandlerRevokeJWT(&cfg))

	mux.HandleFunc("POST /api/chirps", handlers.HandlerCreateChirp(&cfg))
	mux.HandleFunc("GET /api/chirps", handlers.HandlerGetChirps(&cfg))
	mux.HandleFunc("GET /api/chirps/{chirpID}", handlers.HandlerGetChirpById(&cfg))
	mux.HandleFunc("DELETE /api/chirps/{chirpID}", handlers.HandlerDeleteChirp(&cfg))

	mux.HandleFunc("POST /api/polka/webhooks", handlers.HandlerUpgradeUser(&cfg))

	// Swagger UI endpoint
	mux.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), // The URL pointing to your Swagger docs
	))

	server := &http.Server{
		Handler: mux,
		Addr:    ":" + port,
	}

	log.Printf("Serving on port: %s", port)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to open the server!\n")
	}
}
