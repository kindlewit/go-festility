package main

import "github.com/gin-gonic/gin"

func main() {
  router := gin.Default() // Equivalent to app in express

	// Bind routes & handlers
	router.GET("/", indexHandler);
	router.GET("/movie/:id", getMovieHandler);
	router.GET("/list/:id", moviesFromListHandler);
	router.GET("/movies", readMovies);

	router.POST("/fest", createFestHandler);
	router.GET("/fest/:id", getFestHandler);

	router.POST("/fest/:id/schedule", createScheduleHandler);
	router.GET("/fest/:id/schedule/:sid", getScheduleHandler);

	migrate(); // Migrate to setup db schema
	router.Run(); // Start the server
}
