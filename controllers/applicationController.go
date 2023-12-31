// controllers/applicationController.go
package controllers

import (
	"net/http"

	"github.com/PJDEVEX/OppOrakle/models"
	"github.com/PJDEVEX/OppOrakle/services"
	"github.com/gin-gonic/gin"
)

// ApplicationController handles HTTP requests related to applications.
type ApplicationController struct {
	applicationService *services.ApplicationService
}

// NewApplicationController creates a new ApplicationController instance.
func NewApplicationController(applicationService *services.ApplicationService) *ApplicationController {
	return &ApplicationController{
		applicationService: applicationService,
	}
}

// POST application
func (c *ApplicationController) AddApplication(ctx *gin.Context) {
	var input models.ApplicationModel
	if err := ctx.BindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	application, err := c.applicationService.AddApplication(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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
