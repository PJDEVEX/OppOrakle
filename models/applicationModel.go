package models

import (
	"time"

	"gorm.io/gorm"
)

// ApplicationModel represents a Application entity with GORM support.
type ApplicationModel struct {
	gorm.Model
	Title           string    `gorm:"column:title;not null;varchar(100)" json:"title"`
	Company         string    `gorm:"column:company;not null; varchar(100)" json:"company"`
	Source          string    `gorm:"column:source; varchar(500)" json:"source"`
	Location        string    `gorm:"column:location; varchar(100)" json:"location"`
	WorkArrangement string    `gorm:"column:work_arrangement; varchar(100)" json:"work_arrangement"`
	AdvertPdf       string    `gorm:"column:advert_pdf" json:"advert_pdf"`
	TechStack       string    `gorm:"column:tech_stack; varchar(500)" json:"tech_stack"`
	Softskills      string    `gorm:"column:softskills; varchar(500)" json:"softskills"`
	Recruiter       string    `gorm:"column:recruiter; varchar(100)" json:"recruiter"`
	HiringManager   string    `gorm:"column:hiring_manager; varchar(100)" json:"hiring_manager"`
	DatePosted      time.Time `gorm:"column:date_posted; not null" json:"date_posted"`
	Deadline        time.Time `gorm:"column:deadline" json:"deadline"`
	AppliedDate     time.Time `gorm:"column:applied_date" json:"applied_date"`
	ResponseType    string    `gorm:"column:response_type; varchar(100)" json:"response_type"`
	ResponseDate    string    `gorm:"column:response_date" json:"response_date"`
	TimeElapsed     int64     `gorm:"column:time_elapsed" json:"time_elapsed"`
	Note            string    `gorm:"column:note; varchar(500)" json:"note"`
}

// Tabel name for the JobModel
func (ApplicationModel) TableName() string {
	return "applications"
}
