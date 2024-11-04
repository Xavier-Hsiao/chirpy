package handlers

import "net/http"

//	@Summary		Health check endpoint
//	@Description	Returns OK if the server is healthy
//	@Tags			health
//	@Accept			json
//	@Produce		plain
//	@Success		200	{string}	string	"OK"
//	@Router			/api/healthz [get]
func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}
