package constants

import (
	"errors"
	"fmt"
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
	MsgApiFetch          string = "Error fetching data from external service."
	MsgMissingFestParams string = "Required parameters missing: id/name/from_date/to_date."
	MsgInconsistentId    string = "Record was created but found an inconsistency in record id."
	MsgConversion        string = "Faced an error while converting internal data. Please try again."
	MsgUnauthorized      string = "Unauthorized to commit this action."
	MsgMissingApiKey     string = "Key for external API is missing."
	MsgCriticalVal       string = "Attempt to update a critical value in the DB."
)

var (
	ErrDuplicateRecord   error = errors.New(MsgDuplicateRecord)
	ErrNoSuchRecord      error = errors.New(MsgNoSuchRecord)
	ErrConnection        error = errors.New(MsgConnectionFailure)
	ErrMongo             error = errors.New(MsgDatabaseFailure)
	ErrDataParse         error = errors.New(MsgDataParse)
	ErrEmptyData         error = errors.New(MsgEmptyData)
	ErrNoDefaultSchedule error = errors.New(MsgNoSchedule)
	ErrApiFetch          error = errors.New(MsgApiFetch)
	ErrApiParse          error = errors.New("ErrApiParse")
	ErrConversion        error = errors.New(MsgConversion)
	ErrUnauthorized      error = errors.New(MsgUnauthorized)
	ErrMissingApiKey     error = errors.New(MsgMissingApiKey)
	ErrCriticalVal       error = errors.New(MsgCriticalVal)
)

// Determines which custom error to throw based on error received.
func DetermineInternalErrMsg(err error) error {
	if err == nil {
		return nil
	}

	fmt.Println(err.Error())
	if strings.Contains(err.Error(), "server selection error: context deadline exceeded") {
		// Caught DB connection error
		return ErrConnection
	}
	if strings.Contains(err.Error(), "conversion") || strings.Contains(err.Error(), "InvalidUnmarshalError") {
		return ErrConversion
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
	case ErrConversion:
		{
			c.String(http.StatusInternalServerError, MsgConversion)
		}
	case ErrUnauthorized:
		{
			c.String(http.StatusUnauthorized, MsgUnauthorized)
		}
	case ErrMissingApiKey:
		{
			c.String(http.StatusConflict, MsgMissingApiKey)
		}
	case ErrCriticalVal:
		{
			c.String(http.StatusBadRequest, MsgCriticalVal)
		}
	default:
		{
			c.String(http.StatusInternalServerError, "Unable to process request. Please try again.")
		}
	}
}
