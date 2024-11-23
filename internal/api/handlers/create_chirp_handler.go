package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Xavier-Hsiao/Chirpy/internal/auth"
	"github.com/Xavier-Hsiao/Chirpy/internal/config"
	"github.com/Xavier-Hsiao/Chirpy/internal/database"
	"github.com/Xavier-Hsiao/Chirpy/internal/helpers"
	"github.com/Xavier-Hsiao/Chirpy/internal/models"
)

// @Summary		Create new chirp
// @Description	Create a new chirp message instance
// @Tags			chirp
// @ID				post-create-chirp
// @Accept			json
// @Produce		json
// @Param			body	body		chirpParams				true	"chrip message body and the author's ID"
// @Success		201		{object}	models.Chirp			"created chirp's information"
// @Failure		400		{object}	helpers.ErrorResponse	"Chirp is too long, should be less than 140 chars"
// @Failure		500		{object}	helpers.ErrorResponse	"Internal server error: can not deal with data properly"
// @Router			/api/chirps [post]
func HandlerCreateChirp(cfg *config.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Validate JWT access token
		accessToken, err := auth.GetBearerToken(r.Header)
		if err != nil {
			helpers.RespondWithError(w, http.StatusUnauthorized, "Couldn't find JWT token", err)
			return
		}

		userID, err := auth.ValidateJWT(accessToken, cfg.JWTSecret)
		if err != nil {
			helpers.RespondWithError(w, http.StatusUnauthorized, "Failed to valiate JWT", err)
			return
		}

		// Get the request body
		decoder := json.NewDecoder(r.Body)
		params := chirpParams{}
		err = decoder.Decode(&params)
		if err != nil {
			helpers.RespondWithError(w, http.StatusInternalServerError, err.Error(), err)
			return
		}

		// Validate the length of the chirp body
		const maxLength = 140
		if len(params.Body) > maxLength {
			helpers.RespondWithError(w, http.StatusBadRequest, "Chirp is too long, should be less than 140 chars", err)
			return
		}

		// Check badwords and replace them
		profaneWords := map[string]struct{}{
			"kerfuffle": {},
			"sharbert":  {},
			"fornax":    {},
		}

		cleanedBody := helpers.ReplaceBadWords(params.Body, profaneWords)

		// Create the new chirp in database
		chirp, err := cfg.DBQueries.CreateChirp(r.Context(), database.CreateChirpParams{
			Body:   cleanedBody,
			UserID: userID,
		})
		if err != nil {
			helpers.RespondWithError(w, http.StatusInternalServerError, "Failed to create chirp in database", err)
			return
		}

		// Return the created chirp in JSON
		helpers.RespondWithJson(w, http.StatusCreated, models.Chirp{
			ID:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			Body:      chirp.Body,
			UserID:    chirp.UserID,
		})

		log.Printf("chirp %s was created!\n", chirp.ID)
	}
}
