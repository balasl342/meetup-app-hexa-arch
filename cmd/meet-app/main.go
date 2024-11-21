package main

import (
	"fmt"
	"log"
	"meetings-app/internal/adapters/auth"
	"meetings-app/internal/shared/config"
)

func main() {
	// Load configuration
	config, err := config.LoadConfig("configs/app.yaml")
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	// Initialize services
	_ = auth.NewService(config)

	// Start the server (this will depend on your chosen web framework)
	fmt.Println("Starting Meetings App...")
	// More logic for server setup
}
