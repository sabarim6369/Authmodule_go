package db

import (
	"context"
	"time"
	"log"	
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
const connectionString = "mongodb://localhost:27017/godb"
var Client *mongo.Client
func Dbconnection() *mongo.Client{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Could not ping to MongoDB:", err)
	}

	log.Println("Connected to MongoDB successfully")

	Client = client
	return client

}