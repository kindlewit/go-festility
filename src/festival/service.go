package festival

import (
	"github.com/kindlewit/go-festility/src/constants"
	database "github.com/kindlewit/go-festility/src/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db = database.Database{}

// Creates new festival record & returns inserted ID.
func CreateFestival(data Fest) (bool, error) {
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
func GetFestival(festId string) (doc Fest, err error) {
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
func GetBulkFestivals(optionals ...int64) (docs []Fest, err error) {
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
