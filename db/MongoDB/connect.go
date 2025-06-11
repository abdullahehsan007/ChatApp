package mongodb

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	clientInstance *mongo.Client
	clientOnce     sync.Once
	clientErr      error
)

// Connect initializes the MongoDB client connection
func Connect() (*mongo.Client, error) {
	clientOnce.Do(func() {
		// Load .env file
		if err := godotenv.Load(".env"); err != nil {
			clientErr = fmt.Errorf("error loading .env file: %w", err)
			return
		}

		// Get MongoDB URI from environment
		mongoURI := os.Getenv("MONGO_URI")
		if mongoURI == "" {
			clientErr = fmt.Errorf("MONGO_URI not set in environment variables")
			return
		}

		// Set client options
		clientOptions := options.Client().ApplyURI(mongoURI)

		// Connect to MongoDB
		client, err := mongo.Connect(context.Background(), clientOptions)
		if err != nil {
			clientErr = fmt.Errorf("failed to connect to MongoDB: %w", err)
			return
		}

		// Check the connection
		if err = client.Ping(context.Background(), nil); err != nil {
			clientErr = fmt.Errorf("failed to ping MongoDB: %w", err)
			return
		}

		clientInstance = client
		log.Println("Successfully connected to MongoDB!")
	})

	return clientInstance, clientErr
}

// GetCollection returns a handle to a MongoDB collection
func GetCollection(dbName, collectionName string) (*mongo.Collection, error) {
	if clientInstance == nil {
		return nil, fmt.Errorf("MongoDB client not initialized")
	}

	return clientInstance.Database(dbName).Collection(collectionName), nil
}

// Disconnect safely closes the MongoDB connection
func Disconnect() error {
	if clientInstance == nil {
		return nil
	}
	return clientInstance.Disconnect(context.Background())
}