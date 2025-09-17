package transport

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"proyecto/internal/utils"
	"proyecto/services/authservice/internal/dto"
	"proyecto/services/authservice/internal/service"
)

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
		var body dto.LoginConfirmRequest
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
		case service.ConfirmLoginInvalidCode:
			http.Error(w, "Invalid code", http.StatusUnauthorized)
		case service.ConfirmLoginStatusInvalidToken:
			http.Error(w, "Invalid token", http.StatusUnauthorized)
		case service.ConfirmLoginStatusSuccess:
			utils.SetCookie(w, "access_token", token)
			w.Header().Set("Content-Type", "plain/text")
			w.Write([]byte("Login confirmed successfully"))
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}
