package main

import (
	authservice "chatapp/api/auth_service"
	messageservice "chatapp/api/message_service"
	mongodb "chatapp/db/MongoDB"
	postgresql "chatapp/db/Postgresql"
	"chatapp/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize MongoDB
	_, err := mongodb.Connect()
	if err != nil {
		log.Fatalf("MongoDB connection failed: %v", err)
	}
	defer func() {
		if err := mongodb.Disconnect(); err != nil {
			log.Printf("MongoDB disconnection error: %v", err)
		}
	}()

	// Get MongoDB collection
	messageCollection, err := mongodb.GetCollection("chatapp", "messages")
	if err != nil {
		log.Fatalf("Failed to get MongoDB collection: %v", err)
	}

	// Initialize PostgreSQL
	dbConn := postgresql.Database()
	defer dbConn.Close()

	// Initialize repositories
	dbLayer := postgresql.NewUserRepository(dbConn)
	mongoRepo := mongodb.NewMessageRepository(messageCollection)

	// Initialize services
	authService := authservice.NewAuthService(dbLayer)
	messageservice := messageservice.NewMessageService(mongoRepo)

	// Initialize router
	appService := router.NewService(authService, messageservice)

	// Initialize router
	routerLayer := router.NewRouter(appService)

	// Setup Gin router
	r := gin.Default()
	routerLayer.RoutersSetup(r)

	log.Println("Server running on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
