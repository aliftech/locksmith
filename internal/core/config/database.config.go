package config

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	USER := os.Getenv("USER")
	PASS := os.Getenv("PASSWORD")
	DBNAME := os.Getenv("DBNAME")
	HOST := os.Getenv("HOST")

	dsn := USER + ":" + PASS + "@tcp(" + HOST + ")/" + DBNAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
