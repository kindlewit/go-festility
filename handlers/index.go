package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kindlewit/go-festility/constants"
	"github.com/kindlewit/go-festility/services"
)

// Handles requests to index (/) endpoint.
// The purpose of this endpoint is to ensure if the server is working,
// and if the necessary services are up & running.
func HandleIndex(c *gin.Context) {
	// Check connection to database
	timestamp := time.Now().Unix()
	success := services.StatusCheck()

	var resp gin.H

	if !success {
		// Failed to connect to DB
		resp = gin.H{
			"timestamp": timestamp,
			"success":   false,
			"message":   constants.MsgDatabaseFailure,
		}
	} else {
		resp = gin.H{
			"timestamp": timestamp,
			"success":   true,
			"message":   "Welcome to festility!",
		}
	}

	c.JSON(http.StatusOK, resp)
}
