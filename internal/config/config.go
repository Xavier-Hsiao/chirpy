package config

import (
	"sync/atomic"

	"github.com/Xavier-Hsiao/Chirpy/internal/database"
)

// Hold stateful, in-memory data
// Use atomic std lib to prevent data racing casued by concurrent requests
type ApiConfig struct {
	FileServerHits atomic.Int32
	DBQueries      *database.Queries
	Platform       string
	JWTSecret      string
}
