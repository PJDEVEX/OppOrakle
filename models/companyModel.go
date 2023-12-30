package models

import "gorm.io/gorm"

// CompanyModel represents a Company entity with GORM support.
type CompanyModel struct {
	gorm.Model
	Name     string `gorm:"column:name;not null" json:"name"`
	Industry string `gorm:"column:industry" json:"industry"`
	Size     string `gorm:"column:size" json:"size"`
	Web      string `gorm:"column:web" json:"web"`
	Linkedin string `gorm:"column:linkedin" json:"linkedin"`
	Email    string `gorm:"column:email" json:"email"`
	Phone    string `gorm:"column:phone" json:"phone"`
}

// Table name company for CompanyModel
func (CompanyModel) TableName() string {
	return "companies"
}
