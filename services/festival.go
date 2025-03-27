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

// Creates new festival record & returns inserted ID.
func CreateFestival(data models.Fest) (bool, error) {
	// Ensure no other record has the same ID (duplicate)
	query := bson.M{"id": data.Id}
	count, err := db.Count(constants.TableFestival, query)
	if err != nil {
		defer db.Disconnect()
		return false, constants.DetermineInternalErrMsg(err)
	}

	if count > 0 {
		defer db.Disconnect()
		return false, constants.ErrDuplicateRecord // Record already present
	}

	result, err := db.Insert(constants.TableFestival, data)
	defer db.Disconnect()
	return result.InsertedID != nil, constants.DetermineInternalErrMsg(err)
}

// Fetches one festival record by fest id.
func GetFestival(festId string) (doc models.Fest, err error) {
	query := bson.M{"id": festId}
	data, err := db.Retrieve(constants.TableFestival, query)
	defer db.Disconnect()
	if err != nil {
		return doc, constants.DetermineInternalErrMsg(err)
	}

	err = bson.Unmarshal(data, &doc)
	return doc, constants.DetermineInternalErrMsg(err)
}

// Fetches all festival records.
func GetBulkFestivals(optionals ...int64) (docs []models.Fest, err error) {
	var limit int64 = int64(constants.SlotPageLimit)
	var skip int64 = int64(0)

	if len(optionals) == 1 {
		limit = optionals[0]
	} else if len(optionals) == 2 {
		limit = optionals[0]
		skip = optionals[1]
	}

	query := bson.M{}
	opts := options.Find().SetLimit(limit).SetSkip(skip)

	cursor, err := db.RetrieveMany(constants.TableFestival, query, opts)
	defer db.Disconnect()

	if err != nil {
		return docs, constants.DetermineInternalErrMsg(err)
	}

	docs, err = _getFestivalsFromCursor(cursor)
	return docs, constants.DetermineInternalErrMsg(err)
}

func _getFestivalsFromCursor(cursor *mongo.Cursor) (docs []models.Fest, err error) {
	// Create temp context for cursor
	ctx, cancel := context.WithTimeout(context.Background(), constants.CursorTimeout)
	defer cancel()
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var d models.Fest

		err = cursor.Decode(&d)
		if err != nil {
			fmt.Println(err.Error())
			return docs, constants.ErrDataParse
		}

		docs = append(docs, d)
	}
	if err = cursor.Err(); err != nil {
		fmt.Println(err.Error())
	}

	return docs, err
}
