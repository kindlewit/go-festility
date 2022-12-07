package handlers

import (
  "net/http"

  "github.com/gin-gonic/gin"
	"festility/services"
	"festility/constants"
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
