package services

import (
	"github.com/kindlewit/go-festility/constants"
	"github.com/kindlewit/go-festility/models"
	"go.mongodb.org/mongo-driver/bson"
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
