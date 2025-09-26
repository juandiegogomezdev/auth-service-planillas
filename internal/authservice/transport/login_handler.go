package transport

import (
	"encoding/json"
	"log"
	"net/http"
	"proyecto/internal/authservice/dtoauth"
	"proyecto/internal/shared/apperrors"
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

		token, err := h.service.Login(body.Email, body.Password)
		if err != nil {
			if appErr, ok := err.(*apperrors.SError); ok {
				switch appErr.Code {
				case "user_not_found":
					http.Error(w, "User not found", http.StatusNotFound)
				case "password_incorrect":
					http.Error(w, "Password incorrect", http.StatusUnauthorized)
				case "token_generation":
					http.Error(w, "Token generation error", http.StatusInternalServerError)
				default:
					http.Error(w, "Internal server error", http.StatusInternalServerError)
				}
				return

			}
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			log.Println("Unexpected error type:", err)
			return
		}

		utils.SetCookie(w, "auth_token", token)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("Login successful"))

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

		tokenConfirmLogin, err := utils.GetCookie(r, "auth_token")
		if err != nil {
			http.Error(w, "Missing or invalid token", http.StatusUnauthorized)
			return
		}

		tokenOrgSelect, err := h.service.ConfirmLoginCode(tokenConfirmLogin, body.Code)
		if err != nil {
			if appErr, ok := err.(*apperrors.SError); ok {
				switch appErr.Code {
				case "token_parsing":
					http.Error(w, "Invalid token", http.StatusUnauthorized)
				case "verification_not_found":
					http.Error(w, "Verification info not found", http.StatusNotFound)
				case "invalid_code":
					http.Error(w, "Invalid code", http.StatusUnauthorized)
				case "code_expired":
					http.Error(w, "Code has expired", http.StatusUnauthorized)
				case "token_generation":
					http.Error(w, "Error generating access, try later.", http.StatusInternalServerError)
				default:
					http.Error(w, "Internal server error", http.StatusInternalServerError)
				}
				return
			}

			http.Error(w, "Error confirming login code", http.StatusInternalServerError)
			return
		}

		utils.SetCookie(w, "auth_token", tokenOrgSelect)

		w.Header().Set("Content-Type", "plain/text")
		w.Write([]byte("Login confirmed successfully"))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}
