package models

import (
	"gorm.io/gorm"
)

// ContactModel represents a contact entity with GORM support.
type ContactModel struct {
	gorm.Model
	Name        string `gorm:"column:name;not null" json:"name"`
	Company     string `gorm:"column:company;not null" json:"company"`
	Designation string `gorm:"column:designation" json:"designation"`
	Dept        string `gorm:"column:dept" json:"dept"`
	Linkedin    string `gorm:"column:linkedin" json:"linkedin"`
	Connected   bool   `gorm:"column:connected" json:"connected"`
	Email       string `gorm:"column:email" json:"email"`
	Mobile      string `gorm:"column:mobile" json:"mobile"`
	DmuRole     string `gorm:"column:dum_role" json:"dmu_role"`
}

// Table name for the ContactModel
func (ContactModel) TableName() string {
	return "contacts"
}
