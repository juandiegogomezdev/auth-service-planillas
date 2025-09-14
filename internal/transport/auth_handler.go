package transport

import (
	"io"
	"net/http"
	"proyecto/internal/service"
)

func (h *AuthHandler) SetupAuthRoutes() {
	http.HandleFunc("/auth/login", h.handlerLogin)
	http.HandleFunc("/auth/login/confirm-code", h.handlerLoginConfirmCode)
	http.HandleFunc("/auth/register", h.handlerRegister)
	http.HandleFunc("/auth/register/confirm", h.handlerRegisterConfirm)

}

type AuthHandler struct {
	service *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{service: s}
}

func (h *AuthHandler) handlerLogin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}
		r.Body.Close()

		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Login endpoint with POST method"))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

func (h *AuthHandler) handlerLoginConfirmCode(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Confirm login endpoint with POST method"))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

func (h *AuthHandler) handlerRegister(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Register endpoint with POST method"))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *AuthHandler) handlerRegisterConfirm(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Register confirm endpoint with POST method"))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
