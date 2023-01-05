package constants

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	MsgDuplicateRecord   string = "Request to create duplicate record. Please check the record id."
	MsgNoSuchRecord      string = "No such record exists. Please check the record id."
	MsgConnectionFailure string = "Failed to connect to database. Please try again later."
	MsgDatabaseFailure   string = "Faced a database error. Please try again later."
	MsgDataParse         string = "Faced an error while parsing internal data. Please try again."
	MsgEmptyData         string = "Data being requested does not exist."
	MsgNoSchedule        string = "No schedule available for this fest yet."
	MsgMissingFestParams string = "Required parameters missing: id/name/from_date/to_date."
	MsgInconsistentId    string = "Record was created but found an inconsistency in record id."
)

var (
	ErrDuplicateRecord   error = errors.New(MsgDuplicateRecord)
	ErrNoSuchRecord      error = errors.New(MsgNoSuchRecord)
	ErrConnection        error = errors.New(MsgConnectionFailure)
	ErrMongo             error = errors.New(MsgDatabaseFailure)
	ErrDataParse         error = errors.New(MsgDataParse)
	ErrEmptyData         error = errors.New(MsgEmptyData)
	ErrNoDefaultSchedule error = errors.New(MsgNoSchedule)
	ErrApiFetch          error = errors.New("ErrApiFetch")
	ErrApiParse          error = errors.New("ErrApiParse")
)

// Determines which custom error to throw based on error received.
func DetermineInternalErrMsg(err error) error {
	if strings.Contains(err.Error(), "server selection error: context deadline exceeded") {
		// Caught DB connection error
		return ErrConnection
	}
	if err.Error() == "mongo: no documents in result" {
		// Caught a "mongo: no documents in result" error
		return ErrNoSuchRecord
	}

	return ErrMongo
}

// Handles error by returning appropriate HTTP response code.
func HandleError(c *gin.Context, err error) {
	switch err {
	case ErrNoSuchRecord:
		{
			c.String(http.StatusNotFound, MsgNoSuchRecord)
		}
	case ErrDuplicateRecord:
		{
			c.String(http.StatusConflict, MsgDuplicateRecord)
		}
	case ErrConnection:
		{
			c.String(http.StatusInternalServerError, MsgConnectionFailure)
		}
	case ErrMongo:
		{
			c.String(http.StatusInternalServerError, MsgDatabaseFailure)
		}
	case ErrDataParse:
		{
			c.String(http.StatusInternalServerError, MsgDataParse)
		}
	case ErrEmptyData:
		{
			c.String(http.StatusNotFound, MsgEmptyData)
		}
	case ErrNoDefaultSchedule:
		{
			c.String(http.StatusNotFound, MsgNoSchedule)
		}
	default:
		{
			c.String(http.StatusInternalServerError, "Unable to process request. Please try again.")
		}
	}
}
