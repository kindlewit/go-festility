package services

import (
	"fmt"

	"github.com/kindlewit/go-festility/constants"
	"github.com/kindlewit/go-festility/models"
	"go.mongodb.org/mongo-driver/bson"
)

// Creates new schedule.
func CreateSchedule(data models.Schedule) (bool, error) {
	// data.Id is already checked to be unique in handler
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
func GetDefaultScheduleId(festId string) (string, error) {
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

func IsUniqueScheduleId(id string) bool {
	query := bson.M{"id": id}

	count, err := db.Count(constants.TableSchedule, query)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return count < 1
}
