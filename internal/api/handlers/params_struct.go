package handlers

import "github.com/google/uuid"

type userParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type chirpParams struct {
	Body   string    `json:"body"`
	UserID uuid.UUID `json:"user_id"`
}
