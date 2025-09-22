package main

import (
	"fmt"
	"log"
	"net/http"
	"proyecto/config"
	"proyecto/internal/authservice"
	"proyecto/internal/shared/mailer"
	"proyecto/internal/shared/middleware"
)

func main() {

	log.SetFlags(log.Lshortfile)
	// Connect to DB, and mailer
	db := config.ConnectDB()
	defer db.Close()
	mailersmtp := mailer.NewResendMailer(config.Resend.APIKey, "Acme <onboarding@resend.dev>")

	// --- CREATE ROUTERS ---

	// Create router for auth service
	muxAuth := http.NewServeMux()
	authservice.RunAuth(db, mailersmtp, muxAuth)
	// // Create router for org service
	// muxOrg := http.NewServeMux()
	// orgservice.RunOrg(db, mailersmtp, muxOrg)
	// Create router for vehicle

	// create router for

	mainMux := http.NewServeMux()
	mainMux.Handle("/auth/", middleware.CorsMiddleware(muxAuth))

	// Serve static of login, register and confirm pages
	fs := http.FileServer(http.Dir("./internal/authservice/static"))
	mainMux.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Services running in http://localhost:8080")
	http.ListenAndServe(":8080", mainMux)

}
