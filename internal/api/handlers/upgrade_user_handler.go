package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Xavier-Hsiao/Chirpy/internal/auth"
	"github.com/Xavier-Hsiao/Chirpy/internal/config"
	"github.com/Xavier-Hsiao/Chirpy/internal/helpers"
)

func HandlerUpgradeUser(cfg *config.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Validate ApiKey
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			helpers.RespondWithError(w, http.StatusUnauthorized, "Found no API key", err)
			return
		}
		if apiKey != cfg.PolkaKey {
			helpers.RespondWithError(w, http.StatusUnauthorized, "Unmatched API key", err)
			return
		}

		decoder := json.NewDecoder(r.Body)
		params := PolkaParams{}
		err = decoder.Decode(&params)
		if err != nil {
			helpers.RespondWithError(w, http.StatusInternalServerError, "Failed to parse request body", err)
			return
		}

		event := params.Event
		if event != "user.upgraded" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		userID := params.Data.UserID
		dbUser, err := cfg.DBQueries.UpgradeUser(r.Context(), userID)
		if err != nil {
			helpers.RespondWithError(w, http.StatusInternalServerError, "Failed to upgrade the user in DB", err)
			return
		}

		w.WriteHeader(http.StatusNoContent)
		fmt.Printf("User %s is chirpy red: %v\n", dbUser.Email, dbUser.IsChirpyRed)
	}
}
