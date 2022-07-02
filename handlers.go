/* All request handler functions go here */
package main

import (
	"time"
  "net/http"

	"github.com/gin-gonic/gin"
)

// Handles requests to / endpoint.
func indexHandler(c *gin.Context) {
	var response = gin.H{
		"timestamp": time.Now().Unix(),
		"data": "Welcome to festility!",
	}

	c.JSON(http.StatusOK, response);
}

// Handles requests to get one movie.
func getMovieHandler(c *gin.Context) {
	id := c.Param("id");

	movieDoc, success := getMovie(id);
	if !success {
		c.String(http.StatusInternalServerError, "");
	}
	directors := getDirector(id);

	resp := reformat(movieDoc, directors);

	c.JSON(http.StatusOK, resp);
}

// Handles request to get movies of a TMDB list.
func moviesFromListHandler(c *gin.Context) {
	listId := c.Param("id");

	movieIds := getListMovieIds(listId);
	resp := []Movie{};

	for _, id := range movieIds { // Doing this efficiently requires routine & waitgroups
		mDoc, success := getMovie(id);
		if !success {
			panic("Oh no!");
		}
		dir := getDirector(id);

		resp = append(resp, reformat(mDoc, dir));
	}

	// Removing insert multiple movies into db
	// client := connect();
	// bulkInsertMovies(client, resp);

	c.JSON(http.StatusOK, resp);
}

// Handles request to read all movie documents in mongodb.
// func readMovies(c *gin.Context) {
// 	client := connect();
// 	resp := allMovies(client);
// 	defer disconnect(client);

// 	c.JSON(http.StatusOK, resp);
// 	// c.Writer.WriteHeader(204);
// }

// Handles request to create one festival.
func createFestHandler(c *gin.Context) {
	var body Fest;

	if err := c.BindJSON(&body); err != nil {
		c.String(http.StatusBadRequest, "Required parameters missing: name/from_date/to_date.");
		return;
	}

	client := connect();
	festId := createFest(client, body);
	if festId == "" {
		c.String(http.StatusInternalServerError, "Unable to create record.");
		return;
	}
	if festId == DuplicateRecord {
		c.String(http.StatusConflict, "Request to create duplicate record.");
		return;
	}
	defer disconnect(client);

	c.JSON(http.StatusCreated, gin.H{ "id": body.Id });
}

// Handles request to get details of one festival.
func getFestHandler(c *gin.Context) {
	festId := c.Param("id");

	client := connect();
	resp := getFest(client, festId);
	defer disconnect(client);

	c.JSON(http.StatusOK, resp);
}

// Handles request to create a schedule for an existing fest.
func createScheduleHandler(c *gin.Context) {
	festId := c.Param("id");

	client := connect();

	// Get slots from request body
	var body []Slot;
	if err := c.ShouldBindJSON(&body); err != nil {
		c.String(http.StatusBadRequest, "Request body not matching slot structure.");
		return;
	}

	newlyGenSchId := generateRandomHash(ScheduleIDLen); // Generate new schedule ID
	// Ensure new ID is not present in db
	for ensureUniqueScheduleId(client, newlyGenSchId) != true || len(newlyGenSchId) == 0 {
		// while(newId is not unique or is empty)
		newlyGenSchId = generateRandomHash(ScheduleIDLen);
	}

	// Store schedule
	sch := Schedule{
		Id: newlyGenSchId,
		Fest:	festId,
		Custom: false, // by default
	}
	createSchedule(client, sch);

	// Store slots into db under generated ID
	for i := 0; i < len(body); i++ {
		body[i].ScheduleID = newlyGenSchId;
	}
	bulkCreateSlot(client, body);
	defer disconnect(client);


	// Return generated ID
	c.JSON(http.StatusCreated, gin.H{
		"schedule_id": newlyGenSchId,
		"number_of_slots": len(body),
	});
}

// Handles request to fetch one schedule by schedule id.
func getScheduleHandler(c *gin.Context) {
	festId := c.Param("id");
	scheduleId := c.Param("sid");

	client := connect();
	doc := getSchedule(client, festId, scheduleId);
	slots := getSlotsOfSchedule(client, scheduleId);
	defer disconnect(client);

	var resp struct {
		Id					string		`bson:"id" json:"id"`
		Fest				string		`bson:"fest_id" json:"fest_id"`
		Custom			bool			`bson:"custom" json:"custom"`
		Username		string		`bson:"username" json:"username"`
		Slots				[]Slot		`json:"slots"`
	}
	resp.Id = doc.Id;
	resp.Fest = doc.Fest;
	resp.Custom = doc.Custom;
	resp.Username = doc.Username;
	resp.Slots = slots;

	c.JSON(http.StatusOK, resp);
}

// Handles request to fetch default schedule for a particular date.
func getDailyScheduleHandler(c *gin.Context) {
	festId := c.Param("id");
	date := c.DefaultQuery("date", time.Now().Format(DateInputParse)); // Format: YYYY-MM-DD

	startOfDay, _ := time.Parse(time.RFC3339, date + "T00:00:00+05:30"); // 00:00:00 IST in ISO
	endOfDay, _ := time.Parse(time.RFC3339, date + "T23:59:59+05:30"); // 23:59:59 IST in ISO

	client := connect();
	schId := getDefaultScheduleId(client, festId);
	if (schId == "") {
		c.String(http.StatusNotFound, "No schedule present for this fest yet.");
		return;
	}

	resp := findSlotsByTime(client, schId, int(startOfDay.Unix()), int(endOfDay.Unix()));
	c.JSON(http.StatusOK, resp);
}
