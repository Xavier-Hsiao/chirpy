package handlers

import (
	"log"
	"net/http"

	"github.com/Xavier-Hsiao/Chirpy/internal/config"
	"github.com/Xavier-Hsiao/Chirpy/internal/helpers"
	"github.com/Xavier-Hsiao/Chirpy/internal/models"
	"github.com/google/uuid"
)

// @Summary		Get all chirps
// @Description	Retreive all chirps from database in ascendent order of created_at time
// @Tags			chirp
// @ID				get-chirps
// @Produce		json
// @Success		200	{array}		models.Chirp			"created chirp's information"
// @Failure		500	{object}	helpers.ErrorResponse	"Internal server error: can not deal with data properly"
// @Router			/api/chirps [get]
func HandlerGetChirps(cfg *config.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//  Get the slice of Chirps from database
		dbChirps, err := cfg.DBQueries.GetChirps(r.Context())
		if err != nil {
			helpers.RespondWithError(w, http.StatusInternalServerError, "Failed to get chirps from database", err)
			return
		}

		authorIDString := r.URL.Query().Get("author_id")
		authorID := uuid.Nil
		if authorIDString != "" {
			authorID, err = uuid.Parse(authorIDString)
			if err != nil {
				helpers.RespondWithError(w, http.StatusBadRequest, "Invalid author ID", err)
				return
			}
		}

		chirps := []models.Chirp{}
		for _, dbChirp := range dbChirps {
			if authorID != uuid.Nil && dbChirp.UserID != authorID {
				continue
			}

			chirps = append(chirps, models.Chirp{
				ID:        dbChirp.ID,
				CreatedAt: dbChirp.CreatedAt,
				UpdatedAt: dbChirp.UpdatedAt,
				UserID:    dbChirp.UserID,
				Body:      dbChirp.Body,
			})
		}

		// Respond with json data
		helpers.RespondWithJson(w, http.StatusOK, chirps)

		log.Println("Retreive required chirps!")
	}
}
