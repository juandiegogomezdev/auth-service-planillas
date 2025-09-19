package dtoauth

// DTOs for registration
type RegisterRequest struct {
	Email string `json:"email"`
}

// DTOs for registration confirmation
type RegisterConfirmRequest struct {
	Password string `json:"password"`
	Token    string `json:"token"`
}

// DTOs for login
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// DTOs for login confirmation
type LoginConfirmRequest struct {
	Code string `json:"code"`
}
