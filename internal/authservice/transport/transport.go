package transport

import (
	"net/http"
	"proyecto/internal/authservice/serviceauth"
)

func (h *Handler) SetupAuthRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/auth/login", h.handlerLogin)
	mux.HandleFunc("/auth/login/confirm-code", h.handlerLoginConfirmCode)
	mux.HandleFunc("/auth/register", h.handlerRegister)
	mux.HandleFunc("/auth/register/confirm", h.handlerRegisterConfirm)
}

type Handler struct {
	service *serviceauth.ServiceAuth
}

func NewAuthHandler(s *serviceauth.ServiceAuth) *Handler {
	return &Handler{service: s}
}
