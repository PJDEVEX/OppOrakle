// controllers/applicationController.go
package controllers

import (
	"net/http"

	"github.com/PJDEVEX/OppOrakle/models"
	"github.com/PJDEVEX/OppOrakle/services"
	"github.com/gin-gonic/gin"
)

// Handle HTTP requests related to applications.
type ApplicationController struct {
	applicationService *services.ApplicationService
}

// Create a new ApplicationController instance.
func NewApplicationController(applicationService *services.ApplicationService) *ApplicationController {
	return &ApplicationController{
		applicationService: applicationService,
	}
}

// POST application
func (c *ApplicationController) AddApplication(ctx *gin.Context) {
	// Bind JSON input
	var input models.ApplicationModel
	// Error handling
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate title
	if input.Title == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Title is required!"})
		return
	} else if len(input.Title) > 100 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Title exceeds 100 characters limit!"})
	}

	// Validate Company
	if input.Company == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Company is required!"})
		return
	} else if len(input.Company) > 100 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Company name exceeds 100 characters limit!"})
		return
	}

	// Validate Source
	if len(input.Source) > 500 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Source exceeds 500 characters limit!"})
		return
	}

	// Validate Location
	if len(input.Location) > 100 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Location exceeds 100 characters limit!"})
		return
	}

	// Validate WorkArrangement
	if len(input.WorkArrangement) > 100 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "WorkArrangement exceeds 100 characters limit!"})
		return
	}

	// Validate TechStack
	if len(input.TechStack) > 500 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "TechStack exceeds 500 characters limit!"})
		return
	}

	// Validate Softskills
	if len(input.Softskills) > 500 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Softskills exceeds 500 characters limit!"})
		return
	}

	// Validate Recruiter
	if len(input.Recruiter) > 100 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Recruiter exceeds 100 characters limit!"})
		return

	}

	// Validate HiringManager
	if len(input.HiringManager) > 100 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "HiringManager exceeds 100 characters limit!"})
		return
	}

	// Validate ResponseType
	if len(input.ResponseType) > 100 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ResponseType exceeds 100 characters limit!"})
		return
	}

	// Validate Note
	if len(input.Note) > 500 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Note exceeds 500 characters limit!"})
		return
	}

	// Add application
	application, err := c.applicationService.AddApplication(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Respond with created application
	ctx.JSON(http.StatusCreated, gin.H{"application": application})
}

// GET all applications
func (c *ApplicationController) ListApplications(ctx *gin.Context) {
	applications, err := c.applicationService.ListApplications()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"applications": applications})
}
