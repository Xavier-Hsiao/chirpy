package handlers

import (
	"fmt"
	"net/http"

	"github.com/Xavier-Hsiao/Chirpy/internal/config"
)

//	@Summary		Get Metrics
//	@Description	Returns an HTML page showing the number of times Chirpy has been visited
//	@Tags			metrics
//	@ID				get-metrics
//	@Produce		html
//	@Success		200	{string}	string	"HTML content displaying the visit count"
//	@Example		200 "<html><body><h1>Welcome, Chirpy Admin</h1><p>Chirpy has been visited 42 times!</p></body></html>"
//	@Router			/admin/metrics [get]
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
