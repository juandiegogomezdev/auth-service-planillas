package dto

// DTOs for registration
type AuthRegisterRequest struct {
	Email string `json:"email"`
}

// DTOs for registration confirmation
type AuthRegisterConfirmRequest struct {
	Password string `json:"password"`
	Token    string `json:"token"`
}

// DTOs for login
type AuthLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthLoginResponse struct {
	Token string `json:"token"`
}

// DTOs for login confirmation
type AuthLoginConfirmRequest struct {
	Code  string `json:"code"`
	Token string `json:"token"`
}
