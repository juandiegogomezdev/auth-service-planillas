package middleware

import (
	"context"
	"net/http"
	"proyecto/internal/shared/tokenizer"
	"proyecto/internal/shared/utils"

	"github.com/google/uuid"
)

type contextKey string

const (
	UserIDKey contextKey = "userID"
	OrgIDKey  contextKey = "orgID"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := utils.GetCookie(r, "auth_token")
		if err != nil {
			http.Error(w, "Missing or invalid token", http.StatusUnauthorized)
			return
		}

		tokenPayload, err := tokenizer.JWTParseMembershipAccessToken(token)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// save userID and orgID as uuid.UUID

		userID, err := uuid.Parse(tokenPayload.UserUUID)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusUnauthorized)
			return
		}

		orgID, err := uuid.Parse(tokenPayload.MembershipUUID)
		if err != nil {
			http.Error(w, "Invalid organization ID", http.StatusUnauthorized)
			return
		}

		// Attach the user ID and organization ID from the token to the request context
		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		ctx = context.WithValue(ctx, OrgIDKey, orgID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
