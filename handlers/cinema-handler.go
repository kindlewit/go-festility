package handlers

import (
  "fmt"
  "net/http"
  "math/rand"

  "github.com/gin-gonic/gin"
  "festility/models"
  "festility/services"
  "festility/constants"
)

// Handles request to create a cinema.
func CreateCinemaHandler(c *gin.Context) {
  var body models.Cinema;
  var success bool;
  var err error;

  if err = c.BindJSON(&body); err != nil {
    c.String(http.StatusBadRequest, "Required parameters missing: name/city.");
    return;
  }

  // Add Cinema ID
  body.Id = fmt.Sprintf("%d", rand.Intn(1000)); // Random id b/w 0 - 1000

  client := services.Connect();
  success, err = services.CreateCinema(client, body);
  defer services.Disconnect(client);

  if err != nil {
    constants.HandleError(c, err);
    return;
  }
  if (!success) {
    // Record created but with no InsertedID
    c.JSON(http.StatusConflict, "Record was created but found an inconsistency in record id.");
    return;
  }

  c.JSON(http.StatusCreated, gin.H{ "id": body.Id, "name": body.Name });
}

// Handles request to fetch one cinema.
func GetCinemaHandler(c *gin.Context) {
  var record models.Cinema;
  var err error;

  cinemaID := c.Param("id");
  if (cinemaID == "" || cinemaID == "null") {
    // Missing cinema ID param
    c.JSON(http.StatusBadRequest, "Missing valid cinema id. Please check the parameter.");
    return;
  }

  client := services.Connect();
  record, err = services.GetCinema(client, cinemaID);
  defer services.Disconnect(client);

  if err != nil {
    constants.HandleError(c, err);
    return;
  }

  c.JSON(http.StatusOK, record);
  return;
}