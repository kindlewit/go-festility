package schedule

import (
	"fmt"

	"github.com/kindlewit/go-festility/src/constants"
	database "github.com/kindlewit/go-festility/src/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db = database.Database{}

// Creates new schedule.
func CreateSchedule(data Schedule) (bool, error) {
	// data.Id is already checked to be unique in handler
	success, err := db.Insert(constants.TableSchedule, data)
	defer db.Disconnect()
	return success.InsertedID != nil, constants.DetermineInternalErrMsg(err)
}

// Fetches the schedule by fest id & schedule id.
func GetSchedule(festId string, scheduleId string) (doc Schedule, err error) {
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
	var doc Schedule

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

// Creates new slot records & returns success.
func CreateSlots(slots []Slot) (bool, error) {
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
func GetScheduleSlots(scheduleId string, optionals ...int64) (docs []Slot, err error) {
	var limit int64 = int64(constants.SlotPageLimit)
	var skip int64 = int64(0)

	if len(optionals) == 1 {
		limit = optionals[0]
	} else if len(optionals) == 2 {
		limit = optionals[0]
		skip = optionals[1]
	}

	query := bson.M{"schedule_id": scheduleId}
	opts := options.Find().SetLimit(limit).SetSkip(skip)

	cursor, err := db.RetrieveMany(constants.TableScreen, query, opts)
	defer db.Disconnect()

	if err != nil {
		return docs, constants.DetermineInternalErrMsg(err)
	}

	docs, err = _getSlotsFromCursor(cursor)
	return docs, constants.DetermineInternalErrMsg(err)
}

// Fetches screen ID list of slots for a given schedule id.
func GetScheduleScreenList(scheduleId string, optionals ...int64) (docs []string, err error) {
	query := bson.M{"schedule_id": scheduleId}
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

	docs, err = _getScreenIdFromCursor(cursor)
	return docs, constants.DetermineInternalErrMsg(err)
}

// Fetches all slots between given from and to time.
func GetScheduleSlotsByTime(scheduleId string, from int, to int) (docs []Slot, err error) {
	query := bson.M{
		"schedule_id": scheduleId,
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
