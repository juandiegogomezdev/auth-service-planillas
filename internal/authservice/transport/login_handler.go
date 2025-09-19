package transport

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"proyecto/internal/authservice/dtoauth"
	"proyecto/internal/authservice/serviceauth"
	"proyecto/internal/shared/utils"
)

func (h *Handler) handlerLogin(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var body dtoauth.LoginRequest
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
		case serviceauth.LoginStatusInvalidCredentials:
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		case serviceauth.LoginStatusUserNotFound:
			http.Error(w, "User not found", http.StatusNotFound)
		case serviceauth.LoginStatusSuccess:
			utils.SetCookie(w, "access_token", token)
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte("Login successful"))
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}

func (h *Handler) handlerLoginConfirmCode(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var body dtoauth.LoginConfirmRequest
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		r.Body.Close()

		token, err := utils.GetCookie(r, "access_token")
		if err != nil {
			http.Error(w, "Missing or invalid token", http.StatusUnauthorized)
			return
		}
		fmt.Println("Token from cookie:", token)

		token, status, err := h.service.ConfirmLoginCode(token, body.Code)
		if err != nil {
			log.Println("Error confirming login code:", err)
			http.Error(w, "Error confirming login code", http.StatusInternalServerError)
			return
		}

		switch status {
		case serviceauth.ConfirmLoginInvalidCode:
			http.Error(w, "Invalid code", http.StatusUnauthorized)
		case serviceauth.ConfirmLoginStatusInvalidToken:
			http.Error(w, "Invalid token", http.StatusUnauthorized)
		case serviceauth.ConfirmLoginStatusSuccess:
			utils.SetCookie(w, "access_token", token)
			w.Header().Set("Content-Type", "plain/text")
			w.Write([]byte("Login confirmed successfully"))
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}
