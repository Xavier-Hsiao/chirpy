package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Xavier-Hsiao/Chirpy/internal/config"
	"github.com/Xavier-Hsiao/Chirpy/internal/helpers"
)

// @Summary		Reset Metrics
// @Description	Delete all users in database and reset the number of times Chirpy has been visisted to 0
// @Tags			user
// @ID				post-reset
// @Produce		plain
// @Success		200	{string}	string					"Hits set to 0"
// @Failure		500	{object}	helpers.ErrorResponse	"Failed to delete users in db"
// @Router			/admin/reset [post]
func HandlerReset(cfg *config.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check reset authorization
		if cfg.Platform != "dev" {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Rest endpoint is only allowed in dev environment"))
			return
		}

		// Delete all users in database
		err := cfg.DBQueries.DeleteUsers(r.Context())
		if err != nil {
			helpers.RespondWithError(w, http.StatusInternalServerError, "Failed to delete users in db", err)
			return
		}
		log.Println("Success: all users are deleted!")

		// Reset the vistied number
		cfg.FileServerHits.Store(0)
		w.Header().Add("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Hits set to %d", cfg.FileServerHits.Load())))
	}
}
