package main

import (
	"fmt"
	"log"
	"net/http"
	"proyecto/config"
	"proyecto/internal/authservice"
	"proyecto/internal/shared/mailer"
	"proyecto/internal/shared/tokenizer"
)

func main() {

	log.SetFlags(log.Lshortfile)

	// Serve static of login, register and confirm pages
	fs := http.FileServer(http.Dir("./internal/authservice/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Connect to DB, tokenizer and mailer
	db := config.ConnectDB()
	tokenizerJWT := tokenizer.NewTokenizerJWT()
	mailersmtp := mailer.NewSMTPMailer("host", 1234, "username", "password", "from@example.com")
	defer db.Close()

	// Initialize store, service and handler for each "micro" service
	authService, authHandler := authservice.RunAuth(db, tokenizerJWT, mailersmtp)
	authOrgService, authOrgHandler := authservice.RunAuthOrg(db, tokenizerJWT, mailersmtp)

	// Setup routes
	authHandler.SetupAuthRoutes()

	http.ListenAndServe(":8080", nil)

	// Run "micro" services
	fmt.Println("Services running in http://localhost:8080")
}
