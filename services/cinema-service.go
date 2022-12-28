package services

import (
  "fmt"
  "context"

  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/bson"
  "github.com/kindlewit/go-festility/models"
  "github.com/kindlewit/go-festility/constants"
)

// Creates new cinema record.
func CreateCinema(client *mongo.Client, data models.Cinema) (bool, error) {
  collection := client.Database("festility").Collection("cinema"); // Collection to use

  // Create context
  ctx, cancel := context.WithTimeout(context.Background(), constants.QueryTimeout);
  defer cancel();

  // Ensure no other record has the same ID or Name (duplicate)
  query := bson.M{ "$or": []bson.M{ bson.M{ "id": data.Id }, bson.M{ "name": data.Name } } };
  count, err := collection.CountDocuments(ctx, query);
  if err != nil {
    fmt.Println(err.Error());
    return false, constants.DetermineError(err);
  }
  if count > 0 {
    return false, constants.ErrDuplicateRecord; // Record already present
  }

  result, err := collection.InsertOne(ctx, data); // Assuming data is already sanitized
  if err != nil {
    fmt.Println(err.Error());
    return false, constants.DetermineError(err);
  }
  return result.InsertedID != nil, nil;
}

// Fetches the cinema record by id.
func GetCinema(client *mongo.Client, cinemaID string) (data models.Cinema, err error) {
  collection := client.Database("festility").Collection("cinema"); // Collection to use
  query := bson.M{ "id": cinemaID };

  // Create context
  ctx, cancel := context.WithTimeout(context.Background(), constants.QueryTimeout);
  defer cancel();


  err = collection.FindOne(ctx, query).Decode(&data);
  if err != nil {
    fmt.Println(err.Error());
    return data, constants.DetermineError(err);
  }

  return data, nil;
}

// Fetches multiple cinema records.
func GetCinemasInBulk(client *mongo.Client, cinemaIDlist []string) (data []models.Screen, err error) {
  collection := client.Database("festility").Collection("cinema"); // Collection to use
  query := bson.M{ "id": bson.M{ "$in": cinemaIDlist } };

  // Create context
  ctx, cancel := context.WithTimeout(context.Background(), constants.QueryTimeout);
  defer cancel();

  cursor, err := collection.Find(ctx, query);
  if err != nil {
    fmt.Println(err.Error());
    return data, constants.DetermineError(err);
  }
  defer cursor.Close(ctx);

  for cursor.Next(ctx) { // Iterate cursor
    var d models.Screen;

    err = cursor.Decode(&d); // Decode cursor data into model
    if err != nil {
      fmt.Println(err.Error());
      return data, constants.ErrDataParse;
    }

    data = append(data, d); // Push data record into array
  }
  if err = cursor.Err(); err != nil {
    fmt.Println(err.Error());
    return data, constants.DetermineError(err);
  }

  return data, nil;
}

// Updates a cinema record by id.
func ReplaceCinema(client *mongo.Client, cinemaID string, replacement models.Cinema) (success bool, err error) {
  collection := client.Database("festility").Collection("cinema"); // Collection to use
  query := bson.M{ "id": cinemaID };

  // Create context
  ctx, cancel := context.WithTimeout(context.Background(), constants.QueryTimeout);
  defer cancel();


  _, err = collection.ReplaceOne(ctx, query, replacement);
  if err != nil {
    fmt.Println(err.Error());
    return false, constants.DetermineError(err);
  }
  return true, nil;
}

// Deletes the cinema record by id.
func DeleteCinema(client *mongo.Client, cinemaID string) (success bool, err error) {
  collection := client.Database("festility").Collection("cinema"); // Collection to use
  query := bson.M{ "id": cinemaID };

  // Create context
  ctx, cancel := context.WithTimeout(context.Background(), constants.QueryTimeout);
  defer cancel();


  _, err = collection.DeleteOne(ctx, query);
  if err != nil {
    fmt.Println(err.Error());
    return false, constants.DetermineError(err);
  }
  return true, nil;
}
