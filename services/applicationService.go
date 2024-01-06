package services

import (
	"github.com/PJDEVEX/OppOrakle/initializers"
	"github.com/PJDEVEX/OppOrakle/models"
)

// Handles business logic
type ApplicationService struct{}

// Creates new ApplicationService instance
func NewApplicationService() *ApplicationService {
	return &ApplicationService{}
}

// Creates a new application in DB
func (s *ApplicationService) AddApplication(input *models.ApplicationModel) (*models.ApplicationModel, error) {
	application := &models.ApplicationModel{
		Title:           input.Title,
		Company:         input.Company,
		Source:          input.Source,
		Location:        input.Location,
		WorkArrangement: input.WorkArrangement,
		AdvertPdfUrl:    input.AdvertPdfUrl,
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

	result := initializers.DB.Create(application)

	if err := result.Error; err != nil {
		return nil, err
	}

	return application, nil
}

// Retrieve all applications from the DB
func (s *ApplicationService) ListApplications() ([]models.ApplicationModel, error) {
	var applications []models.ApplicationModel
	if err := initializers.DB.Find(&applications).Error; err != nil {
		return nil, err
	}

	return applications, nil
}
