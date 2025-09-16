package main

import (
	"fmt"
	"log"
	"net/http"
	"proyecto/authservice"
)

func main() {

	log.SetFlags(log.Lshortfile)

	// Run "micro" services
	authservice.RunAuth()

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
