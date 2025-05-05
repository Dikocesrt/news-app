package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Password string `gorm:"type:varchar(255);not null"`
	Username string `gorm:"type:varchar(255);not null;unique"`
}

func (User) TableName() string {
	return "users"
}

func NewUser (password string, username string) User {
	return User{
		Password: password,
		Username: username,
	}
}