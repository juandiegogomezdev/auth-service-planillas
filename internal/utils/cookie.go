package utils

import (
	"fmt"
	"net/http"
)

func SetCookie(w http.ResponseWriter, name string, value string) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		HttpOnly: true,
		// Secure:   true,
	}
	http.SetCookie(w, cookie)
}

func GetCookie(r *http.Request, name string) (string, error) {
	cookie, err := r.Cookie(name)
	if err != nil {
		return "", fmt.Errorf("error getting cookie: %w", err)
	}
	return cookie.Value, nil
}
