package handlers

import (
	"net/http"
	"time"

	"github.com/Xavier-Hsiao/Chirpy/internal/auth"
	"github.com/Xavier-Hsiao/Chirpy/internal/config"
	"github.com/Xavier-Hsiao/Chirpy/internal/helpers"
)

type refreshRsp struct {
	Token string `json:"token"`
}

func HandlerRefreshJWT(cfg *config.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get refresh token from request header
		refreshToken, err := auth.GetBearerToken(r.Header)
		if err != nil {
			helpers.RespondWithError(w, http.StatusBadRequest, "Miss refresh token", err)
			return
		}

		// Generate new access token (refreshing)
		// Wait... we need a query to get user ID from refresh token...
		user, err := cfg.DBQueries.GetUserFromRefreshToken(r.Context(), refreshToken)
		if err != nil {
			helpers.RespondWithError(w, http.StatusUnauthorized, "Failed to get user by refresh token", err)
			return
		}

		newAccessToken, err := auth.MakeJWT(user.ID, cfg.JWTSecret, time.Hour)
		if err != nil {
			helpers.RespondWithError(w, http.StatusUnauthorized, "Failed to generate access token", err)
			return
		}

		helpers.RespondWithJson(w, http.StatusOK, refreshRsp{
			Token: newAccessToken,
		})
	}
}
