package main

import (
	"log"
	"os"

	"github.com/PJDEVEX/OppOrakle/initializers"
	"github.com/PJDEVEX/OppOrakle/models"
)

// init before the main
func init() {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}
	log.Printf("Current working directory: %s", currentDir)

	initializers.LoadEnvVarialbles()
	initializers.ConnectToDB()
}

func main() {
	// Automigrate DB tables
	modelsToMigrate := []interface{}{
		&models.ApplicationModel{},
		&models.CompanyModel{},
		&models.ContactModel{},
	}

	for _, model := range modelsToMigrate {
		if err := initializers.DB.AutoMigrate(model); err != nil {
			log.Printf("Error auto-migrateing model %T: %v", model, err)
		}
		log.Printf("Model %T auto-migrated successfully", model)
	}
	log.Println("Database migration completed successfully")
}
