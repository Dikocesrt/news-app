package repositories

import (
	"test-indonesia-cakap-digital/entities"
	"test-indonesia-cakap-digital/models"
	"test-indonesia-cakap-digital/utils"

	"gorm.io/gorm"
)

type TagRepository struct {
	DB *gorm.DB
}

func NewTagRepository(db *gorm.DB) TagRepository {
	return TagRepository{
		DB: db,
	}
}

func (c TagRepository) CreateTag(tag entities.Tag) (entities.Tag, error) {
	tagDB := models.NewTag(tag.Name)

	if err := c.DB.Where("name = ?", tag.Name).First(&tagDB).Error; err == nil {
		return entities.Tag{}, utils.ErrTagAlreadyExists
	}

	if err := c.DB.Create(&tagDB).Error; err != nil {
		return entities.Tag{}, err
	}

	return entities.Tag{
		ID:   tagDB.ID,
		Name: tagDB.Name,
	}, nil
}