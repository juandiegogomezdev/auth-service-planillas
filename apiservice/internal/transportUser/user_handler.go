package transportApi

import (
	"encoding/json"
	"net/http"
	"proyecto/auth-service/service"
)

func (h *UserHandler) SetupUserRoutes() {
	http.HandleFunc("/users", h.handlerGetAllUsers)
}

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) handlerGetAllUsers(w http.ResponseWriter, r *http.Request) {
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

// func (h *UserHandler) RegisterRouter(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodGet:
// 		users, err := h.service.GetAllUsers()
// 		if err != nil {
// 			http.Error(w, "Failed to get users", http.StatusInternalServerError)
// 			return
// 		}
// 		w.Header().Set("Content-Type", "application/json")
// 		json.NewEncoder(w).Encode(users)
// 	default:
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 	}

// }
