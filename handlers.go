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

	client := connect();
	bulkInsertMovies(client, resp);

	c.JSON(http.StatusOK, resp);
}

// Handles request to read all movie documents in mongodb.
func readMovies(c *gin.Context) {
	client := connect();

	resp := allMovies(client);

	c.JSON(http.StatusOK, resp);
	// c.Writer.WriteHeader(204);
}

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

	c.JSON(http.StatusCreated, gin.H{ "id": body.Id });
}

// Handles request to get details of one festival.
func getFestHandler(c *gin.Context) {
	id := c.Param("id");

	client := connect();
	resp := getFest(client, id);

	c.JSON(http.StatusOK, resp);
}
