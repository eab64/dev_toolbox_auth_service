package dto

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
}

type RegisterResponse struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
}
