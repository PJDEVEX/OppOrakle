package main

import (
	"os"

	"github.com/PJDEVEX/OppOrakle/controllers"
	"github.com/PJDEVEX/OppOrakle/initializers"
	"github.com/PJDEVEX/OppOrakle/services"
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

	applicationService := services.NewApplicationService()
	applicationController := controllers.NewApplicationController(applicationService)

	// Route paths
	r.POST("/applications/", applicationController.AddApplication)
	r.GET("/applications/", applicationController.ListApplications)

	// Set local port
	r.Run(":" + os.Getenv("PORT"))
}
