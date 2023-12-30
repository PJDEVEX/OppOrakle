package models

import (
	"time"

	"gorm.io/gorm"
)

// ApplicationModel represents a Application entity with GORM support.
type ApplicationModel struct {
	gorm.Model
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

// Tabel name for the JobModel
func (ApplicationModel) TableName() string {
	return "applications"
}
