package main

import (
	"log"
	"os"

	"github.com/PJDEVEX/OppOrakle/controllers"
	"github.com/PJDEVEX/OppOrakle/initializers"
	"github.com/PJDEVEX/OppOrakle/services"
	"github.com/PJDEVEX/OppOrakle/utils"
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

	// Call s3 client from utils package
	s3Client, err := utils.CreateS3Client()
	if err != nil {
		log.Fatalf("Error crateing S3 client: v%", err)
	}
	log.Println("S3 client created successfully:", s3Client)

	// Route paths
	r.POST("/applications/", applicationController.AddApplication)
	r.GET("/applications/", applicationController.ListApplications)

	// Set local port
	r.Run(":" + os.Getenv("PORT"))
}
