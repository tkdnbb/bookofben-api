package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoClient *mongo.Client
	database    *mongo.Database
)

// InitMongoDB initializes the MongoDB connection
func InitMongoDB(connectionString, databaseName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Test connection
	err = client.Ping(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	mongoClient = client
	database = client.Database(databaseName)

	fmt.Println("Connected to MongoDB successfully")
	return nil
}

// GetDatabase returns the database instance
func GetDatabase() *mongo.Database {
	return database
}

// Close closes the MongoDB connection
func Close() error {
	if mongoClient != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := mongoClient.Disconnect(ctx); err != nil {
			return fmt.Errorf("error disconnecting from MongoDB: %w", err)
		}
		log.Println("Disconnected from MongoDB")
	}
	return nil
}
