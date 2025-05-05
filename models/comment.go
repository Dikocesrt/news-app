package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Name string `gorm:"type:varchar(255)"`
	Comment string `gorm:"type:text;not null"`
	NewsID uint `gorm:"column:news_id;not null;index"`
	News News `gorm:"foreignKey:NewsID"`
}

func (Comment) TableName() string {
	return "news_comments"
}

func NewComment(name string, comment string, newsID uint) Comment {
	if name == "" {
		name = "Anonymous"
	}
	return Comment{
		Name:    name,
		Comment: comment,
		NewsID:  newsID,
	}
}