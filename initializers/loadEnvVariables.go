package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

// Load env variables from .env
func LoadEnvVarialbles() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error in loading .env file")
	}
}
