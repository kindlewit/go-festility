package services

import (
	"context"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Creates indexes for necessary collections.
func Migrate() bool {
	client := Database.GetConnection(Database{})
	festCollection := client.Database("festility").Collection("festival")
	cinemaCollection := client.Database("festility").Collection("cinema")

	_, err := festCollection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "id", Value: 1}},
			Options: options.Index().SetUnique(true), // Fest IDs are unique
		},
	)
	if err != nil {
		return false
	}

	_, err = cinemaCollection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "id", Value: 1}},
			Options: options.Index().SetUnique(true), // Cinema IDs are unique
		},
	)
	if err != nil {
		return false
	}
	return true
}

// Clears db data for testing purpose only.
func Clear(name string) bool {
	if os.Getenv("GIN_MODE") == "release" || strings.Contains(os.Getenv("GOENV"), "prod") {
		// Not in debug/testing environment. Cannot clear the db
		return false
	}

	client := Database.GetConnection(Database{})
	collection := client.Database("festility").Collection(name) // Whichever database to clear

	// Create context
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := collection.DeleteMany(ctx, bson.M{})
	if err != nil {
		return false
	}
	return true
}
