package models

import (
	"github.com/Lincoln-a-Bot/GoPost/internal/utils"
	"gorm.io/gorm"
)

// User struct

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
}

// Function to create a user, Hashing the password and saving to DB
func CreateUser(db *gorm.DB, Username string, Password string, Email string) error {

	HashPassword, err := utils.HashPassword(Password)
	if err != nil {
		return err
	}

	user := User{
		Username: Username,
		Password: HashPassword,
		Email:    Email,
	}

	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

// Function to get a user from the DB
func GetUser(db *gorm.DB, Username string) (*User, error) {
	var user User
	if err := db.Where("username = ?", Username).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
