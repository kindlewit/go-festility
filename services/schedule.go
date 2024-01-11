package services

import (
	"fmt"

	"github.com/kindlewit/go-festility/constants"
	"github.com/kindlewit/go-festility/models"
	"go.mongodb.org/mongo-driver/bson"
)

// Creates new schedule.
func CreateSchedule(data models.Schedule) (bool, error) {
	// Ensure no other record has the same ID (duplicate)
	query := bson.M{"id": data.Id}
	if !_isUnique(constants.TableSchedule, query) {
		defer db.Disconnect()
		return false, constants.ErrDuplicateRecord // Record already present
	}

	success, err := db.Insert(constants.TableSchedule, data)
	defer db.Disconnect()
	return success.InsertedID != nil, constants.DetermineInternalErrMsg(err)
}

// Fetches the schedule by fest id & schedule id.
func GetSchedule(festId string, scheduleId string) (doc models.Schedule, err error) {
	query := bson.M{"id": scheduleId, "fest_id": festId}
	data, err := db.Retrieve(constants.TableSchedule, query)
	defer db.Disconnect()
	if err != nil {
		return doc, constants.DetermineInternalErrMsg(err)
	}

	err = bson.Unmarshal(data, &doc)
	return doc, constants.DetermineInternalErrMsg(err)
}

// Fetches the default schedule ID for a festival.
func GetDefaultScheduleID(festId string) (string, error) {
	var doc models.Schedule

	query := bson.M{"fest_id": festId, "custom": false} // Default schedule will have Custom=false
	data, err := db.Retrieve(constants.TableSchedule, query)
	defer db.Disconnect()
	if err != nil {
		return "", constants.DetermineInternalErrMsg(err)
	}

	err = bson.Unmarshal(data, &doc)
	return doc.Id, constants.DetermineInternalErrMsg(err)
}

// Checks if a record already exists in db.
func _isUnique(tableName string, query bson.M) bool {
	count, err := db.Count(tableName, query)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return count < 1 // No docs matching the given query
}
