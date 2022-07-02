package main

import (
	"os"
	"fmt"
	"time"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
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

func disconnect(client *mongo.Client) {
	if client == nil {
		return;
	}
	err := client.Disconnect(context.Background());
	if err != nil {
		panic(err);
	}
}

func migrate() bool {
	client := connect();
	collection := client.Database("festility").Collection("festival");

	_, err := collection.Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys: bson.D{{ Key: "id", Value: 1 }},
			Options: options.Index().SetUnique(true), // Fest IDs are unique
		},
	);
	if err != nil { return false; }
	return true;
}

// Bulk insert movie records into mongodb.
// func bulkInsertMovies(client *mongo.Client, movies []Movie) {
// 	collection := client.Database("festility").Collection("movies"); // Collection to use

// 	data := make([]interface{}, len(movies));
// 	for i, m := range movies {
// 		data[i] = m;
// 	}

// 	_, err := collection.InsertMany(context.TODO(), data, options.InsertMany().SetOrdered(false));
// 	if err != nil {
// 		panic(err);
// 	}
// }

// Fetches all movie documents in mongodb
// func allMovies(client *mongo.Client) []Movie {
// 	// Create context
// 	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second);
// 	defer cancel();

// 	collection := client.Database("festility").Collection("movies"); // Collection to use

// 	cur, err := collection.Find(ctx, bson.D{})
// 	if err != nil { panic(err) }
// 	defer cur.Close(ctx);

// 	var res []Movie;

// 	// TODO: Add pagination

// 	for cur.Next(ctx) { // Iterate cursor
// 		var result Movie;

// 		err := cur.Decode(&result);
// 		if err != nil { panic(err) }

// 		res = append(res, result);
// 		// do something with result....
// 	}
// 	if err := cur.Err(); err != nil {
// 		panic(err);
// 	}
// 	return res; // Return document slice
// }

// Creates new festival record & returns inserted ID.
func createFest(client *mongo.Client, data Fest) string {
	collection := client.Database("festility").Collection("festival"); // Collection to use

	// Ensure no duplication
	count, err := collection.CountDocuments(context.TODO(), bson.M{ "id": data.Id });
	if err != nil {
		panic(err);
	}
	if count > 0 {
		return DuplicateRecord; // Record already present
	}

	result, err := collection.InsertOne(context.TODO(), data);
	if err != nil {
		panic(err);
	}
	return fmt.Sprintf("%v", result.InsertedID);
}

// Fetches one festival record by fest id.
func getFest(client *mongo.Client, fid string) Fest {
	collection := client.Database("festility").Collection("festival"); // Collection to use

	query := bson.M{ "id": fid };

	var data Fest;
	err := collection.FindOne(context.TODO(),query).Decode(&data);
	// Throwing mongo: no documents in error
	if err != nil {
		panic(err);
	}

	return data;
}

// Creates new slot records & returns success.
func bulkCreateSlot(client *mongo.Client, slots []Slot) bool {
	collection := client.Database("festility").Collection("slot"); // Collection to use
	// data should already be sanitized
	data := make([]interface{}, len(slots))
	for i, s := range slots {
		data[i] = s;
	}

	_, err := collection.InsertMany(context.TODO(), data);
	if err != nil {
		panic(err);
		return false;
	}
	return true;
}

// Fetches all slots for a given schedule id.
func getSlotsOfSchedule(client *mongo.Client, schId string) []Slot {
	ctx, cancel := context.WithTimeout(context.TODO(), 30 * time.Second);
	defer cancel();

	collection := client.Database("festility").Collection("slot"); // Collection to use
	query := bson.M{ "schedule_id": schId };

	cur, err := collection.Find(ctx, query);
	if err != nil { panic(err); }
	defer cur.Close(ctx);

	var docs []Slot; // Array

	for cur.Next(ctx) { // Iterate cursor
		var d Slot;

		err := cur.Decode(&d); // Decode
		if err != nil { panic(err); }
		docs = append(docs, d); // Push into array
	}
	if err := cur.Err(); err != nil { panic(err); }

	return docs;
}

// Fetches all slots between given from and to time.
func findSlotsByTime(client *mongo.Client, schId string, from int, to int) []Slot {
	ctx, cancel := context.WithTimeout(context.TODO(), 30 * time.Second);
	defer cancel();

	collection := client.Database("festility").Collection("slot"); // Collection to use
	query := bson.M{
		"schedule_id": schId,
		"start_time": bson.M{ "$gte": from, "$lt": to }, // Start time between dates
	};
	// Omission options
	opts := options.Find().SetProjection(bson.M{
		"directors": 0,
		"original_title": 0,
		"genres": 0,
		"languages": 0,
		"countries": 0,
	});

	cur, err := collection.Find(ctx, query, opts);
	if err != nil { panic(err); }
	defer cur.Close(ctx);

	var docs []Slot; // Array

	for cur.Next(ctx) { // Iterate cursor
		var d Slot;

		err := cur.Decode(&d); // Decode
		if err != nil { panic(err); }
		docs = append(docs, d); // Push into array
	}
	if err := cur.Err(); err != nil { panic(err); }

	return docs;

}

// Creates new schedule from slots list.
func createSchedule(client *mongo.Client, data Schedule) bool {
	collection := client.Database("festility").Collection("schedule"); // Collection to use
	_, err := collection.InsertOne(context.TODO(), data);
	if err != nil {
		panic(err);
		return false;
	}
	return true;
}

// Fetches the default schedule id for a fest.
func getDefaultScheduleId(client *mongo.Client, fid string) string {
	collection := client.Database("festility").Collection("schedule"); // Collection to use
	query := bson.M{ "fest_id": fid, "custom": false }; // Default schedule will have Custom=false

	var doc Schedule;
	err := collection.FindOne(context.TODO(), query).Decode(&doc);
	if err != nil {
		panic(err);
		return "";
	}
	return doc.Id;
}

// Fetches the schedule record by fest id & schedule id.
func getSchedule(client *mongo.Client, fid string, schId string) Schedule {
	collection := client.Database("festility").Collection("schedule"); // Collection to use
	query := bson.M{ "id": schId, "fest_id": fid };

	var doc Schedule;
	err := collection.FindOne(context.TODO(), query).Decode(&doc);
	if err != nil {
		panic(err);
		return Schedule{};
	}

	return doc;
}

// Checks if newId already exists in db.
func ensureUniqueScheduleId(client *mongo.Client, newId string) bool {
	collection := client.Database("festility").Collection("schedule"); // Collection to use
	query := bson.M{ "id": newId }; // Docs with the same id
	count, err := collection.CountDocuments(context.TODO(), query); // Count query
	if err != nil {
		panic(err);
		return false;
	}
	return count < 1; // No docs under query
}
