package schedule

import (
	"context"
	"fmt"

	"github.com/kindlewit/go-festility/src/constants"
	"go.mongodb.org/mongo-driver/mongo"
)

func _getSlotsFromCursor(cursor *mongo.Cursor) (docs []Slot, err error) {
	// Create temp context for cursor close
	ctx, cancel := context.WithTimeout(context.Background(), constants.CursorTimeout)
	defer cancel()
	defer cursor.Close(ctx)

	for cursor.Next(ctx) { // Iterate cursor
		var d Slot

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

func _getScreenIdFromCursor(cursor *mongo.Cursor) (docs []string, err error) {
	// Create temp context for cursor close
	ctx, cancel := context.WithTimeout(context.Background(), constants.CursorTimeout)
	defer cancel()
	defer cursor.Close(ctx)

	for cursor.Next(ctx) { // Iterate cursor
		var d Slot

		err = cursor.Decode(&d) // Decode cursor data into model
		if err != nil {
			fmt.Println(err.Error())
			return docs, constants.ErrDataParse
		}
		if d.ScreenId != "" {
			docs = append(docs, d.ScreenId)
		} // Do not break if slots are missing screen ID (might interview/other)
	}
	if err = cursor.Err(); err != nil {
		fmt.Println(err.Error())
	}
	return docs, err
}
