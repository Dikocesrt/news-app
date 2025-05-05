package models

import "gorm.io/gorm"

type CustomPage struct {
	gorm.Model
	CustomURL string `gorm:"type:varchar(255);not null;column:custom_url"`
	Content string `gorm:"type:text;not null"`
	UserID uint `gorm:"column:user_id;not null;index"`
	User User `gorm:"foreignKey:UserID"`
}

func (CustomPage) TableName() string {
	return "custom_pages"
}

func NewCustomPage(customURL string, content string, userID uint) CustomPage {
	return CustomPage{
		CustomURL: customURL,
		Content: content,
		UserID: userID,
	}
}