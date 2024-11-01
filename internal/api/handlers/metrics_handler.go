package handlers

import (
	"fmt"
	"net/http"

	"github.com/Xavier-Hsiao/Chirpy/internal/config"
)

func HandlerMetrics(cfg *config.ApiConfig) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		requestsCount := cfg.FileServerHits.Load()
		w.Header().Add("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf(`
		<html>
    		<body>
    			<h1>Welcome, Chirpy Admin</h1>
    			<p>Chirpy has been visited %d times!</p>
  			</body>
		</html>
		`, requestsCount)))
	}
}
