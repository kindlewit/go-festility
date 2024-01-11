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

// Creates multiple new screen records.
func CreateCinemaScreens(screens []models.Screen) (success bool, err error) {
	// TODO: data should be sanitized
	data := make([]interface{}, len(screens))
	for i, s := range screens {
		data[i] = s
	}

	result, err := db.InsertMany(constants.TableScreen, data)
	defer db.Disconnect()
	return result.InsertedIDs != nil, constants.DetermineInternalErrMsg(err)
}

// Fetches one screen record by screen id.
func GetScreen(screenID string) (doc models.Screen, err error) {
	query := bson.M{"id": screenID}
	data, err := db.Retrieve(constants.TableScreen, query)
	defer db.Disconnect()
	if err != nil {
		return doc, constants.DetermineInternalErrMsg(err)
	}

	err = bson.Unmarshal(data, &doc)
	return doc, constants.DetermineInternalErrMsg(err)
}

// Fetches multiple screen records.
func GetScreensInBulk(screenList []string) (docs []models.Screen, err error) {
	query := bson.M{"id": bson.M{"$in": screenList}}
	opts := options.Find() // TODO: optimize search using options
	cursor, err := db.RetrieveMany(constants.TableScreen, query, opts)
	defer db.Disconnect()

	if err != nil {
		return docs, constants.DetermineInternalErrMsg(err)
	}

	docs, err = _getScreensFromCursor(cursor)
	return docs, constants.DetermineInternalErrMsg(err)
}

// Fetches screens for a cinema by cinema id.
func GetCinemaScreens(cinemaID string) (docs []models.Screen, err error) {
	query := bson.M{"cinema_id": cinemaID}
	opts := options.Find() // TODO: optimize search using options
	cursor, err := db.RetrieveMany(constants.TableScreen, query, opts)
	defer db.Disconnect()

	if err != nil {
		return docs, constants.DetermineInternalErrMsg(err)
	}

	docs, err = _getScreensFromCursor(cursor)
	return docs, constants.DetermineInternalErrMsg(err)
}

// Updates a screen record by id.
func ReplaceScreen(screenID string, replacement models.Screen) (success bool, err error) {
	query := bson.M{"id": screenID}
	result, err := db.Replace(constants.TableScreen, query, replacement)
	defer db.Disconnect()

	if err != nil {
		return false, constants.DetermineInternalErrMsg(err)
	}

	return result.ModifiedCount > 0, constants.DetermineInternalErrMsg(err)
}

func _getScreensFromCursor(cursor *mongo.Cursor) (docs []models.Screen, err error) {
	// Create temp context for cursor close
	ctx, cancel := context.WithTimeout(context.Background(), constants.CursorTimeout)
	defer cancel()
	defer cursor.Close(ctx)

	for cursor.Next(ctx) { // Iterate cursor
		var d models.Screen

		err = cursor.Decode(&d) // Decode cursor data into screen model
		if err != nil {
			fmt.Println(err.Error())
			return docs, constants.ErrDataParse
		}
		docs = append(docs, d) // Push data record into array
	}
	if err = cursor.Err(); err != nil {
		fmt.Println(err.Error())
	}
	return docs, err
}
