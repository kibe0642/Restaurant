package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DBinstance() (*mongo.Client, error) {
	mongoDbURI := os.Getenv("MONGODB_URI")
	if mongoDbURI == "" {
		mongoDbURI = "mongodb://localhost:27017"
	}

	fmt.Println("Connecting to MongoDB at:", mongoDbURI)

	clientOptions := options.Client().ApplyURI(mongoDbURI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to create client: %w", err)
	}
	fmt.Println("connected to mongodb")
	return client, nil
}

var Client *mongo.Client

func Init() {
	var err error
	Client, err = DBinstance()
	if err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}
}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("restaurant").Collection(collectionName)

}
func Disconnect() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := Client.Disconnect(ctx); err != nil {
		log.Fatalf("failed to disconnect from MongoDB: %v", err)
	}
	fmt.Println("Disconnected from MongoDb")
}
