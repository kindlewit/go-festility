package constants

import (
  "errors"
  "strings"
  "net/http"

  "github.com/gin-gonic/gin"
)

var (
  ErrDuplicateRecord error = errors.New("ErrDuplicateRecord");
  ErrNoSuchRecord error = errors.New("ErrNoSuchRecord");
  ErrConnection error = errors.New("ErrConnection");
  ErrMongo error = errors.New("ErrMongo");
  ErrDataParse error = errors.New("ErrDataParse");
  ErrEmptyData error = errors.New("ErrEmptyData");
  ErrNoDefaultSchedule error = errors.New("ErrNoDefaultSchedule");
  ErrApiFetch error = errors.New("ErrApiFetch");
  ErrApiParse error = errors.New("ErrApiParse");
)

// Determines which custom error to throw based on error received.
func DetermineError(err error) (error) {
  if (strings.Contains(err.Error(), "server selection error: context deadline exceeded")) {
    // Caught DB connection error
    return ErrConnection;
  }
  if (err.Error() == "mongo: no documents in result") {
    // Caught a "mongo: no documents in result" error
    return ErrNoSuchRecord;
  }

  return ErrMongo;
}

// Handles error by returning appropriate HTTP response code.
func HandleError(c *gin.Context, err error) {
  switch err {
  case ErrNoSuchRecord: {
    c.JSON(http.StatusNotFound, "No such record exists. Please check the record id.");
  }
  case ErrDuplicateRecord: {
    c.JSON(http.StatusConflict, "Request to create duplicate record. Please check the record id.");
  }
  case ErrConnection: {
    c.JSON(http.StatusInternalServerError, "Failed to connect to database. Please try again later.");
  }
  case ErrMongo: {
    c.JSON(http.StatusInternalServerError, "Faced a database error. Please try again later.");
  }
  case ErrDataParse: {
    c.JSON(http.StatusInternalServerError, "Faced an error while parsing internal data. Please try again.");
  }
  case ErrEmptyData: {
    c.JSON(http.StatusNotFound, "Data being requested does not exist.");
  }
  case ErrNoDefaultSchedule: {
    c.JSON(http.StatusNotFound, "No schedule available for this fest yet.")
  }
  default: {
    c.JSON(http.StatusInternalServerError, "Unable to process request. Please try again.");
  }
  }
}
