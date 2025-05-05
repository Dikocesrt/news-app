package models

import "gorm.io/gorm"

type News struct {
	gorm.Model
	Content string `gorm:"type:varchar(255);not null"`
	CategoryID uint `gorm:"column:category_id;not null;index"`
	Category Category `gorm:"foreignKey:CategoryID"`
	UserID uint `gorm:"column:user_id;not null;index"`
	User User `gorm:"foreignKey:UserID"`
}

func (News) TableName() string {
	return "news"
}

func NewNews(content string, categoryID uint, userID uint) News {
	return News{
		Content: content,
		CategoryID: categoryID,
		UserID: userID,
	}
}