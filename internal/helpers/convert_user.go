package helpers

import (
	"github.com/Xavier-Hsiao/Chirpy/internal/database"
	"github.com/Xavier-Hsiao/Chirpy/internal/models"
)

// We can not add json tags for database.Chirp
// So need to use this helper to convert database.Chirp into models.Chirp
func ConvertUser(user database.User) models.User {
	converted := models.User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	}

	return converted
}
