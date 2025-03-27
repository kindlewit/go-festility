package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kindlewit/go-festility/constants"
	"github.com/kindlewit/go-festility/models"
	"github.com/kindlewit/go-festility/services"
	"github.com/kindlewit/go-festility/utils"
)

// Handles request to create a schedule for an existing fest.
func HandleCreateSchedule(c *gin.Context) {
	var err error
	var success bool

	festId := c.Param("id")

	// Get slots from request body
	var body []models.Slot
	if err = c.ShouldBindJSON(&body); err != nil {
		c.String(http.StatusBadRequest, "Request body not matching slot structure.")
		return
	}

	// Generate new schedule ID
	newScheduleId := utils.GenerateRandomHash(constants.ScheduleIdHashLength) // Generate new schedule ID

	// Ensure new ID is not present in db
	for (constants.ScheduleIdHashLength != 0 && len(newScheduleId) == 0) || (services.IsUniqueScheduleId(newScheduleId)) {
		// while(new ID is not unique or is empty)'
		newScheduleId = utils.GenerateRandomHash(constants.ScheduleIdHashLength)
	}

	// Store schedule
	newSchedule := models.Schedule{
		Id:     newScheduleId,
		Fest:   festId,
		Custom: false, // by default
	}
	success, err = services.CreateSchedule(newSchedule)
	// TODO: handle this err
	if !success {
		constants.HandleError(c, err)
		return
	}

	// Store slots into db under generated schedule ID
	for i := 0; i < len(body); i++ {
		// TODO: check slot id
		body[i].ScheduleId = newScheduleId
	}
	services.CreateSlots(body)

	// Return success with newly generated schedule ID
	c.JSON(http.StatusCreated, gin.H{
		"schedule_id":     newScheduleId,
		"number_of_slots": len(body),
	})
}

// Handles request to fetch one schedule by schedule id.
func HandleGetSchedule(c *gin.Context) {
	var err error
	var schedule models.Schedule
	var resp models.ScheduleApiRespWithSlots
	var slots []models.Slot

	festId := c.Param("id")
	scheduleId := c.Param("sid")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1")) // Pagination

	limit := int64(constants.SlotPageLimit)
	skip := int64(page*constants.SlotPageLimit - constants.SlotPageLimit)

	schedule, err = services.GetSchedule(festId, scheduleId)
	if err != nil {
		constants.HandleError(c, err)
		return
	}

	slots, err = services.GetScheduleSlots(scheduleId, limit, skip)
	if err != nil {
		constants.HandleError(c, err)
		return
	}

	resp = models.ScheduleApiRespWithSlots{
		Id:     schedule.Id,
		Fest:   schedule.Fest,
		Custom: schedule.Custom,
		Slots:  slots,
	}

	if page > 1 {
		resp.PrevPage = fmt.Sprintf("http://%s%s?page=%d", c.Request.Host, c.Request.URL.Path, page-1)
	}
	resp.NextPage = fmt.Sprintf("http://%s%s?page=%d", c.Request.Host, c.Request.URL.Path, 1+page)
	// TODO: custom schedule handler

	c.JSON(http.StatusOK, resp)
}

// Handles request to fetch default schedule for a particular date.
func HandleGetDailySchedule(c *gin.Context) {
	var err error
	var scheduleId string
	var resp []models.Slot

	festId := c.Param("id")
	date := c.DefaultQuery("date", time.Now().Format(constants.DateInputFormat)) // Format: YYYY-MM-DD

	startOfDay, _ := time.Parse(time.RFC3339, date+"T00:00:00+05:30") // 00:00:00 IST in ISO
	endOfDay, _ := time.Parse(time.RFC3339, date+"T23:59:59+05:30")   // 23:59:59 IST in ISO
	startOfDayAsUnix := int(startOfDay.Unix())
	endOfDayAsUnix := int(endOfDay.Unix())

	scheduleId, err = services.GetDefaultScheduleId(festId)
	if err != nil {
		constants.HandleError(c, err)
		return
	}

	if scheduleId == "" {
		c.String(http.StatusNotFound, "No schedule available for this fest yet.")
		return
	}

	resp, err = services.GetScheduleSlotsByTime(scheduleId, startOfDayAsUnix, endOfDayAsUnix)
	if err != nil {
		constants.HandleError(c, err)
		return
	}
	c.JSON(http.StatusOK, resp)
	return
}
