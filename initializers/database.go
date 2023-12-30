package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Connecting to the DB
func ConnectToDB() {
	var err error
	// Retrieve the DB from env veriables.
	dsn := os.Getenv("VMDATABASE_URL")
	// Trying to open a conncection to the DB.
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		// log error
		log.Printf("Fail to connect to database: %v", err)
		return
	}
	// log success
	log.Println("Connected to the database successfully")

}
