package constants

import (
  "errors"
  "net/http"

  "github.com/gin-gonic/gin"
)

var DuplicateRecordError error = errors.New("DuplicateRecordError");
var NoSuchRecordError error = errors.New("NoSuchRecordError");
var MongoReadError error = errors.New("MongoReadError");
var MongoWriteError error = errors.New("MongoWriteError");
var DataParsingError error = errors.New("DataParsingError");
var EmptyDataError error = errors.New("EmptyDataError");
var NonExistantDefaultSchedule error = errors.New("NonExistantDefaultSchedule");
var ApiFetchError error = errors.New("ApiFetchError");
var ApiParsingError error = errors.New("ApiParsingError");

func HandleError(c *gin.Context, err error) {
  switch err {
  case NoSuchRecordError: {
    c.JSON(http.StatusNotFound, "No such record exists. Please check the record id.");
  }
  case DuplicateRecordError: {
    c.JSON(http.StatusConflict, "Request to create duplicate record. Please check the record id.");
  }
  case MongoReadError: {
    c.JSON(http.StatusInternalServerError, "Faced a database error. Please try again later.");
  }
  case MongoWriteError: {
    c.JSON(http.StatusInternalServerError, "Faced a database error. Please try again later.");
  }
  case DataParsingError: {
    c.JSON(http.StatusInternalServerError, "Faced an error while parsing internal data. Please try again.");
  }
  case EmptyDataError: {
    c.JSON(http.StatusNotFound, "Data being requested does not exist.");
  }
  case NonExistantDefaultSchedule: {
    c.JSON(http.StatusNotFound, "No schedule available for this fest yet.")
  }
  default: {
    c.JSON(http.StatusInternalServerError, "Unable to process request. Please try again.");
  }
  }
}
