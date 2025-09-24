package transport

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"proyecto/internal/authservice/apperrors"
	"proyecto/internal/authservice/dtoauth"
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

		fmt.Println("Received body:", body)

		err = h.service.CreateUser(body.Token, body.Password)
		if err != nil {
			if appErr, ok := err.(*apperrors.SError); ok {
				switch appErr.Code {
				case "invalid_token":
					http.Error(w, "Invalid or expired link", http.StatusBadRequest)
					return
				case "user_exists":
					http.Error(w, "User already exists", http.StatusConflict)
					return
				case "user_creation":
					http.Error(w, "Error creating user", http.StatusInternalServerError)
					return
				}
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

		}

		w.Write([]byte("User created successfully"))
		return
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
