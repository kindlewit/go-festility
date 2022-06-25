package main

import (
	"os"
	"fmt"
	"time"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Connect to mongodb.
func connect() *mongo.Client {
	var MONGO_URI = os.Getenv("MONGO_URI");

	// Create context
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second);
	defer cancel();

	// Create the client
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MONGO_URI));
	if err != nil {
		panic(err);
	}

	return client; // Return mongo client
}

// Bulk insert movie records into mongodb.
func bulkInsertMovies(client *mongo.Client, movies []Movie) {
	collection := client.Database("festility").Collection("movies"); // Collection to use

	data := make([]interface{}, len(movies));

	for i, m := range movies {
		data[i] = m;
	}

	_, err := collection.InsertMany(context.TODO(), data, options.InsertMany().SetOrdered(false));
	if err != nil {
		panic(err);
	}
}

// Fetch all movie documents in mongodb
func allMovies(client *mongo.Client) []Movie {
	// Create context
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second);
	defer cancel();

	collection := client.Database("festility").Collection("movies"); // Collection to use

	cur, err := collection.Find(ctx, bson.D{})

	if err != nil { panic(err) }

	defer cur.Close(ctx);

	var res []Movie;

	for cur.Next(ctx) { // Iterate cursor
			var result Movie;

			err := cur.Decode(&result);
			if err != nil { panic(err) }

			res = append(res, result);
			// do something with result....
	}
	if err := cur.Err(); err != nil {
			panic(err);
	}

	return res; // Return document slice
}

// Creates new festival record & returns success.
func createFest(client *mongo.Client, data Fest) string {
	collection := client.Database("festility").Collection("festival"); // Collection to use

	result, err := collection.InsertOne(context.TODO(), data);
	if err != nil {
		panic(err);
		return "";
	}

	return fmt.Sprintf("%v", result.InsertedID);
}

// Fetches one festival record.
func getFest(client *mongo.Client, id string) Fest {
	collection := client.Database("festility").Collection("festival"); // Collection to use

	objectId, err := primitive.ObjectIDFromHex(id);
	if err != nil {
		panic(err);
	}

	query := bson.M{ "_id": objectId };

	var data Fest;
	err = collection.FindOne(context.TODO(), query).Decode(&data); // Throwing mongo: no documents in result
	if err != nil {
		panic(err);
	}

	return data;
}
