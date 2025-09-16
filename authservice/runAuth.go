package authservice

import (
	"net/http"

	"proyecto/authservice/internal/service"
	"proyecto/authservice/internal/store"
	"proyecto/authservice/internal/transport"
	"proyecto/config"
	"proyecto/internal/mailer"
	"proyecto/internal/tokenizer"
)

func RunAuth() {
	db := config.ConnectDB()
	tokenizerJWT := tokenizer.NewTokenizerJWT()
	mailersmtp := mailer.NewSMTPMailer("host", 1234, "username", "password", "from@example.com")
	defer db.Close()

	// Serve static files from the "static" directory
	fs := http.FileServer(http.Dir("./authservice/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	authStore := store.NewAuthStore(db)
	authService := service.NewAuthService(authStore, tokenizerJWT, mailersmtp)
	authHandler := transport.NewAuthHandler(authService)

	authHandler.SetupAuthRoutes()

}
