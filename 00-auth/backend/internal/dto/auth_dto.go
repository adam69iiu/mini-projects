package dto



type AuthRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	ID string `json:"id"`
	Username string `json:"Username"`
	Email string `json:"email"`
	Token string `json:"token"`
}
