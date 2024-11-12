package handlers

import (
	"log"
	"net/http"

	"github.com/Xavier-Hsiao/Chirpy/internal/config"
	"github.com/Xavier-Hsiao/Chirpy/internal/helpers"
	"github.com/google/uuid"
)

// @Summary		Get a specific chirp
// @Description	Retreive the chirp from database by ID
// @Tags			chirp
// @ID				get-chirp-by-id
// @Produce		json
// @Success		200	{object}	models.Chirp			"created chirp's information"
// @Failure		400	{object}	helpers.ErrorResponse	"Internal server error: can not deal with data properly"
// @Failure		404	{object}	helpers.ErrorResponse	"Chirp not found in database"
// @Router			/api/chirps/{chirpID} [get]
func HandlerGetChirpById(cfg *config.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the chirp's id from path and convert it to uuid
		id, err := uuid.Parse(r.PathValue("chirpID"))
		if err != nil {
			helpers.RespondWithError(w, http.StatusBadRequest, "Failed to convert id string to uuid", err)
			return
		}

		//  Get the chirp from database
		dbChirp, err := cfg.DBQueries.GetChirpById(r.Context(), id)
		if err != nil {
			helpers.RespondWithError(w, http.StatusNotFound, "Chirp not found in database", err)
		}
		chirp := helpers.ConvertChirp(dbChirp)

		// Respond with json data
		helpers.RespondWithJson(w, http.StatusOK, chirp)

		log.Printf("Retreived the chirp: %s\n", chirp.ID)
	}
}
