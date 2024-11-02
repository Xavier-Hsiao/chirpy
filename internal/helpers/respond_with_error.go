package helpers

import (
	"log"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, msg string, err error) {
	// Server-side logging
	if err != nil {
		log.Println(err)
	}

	// Client-side user friendly error message
	type errorResponse struct {
		Error string `json:"error"`
	}
	RespondWithJson(w, code, errorResponse{
		Error: msg,
	})
}
