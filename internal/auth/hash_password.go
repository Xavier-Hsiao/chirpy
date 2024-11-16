package auth

import (
	"golang.org/x/crypto/bcrypt"
)

// Hash the password with bcrpt std lib
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

// Compare the password in HTTP request with the one stored in database
func CheckPasswordHash(password, hash string) error {
	plainPassword := []byte(password)
	hashedPassword := []byte(hash)

	// Return nil on success, or an error on failure
	err := bcrypt.CompareHashAndPassword(hashedPassword, plainPassword)
	if err != nil {
		return err
	}

	return nil
}
