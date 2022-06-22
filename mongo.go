package main

import (
	"os"
	"time"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)

// Connect to mongodb
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

// Bulk insert movie records into mongodb
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
