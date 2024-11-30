package handlers

type userParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type chirpParams struct {
	Body string `json:"body"`
}
