package transportorg

import (
	"encoding/json"
	"net/http"
	"proyecto/internal/orgservice/serviceorg"
	"proyecto/internal/shared/tokenizer"
	"proyecto/internal/shared/utils"

	"github.com/google/uuid"
)

// The user must access to a organization after login
// This handler will return all the organizations of the user
// and will allow to select one organization to login
// The selected organization will be stored in a cookie
// The cookie will be used to access to the organization resources

type loginMemRequest struct {
	MemID string `json:"memID"`
}

func (h *HandlerOrg) HandlerSessionOrg(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		token, err := utils.GetCookie(r, "auth_token")
		if err != nil {
			http.Error(w, "Missing or invalid token", http.StatusUnauthorized)
			return
		}

		payload, err := tokenizer.JWTParseOrgSelectToken(token)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		userID, err := uuid.Parse(payload.UserUUID)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusUnauthorized)
			return
		}

		mem, status, err := h.service.GetAllUserMemberships(userID)
		if err != nil {
			http.Error(w, "Failed to retrieve memberships", http.StatusInternalServerError)
			return
		}

		switch status {
		case serviceorg.NoUserMembershipsFound:
			http.Error(w, "No memberships found", http.StatusNotFound)
			return
		case serviceorg.ErrorGettingUserMemberships:
			http.Error(w, "Error retrieving memberships", http.StatusInternalServerError)
			return
		case serviceorg.UserMembershipsFound:
			// Continue to return memberships
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(mem); err != nil {
				http.Error(w, "Failed to encode response", http.StatusInternalServerError)
				return
			}
			return
		}
	case http.MethodPost:
		token, err := utils.GetCookie(r, "auth_token")
		if err != nil {
			http.Error(w, "Missing or invalid token", http.StatusUnauthorized)
			return
		}

		payload, err := tokenizer.JWTParseOrgSelectToken(token)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		userID, err := uuid.Parse(payload.UserUUID)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusUnauthorized)
			return
		}

		// Read organization ID from request body
		var req loginMemRequest
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		memID, err := uuid.Parse(req.MemID)
		if err != nil {
			http.Error(w, "Invalid membership ID", http.StatusBadRequest)
			return
		}

		// Check if the user is a member of the organization
		memID, status, err := h.service.CheckUserMembership(userID, memID)

		switch status {
		case serviceorg.NotFoundMembership:
			http.Error(w, "Membership not found or inactive", http.StatusForbidden)
			return
		case serviceorg.ErrorCheckingMembership:
			http.Error(w, "Error checking membership", http.StatusInternalServerError)
			return
		case serviceorg.FoundMembership:
			// Create a new token with the membershipID
			newToken, err := tokenizer.JWTGenerateAppAccessToken(userID, memID)
			if err != nil {
				http.Error(w, "Error generating access", http.StatusInternalServerError)
				return
			}
			// Set the new token in a cookie
			utils.SetCookie(w, "auth_token", newToken)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Membership selected successfully"))
			return
		}

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
