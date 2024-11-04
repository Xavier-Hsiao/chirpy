package helpers

import (
	"log"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func RespondWithError(w http.ResponseWriter, code int, msg string, err error) {
	// Server-side logging
	if err != nil {
		log.Println(err)
	}

	RespondWithJson(w, code, ErrorResponse{
		Error: msg,
	})
}
