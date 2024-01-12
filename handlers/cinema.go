package handlers

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kindlewit/go-festility/constants"
	"github.com/kindlewit/go-festility/models"
	"github.com/kindlewit/go-festility/services"
)

// Handles request to create a cinema.
func HandleCreateCinema(c *gin.Context) {
	var body models.Cinema
	var success bool
	var err error

	if err = c.BindJSON(&body); err != nil {
		c.String(http.StatusBadRequest, "Required parameters missing: name/city.")
		return
	}

	// Add Cinema ID
	body.Id = fmt.Sprintf("%d", rand.Intn(1000)) // Random id b/w 0 - 1000

	success, err = services.CreateCinema(body)
	if err != nil {
		constants.HandleError(c, err)
		return
	}
	if !success {
		// Record created but with no InsertedID
		c.JSON(http.StatusConflict, "Record was created but found an inconsistency in record id.")
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": body.Id, "name": body.Name})
}

// Handles request to fetch one cinema.
func HandleGetCinema(c *gin.Context) {
	var doc models.Cinema
	var err error

	cinemaID := c.Param("id")
	if cinemaID == "" || cinemaID == "null" {
		// Missing cinema ID param
		c.JSON(http.StatusBadRequest, "Missing valid cinema id. Please check the parameter.")
		return
	}

	doc, err = services.GetCinema(cinemaID)
	if err != nil {
		constants.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, doc)
	return
}

// Handles request to insert screens to a cinema.
func HandleAddCinemaScreens(c *gin.Context) {
	var body []models.Screen
	var success bool
	var err error

	cinemaID := c.Param("id")
	if cinemaID == "" || cinemaID == "null" {
		// Missing cinema ID param
		c.JSON(http.StatusBadRequest, "Missing valid cinema id. Please check the parameter.")
		return
	}

	if err = c.BindJSON(&body); err != nil {
		c.String(http.StatusBadRequest, "Request body is of invalid structure.")
		return
	}

	// Ensure all records have same cinema ID.
	for i := 0; i < len(body); i++ {
		body[i].CinemaID = cinemaID
		// Add screen ID
		body[i].Id = fmt.Sprintf("%d", rand.Intn(1000)) // Random id b/w 0 - 1000
	}

	success, err = services.CreateCinemaScreens(body)
	if err != nil {
		constants.HandleError(c, err)
		return
	}
	if !success {
		c.JSON(http.StatusInternalServerError, "Faced an error in record creation. Please try again.")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"cinema_id":         cinemaID,
		"number_of_screens": len(body),
	})
}

// Handles request to fetch cinema screens.
func HandleGetCinemaScreens(c *gin.Context) {
	var resp []models.Screen
	var err error

	cinemaID := c.Param("id")
	if cinemaID == "" || cinemaID == "null" {
		// Missing cinema ID param
		c.JSON(http.StatusBadRequest, "Missing valid cinema id. Please check the parameter.")
		return
	}

	resp, err = services.GetCinemaScreens(cinemaID)
	if err != nil {
		constants.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, resp)
	return
}
