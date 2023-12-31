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
	var input struct {
		Title           string        `gorm:"column:title;not null" json:"title"`
		Company         string        `gorm:"column:company;not null" json:"company"`
		Source          string        `gorm:"column:source" json:"source"`
		Location        string        `gorm:"column:location" json:"location"`
		WorkArrangement string        `gorm:"column:work_arrangement" json:"work_arrangement"`
		AdvertPdf       string        `gorm:"column:advert_pdf" json:"advert_pdf"`
		TechStack       string        `gorm:"column:tech_stack" json:"tech_stack"`
		Softskills      string        `gorm:"column:softskills" json:"softskills"`
		Recruiter       string        `gorm:"column:recruiter" json:"recruiter"`
		HiringManager   string        `gorm:"column:hiring_manager" json:"hiring_manager"`
		DatePosted      time.Time     `gorm:"column:date_posted;not null" json:"date_posted"`
		Deadline        time.Time     `gorm:"column:deadline" json:"deadline"`
		AppliedDate     time.Time     `gorm:"column:applied_date" json:"applied_date"`
		ResponseType    string        `gorm:"column:response_type" json:"response_type"`
		ResponseDate    string        `gorm:"column:response_date" json:"response_date"`
		TimeElapsed     time.Duration `gorm:"column:time_elapsed" json:"time_elapsed"`
		Note            string        `gorm:"column:note" json:"note"`
	}

	c.Bind(&input)

	application := models.ApplicationModel{
		Title:           input.Title,
		Company:         input.Company,
		Source:          input.Source,
		Location:        input.Location,
		WorkArrangement: input.WorkArrangement,
		AdvertPdf:       input.AdvertPdf,
		TechStack:       input.TechStack,
		Softskills:      input.Softskills,
		Recruiter:       input.Recruiter,
		HiringManager:   input.HiringManager,
		DatePosted:      input.DatePosted,
		Deadline:        input.Deadline,
		AppliedDate:     input.AppliedDate,
		ResponseType:    input.ResponseType,
		ResponseDate:    input.ResponseDate,
		TimeElapsed:     input.TimeElapsed,
		Note:            input.Note,
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
