package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Xavier-Hsiao/Chirpy/internal/config"
	"github.com/Xavier-Hsiao/Chirpy/internal/helpers"
	"github.com/Xavier-Hsiao/Chirpy/internal/models"
)

type userParams struct {
	Email string
}

// @Summary		Create new user
// @Description	Create a new Chirpy user
// @Tags			user
// @ID				post-create-user
// @Accept			json
// @Produce		json
// @Param			body	body		userParams				true	"user email to get new user created"
// @Success		200		{object}	models.User				"created user's information"
// @Failure		500		{object}	helpers.ErrorResponse	"Failed to decode parameters"
// @Failure		500		{object}	helpers.ErrorResponse	"Failed to create user in DB"
// @Router			/api/users [post]
func HandlerCreateUser(cfg *config.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		decoder := json.NewDecoder(r.Body)
		params := userParams{}
		err := decoder.Decode(&params)
		if err != nil {
			helpers.RespondWithError(
				w,
				http.StatusInternalServerError,
				"Failed to decode parameters",
				err,
			)
			return
		}

		user, err := cfg.DBQueries.CreateUser(context.Background(), params.Email)
		if err != nil {
			helpers.RespondWithError(
				w,
				http.StatusInternalServerError,
				"Failed to create user in DB",
				err,
			)
			return
		}

		helpers.RespondWithJson(w, http.StatusCreated, models.User{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Email:     user.Email,
		})

		log.Printf("user %s was created!\n", user.Email)
	}

}
