package middleware

import (
	"net/http"

	"github.com/Xavier-Hsiao/Chirpy/internal/config"
)

func MiddlewareMetricsInc(cfg *config.ApiConfig, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.FileServerHits.Add(1)
		next.ServeHTTP(w, r)
	})
}
