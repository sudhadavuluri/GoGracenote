package main

//Import dependencies in the defined order to avoid conflicts. These initiate before this main executes
import (
  "github.com/gin-gonic/gin"
  
)

var router *gin.Engine

func main() {


  // Set the router as the default one provided by Gin
  router = gin.Default()
  
  // Initialize the routes
  initializeRoutes()

  // Start serving the application
  router.Run()
}

