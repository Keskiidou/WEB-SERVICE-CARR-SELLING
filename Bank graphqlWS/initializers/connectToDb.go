package initializers

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func ConnectToDb() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		panic("Failed to load .env file")
	}

	dsn := os.Getenv("DB")
	if dsn == "" {
		panic("DB environment variable is not set")
	}

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	println("Connected to database successfully!")
}
