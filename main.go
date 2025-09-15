package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"proyecto/config"
	"proyecto/internal/mailer"
	"proyecto/internal/service"
	"proyecto/internal/store"
	"proyecto/internal/tokenizer"
	"proyecto/internal/transport"
)

func main() {

	log.SetFlags(log.Lshortfile)

	db := config.ConnectDB()
	defer db.Close()

	// Serve static files from the "static" directory
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Setup API routes
	handlers := NewHandlers(db)
	handlers.authHandler.SetupAuthRoutes()
	handlers.userHandler.SetupUserRoutes()

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

type handlers struct {
	authHandler *transport.AuthHandler
	userHandler *transport.UserHandler
}

func NewHandlers(db *sql.DB) *handlers {
	tokenizerJWT := tokenizer.NewTokenizerJWT()
	mailersmtp := mailer.NewSMTPMailer("host", 1234, "username", "password", "from@example.com")

	userStore := store.NewUserStore(db)
	userService := service.NewUserService(userStore)
	userHandler := transport.NewUserHandler(userService)

	authStore := store.NewAuthStore(db)
	authService := service.NewAuthService(authStore, tokenizerJWT, mailersmtp)
	authHandler := transport.NewAuthHandler(authService)

	return &handlers{
		userHandler: userHandler,
		authHandler: authHandler,
	}
}
