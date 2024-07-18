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

// Creates new cinema record.
func CreateCinema(data models.Cinema) (bool, error) {
	// Ensure no other record has the same ID or Name (duplicate)
	query := bson.M{"$or": []bson.M{{"id": data.Id}, {"name": data.Name}}}
	if !(_isUnique(constants.TableCinema, query)) {
		defer db.Disconnect()
		return false, constants.ErrDuplicateRecord // Record already present
	}

	success, err := db.Insert(constants.TableCinema, data)
	defer db.Disconnect()
	return success.InsertedID != nil, constants.DetermineInternalErrMsg(err)
}

// Fetches the cinema record by id.
func GetCinema(cinemaId string) (doc models.Cinema, err error) {
	query := bson.M{"id": cinemaId}
	data, err := db.Retrieve(constants.TableCinema, query)
	defer db.Disconnect()
	if err != nil {
		return doc, constants.DetermineInternalErrMsg(err)
	}

	err = bson.Unmarshal(data, &doc)
	return doc, constants.DetermineInternalErrMsg(err)
}

// Fetches multiple cinema records.
func GetCinemasInBulk(cinemaIdList []string) (docs []models.Cinema, err error) {
	// TODO: include limit to reduce DB querying
	query := bson.M{"id": bson.M{"$in": cinemaIdList}}
	opts := options.Find() // TODO: optimize search using options
	cursor, err := db.RetrieveMany(constants.TableCinema, query, opts)
	defer db.Disconnect()

	if err != nil {
		return docs, constants.DetermineInternalErrMsg(err)
	}

	docs, err = _getCinemasFromCursor(cursor)
	return docs, constants.DetermineInternalErrMsg(err)
}

// Replaces a cinema record by id.
func ReplaceCinema(cinemaId string, replacement models.Cinema) (success bool, err error) {
	if replacement.Id != cinemaId {
		// Trying to update screen ID
		return false, constants.ErrCriticalVal
	}
	query := bson.M{"id": cinemaId}
	result, err := db.Replace(constants.TableCinema, query, replacement)
	defer db.Disconnect()

	if err != nil {
		return false, constants.DetermineInternalErrMsg(err)
	}

	return result.ModifiedCount > 0, constants.DetermineInternalErrMsg(err)
}

// Deletes the cinema record by id.
func DeleteCinema(cinemaId string) (bool, error) {
	query := bson.M{"id": cinemaId}
	success, err := db.DeleteOne(constants.TableCinema, query)
	defer db.Disconnect()
	return success.DeletedCount > 0, constants.DetermineInternalErrMsg(err)
}

func _getCinemasFromCursor(cursor *mongo.Cursor) (docs []models.Cinema, err error) {
	// Create temp context for cursor close
	ctx, cancel := context.WithTimeout(context.Background(), constants.CursorTimeout)
	defer cancel()
	defer cursor.Close(ctx)

	for cursor.Next(ctx) { // Iterate cursor
		var d models.Cinema

		err = cursor.Decode(&d) // Decode cursor data into model
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
