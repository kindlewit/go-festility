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

  // Create context
  ctx, cancel := context.WithTimeout(context.Background(), constants.QueryTimeout);
  defer cancel();

  // Ensure no other record has the same ID (duplicate)
  count, err := collection.CountDocuments(ctx, bson.M{ "id": data.Id });
  if err != nil {
    fmt.Println(err.Error());
    return false, constants.DetermineError(err);
  }
  if count > 0 {
    return false, constants.ErrDuplicateRecord; // Record already present
  }

  _, err = collection.InsertOne(ctx, data);
  if err != nil {
    fmt.Println(err.Error());
    return false, constants.DetermineError(err);
  }
  return true, nil;
}

// Fetches one festival record by fest id.
func GetFestival(client *mongo.Client, fid string) (data models.Fest, err error) {
  collection := client.Database("festility").Collection("festival"); // Collection to use

  // Create context
  ctx, cancel := context.WithTimeout(context.Background(), constants.QueryTimeout);
  defer cancel();

  query := bson.M{ "id": fid };

  err = collection.FindOne(ctx, query).Decode(&data);
  if err != nil {
    fmt.Println(err.Error());
    return data, constants.DetermineError(err);
  }

  return data, nil;
}
