package main

import (
	"os"

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
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Set local port
	r.Run(":" + os.Getenv("PORT"))
}
