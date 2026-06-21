package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kindlewit/go-festility/src/api"
)

func main() {
	app := gin.Default() // Equivalent to app in express

	// services.Migrate(); // Migrate to setup db schema
	api.SetupRouter(app)

	app.Run() // Start the server
}
