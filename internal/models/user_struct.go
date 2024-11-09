package models

import (
	"time"

	"github.com/google/uuid"
)

// API user struct to add json tags
type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Email     string    `json:"email"`
}