package services

import (
  "fmt"
  "context"

  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/bson"
  "festility/models"
  "festility/constants"
)


// Creates new festival record & returns inserted ID.
func CreateFestival(client *mongo.Client, data models.Fest) (bool, error) {
  collection := client.Database("festility").Collection("festival"); // Collection to use

  // Ensure no other record has the same ID (duplicate)
  count, err := collection.CountDocuments(context.TODO(), bson.M{ "id": data.Id });
  if err != nil {
    fmt.Println(err.Error());
    return false, constants.MongoReadError;
  }
  if count > 0 {
    return false, constants.DuplicateRecordError; // Record already present
  }

  _, err = collection.InsertOne(context.TODO(), data);
  if err != nil {
    fmt.Println(err.Error());
    return false, constants.MongoWriteError;
  }
  return true, nil;
}

// Fetches one festival record by fest id.
func GetFestival(client *mongo.Client, fid string) (models.Fest, error) {
  collection := client.Database("festility").Collection("festival"); // Collection to use

  query := bson.M{ "id": fid };

  var data models.Fest;
  err := collection.FindOne(context.TODO(),query).Decode(&data);
  // Throwing "mongo: no documents in result" error
  if err != nil {
    if err.Error() == "mongo: no documents in result" { // errors.Is comparison does not work
      return data, constants.NoSuchRecordError
    }
    panic(err);
  }

  return data, nil;
}
