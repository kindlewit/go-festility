package handlers

import (
  "time"
  "net/http"

  "github.com/gin-gonic/gin"
  "festility/services"
  "festility/constants"
)

// Handles requests to index (/) endpoint.
// The purpose of this endpoint is to ensure if the server is working,
// and if the necessary services are up & running.
func IndexHandler(c *gin.Context) {
  // Check connection to database
  client := services.Connect();
  defer services.Disconnect(client);

  var resp gin.H;

  _, err := services.GetFestival(client, "123");
  if (err != nil && err.Error() == constants.ErrConnection.Error()) {
    // Failed to connect to DB
    resp = gin.H {
      "timestamp": time.Now().Unix(),
      "success": false,
      "message": err.Error(),
    };
  } else {
    resp = gin.H{
      "timestamp": time.Now().Unix(),
      "success": true,
      "message": "Welcome to festility!",
    }
  }

  c.JSON(http.StatusOK, resp);
}
