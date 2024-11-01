package helper

import (
	"log"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, msg string, err error) {
	// Still want to print the error in terminal
	if err != nil {
		log.Println(err)
	}

	type errorResponse struct {
		Error string `json:"error"`
	}
	RespondWithJson(w, code, errorResponse{
		Error: msg,
	})
}
