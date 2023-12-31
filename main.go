package main

import (
	"os"

	"github.com/PJDEVEX/OppOrakle/controllers"
	"github.com/PJDEVEX/OppOrakle/initializers"
	"github.com/gin-gonic/gin"
)

// Initialize key comp before the main
func init() {
	initializers.LoadEnvVarialbles()
	initializers.ConnectToDB()
}

func main() {
	// Gin router
	r := gin.Default()

	// Route paths
	r.GET("/applications/", controllers.ListApplication)
	r.POST("/applications/", controllers.AddApplication)

	// Set local port
	r.Run(":" + os.Getenv("PORT"))
}
