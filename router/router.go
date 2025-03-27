package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kindlewit/go-festility/handlers"
)

// Helps bind the API path routes to the appropriate handlers.
func SetupRouter(router *gin.Engine) {
	// router.GET("/", handlers.HandleIndex)

	router.GET("/movie/:id", handlers.HandleGetMovie)

	router.POST("/fest", handlers.HandleCreateFest)
	router.GET("/fest/:id", handlers.HandleGetFest)

	// router.GET("/fest/:id/screen", handlers.HandleGetFestScreens)

	router.POST("/fest/:id/schedule", handlers.HandleCreateSchedule)
	router.GET("/fest/:id/schedule", handlers.HandleGetDailySchedule)
	router.GET("/fest/:id/schedule/:sid", handlers.HandleGetSchedule)

	router.POST("/cinema", handlers.HandleCreateCinema)
	router.GET("/cinema/:id", handlers.HandleGetCinema)
	router.POST("/cinema/:id/screen", handlers.HandleAddCinemaScreens)
	router.GET("/cinema/:id/screen", handlers.HandleGetCinemaScreens)
	// Avoid return as incoming pointer.
}
