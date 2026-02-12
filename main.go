package main

import (
	"log"
	"mywallet/config"
	"mywallet/server"
	"mywallet/server/http"
)

func main() {
	// Initialize the server and defer the closing of resources
	config := config.LoadConfig()

	err := server.Init(config)
	if err != nil {
		log.Fatalf("Failed to initialize server: %v", err)
	}
	defer server.Close()

	// Create the Gin router
	router := http.NewServer()

	// Start the HTTP server
	log.Println("Server starting at :" + config.ServerPort)
	router.Run(":" + config.ServerPort)
}
