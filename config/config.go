package config

import (
	"fmt"

	"github.com/Lincoln-a-Bot/GoPost/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Function to connect to the DB or create if none
func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("../database.sqlite"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	// Creating tables for user
	if err := db.AutoMigrate(&models.User{}); err != nil {
		fmt.Println(err)
	}

	return db, nil
}
