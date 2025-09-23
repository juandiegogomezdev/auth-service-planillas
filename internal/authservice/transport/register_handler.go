package transport

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"proyecto/internal/authservice/apperrors"
	"proyecto/internal/authservice/dtoauth"
	"proyecto/internal/authservice/serviceauth"
)

func (h *Handler) handlerRegister(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var reg dtoauth.RegisterRequest
		err := json.NewDecoder(r.Body).Decode(&reg)
		if err != nil {
			log.Println("Error decoding request payload:", err)
			http.Error(w, "Internal Server Error", http.StatusBadRequest)
			return
		}
		r.Body.Close()

		err = h.service.RegisterUser(reg.Email)
		if err != nil {
			if appErr, ok := err.(*apperrors.SError); ok {
				switch appErr.Code {
				case "user_exists":
					http.Error(w, "User already exists", http.StatusConflict)
					return
				}
			}
			log.Println("Error registering user:", err)
			http.Error(w, "Error registering user", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Check your email for confirmation link"))
		return
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) handlerRegisterConfirm(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var body dtoauth.RegisterConfirmRequest
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			fmt.Println("Error decoding request payload:", err)
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
		case serviceauth.CreateUserStatusUserExists:
			http.Error(w, "User already exists", http.StatusConflict)
		case serviceauth.CreateUserStatusUserCreated:
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("User created successfully"))
		default:
			w.WriteHeader(http.StatusOK)
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
