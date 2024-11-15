package models

import "gorm.io/gorm"

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
}

func CreateUser(db *gorm.DB, Username string, Password string, Email string) error {
	user := User{
		Username: Username,
		Password: Password,
		Email:    Email,
	}

	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}
