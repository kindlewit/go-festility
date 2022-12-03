package services

import (
  "fmt"
  "context"

  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/bson"
  "festility/models"
  "festility/constants"
)

// Creates new cinema record.
func CreateCinema(client *mongo.Client, data models.Cinema) (bool, error) {
  collection := client.Database("festility").Collection("cinema"); // Collection to use

  // Ensure no other record has the same ID or Name (duplicate)
  query := bson.M{ "$or": []bson.M{ bson.M{ "id": data.Id }, bson.M{ "name": data.Name } } };
  count, err := collection.CountDocuments(context.TODO(), query);
  if err != nil {
    fmt.Println(err.Error());
    return false, constants.MongoReadError;
  }
  if count > 0 {
    return false, constants.DuplicateRecordError; // Record already present
  }

  result, err := collection.InsertOne(context.TODO(), data); // Assuming data is already sanitized
  if err != nil {
    fmt.Println(err.Error());
    return false, constants.MongoWriteError;
  }
  return result.InsertedID != nil, nil;
}

// Fetches the cinema record by id.
func GetCinema(client *mongo.Client, cinemaID string) (data models.Cinema, err error) {
  collection := client.Database("festility").Collection("cinema"); // Collection to use
  query := bson.M{ "id": cinemaID };

  err = collection.FindOne(context.TODO(), query).Decode(&data);
  if err != nil {
    fmt.Println(err.Error());
    if (err.Error() == "mongo: no documents in result") {
      return data, constants.NoSuchRecordError;
    }
    return data, constants.MongoReadError;
  }

  return data, nil;
}
