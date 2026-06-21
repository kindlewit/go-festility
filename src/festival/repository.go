package festival

import (
	"context"
	"fmt"

	"github.com/kindlewit/go-festility/src/constants"
	"go.mongodb.org/mongo-driver/mongo"
)

func _getFestivalsFromCursor(cursor *mongo.Cursor) (docs []Fest, err error) {
	// Create temp context for cursor
	ctx, cancel := context.WithTimeout(context.Background(), constants.CursorTimeout)
	defer cancel()
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var d Fest

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
