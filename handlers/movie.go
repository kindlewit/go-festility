package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kindlewit/go-festility/constants"
	"github.com/kindlewit/go-festility/services"
	"github.com/kindlewit/go-festility/utils"
)

// Handles requests to get one movie.
func HandleGetMovie(c *gin.Context) {
	id := c.Param("id")

	movieRecord, err := services.GetMovie(id)
	if err != nil {
		constants.HandleError(c, err)
		return
	}
	movieDoc := utils.ConvertTMDBtoMovie(movieRecord)

	dirList, err := services.GetDirector(id)
	if err != nil {
		constants.HandleError(c, err)
	}
	movieDoc.Directors = dirList

	c.JSON(http.StatusOK, movieDoc)
}
