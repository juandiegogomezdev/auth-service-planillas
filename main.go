package main

import (
	"fmt"
	"log"
	"proyecto/services/authservice"
)

func main() {

	log.SetFlags(log.Lshortfile)

	// Run "micro" services
	go authservice.RunAuth()
	fmt.Println("Services running in http://localhost:8080")
	select {}
}
