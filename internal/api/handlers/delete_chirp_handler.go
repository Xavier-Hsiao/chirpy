package handlers

import (
	"net/http"

	"github.com/Xavier-Hsiao/Chirpy/internal/auth"
	"github.com/Xavier-Hsiao/Chirpy/internal/config"
	"github.com/Xavier-Hsiao/Chirpy/internal/helpers"
	"github.com/google/uuid"
)

func HandlerDeleteChirp(cfg *config.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		chirpIDString := r.PathValue("chirpID")
		chirpID, err := uuid.Parse(chirpIDString)
		if err != nil {
			helpers.RespondWithError(w, http.StatusBadRequest, "Malfromed chirp ID", err)
			return
		}

		accessToken, err := auth.GetBearerToken(r.Header)
		if err != nil {
			helpers.RespondWithError(w, http.StatusUnauthorized, "Miss JWT token", err)
			return
		}
		userID, err := auth.ValidateJWT(accessToken, cfg.JWTSecret)
		if err != nil {
			helpers.RespondWithError(w, http.StatusUnauthorized, "Failed to validate JWT token", err)
			return
		}

		dbChirp, err := cfg.DBQueries.GetChirpById(r.Context(), chirpID)
		if err != nil {
			helpers.RespondWithError(w, http.StatusNotFound, "Failed to find the chirp based on your account", err)
			return
		}
		if dbChirp.UserID != userID {
			helpers.RespondWithError(w, http.StatusForbidden, "You're not allowed to delete this chirp", err)
			return
		}

		err = cfg.DBQueries.DeleteChirp(r.Context(), chirpID)
		if err != nil {
			helpers.RespondWithError(w, http.StatusInternalServerError, "Failed to delete the chirp", err)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}

}
