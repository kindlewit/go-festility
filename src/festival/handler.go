package festival

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kindlewit/go-festility/src/constants"
)

// Handles request to create one festival.
func HandleCreateFest(c *gin.Context) {
	var body Fest
	var success bool
	var err error

	if err = c.BindJSON(&body); err != nil {
		c.String(http.StatusBadRequest, constants.MsgMissingFestParams)
		return
	}

	success, err = CreateFestival(body)
	if err != nil {
		constants.HandleError(c, err)
		return
	}
	if !success {
		c.String(http.StatusConflict, constants.MsgInconsistentId)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": body.Id})
}

// Handles request to get details of one festival.
func HandleGetFest(c *gin.Context) {
	festId := c.Param("id")

	resp, err := GetFestival(festId)
	if err != nil {
		constants.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
	return
}

func HandleGetBulkFestivals(c *gin.Context) {
	var err error
	var resp []Fest

	resp, err = GetBulkFestivals()
	if err != nil {
		constants.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
	return
}
