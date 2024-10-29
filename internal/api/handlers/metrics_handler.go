package handlers

import (
	"fmt"
	"net/http"

	"github.com/Xavier-Hsiao/Chirpy/internal/config"
)

func HandlerMetrics(cfg *config.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestsCount := cfg.FileServerHits.Load()
		w.Header().Add("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Hits: %d", requestsCount)))
	}
}
