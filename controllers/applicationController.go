package controllers

import (
	"time"

	"github.com/PJDEVEX/OppOrakle/initializers"
	"github.com/PJDEVEX/OppOrakle/models"
	"github.com/gin-gonic/gin"
)

// POST application
func AddApplication(c *gin.Context) {
	// Sample for testing
	application := models.ApplicationModel{
		Title:           "Software Engineer",
		Company:         "Example Corp",
		Source:          "LinkedIn",
		Location:        "New York",
		WorkArrangement: "Remote",
		AdvertPdf:       "https://example.com/job_ad.pdf",
		TechStack:       "Go, React, PostgreSQL",
		Softskills:      "Communication, Teamwork",
		Recruiter:       "John Doe",
		HiringManager:   "Jane Smith",
		DatePosted:      time.Now(),
		Deadline:        time.Now().AddDate(0, 0, 14), // Two weeks from now
		AppliedDate:     time.Now().AddDate(0, 0, -7), // One week ago
		ResponseType:    "Pending",
		ResponseDate:    "",
		TimeElapsed:     7 * 24 * time.Hour, // One week in duration
		Note:            "Great potential candidate",
	}

	// Create data in DB
	result := initializers.DB.Create(&application)

	// Error handling
	if result.Error != nil {
		c.Status(400)
		return
	}

	// Confirm the creation
	c.JSON(201, gin.H{
		"application": application,
	})

}

// Get all applications
func ListApplication(C *gin.Context) {
	// Slice to store data
	var applications []models.ApplicationModel
	// Fetch data from DB
	if err := initializers.DB.Find(&applications).Error; err != nil {
		// handling DB error
		C.JSON(500, gin.H{"error": err.Error()})
		return

	}
	// Respond with app data
	C.JSON(200, gin.H{
		"applications": applications,
	})
}
