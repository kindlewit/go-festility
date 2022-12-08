package services

import (
  "fmt"
  "time"
  "context"

  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo/options"
  "festility/models"
  "festility/constants"
)

// Creates new slot records & returns success.
func CreateSlots(client *mongo.Client, slots []models.Slot) bool {
  collection := client.Database("festility").Collection("slot"); // Collection to use
  // Create context
  ctx, cancel := context.WithTimeout(context.Background(), constants.QueryTimeout);
  defer cancel();

  // data should already be sanitized
  data := make([]interface{}, len(slots))
  for i, s := range slots {
    data[i] = s;
  }

  _, err := collection.InsertMany(ctx, data);
  if err != nil {
    fmt.Println(err.Error());
    return false;
  }
  return true;
}


// Fetches all slots for a given schedule id.
func GetScheduleSlots(client *mongo.Client, scheduleId string, optionals ...int64) (records []models.Slot, err error) {
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

  // New context for find query
  ctx, cancel := context.WithTimeout(context.Background(), constants.QueryTimeout);
  defer cancel();

  cursor, err := collection.Find(ctx, query, opts);
  if err != nil {
    fmt.Println(err.Error());
    return records, constants.DetermineError(err);
  }
  defer cursor.Close(ctx);


  for cursor.Next(ctx) { // Iterate cursor
    var d models.Slot;

    err = cursor.Decode(&d); // Decode cursor data into model
    if err != nil {
      fmt.Println(err.Error());
      return records, constants.ErrDataParse;
    }

    // Add data if movie.
    if (d.Type == constants.SlotTypeMovie && d.MovieId != 0) {
      var movieData models.TMDBmovie;
      movieData, err = GetMovie(fmt.Sprintf("%d", d.MovieId));
      if (err != nil) {
        return records, err; // Error already determined in movie service
      }
      d.Title = movieData.Title;
      d.Synopsis = movieData.Synopsis;
      d.Duration = movieData.Runtime;
    }

    records = append(records, d); // Push data record into array
  }
  if err = cursor.Err(); err != nil {
    fmt.Println(err.Error());
    return records, constants.DetermineError(err);
   }

  return records, nil;
}


// Fetches all slots between given from and to time.
func GetScheduleSlotsByTime(client *mongo.Client, scheduleId string, from int, to int) (records []models.Slot, err error) {
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
  ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second); // New context for find query
  defer cancel();

  cur, err := collection.Find(ctx, query, opts);
  if err != nil {
    fmt.Println(err.Error());
    return records, constants.DetermineError(err);
  }
  defer cur.Close(ctx);

  for cur.Next(ctx) { // Iterate cursor
    var d models.Slot;

    err := cur.Decode(&d); // Decode cursor data into model
    if err != nil {
      fmt.Println(err.Error());
      return records, constants.ErrDataParse;
    }
    records = append(records, d); // Push data into array
  }
  if err := cur.Err(); err != nil {
    fmt.Println(err.Error());
    return records, constants.DetermineError(err);
  }

  return records, nil;
}
