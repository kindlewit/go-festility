package services

import (
	"github.com/kindlewit/go-festility/constants"
	"github.com/kindlewit/go-festility/models"
	"go.mongodb.org/mongo-driver/bson"
)

// Creates new cinema record.
func CreateCinema(data models.Cinema) (bool, error) {
	// Ensure no other record has the same ID or Name (duplicate)
	query := bson.M{"$or": []bson.M{{"id": data.Id}, {"name": data.Name}}}
	if !(_isUnique(constants.TableCinema, query)) {
		defer db.Disconnect()
		return false, constants.ErrDuplicateRecord // Record already present
	}

	success, err := db.Insert(constants.TableSchedule, data)
	defer db.Disconnect()
	return success.InsertedID != nil, constants.DetermineInternalErrMsg(err)
}

// Fetches the cinema record by id.
func GetCinema(cinemaID string) (doc models.Cinema, err error) {
	query := bson.M{"id": cinemaID}
	data, err := db.Retrieve(constants.TableCinema, query)
	defer db.Disconnect()
	if err != nil {
		return doc, constants.DetermineInternalErrMsg(err)
	}

	err = bson.Unmarshal(data, &doc)
	return doc, constants.DetermineInternalErrMsg(err)
}

// // Fetches multiple cinema records.
// func GetCinemasInBulk(client *mongo.Client, cinemaIDlist []string) (data []models.Screen, err error) {
// 	collection := client.Database("festility").Collection("cinema") // Collection to use
// 	query := bson.M{"id": bson.M{"$in": cinemaIDlist}}

// 	// Create context
// 	ctx, cancel := context.WithTimeout(context.Background(), constants.QueryTimeout)
// 	defer cancel()

// 	cursor, err := collection.Find(ctx, query)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return data, constants.DetermineInternalErrMsg(err)
// 	}
// 	defer cursor.Close(ctx)

// 	for cursor.Next(ctx) { // Iterate cursor
// 		var d models.Screen

// 		err = cursor.Decode(&d) // Decode cursor data into model
// 		if err != nil {
// 			fmt.Println(err.Error())
// 			return data, constants.ErrDataParse
// 		}

// 		data = append(data, d) // Push data record into array
// 	}
// 	if err = cursor.Err(); err != nil {
// 		fmt.Println(err.Error())
// 		return data, constants.DetermineInternalErrMsg(err)
// 	}

// 	return data, nil
// }

// // Updates a cinema record by id.
// func ReplaceCinema(client *mongo.Client, cinemaID string, replacement models.Cinema) (success bool, err error) {
// 	collection := client.Database("festility").Collection("cinema") // Collection to use
// 	query := bson.M{"id": cinemaID}

// 	// Create context
// 	ctx, cancel := context.WithTimeout(context.Background(), constants.QueryTimeout)
// 	defer cancel()

// 	_, err = collection.ReplaceOne(ctx, query, replacement)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return false, constants.DetermineInternalErrMsg(err)
// 	}
// 	return true, nil
// }

// // Deletes the cinema record by id.
// func DeleteCinema(client *mongo.Client, cinemaID string) (success bool, err error) {
// 	collection := client.Database("festility").Collection("cinema") // Collection to use
// 	query := bson.M{"id": cinemaID}

// 	// Create context
// 	ctx, cancel := context.WithTimeout(context.Background(), constants.QueryTimeout)
// 	defer cancel()

// 	_, err = collection.DeleteOne(ctx, query)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return false, constants.DetermineInternalErrMsg(err)
// 	}
// 	return true, nil
// }
