package services

import (
  "fmt"
  "context"

  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/bson"
  "github.com/kindlewit/go-festility/models"
  "github.com/kindlewit/go-festility/constants"
)

// Creates new schedule.
func CreateSchedule(client *mongo.Client, data models.Schedule) (bool, error) {
  collection := client.Database("festility").Collection("schedule"); // Collection to use

  // Create context
  ctx, cancel := context.WithTimeout(context.Background(), constants.QueryTimeout);
  defer cancel();

  // Ensure no other record has the same ID (duplicate)
  count, err := collection.CountDocuments(ctx, bson.M{ "id": data.Id });
  if err != nil {
    fmt.Println(err.Error());
    return false, constants.DetermineInternalErrMsg(err);
  }
  if count > 0 {
    return false, constants.ErrDuplicateRecord; // Record already present
  }

  _, err = collection.InsertOne(ctx, data);
  if err != nil {
    fmt.Println(err.Error());
    return false, constants.DetermineInternalErrMsg(err);
  }
  return true, nil;
}

// Fetches the schedule record by fest id & schedule id.
func GetSchedule(client *mongo.Client, festId string, scheduleId string) (data models.Schedule, err error) {
  collection := client.Database("festility").Collection("schedule"); // Collection to use
  query := bson.M{ "id": scheduleId, "fest_id": festId };

  // Create context
  ctx, cancel := context.WithTimeout(context.Background(), constants.QueryTimeout);
  defer cancel();

  err = collection.FindOne(ctx, query).Decode(&data);
  if err != nil {
    fmt.Println(err.Error());
    return data, constants.DetermineInternalErrMsg(err);
  }

  return data, nil;
}

// Fetches the default schedule of a festival
func GetDefaultScheduleID(client *mongo.Client, festId string) (string, error) {
  var record models.Schedule;

  collection := client.Database("festility").Collection("schedule"); // Collection to use
  query := bson.M{ "fest_id": festId, "custom": false }; // Default schedule will have Custom=false
  // Create context
  ctx, cancel := context.WithTimeout(context.Background(), constants.QueryTimeout);
  defer cancel();

  err := collection.FindOne(ctx, query).Decode(&record);
  if err != nil {
    fmt.Println(err.Error());
    return "", constants.DetermineInternalErrMsg(err);
  }
  return record.Id, nil;
}

// Checks if id already exists in db.
func IsUniqueScheduleID(client *mongo.Client, id string) bool {
  collection := client.Database("festility").Collection("schedule"); // Collection to use
  query := bson.M{ "id": id }; // Docs with the same id

  // Create context
  ctx, cancel := context.WithTimeout(context.Background(), constants.QueryTimeout);
  defer cancel();

  count, err := collection.CountDocuments(ctx, query); // Count query
  if err != nil {
    fmt.Println(err.Error());
    return false;
  }
  return count < 1; // No docs with the id
}
