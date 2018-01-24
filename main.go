package main

//Import dependencies in the defined order to avoid conflicts. These initiate before this main executes
import (
  "github.com/gin-gonic/gin"
  "fmt"
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

func ErrDefaultHandler(er error) {
	fmt.Println("handle using your own error handling or logging solution","\n")
	panic(er)
	}

