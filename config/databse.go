package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client

func ConnectDatabase() {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Sets a timeout of 10 seconds for the connection process.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // Ensures the timeout context is released after the function finishes.

	// Connects to MongoDB using the specified options.
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err) // Logs an error and exits the program if the connection fails.
	}

	// Pings MongoDB to ensure it's reachable.
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("MongoDB not reachable: %v", err)
	}

	log.Println("Connected to MongoDB") // Logs a success message.
	DB = client                         // Assigns the connected client to the global variable `DB`.
}

// retrieves a collection from userdb databse
func GetCollection(collectionName string) *mongo.Collection {
	return DB.Database("usersdb").Collection(collectionName)
}
