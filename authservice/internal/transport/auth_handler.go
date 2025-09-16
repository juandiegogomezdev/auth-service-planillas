package transport

import (
	"encoding/json"
	"log"
	"net/http"
	"proyecto/authservice/internal/dto"
	"proyecto/authservice/internal/service"
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

func (h *Handler) handlerRegister(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var reg dto.RegisterRequest
		err := json.NewDecoder(r.Body).Decode(&reg)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		r.Body.Close()

		status, err := h.service.RegisterUser(reg.Email)
		if err != nil {
			log.Println("Error registering user:", err)
			http.Error(w, "Error registering user", http.StatusInternalServerError)
			return
		}

		switch status {
		case service.RegisterUserStatusUserExists:
			http.Error(w, "User already exists", http.StatusConflict)
			return
		case service.RegisterUserStatusConfirmationSent:
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Confirmation email sent"))
			return
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) handlerRegisterConfirm(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var body dto.RegisterConfirmRequest
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		r.Body.Close()

		status, err := h.service.CreateUser(body.Token, body.Password)
		if err != nil {
			log.Println("Error creating user:", err)
			http.Error(w, "Error creating user", http.StatusInternalServerError)
			return
		}

		switch status {
		case service.CreateUserStatusUserExists:
			http.Error(w, "User already exists", http.StatusConflict)
		case service.CreateUserStatusUserCreated:
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("User created successfully"))
		default:
			w.WriteHeader(http.StatusOK)
		}

		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Register confirm endpoint with POST method"))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) handlerLogin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var body dto.LoginRequest
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		r.Body.Close()

		token, status, err := h.service.Login(body.Email, body.Password)
		if err != nil {
			log.Println("Error logging in. ", err)
			http.Error(w, "Error logging in", http.StatusInternalServerError)
			return
		}

		switch status {
		case service.LoginStatusInvalidCredentials:
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		case service.LoginStatusUserNotFound:
			http.Error(w, "User not found", http.StatusNotFound)
		case service.LoginStatusSuccess:
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(dto.LoginResponse{Token: token})
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

func (h *Handler) handlerLoginConfirmCode(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var body dto.LoginConfirmRequest
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		r.Body.Close()

		token, status, err := h.service.ConfirmLoginCode(body.Token, body.Code)
		if err != nil {
			log.Println("Error confirming login code:", err)
			http.Error(w, "Error confirming login code", http.StatusInternalServerError)
			return
		}

		switch status {
		case service.ConfirmLoginInvalidCode:
			http.Error(w, "Invalid code", http.StatusUnauthorized)
		case service.ConfirmLoginStatusInvalidToken:
			http.Error(w, "Invalid token", http.StatusUnauthorized)
		case service.ConfirmLoginStatusSuccess:
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(dto.LoginResponse{Token: token})
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}
