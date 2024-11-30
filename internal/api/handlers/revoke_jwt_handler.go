package handlers

import (
	"net/http"

	"github.com/Xavier-Hsiao/Chirpy/internal/auth"
	"github.com/Xavier-Hsiao/Chirpy/internal/config"
	"github.com/Xavier-Hsiao/Chirpy/internal/helpers"
)

func HandlerRevokeJWT(cfg *config.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get refresh token from the request
		refreshToken, err := auth.GetBearerToken(r.Header)
		if err != nil {
			helpers.RespondWithError(w, http.StatusBadRequest, "Miss refresh token", err)
			return
		}

		// Operate refresh token revoking
		_, err = cfg.DBQueries.RevokeRefreshToken(r.Context(), refreshToken)
		if err != nil {
			helpers.RespondWithError(w, http.StatusUnauthorized, "Failed to revoke the token", err)
			return
		}

		// Respond with 204 status code
		w.WriteHeader(204)
	}
}
