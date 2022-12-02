package services

import (
  "fmt"
  "time"
  "context"

  "go.mongodb.org/mongo-driver/mongo"
  "festility/models"
  "festility/constants"
  "go.mongodb.org/mongo-driver/mongo/options"
  "go.mongodb.org/mongo-driver/bson"
)

// Creates new slot records & returns success.
func CreateSlots(client *mongo.Client, slots []models.Slot) bool {
  collection := client.Database("festility").Collection("slot"); // Collection to use
  // data should already be sanitized
  data := make([]interface{}, len(slots))
  for i, s := range slots {
    data[i] = s;
  }

  _, err := collection.InsertMany(context.TODO(), data);
  if err != nil {
    fmt.Println(err.Error());
    return false;
  }
  return true;
}


// Fetches all slots for a given schedule id.
func GetScheduleSlots(client *mongo.Client, scheduleId string, optionals ...int64) ([]models.Slot, error) {
  var records []models.Slot; // Array of db records
  var err error;
  var limit int64 = int64(constants.SlotPageLimit);
  var skip int64 = int64(0);

  if (len(optionals) == 1) {
    limit = optionals[0];
  } else if (len(optionals) == 2) {
    limit = optionals[0];
    skip = optionals[1];
  }

  collection := client.Database("festility").Collection("slot"); // Collection to use
  query := bson.M{ "schedule_id": scheduleId };
  opts := options.Find().SetLimit(limit).SetSkip(skip);


  ctx, cancel := context.WithTimeout(context.TODO(), 30 * time.Second); // New context for find query
  defer cancel();

  cursor, err := collection.Find(ctx, query, opts);
  if err != nil {
    fmt.Println(err.Error());
    return records, constants.MongoReadError;
  }
  defer cursor.Close(ctx);


  for cursor.Next(ctx) { // Iterate cursor
    var d models.Slot;

    err = cursor.Decode(&d); // Decode cursor data into model
    if err != nil {
      fmt.Println(err.Error());
      return records, constants.DataParsingError;
    }

    // Add data if movie.
    if (d.Type == constants.SlotTypeMovie && d.MovieId != 0) {
      var movieData models.TMDBmovie;
      movieData, err = GetMovie(d.MovieId);
      if (err != nil) {
        return records, err;
      }
      d.Title = movieData.Title;
      d.Synopsis = movieData.Synopsis;
      d.Duration = movieData.Runtime;
    }

    records = append(records, d); // Push data record into array
  }
  if err = cursor.Err(); err != nil {
    fmt.Println(err.Error());
    return records, constants.MongoReadError;
   }

  return records, nil;
}


// Fetches all slots between given from and to time.
func GetScheduleSlotsByTime(client *mongo.Client, scheduleId string, from int, to int) ([]models.Slot, error) {
  var records []models.Slot;
  var err error;

  ctx, cancel := context.WithTimeout(context.TODO(), 30 * time.Second); // New context for find query
  defer cancel();

  collection := client.Database("festility").Collection("slot"); // Collection to use
  query := bson.M{
    "schedule_id": scheduleId,
    "start_time": bson.M{ "$gte": from, "$lt": to }, // Start time between dates
  };
  // Omission options (keys to include/omit)
  opts := options.Find().SetProjection(bson.M{
    "directors": 0,
    "original_title": 0,
    "genres": 0,
    "languages": 0,
    "countries": 0,
  });

  cur, err := collection.Find(ctx, query, opts);
  if err != nil {
    fmt.Println(err.Error());
    return records, constants.MongoReadError;
  }
  defer cur.Close(ctx);

  for cur.Next(ctx) { // Iterate cursor
    var d models.Slot;

    err := cur.Decode(&d); // Decode cursor data into model
    if err != nil {
      fmt.Println(err.Error());
      return records, constants.DataParsingError;
    }
    records = append(records, d); // Push data into array
  }
  if err := cur.Err(); err != nil {
    fmt.Println(err.Error());
    return records, constants.MongoReadError;
  }

  return records, nil;
}
