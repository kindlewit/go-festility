package services

import (
	"context"
	"fmt"

	"github.com/kindlewit/go-festility/constants"
	"github.com/kindlewit/go-festility/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Creates new slot records & returns success.
func CreateSlots(slots []models.Slot) (bool, error) {
	// TODO: data should be sanitized
	data := make([]interface{}, len(slots))
	for i, s := range slots {
		data[i] = s
	}

	result, err := db.InsertMany(constants.TableSlot, data)
	defer db.Disconnect()
	return result.InsertedIDs != nil, constants.DetermineInternalErrMsg(err)
}

// Fetches all slots for a given schedule id.
// optionals: [limit, skip]
func GetScheduleSlots(scheduleID string, optionals ...int64) (docs []models.Slot, err error) {
	var limit int64 = int64(constants.SlotPageLimit)
	var skip int64 = int64(0)

	if len(optionals) == 1 {
		limit = optionals[0]
	} else if len(optionals) == 2 {
		limit = optionals[0]
		skip = optionals[1]
	}

	query := bson.M{"schedule_id": scheduleID}
	opts := options.Find().SetLimit(limit).SetSkip(skip)

	cursor, err := db.RetrieveMany(constants.TableScreen, query, opts)
	defer db.Disconnect()

	if err != nil {
		return docs, constants.DetermineInternalErrMsg(err)
	}

	docs, err = _getSlotsFromCursor(cursor)
	return docs, constants.DetermineInternalErrMsg(err)
}

// // Fetches screen ID list of slots for a given schedule id.
func GetScheduleScreenList(scheduleID string, optionals ...int64) (docs []string, err error) {
	query := bson.M{"schedule_id": scheduleID}
	// Omission options (keys to include/omit)
	opts := options.Find().SetProjection(bson.M{
		"duration":       0,
		"start":          0,
		"title":          0,
		"synopsis":       0,
		"directors":      0,
		"original_title": 0,
		"year":           0,
		"genres":         0,
		"languages":      0,
		"countries":      0,
	})

	cursor, err := db.RetrieveMany("slot", query, opts)
	defer db.Disconnect()

	if err != nil {
		return docs, constants.DetermineInternalErrMsg(err)
	}

	docs, err = _getScreenIDFromCursor(cursor)
	return docs, constants.DetermineInternalErrMsg(err)
}

// Fetches all slots between given from and to time.
func GetScheduleSlotsByTime(scheduleID string, from int, to int) (docs []models.Slot, err error) {
	query := bson.M{
		"schedule_id": scheduleID,
		"start_time":  bson.M{"$gte": from, "$lt": to}, // Show start between given params
	}
	// Omission options (keys to include/omit)
	opts := options.Find().SetProjection(bson.M{
		"directors":      0,
		"original_title": 0,
		"genres":         0,
		"languages":      0,
		"countries":      0,
	})

	cursor, err := db.RetrieveMany("slot", query, opts)
	defer db.Disconnect()

	if err != nil {
		return docs, constants.DetermineInternalErrMsg(err)
	}

	docs, err = _getSlotsFromCursor(cursor)
	return docs, constants.DetermineInternalErrMsg(err)
}

func _getSlotsFromCursor(cursor *mongo.Cursor) (docs []models.Slot, err error) {
	// Create temp context for cursor close
	ctx, cancel := context.WithTimeout(context.Background(), constants.CursorTimeout)
	defer cancel()
	defer cursor.Close(ctx)

	for cursor.Next(ctx) { // Iterate cursor
		var d models.Slot

		err = cursor.Decode(&d) // Decode cursor data into model
		if err != nil {
			fmt.Println(err.Error())
			return docs, constants.ErrDataParse
		}

		// 	// A hashmap to eliminate fetching same movie details multiple times
		// 	movieHashMap := make(map[int]models.TMDBmovie)

		// TODO: Add data if type = movie.
		// if d.Type == constants.SlotTypeMovie && d.MovieId != 0 {
		// if movieData, isPresent := movieHashMap[d.MovieId]; isPresent {

		// 				// The movie details is already present in the hashmap,
		// 				// so we use the same data, instead of calling the external API
		// 				d = utils.BindMovieToSlot(d, movieData)

		// 			} else {

		// 				// The data is not present in the hashmap, so we call the external API
		// 				movieData, err := GetMovie(fmt.Sprintf("%d", d.MovieId)) // Service calls external API
		// 				if err != nil {
		// 					return records, err // Error already determined in movie service
		// 				}

		// 				d = utils.BindMovieToSlot(d, movieData)
		// 				movieHashMap[d.MovieId] = movieData

		// 			}

		// 	movieData, err := GetMovie(fmt.Sprintf("%d", d.MovieId))
		// 	if err != nil {
		// 		return docs, err // Error already determined in movie service
		// 	}
		// 	d = utils.BindMovieToSlot(d, movieData)

		// 	for k := range movieHashMap {
		// 		delete(movieHashMap, k) // Force clear the hashmap
		// 	}
		// }

		docs = append(docs, d) // Push data record into array
	}
	if err = cursor.Err(); err != nil {
		fmt.Println(err.Error())
	}
	return docs, err
}

func _getScreenIDFromCursor(cursor *mongo.Cursor) (docs []string, err error) {
	// Create temp context for cursor close
	ctx, cancel := context.WithTimeout(context.Background(), constants.CursorTimeout)
	defer cancel()
	defer cursor.Close(ctx)

	for cursor.Next(ctx) { // Iterate cursor
		var d models.Slot

		err = cursor.Decode(&d) // Decode cursor data into model
		if err != nil {
			fmt.Println(err.Error())
			return docs, constants.ErrDataParse
		}
		if d.ScreenID != "" {
			docs = append(docs, d.ScreenID)
		} // Do not break if slots are missing screen ID (might interview/other)
	}
	if err = cursor.Err(); err != nil {
		fmt.Println(err.Error())
	}
	return docs, err
}
