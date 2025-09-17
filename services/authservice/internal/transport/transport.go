package transport

import (
	"net/http"
	"proyecto/services/authservice/internal/service"
)

func (h *Handler) SetupAuthRoutes() {
	http.HandleFunc("/auth/login", h.handlerLogin)
	http.HandleFunc("/auth/login/confirm-code", h.handlerLoginConfirmCode)
	http.HandleFunc("/auth/register", h.handlerRegister)
	http.HandleFunc("/auth/register/confirm", h.handlerRegisterConfirm)

}

type Handler struct {
	service *service.Service
}

func NewAuthHandler(s *service.Service) *Handler {
	return &Handler{service: s}
}
