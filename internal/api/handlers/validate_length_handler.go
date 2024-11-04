package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Xavier-Hsiao/Chirpy/internal/helpers"
)

type parameters struct {
	Body string `json:"body"`
}
type returnValues struct {
	CleanedBody string `json:"cleaned_body"`
}

//	@Summary		Validate a chirp
//	@Description	Validate a chirp's length (should be less than 14 characters) and replace profane words with ****
//	@Tags			validation
//	@ID				post-validation
//	@Accept			json
//	@Produce		json
//	@Param			body	body		parameters				true	"Chirp content to validate"
//	@Success		200		{object}	returnValues			"Cleaned chirp body with profanity removed"
//	@Failure		400		{object}	helpers.ErrorResponse	"Chirp is too long, should be less than 140 chars"
//	@Failure		500		{object}	helpers.ErrorResponse	"Failed to decode parameters"
//	@Router			/app/validate_chirp [post]
func HandlerValidateLength(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
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

	// Validate the length of the chirp
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

	helpers.RespondWithJson(w, http.StatusOK, returnValues{CleanedBody: cleanedBody})
}
