package main

import (
	"github.com/PJDEVEX/OppOrakle/initializers"
	"github.com/PJDEVEX/OppOrakle/models"
)

// init before the main
func init() {
	initializers.LoadEnvVarialbles()
	initializers.ConnectToDB()
}

func main() {
	// Automigrate DB tables
	initializers.DB.AutoMigrate(&models.ContactModel{})
}
