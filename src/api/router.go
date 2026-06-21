package api

import (
	"github.com/gin-gonic/gin"
	"github.com/kindlewit/go-festility/src/cinema"
	"github.com/kindlewit/go-festility/src/festival"
	"github.com/kindlewit/go-festility/src/movie"
	"github.com/kindlewit/go-festility/src/schedule"
)

// Helps bind the API path routes to the appropriate handlers.
func SetupRouter(router *gin.Engine) {
	// router.GET("/", handlers.HandleIndex)

	router.GET("/movie/:id", movie.HandleGetMovie)

	router.POST("/fest", festival.HandleCreateFest)
	router.GET("/fest", festival.HandleGetBulkFestivals)
	router.GET("/fest/:id", festival.HandleGetFest)

	// router.GET("/fest/:id/screen", handlers.HandleGetFestScreens)

	router.POST("/fest/:id/schedule", schedule.HandleCreateSchedule)
	router.GET("/fest/:id/schedule", schedule.HandleGetDailySchedule)
	router.GET("/fest/:id/schedule/:sid", schedule.HandleGetSchedule)

	router.POST("/cinema", cinema.HandleCreateCinema)
	router.GET("/cinema/:id", cinema.HandleGetCinema)
	router.POST("/cinema/:id/screen", cinema.HandleAddCinemaScreens)
	router.GET("/cinema/:id/screen", cinema.HandleGetCinemaScreens)
}
