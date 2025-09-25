package middleware

import (
	"fmt"
	"net/http"
	"proyecto/internal/shared/tokenizer"
	"strings"
)

// Check if has a org o app access token
// If has, redirect to the login page
// If not, allow to access to static files
func RedirectBeforeSendStaticFiles(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Know what file is being requested
		page := r.URL.Path

		fmt.Println("page: ", page)

		// Get the token from cookie
		cookie, err := r.Cookie("auth_token")
		token := ""
		if err == nil {
			token = cookie.Value
		}

		// Identify the token type
		tokenType := tokenizer.TokenTypeUnkown
		if token != "" {
			tokenType, _ = tokenizer.IdentifyTokenType(token)
		}

		fmt.Println("tokenType: ", tokenType)

		switch tokenType {
		case tokenizer.TokenTypeUnkown:
			if strings.HasPrefix(page, "org-select") || strings.HasPrefix(page, "app") {
				http.Redirect(w, r, "http://localhost:8080/static/login/", http.StatusFound)
				return
			}
		case tokenizer.TokenTypeOrgSelect:
			if !strings.HasPrefix(page, "org-select") {
				http.Redirect(w, r,
					"http://localhost:8080/static/org-select/", http.StatusFound)
				return
			}

		case tokenizer.TokenTypeAppAccess:
			if !strings.HasPrefix(page, "app") {
				http.Redirect(w, r, "http://localhost:8080/static/app/", http.StatusFound)
				return
			}

		}

		next.ServeHTTP(w, r)
	})
}
