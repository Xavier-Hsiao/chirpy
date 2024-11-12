package helpers

import (
	"github.com/Xavier-Hsiao/Chirpy/internal/database"
	"github.com/Xavier-Hsiao/Chirpy/internal/models"
)

// We can not add json tags for database.Chirp
// So need to use this helper to convert database.Chirp into models.Chirp
func ConvertChirp(chirp database.Chirp) models.Chirp {
	converted := models.Chirp{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	}

	return converted
}
