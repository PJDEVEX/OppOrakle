package models

import (
	"gorm.io/gorm"
)

// GORM Contact Entity
type ContactModel struct {
	gorm.Model
	Name        string `gorm:"column:name"`
	Company     string `gorm:"column:company"`
	Designation string `gorm:"column:designation"`
	Dept        string `gorm:"column:dept"`
	Linkedin    string `gorm:"column:linkedin"`
	Connected   bool   `gorm:"column:connected"`
	Email       string `gorm:"column:email"`
	Mobile      string `gorm:"column:mobile"`
	DumRole     string `gorm:"column:dum_role"`
}
