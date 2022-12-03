package handlers

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "festility/models"
  "festility/services"
  "festility/constants"
)

// Handles request to create one festival.
func CreateFestivalHandler(c *gin.Context) {
  var body models.Fest;
  var success bool;
  var err error;

  if err = c.BindJSON(&body); err != nil {
    c.String(http.StatusBadRequest, "Required parameters missing: id/name/from_date/to_date.");
    return;
  }

  client := services.Connect();
  success, err = services.CreateFestival(client, body);
  defer services.Disconnect(client);

  if err != nil {
    constants.HandleError(c, err);
    return;
  }
  if (!success) {
    c.String(http.StatusConflict, "Record was created but found an inconsistency in record id.");
    return;
  }

  c.JSON(http.StatusCreated, gin.H{ "id": body.Id });
}

// Handles request to get details of one festival.
func GetFestHandler(c *gin.Context) {
  festId := c.Param("id");

  client := services.Connect();
  resp, err := services.GetFestival(client, festId);
  defer services.Disconnect(client);

  if err != nil {
    constants.HandleError(c, err);
    return;
  }

  c.JSON(http.StatusOK, resp);
  return;
}
