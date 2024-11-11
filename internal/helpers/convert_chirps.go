package helpers

import (
	"github.com/Xavier-Hsiao/Chirpy/internal/database"
	"github.com/Xavier-Hsiao/Chirpy/internal/models"
)

// We can not add json tags for database.Chirp
// So need to use this helper to convert database.Chirp into models.Chirp
func ConvertChirps(chirps []database.Chirp) []models.Chirp {
	converted := make([]models.Chirp, len(chirps))
	for i, chirp := range chirps {
		converted[i] = models.Chirp{
			ID:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			Body:      chirp.Body,
			UserID:    chirp.UserID,
		}
	}

	return converted
}
