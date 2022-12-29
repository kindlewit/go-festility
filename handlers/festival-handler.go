package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kindlewit/go-festility/constants"
	"github.com/kindlewit/go-festility/models"
	"github.com/kindlewit/go-festility/services"
	"github.com/kindlewit/go-festility/utils"
)

// Handles request to create one festival.
func CreateFestivalHandler(c *gin.Context) {
	var body models.Fest
	var success bool
	var err error

	if err = c.BindJSON(&body); err != nil {
		c.String(http.StatusBadRequest, constants.MsgMissingFestParams)
		return
	}

	client := services.Connect()
	success, err = services.CreateFestival(client, body)
	defer services.Disconnect(client)

	if err != nil {
		constants.HandleError(c, err)
		return
	}
	if !success {
		c.String(http.StatusConflict, MsgInconsistentId)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": body.Id})
}

// Handles request to get details of one festival.
func GetFestHandler(c *gin.Context) {
	festId := c.Param("id")

	client := services.Connect()
	resp, err := services.GetFestival(client, festId)
	defer services.Disconnect(client)

	if err != nil {
		constants.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
	return
}

// Handles request to get all screens of one festival.
func GetFestScreensHandler(c *gin.Context) {
	var err error
	var resp []models.CinemaScreen

	festID := c.Param("id")

	client := services.Connect()
	scheduleID, err := services.GetDefaultScheduleID(client, festID)
	if err != nil {
		constants.HandleError(c, err)
		return
	}
	screenIDList, err := services.GetSlotScreensOfSchedule(client, scheduleID)
	if err != nil {
		constants.HandleError(c, err)
		return
	}
	screenList, err := services.GetScreensInBulk(client, screenIDList)
	if err != nil {
		constants.HandleError(c, err)
		return
	}

	cinemaHashMap := make(map[string]models.Cinema)

	for i := 0; i < len(screenList); i++ {
		cID := screenList[i].CinemaID
		if cinemaData, isPresent := cinemaHashMap[cID]; isPresent {
			resp = append(resp, utils.BindCinemaToScreen(screenList[i], cinemaData))
		} else {
			cinemaHashMap[cID], err = services.GetCinema(client, cID)
			resp = append(resp, utils.BindCinemaToScreen(screenList[i], cinemaHashMap[cID]))
		}
	}
	// cinemaList, err := services.GetCinemasInBulk(client, cinemaHashMap.Keys());
	defer services.Disconnect(client)

	c.JSON(http.StatusOK, resp)
}
