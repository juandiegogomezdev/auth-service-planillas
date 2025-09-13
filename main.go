package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"proyecto/config"
	"proyecto/internal/service"
	"proyecto/internal/store"
	"proyecto/internal/transport"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	handlers := NewHandlers(db)

	http.HandleFunc("/auth/", handlers.userHandler)
	http.HandleFunc("/users", handlers.userHandler.RegisterRouter)
	http.HandleFunc("/users/", handlers.userHandler.RegisterRouter)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

type handlers struct {
	userHandler *transport.UserHandler
}

func NewHandlers(db *sql.DB) *handlers {
	userStore := store.NewUserStore(db)
	userService := service.NewUserService(userStore)
	userHandler := transport.NewUserRoutes(userService)
	return &handlers{userHandler: userHandler}
}
