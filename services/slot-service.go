package services

import (
  "fmt"
  "time"
  "context"

  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo/options"
  "github.com/kindlewit/go-festility/utils"
  "github.com/kindlewit/go-festility/models"
  "github.com/kindlewit/go-festility/constants"
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
func GetScheduleSlots(client *mongo.Client, scheduleID string, optionals ...int64) (records []models.Slot, err error) {
  var limit int64 = int64(constants.SlotPageLimit);
  var skip int64 = int64(0);

  if (len(optionals) == 1) {
    limit = optionals[0];
  } else if (len(optionals) == 2) {
    limit = optionals[0];
    skip = optionals[1];
  }

  collection := client.Database("festility").Collection("slot"); // Collection to use
  query := bson.M{ "schedule_id": scheduleID };
  opts := options.Find().SetLimit(limit).SetSkip(skip);

  // New context for find query
  ctx, cancel := context.WithTimeout(context.Background(), constants.QueryTimeout);
  defer cancel();

  cursor, err := collection.Find(ctx, query, opts);
  if err != nil {
    fmt.Println(err.Error());
    return records, constants.DetermineInternalErrMsg(err);
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

      movieData, err := GetMovie(fmt.Sprintf("%d", d.MovieId));
      if (err != nil) {
        return records, err; // Error already determined in movie service
      }
      d = utils.BindMovieToSlot(d, movieData);

    }

    records = append(records, d); // Push data record into array
  }
  if err = cursor.Err(); err != nil {
    fmt.Println(err.Error());
    return records, constants.DetermineInternalErrMsg(err);
   }

  return records, nil;
}

// Fetches screen ID list of slots for a given schedule id.
func GetSlotScreensOfSchedule(client *mongo.Client, scheduleID string) (data []string, err error) {
  collection := client.Database("festility").Collection("slot"); // Collection to use
  query := bson.M{ "schedule_id": scheduleID };
  opts := options.Find().SetProjection(bson.M{
    "duration": 0,
    "start":0,
    "title": 0,
    "synopsis": 0,
    "directors": 0,
    "original_title": 0,
    "year": 0,
    "genres": 0,
    "languages": 0,
    "countries": 0,
  });

  // New context for find query
  ctx, cancel := context.WithTimeout(context.Background(), constants.QueryTimeout);
  defer cancel();

  cur, err := collection.Find(ctx, query, opts);
  if err != nil {
    fmt.Println(err.Error());
    return data, constants.DetermineInternalErrMsg(err);
  }
  defer cur.Close(ctx);

  for cur.Next(ctx) { // Iterate cursor
    var d models.Slot;

    err := cur.Decode(&d); // Decode cursor data into model
    if err != nil {
      fmt.Println(err.Error());
      return data, constants.ErrDataParse;
    }

    data = append(data, d.ScreenID);
  }
  if err := cur.Err(); err != nil {
    fmt.Println(err.Error());
    return data, constants.DetermineInternalErrMsg(err);
  }

  return data, nil;
}

// Fetches all slots between given from and to time.
func GetScheduleSlotsByTime(client *mongo.Client, scheduleID string, from int, to int) (records []models.Slot, err error) {
  collection := client.Database("festility").Collection("slot"); // Collection to use
  query := bson.M{
    "schedule_id": scheduleID,
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

  // New context for find query
  ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second);
  defer cancel();

  cur, err := collection.Find(ctx, query, opts);
  if err != nil {
    fmt.Println(err.Error());
    return records, constants.DetermineInternalErrMsg(err);
  }
  defer cur.Close(ctx);

  // A hashmap to eliminate fetching same movie details multiple times
  movieHashMap := make(map[int]models.TMDBmovie);

  for cur.Next(ctx) { // Iterate cursor
    var d models.Slot;

    err := cur.Decode(&d); // Decode cursor data into model
    if err != nil {
      fmt.Println(err.Error());
      return records, constants.ErrDataParse;
    }

    // Add data if movie.
    if (d.Type == constants.SlotTypeMovie && d.MovieId != 0) {

      if movieData, isPresent := movieHashMap[d.MovieId]; isPresent {

        // The movie details is already present in the hashmap,
        // so we use the same data, instead of calling the external API
        d = utils.BindMovieToSlot(d, movieData);

      } else {

        // The data is not present in the hashmap, so we call the external API
        movieData, err := GetMovie(fmt.Sprintf("%d", d.MovieId)); // Service calls external API
        if (err != nil) {
          return records, err; // Error already determined in movie service
        }

        d = utils.BindMovieToSlot(d, movieData);
        movieHashMap[d.MovieId] = movieData;

      }

    }

    records = append(records, d); // Push data into array
  }
  if err := cur.Err(); err != nil {
    fmt.Println(err.Error());
    return records, constants.DetermineInternalErrMsg(err);
  }

  fmt.Println(len(movieHashMap));
  for k := range movieHashMap {
    delete(movieHashMap, k); // Force clear the hashmap
  }

  return records, nil;
}
