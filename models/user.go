package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Password string `gorm:"type:varchar(255);not null"`
	Username string `gorm:"type:varchar(255);not null;unique"`
}

func NewUser (password string, username string) User {
	return User{
		Password: password,
		Username: username,
	}
}