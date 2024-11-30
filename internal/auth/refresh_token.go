package auth

import (
	"crypto/rand"
	"encoding/hex"
)

// Generate a 32 bytes random string
func MakeRefreshToken() (string, error) {
	token := make([]byte, 32)
	// Fill the token bytes slice with random bytes
	_, err := rand.Read(token)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(token), nil

}
