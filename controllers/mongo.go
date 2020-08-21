package controllers

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// DB connection string
// const connectionString = "mongodb://localhost:27017" or Atlas Connection String
const connectionString = "mongodb://localhost:27017"

// Database Name
const dbName = "todos"


// collection object/instance
var todoscol *mongo.Collection

// create connection with mongo db
func init() {

	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Add any collections needed here
	todoscol = client.Database(dbName).Collection("todos")

	// If you need indexes, define them like this per collection
	todosIndexes := mongo.IndexModel{
		Keys: bson.M{
			"title": 1, // index in ascending order
		}, Options: nil,
	}
	_, err = todoscol.Indexes().CreateOne(context.TODO(), todosIndexes)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Collection instance created!")
}