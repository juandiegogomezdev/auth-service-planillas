package transport

import (
	"encoding/json"
	"log"
	"net/http"
	"proyecto/services/authservice/internal/dto"
	"proyecto/services/authservice/internal/service"
)

func (h *Handler) handlerRegister(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var reg dto.RegisterRequest
		err := json.NewDecoder(r.Body).Decode(&reg)
		if err != nil {
			log.Println("Error decoding request payload:", err)
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
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
