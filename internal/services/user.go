package services

import (
	"github.com/Lincoln-a-Bot/GoPost/internal/models"
	"github.com/Lincoln-a-Bot/GoPost/internal/utils"
	"gorm.io/gorm"
)

// Function to check if the login is correct
func CheckLogin(db *gorm.DB, Username string, Password string) bool {
	user, err := models.GetUser(db, Username)
	if err != nil {
		return false
	}

	if utils.ComparePassword(user.Password, Password) {
		return true
	}
	return false
}
