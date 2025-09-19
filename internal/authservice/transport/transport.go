package transport

import (
	"net/http"
	"proyecto/internal/authservice/serviceauth"
)

func (h *Handler) SetupAuthRoutes() {
	http.HandleFunc("/auth/login", h.handlerLogin)
	http.HandleFunc("/auth/login/confirm-code", h.handlerLoginConfirmCode)
	http.HandleFunc("/auth/register", h.handlerRegister)
	http.HandleFunc("/auth/register/confirm", h.handlerRegisterConfirm)

}

type Handler struct {
	service *serviceauth.ServiceAuth
}

func NewAuthHandler(s *serviceauth.ServiceAuth) *Handler {
	return &Handler{service: s}
}
