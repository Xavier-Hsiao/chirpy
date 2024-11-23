package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Xavier-Hsiao/Chirpy/internal/auth"
	"github.com/Xavier-Hsiao/Chirpy/internal/config"
	"github.com/Xavier-Hsiao/Chirpy/internal/helpers"
	"github.com/Xavier-Hsiao/Chirpy/internal/models"
)

type loginResp struct {
	User  models.User `json:"user"`
	Token string      `json:"token"`
}

// @Summary		Login users
// @Description	Check if the users are who they claimed
// @Tags			user
// @ID				post-user-login
// @Accept			json
// @Produce		json
// @Param			body	body		userParams				true	"user email and passowrd"
// @Success		200		{object}	models.User				"user's information"
// @Failure		500		{object}	helpers.ErrorResponse	"Internal server error: can not deal with data properly"
// @Failure		401		{object}	helpers.ErrorResponse	"Incorrect email or password"
// @Router			/api/login [post]
func HandlerLogin(cfg *config.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get the request body
		decoder := json.NewDecoder(r.Body)
		params := userParams{}
		err := decoder.Decode(&params)
		if err != nil {
			helpers.RespondWithError(w, http.StatusInternalServerError, "Failed to parse request data", err)
			return
		}

		// Get hashed password of the user from database
		dbUser, err := cfg.DBQueries.GetUserByEmail(r.Context(), params.Email)
		if err != nil {
			helpers.RespondWithError(w, http.StatusInternalServerError, "Failed to get user by email from db", err)
			return
		}

		// Check password
		err = auth.CheckPasswordHash(params.Password, dbUser.HashedPassword)
		if err != nil {
			helpers.RespondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
			return
		}

		// Handle JWT expiration time, default is 1 hour
		expirationTime := time.Hour
		if params.ExpiresInSeconds > 0 && params.ExpiresInSeconds < 3600 {
			expirationTime = time.Duration(params.ExpiresInSeconds)
		}

		// Generate JWT access token
		accessToken, err := auth.MakeJWT(dbUser.ID, cfg.JWTSecret, expirationTime)
		if err != nil {
			helpers.RespondWithError(w, http.StatusInternalServerError, "Faile to generate JWT access token", err)
			return
		}

		// Return user json once password check passed
		user := helpers.ConvertUser(dbUser)
		err = helpers.RespondWithJson(w, http.StatusOK, loginResp{
			User:  user,
			Token: accessToken,
		})
		if err != nil {
			helpers.RespondWithError(w, http.StatusInternalServerError, "Failed to parse resp data", err)
			return
		}

		log.Printf("User %s logged in!\n", user.Email)
	}
}
