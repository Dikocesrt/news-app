package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);not null"`
}

func (Category) TableName() string {
	return "categories"
}

func NewCategory(name string) Category {
	return Category{
		Name: name,
	}
}