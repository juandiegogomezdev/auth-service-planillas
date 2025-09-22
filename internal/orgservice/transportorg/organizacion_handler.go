package transportorg

// func (h *HandlerOrg) HandlerOrganization(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodGet:
// 		userIDVal := r.Context().Value(middleware.UserIDKey)
// 		if userIDVal == nil {
// 			http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 			return
// 		}
// 		userID := userIDVal.(uuid.UUID)

// 		organizations, err := h.service.GetAllUserOrganizations(userID)
// 		if err != nil {
// 			http.Error(w, "Failed to retrieve organizations", http.StatusInternalServerError)
// 			return
// 		}

// 		w.Header().Set("Content-Type", "application/json")
// 		if err := json.NewEncoder(w).Encode(organizations); err != nil {
// 			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
// 			return
// 		}
// 		return

// 	default:
// 		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
// 	}

// }
