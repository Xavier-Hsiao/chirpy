package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Xavier-Hsiao/Chirpy/internal/helpers"
)

// It should expect a JSON body of chirp
// If any error occur, it should respond with an appropriate HTTP status code and a JSON body
// If a chirp is valid, respond with a JSON body as well
func HandlerValidateLength(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	type returnValues struct {
		Valid bool `json:"valid"`
	}

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

	helpers.RespondWithJson(w, http.StatusOK, returnValues{Valid: true})
}
