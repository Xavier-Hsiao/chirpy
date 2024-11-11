package handlers

import (
	"log"
	"net/http"

	"github.com/Xavier-Hsiao/Chirpy/internal/config"
	"github.com/Xavier-Hsiao/Chirpy/internal/helpers"
)

//	@Summary		Get all chirps
//	@Description	Retreive all chirps from database in ascendent order of created_at time
//	@Tags			chirp
//	@ID				get-chirps
//	@Produce		json
//	@Success		200	{array}		models.Chirp			"created chirp's information"
//	@Failure		500	{object}	helpers.ErrorResponse	"Internal server error: can not deal with data properly"
//	@Router			/api/chirps [get]
func HandlerGetChirps(cfg *config.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//  Get the slice of Chirps from database
		chirpsDB, err := cfg.DBQueries.GetChirps(r.Context())
		if err != nil {
			helpers.RespondWithError(w, http.StatusInternalServerError, "Failed to get chirps from database", err)
			return
		}
		chirps := helpers.ConvertChirps(chirpsDB)

		// Respond with json data
		helpers.RespondWithJson(w, http.StatusOK, chirps)

		log.Println("Retreive all chirps!")
	}
}
