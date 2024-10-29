package handlers

import (
	"fmt"
	"net/http"

	"github.com/Xavier-Hsiao/Chirpy/internal/config"
)

func HandlerReset(cfg *config.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cfg.FileServerHits.Store(0)
		w.Header().Add("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Hits set to %d", cfg.FileServerHits.Load())))
	}
}
