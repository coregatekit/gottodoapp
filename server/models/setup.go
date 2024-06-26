package models

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	dsn := os.Getenv("DATABASE_URL")
	DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}
}

func MigrateDatabase() {
}
