package transportorg

import (
	"encoding/json"
	"net/http"
)

func (h *HandlerOrg) HandlerPermissionOrg(w http.ResponseWriter, r *http.Request) {
	// Implement permission handling logic here
	switch r.Method {
	case http.MethodGet:
		// Handle GET request for permissions
		permissionsInfo, err := h.service.PermissionsInfo()
		if err != nil {
			http.Error(w, "Failed to retrieve permissions info", http.StatusInternalServerError)
			return
		}
		// Return the permissions info
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(permissionsInfo); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
