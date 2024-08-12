package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Collection *mongo.Collection
	client     *mongo.Client
)

func ConnectDB() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error while loading env in connectDb", err)
	}
	Mongo_uri := os.Getenv("MONGO_URI")

	clientOptions := options.Client().ApplyURI(Mongo_uri)

	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal("Error :", err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Error :", err)
	}

	fmt.Println("Connection success")
	Collection = client.Database("golang_db").Collection("todos")
	if Collection == nil {
		log.Fatal("Failed to initialize collection")
	}

	fmt.Println("MongoDB connected and collection initialized")
	fmt.Println(Collection)
}

func DisconnectDB() {
	if client != nil {
		err := client.Disconnect(context.Background())
		if err != nil {
			log.Fatal("Error disconnecting from MongoDB:", err)
		}
		fmt.Println("MongoDB connection closed")
	}
}
