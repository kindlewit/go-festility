package main

import (
  "github.com/gin-gonic/gin"
  "github.com/kindlewit/go-festility/router"
)

func main() {
  app := gin.Default() // Equivalent to app in express

  // services.Migrate(); // Migrate to setup db schema
  router.SetupRouter(app);

  app.Run() // Start the server
}
