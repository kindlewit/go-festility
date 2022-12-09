package main

import (
  "github.com/gin-gonic/gin"
  "festility/handlers"
  // "festility/services"
)

func main() {
  router := gin.Default() // Equivalent to app in express

  // Bind routes & handlers
  router.GET("/", handlers.IndexHandler);

  router.GET("/movie/:id", handlers.GetMovieHandler);

  router.POST("/fest", handlers.CreateFestivalHandler);
  router.GET("/fest/:id", handlers.GetFestHandler);

  router.POST("/fest/:id/schedule", handlers.CreateScheduleHandler);
  router.GET("/fest/:id/schedule", handlers.GetDailyScheduleHandler);
  router.GET("/fest/:id/schedule/:sid", handlers.GetScheduleHandler);

  router.POST("/cinema", handlers.CreateCinemaHandler);
  router.GET("/cinema/:id", handlers.GetCinemaHandler);
  router.POST("/cinema/:id/screen", handlers.AddCinemaScreensHandler);
  router.GET("/cinema/:id/screen", handlers.GetCinemaScreensHandler);

  // services.Migrate(); // Migrate to setup db schema
  router.Run(); // Start the server
}
