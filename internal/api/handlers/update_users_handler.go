package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Xavier-Hsiao/Chirpy/internal/auth"
	"github.com/Xavier-Hsiao/Chirpy/internal/config"
	"github.com/Xavier-Hsiao/Chirpy/internal/database"
	"github.com/Xavier-Hsiao/Chirpy/internal/helpers"
)

func HandlerUpdateUsers(cfg *config.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accessToken, err := auth.GetBearerToken(r.Header)
		if err != nil {
			helpers.RespondWithError(w, http.StatusUnauthorized, "Miss JWT token", err)
			return
		}
		userID, err := auth.ValidateJWT(accessToken, cfg.JWTSecret)
		if err != nil {
			helpers.RespondWithError(w, http.StatusUnauthorized, "Failed to validat JWT token", err)
			return
		}

		decoder := json.NewDecoder(r.Body)
		params := userParams{}
		err = decoder.Decode(&params)
		if err != nil {
			helpers.RespondWithError(w, http.StatusInternalServerError, "Failed to parse json data", err)
			return
		}

		hashedPassword, err := auth.HashPassword(params.Password)
		if err != nil {
			helpers.RespondWithError(w, http.StatusInternalServerError, "Failed to hash password", err)
			return
		}

		dbUser, err := cfg.DBQueries.UpdateUser(r.Context(), database.UpdateUserParams{
			ID:             userID,
			Email:          params.Email,
			HashedPassword: hashedPassword,
		})
		if err != nil {
			helpers.RespondWithError(w, http.StatusInternalServerError, "Failed to update user in db", err)
			return
		}

		user := helpers.ConvertUser(dbUser)

		helpers.RespondWithJson(w, http.StatusOK, user)
	}
}
