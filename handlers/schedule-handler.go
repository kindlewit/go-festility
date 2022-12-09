package handlers

import (
  "fmt"
  "time"
  "net/http"
  "strconv"

  "github.com/gin-gonic/gin"
  "festility/utils"
  "festility/models"
  "festility/services"
  "festility/constants"
)

// Handles request to create a schedule for an existing fest.
func CreateScheduleHandler(c *gin.Context) {
  var err error;
  var success bool;

  festId := c.Param("id");

  // Get slots from request body
  var body []models.Slot;
  if err = c.ShouldBindJSON(&body); err != nil {
    c.String(http.StatusBadRequest, "Request body not matching slot structure.");
    return;
  }

  // Generate new schedule ID
  newScheduleId := utils.GenerateRandomHash(constants.ScheduleIDHashLength); // Generate new schedule ID

  client := services.Connect();
  // Ensure new ID is not present in db
  for (constants.ScheduleIDHashLength != 0 && len(newScheduleId) == 0) || !(services.IsUniqueScheduleID(client, newScheduleId)) {
    // while(new ID is not unique or is empty)
    newScheduleId = utils.GenerateRandomHash(constants.ScheduleIDHashLength);
  }

  // Store schedule
  newSchedule := models.Schedule{
    Id: newScheduleId,
    Fest: festId,
    Custom: false, // by default
  }
  success, err = services.CreateSchedule(client, newSchedule);
  // TODO: handle this err
  if (!success) {
    defer services.Disconnect(client);
    c.JSON(http.StatusInternalServerError, "Unable to create schedule. Please try again.");
    return;
  }

  // Store slots into db under generated schedule ID
  for i := 0; i < len(body); i++ {
    // TODO: check slot id
    body[i].ScheduleID = newScheduleId;
  }
  services.CreateSlots(client, body);
  defer services.Disconnect(client);

  // Return success with newly generated schedule ID
  c.JSON(http.StatusCreated, gin.H{
    "schedule_id": newScheduleId,
    "number_of_slots": len(body),
  });
}

// Handles request to fetch one schedule by schedule id.
func GetScheduleHandler(c *gin.Context) {
  var err error;

  festId := c.Param("id");
  scheduleId := c.Param("sid");
  page, _ := strconv.Atoi(c.DefaultQuery("page", "1")); // Pagination

  limit := int64(constants.SlotPageLimit);
  skip := int64(page * constants.SlotPageLimit - constants.SlotPageLimit);

  client := services.Connect();
  doc, err := services.GetSchedule(client, festId, scheduleId);
  if err != nil {
    constants.HandleError(c, err);
    return;
  }
  slots, err := services.GetScheduleSlots(client, scheduleId, limit, skip);
  if err != nil {
    constants.HandleError(c, err);
    return;
  }
  defer services.Disconnect(client);

  var resp struct {
    Id        string          `bson:"id" json:"id"`
    Fest      string          `bson:"fest_id" json:"fest_id"`
    Custom    bool            `bson:"custom" json:"custom"`
    // Username  string  `bson:"username" json:"username"`
    PrevPage  string          `json:"prev_page"`
    NextPage  string          `json:"next_page"`
    Slots     []models.Slot   `json:"slots"`
  }
  resp.Id = doc.Id;
  resp.Fest = doc.Fest;
  resp.Custom = doc.Custom;
  if (page > 1) {
    resp.PrevPage = fmt.Sprintf("http://%s%s?page=%d", c.Request.Host, c.Request.URL.Path, page - 1);
  }
  resp.NextPage = fmt.Sprintf("http://%s%s?page=%d", c.Request.Host, c.Request.URL.Path, 1 + page);
  // TODO: custom schedule handler
  // if (doc.Custom) {
  // // This is a custom schedule.
  // }
  // resp.Username = doc.Username;
  resp.Slots = slots;

  c.JSON(http.StatusOK, resp);
}

// Handles request to fetch default schedule for a particular date.
func GetDailyScheduleHandler(c *gin.Context) {
  var err error;
  var schId string;
  var resp []models.Slot;

  festId := c.Param("id");
  date := c.DefaultQuery("date", time.Now().Format(constants.DateInputFormat)); // Format: YYYY-MM-DD

  startOfDay, _ := time.Parse(time.RFC3339, date + "T00:00:00+05:30"); // 00:00:00 IST in ISO
  endOfDay, _ := time.Parse(time.RFC3339, date + "T23:59:59+05:30"); // 23:59:59 IST in ISO
  startOfDayAsUnix := int(startOfDay.Unix());
  endOfDayAsUnix := int(endOfDay.Unix());

  client := services.Connect();
  schId, err = services.GetDefaultScheduleID(client, festId);
  if (err != nil) {
    constants.HandleError(c, err);
    return;
  }
  if (schId == "") {
    c.String(http.StatusNotFound, "No schedule available for this fest yet.");
    return;
  }

  resp, err = services.GetScheduleSlotsByTime(client, schId, startOfDayAsUnix, endOfDayAsUnix);
  if (err != nil) {
    constants.HandleError(c, err);
    return;
  }
  c.JSON(http.StatusOK, resp);
  return;
}
