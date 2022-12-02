package services

import (
  "fmt"
  "context"

  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/bson"
  "festility/models"
  "festility/constants"
)

// Creates new schedule from slots list.
func CreateSchedule(client *mongo.Client, data models.Schedule) bool {
  collection := client.Database("festility").Collection("schedule"); // Collection to use
  _, err := collection.InsertOne(context.TODO(), data);
  if err != nil {
    fmt.Println(err.Error());
    return false;
  }
  return true;
}

// Fetches the schedule record by fest id & schedule id.
func GetSchedule(client *mongo.Client, festId string, scheduleId string) (models.Schedule, error) {
  collection := client.Database("festility").Collection("schedule"); // Collection to use
  query := bson.M{ "id": scheduleId, "fest_id": festId };

  var doc models.Schedule;
  err := collection.FindOne(context.TODO(), query).Decode(&doc);
  if err != nil {
    fmt.Println(err.Error());
    return doc, constants.MongoReadError;
  }

  return doc, nil;
}

// Fetches the default schedule of a festival
func GetDefaultScheduleID(client *mongo.Client, festId string) (string, error) {
  var record models.Schedule;

  collection := client.Database("festility").Collection("schedule"); // Collection to use
  query := bson.M{ "fest_id": festId, "custom": false }; // Default schedule will have Custom=false

  err := collection.FindOne(context.TODO(), query).Decode(&record);
  if err != nil {
    fmt.Println(err.Error());
    if (err.Error() == "mongo: no documents in result") {
      return "", constants.NonExistantDefaultSchedule;
    }
    return "", constants.MongoReadError;
  }
  return record.Id, nil;
}

// Checks if id already exists in db.
func IsUniqueScheduleID(client *mongo.Client, id string) bool {
  collection := client.Database("festility").Collection("schedule"); // Collection to use
  query := bson.M{ "id": id }; // Docs with the same id
  count, err := collection.CountDocuments(context.TODO(), query); // Count query
  if err != nil {
    fmt.Println(err.Error());
    return false;
  }
  return count < 1; // No docs with the id
}
