package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Xavier-Hsiao/Chirpy/internal/config"
	"github.com/Xavier-Hsiao/Chirpy/internal/database"
	"github.com/Xavier-Hsiao/Chirpy/internal/helpers"
	"github.com/Xavier-Hsiao/Chirpy/internal/models"
	"github.com/google/uuid"
)

type chirpParams struct {
	Body   string    `json:"body"`
	UserID uuid.UUID `json:"user_id"`
}

//	@Summary		Create new chirp
//	@Description	Create a new chirp message instance
//	@Tags			chirp
//	@ID				post-create-chirp
//	@Accept			json
//	@Produce		json
//	@Param			body	body		chirpParams				true	"chrip message body and the author's ID"
//	@Success		201		{object}	models.Chirp			"created chirp's information"
//	@Failure		400		{object}	helpers.ErrorResponse	"Chirp is too long, should be less than 140 chars"
//	@Failure		500		{object}	helpers.ErrorResponse	"Internal server error: can not deal with data properly"
//	@Router			/api/chirps [post]
func HandlerCreateChirp(cfg *config.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the request body
		decoder := json.NewDecoder(r.Body)
		params := chirpParams{}
		err := decoder.Decode(&params)
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
		fmt.Printf("User ID in params: %v\n", params.UserID)
		chirp, err := cfg.DBQueries.CreateChirp(r.Context(), database.CreateChirpParams{
			Body:   cleanedBody,
			UserID: params.UserID,
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
