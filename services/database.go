package services

import (
	"context"
	"os"
	"time"

	"github.com/kindlewit/go-festility/constants"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IDatabase interface {
	GetConnection()
	Disconnect()
	Insert()
	Retrieve()
	Count()
}

type Database struct {
	Instance *mongo.Client
}

var db Database

// Creates a connection to the database.
func (d Database) GetConnection() *mongo.Client {
	if d.Instance == nil {
		var MONGO_URI = os.Getenv("MONGO_URI")

		// Create context
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		// Create the connection
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(MONGO_URI))
		if err != nil {
			panic(err)
			// TODO: handle this situation
		}

		d.Instance = client
	}
	return d.Instance
}

// Closes the database connection. Always called from the relevant service.
func (d Database) Disconnect() {
	if d.Instance != nil {
		err := d.Instance.Disconnect(context.Background())
		if err != nil {
			panic(err)
			// TODO: handle this situation
		}
		d.Instance = nil
	}
}

// Counts the records which match the given query.
func (d Database) Count(tableName string, query bson.M) (int64, error) {
	client := d.GetConnection()
	collection := client.Database(constants.DatabaseName).Collection(tableName)

	// Create context
	ctx, cancel := context.WithTimeout(context.Background(), constants.QueryTimeout)
	defer cancel()

	return collection.CountDocuments(ctx, query)
}

// Inserts a record into the given table.
func (d Database) Insert(tableName string, data interface{}) (*mongo.InsertOneResult, error) {
	client := d.GetConnection()
	collection := client.Database(constants.DatabaseName).Collection(tableName)

	// Create context
	ctx, cancel := context.WithTimeout(context.Background(), constants.QueryTimeout)
	defer cancel()

	return collection.InsertOne(ctx, data)
}

// Inserts multiple records into a given table.
func (d Database) InsertMany(tableName string, data []interface{}) (*mongo.InsertManyResult, error) {
	client := d.GetConnection()
	collection := client.Database(constants.DatabaseName).Collection(tableName)

	// Create context
	ctx, cancel := context.WithTimeout(context.Background(), constants.QueryTimeout)
	defer cancel()

	return collection.InsertMany(ctx, data)
}

// Retrieves one record from a table matching the query.
func (d Database) Retrieve(tableName string, query bson.M) (bson.Raw, error) {
	client := d.GetConnection()
	collection := client.Database(constants.DatabaseName).Collection(tableName)

	// Create context
	ctx, cancel := context.WithTimeout(context.Background(), constants.QueryTimeout)
	defer cancel()

	return collection.FindOne(ctx, query).Raw()
}

// Retrieves multiple records from a table matching the query.
func (d Database) RetrieveMany(tableName string, query bson.M, opts *options.FindOptions) (*mongo.Cursor, error) {
	// TODO: optimize search using options parameter
	client := d.GetConnection()
	collection := client.Database(constants.DatabaseName).Collection(tableName)

	// Create context
	ctx, cancel := context.WithTimeout(context.Background(), constants.QueryTimeout)
	defer cancel()

	return collection.Find(ctx, query, opts)
}

// Replaces one record in a table.
func (d Database) Replace(tableName string, query bson.M, replacement interface{}) (*mongo.UpdateResult, error) {
	client := d.GetConnection()
	collection := client.Database(constants.DatabaseName).Collection(tableName)

	// Create context
	ctx, cancel := context.WithTimeout(context.Background(), constants.QueryTimeout)
	defer cancel()

	return collection.ReplaceOne(ctx, query, replacement)
}

// Removes one record from a table.
func (d Database) DeleteOne(tableName string, query bson.M) (*mongo.DeleteResult, error) {
	client := d.GetConnection()
	collection := client.Database(constants.DatabaseName).Collection(tableName)

	// Create context
	ctx, cancel := context.WithTimeout(context.Background(), constants.QueryTimeout)
	defer cancel()

	return collection.DeleteOne(ctx, query)
}
