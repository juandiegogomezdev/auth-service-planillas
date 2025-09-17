package authservice

import (
	"net/http"

	"proyecto/config"
	"proyecto/internal/mailer"
	"proyecto/internal/tokenizer"
	"proyecto/services/authservice/internal/service"
	"proyecto/services/authservice/internal/store"
	"proyecto/services/authservice/internal/transport"
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
	http.ListenAndServe(":8080", nil)

}
