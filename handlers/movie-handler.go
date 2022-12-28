package handlers

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/kindlewit/go-festility/services"
  "github.com/kindlewit/go-festility/constants"
)

// Handles requests to get one movie.
func GetMovieHandler(c *gin.Context) {
  id := c.Param("id");

  movieRecord, err := services.GetMovie(id);
  if (err != nil) {
    constants.HandleError(c, err);
    return;
  }
  // directors := getDirector(id);

  // resp := reformat(movieDoc, directors);

  c.JSON(http.StatusOK, movieRecord);
}
