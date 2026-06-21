package cinema

import (
	"github.com/kindlewit/go-festility/src/constants"
	database "github.com/kindlewit/go-festility/src/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db = database.Database{}

// Creates new cinema record.
func CreateCinema(data Cinema) (bool, error) {
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
func GetCinema(cinemaId string) (doc Cinema, err error) {
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
func GetCinemasInBulk(cinemaIdList []string) (docs []Cinema, err error) {
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
func ReplaceCinema(cinemaId string, replacement Cinema) (success bool, err error) {
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
