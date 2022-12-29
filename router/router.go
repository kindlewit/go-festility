package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kindlewit/go-festility/handlers"
)

func SetupRouter(router *gin.Engine) {
	// Bind routes & handlers
	router.GET("/", handlers.IndexHandler)

	router.GET("/movie/:id", handlers.GetMovieHandler)

	router.POST("/fest", handlers.CreateFestivalHandler)
	router.GET("/fest/:id", handlers.GetFestHandler)

	router.GET("/fest/:id/screen", handlers.GetFestScreensHandler)

	router.POST("/fest/:id/schedule", handlers.CreateScheduleHandler)
	router.GET("/fest/:id/schedule", handlers.GetDailyScheduleHandler)
	router.GET("/fest/:id/schedule/:sid", handlers.GetScheduleHandler)

	router.POST("/cinema", handlers.CreateCinemaHandler)
	router.GET("/cinema/:id", handlers.GetCinemaHandler)
	router.POST("/cinema/:id/screen", handlers.AddCinemaScreensHandler)
	router.GET("/cinema/:id/screen", handlers.GetCinemaScreensHandler)
	// Avoid return as incoming pointer.
}
