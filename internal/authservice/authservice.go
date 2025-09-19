package authservice

import (
	"proyecto/internal/authservice/serviceauth"
	"proyecto/internal/authservice/storeauth"
	"proyecto/internal/authservice/transport"
	"proyecto/internal/shared/mailer"
	"proyecto/internal/shared/tokenizer"

	"github.com/jmoiron/sqlx"
)

func RunAuth(db *sqlx.DB, tokenizerJWT tokenizer.TokenizerJWT, mailersmtp mailer.Mailer) (*serviceauth.ServiceAuth, *transport.Handler) {

	authStore := storeauth.NewAuthStore(db)
	authService := serviceauth.NewAuthService(authStore, tokenizerJWT, mailersmtp)
	authHandler := transport.NewAuthHandler(authService)

	authHandler.SetupAuthRoutes()

	return authService, authHandler

}
