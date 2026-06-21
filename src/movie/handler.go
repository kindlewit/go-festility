package movie

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kindlewit/go-festility/src/constants"
)

// Handles requests to get one movie.
func HandleGetMovie(c *gin.Context) {
	id := c.Param("id")

	movieRecord, err := GetMovie(id)
	if err != nil {
		constants.HandleError(c, err)
		return
	}
	movieDoc := convertTMDBtoMovie(movieRecord)

	dirList, err := GetDirector(id)
	if err != nil {
		constants.HandleError(c, err)
	}
	movieDoc.Directors = dirList

	c.JSON(http.StatusOK, movieDoc)
}
