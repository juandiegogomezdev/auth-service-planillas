package transport

import (
	"encoding/json"
	"net/http"
	"proyecto/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserRoutes(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) RegisterRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		users, err := h.service.GetAllUsers()
		if err != nil {
			http.Error(w, "Failed to get users", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

}
