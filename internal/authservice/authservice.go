package authservice

import (
	"net/http"
	"proyecto/internal/authservice/serviceauth"
	"proyecto/internal/authservice/storeauth"
	"proyecto/internal/authservice/transport"
	"proyecto/internal/shared/mailer"

	"github.com/jmoiron/sqlx"
)

func RunAuth(db *sqlx.DB, mailersmtp mailer.ResendMailer, mux *http.ServeMux) {

	authStore := storeauth.NewAuthStore(db)
	authService := serviceauth.NewAuthService(authStore, mailersmtp)
	authHandler := transport.NewAuthHandler(authService)

	authHandler.SetupAuthRoutes(mux)

}
