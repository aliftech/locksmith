package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// CheckEnv validates environment variables (assuming it exists; if not, remove or define it)
func setup() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

}

func Connect() {
	// Validate environment variables (optional if CheckEnv is called elsewhere)
	setup()
	var err error

	// Retrieve environment variables
	host := os.Getenv("HOST")
	dbName := os.Getenv("DBNAME")
	user := os.Getenv("USER")
	pass := os.Getenv("PASSWORD")

	// Construct DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, dbName)

	// Open database connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panicf("Failed to connect to database: %v", err)
	}

	// Assign to global DB variable
	DB = db

	// Verify connection
	sqlDB, err := db.DB()
	if err != nil {
		log.Panicf("Failed to get database instance: %v", err)
	}
	if pingErr := sqlDB.Ping(); pingErr != nil {
		log.Panicf("Database ping failed: %v", pingErr)
	}

	log.Println("Successfully connected to the database!")
}
