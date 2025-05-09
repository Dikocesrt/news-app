package models

import "gorm.io/gorm"

type NewsTag struct {
	gorm.Model
	NewsID uint `gorm:"column:news_id;not null;index"`
	News News `gorm:"foreignKey:NewsID"`
	TagID uint `gorm:"column:tag_id;not null;index"`
	Tag Tag `gorm:"foreignKey:TagID"`
}

func (NewsTag) TableName() string {
	return "news_tags"
}

func NewNewsTag(newsID uint, tagID uint) NewsTag {
	return NewsTag{
		NewsID: newsID,
		TagID: tagID,
	}
}