package services

import (
  "fmt"
  "context"

  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo/options"
  "festility/models"
  "festility/constants"
)

// Creates multiple new screen records.
func CreateCinemaScreens(client *mongo.Client, screens []models.Screen) (success bool, err error) {
  collection := client.Database("festility").Collection("screen"); // Collection to use

  // Create context
  ctx, cancel := context.WithTimeout(context.Background(), constants.QueryTimeout);
  defer cancel();

  // data should already be sanitized
  data := make([]interface{}, len(screens));
  for i, s := range screens {
    data[i] = s;
  }

  _, err = collection.InsertMany(ctx, data);
  if err != nil {
    fmt.Println(err.Error());
    return false, constants.DetermineError(err);
  }
  return true, nil;
}

// Fetches a record by screen id.
func GetScreen(client *mongo.Client, screenID string) (data models.Screen, err error) {
  collection := client.Database("festility").Collection("screen"); // Collection to use
  query := bson.M{ "id": screenID };

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

// Fetches multiple screen records.
func GetScreensInBulk(client *mongo.Client, screenIDlist []string) (data []models.Screen, err error) {
  collection := client.Database("festility").Collection("screen"); // Collection to use
  query := bson.M{ "id": bson.M{ "$in": screenIDlist } };

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

// Fetches screens for a cinema by cinema id.
func GetCinemaScreens(client *mongo.Client, cinemaID string) (data []models.Screen, err error) {
  collection := client.Database("festility").Collection("screen"); // Collection to use
  query := bson.M{ "cinema_id": cinemaID };
  opts := options.Find();

  // New context for find query
  ctx, cancel := context.WithTimeout(context.Background(), constants.QueryTimeout);
  defer cancel();

  cursor, err := collection.Find(ctx, query, opts);
  if err != nil {
    fmt.Println(err.Error());
    return data, constants.DetermineError(err);
  }
  defer cursor.Close(ctx);


  for cursor.Next(ctx) { // Iterate cursor
    var d models.Screen;

    err = cursor.Decode(&d); // Decode cursor data into screen model
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

// Updates a screen record by id.
func ReplaceScreen(client *mongo.Client, screenID string, replacement models.Screen) (success bool, err error) {
  collection := client.Database("festility").Collection("screen");
  query := bson.M{ "id": screenID };

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

// Deletes a screen record by id.
