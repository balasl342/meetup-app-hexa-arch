package main

import (
	"context"
	"fmt"
	"log"
	"meetup-app-hexa-arch/internal/adapters/auth"
	"meetup-app-hexa-arch/internal/adapters/http"
	"meetup-app-hexa-arch/internal/adapters/http/handlers"
	"meetup-app-hexa-arch/internal/core/user"
	"meetup-app-hexa-arch/internal/shared/config"
	apphtpp "net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Load configuration
	config, err := config.LoadConfig("configs/app.yaml")
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	// Set up MongoDB client
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017") // MongoDB URI
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Initialize MongoDB repositories
	userRepo := user.NewMongoDBUserRepository(client, "meetings_app", "users")

	// Initialize services
	authService := auth.NewJWTService(config.JWTSecretKey) // Pass your secret key
	userService := user.NewUserService(userRepo, authService)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService)

	// Initialize router
	router := http.NewRouter(userHandler, nil, nil, authService, nil) // Pass other handlers as needed

	// Start the server
	fmt.Println("Starting Meetings App...")
	log.Fatal(apphtpp.ListenAndServe(":8080", router))
}
