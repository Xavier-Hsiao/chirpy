package handlers

import (
	"fmt"
	"net/http"

	"github.com/Xavier-Hsiao/Chirpy/internal/config"
)

// @Summary		Reset Metrics
// @Description	Reset the number of times Chirpy has been visisted to 0
// @Tags			metrics
// @ID				post-reset
// @Produce		plain
// @Success		200	{string}	string	"Hits set to 0"
// @Example		200 "<html><body><h1>Welcome, Chirpy Admin</h1><p>Chirpy has been visited 42 times!</p></body></html>"
// @Router			/admin/reset [post]
func HandlerReset(cfg *config.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cfg.FileServerHits.Store(0)
		w.Header().Add("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Hits set to %d", cfg.FileServerHits.Load())))
	}
}
