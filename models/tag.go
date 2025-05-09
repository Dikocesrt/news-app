package models

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);not null;unique"`
}

func (Tag) TableName() string {
	return "tags"
}

func NewTag(name string) Tag {
	return Tag{
		Name: name,
	}
}