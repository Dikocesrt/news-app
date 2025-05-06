package repositories

import (
	"errors"
	"test-indonesia-cakap-digital/entities"
	"test-indonesia-cakap-digital/models"
	"test-indonesia-cakap-digital/utils"

	"gorm.io/gorm"
)

type CustomPageRepository struct {
	DB *gorm.DB
}

func NewCustomPageRepository(db *gorm.DB) CustomPageRepository {
	return CustomPageRepository{
		DB: db,
	}
}

func (c CustomPageRepository) CreateCustomPage(customPage entities.CustomPage) (entities.CustomPage, error) {
	customPageDB := models.NewCustomPage(customPage.CustomURL, customPage.Content, customPage.User.ID)

	if err := c.DB.Where("custom_url = ?", customPage.CustomURL).First(&customPageDB).Error; err == nil {
		return entities.CustomPage{}, utils.ErrCustomPageAlreadyExists
	}

	if err := c.DB.Create(&customPageDB).Error; err != nil {
		return entities.CustomPage{}, errors.New("failed to create custom page")
	}

	return entities.CustomPage{
		ID:        customPageDB.ID,
		CustomURL: customPageDB.CustomURL,
		Content:   customPageDB.Content,
		User: entities.User{
			ID: customPageDB.UserID,
		},
	}, nil
}

func (c CustomPageRepository) GetAllCustomPages(metadata entities.Metadata) ([]entities.CustomPage, error) {
	var customPagesDB []models.CustomPage

	if err := c.DB.Preload("User").Offset(metadata.GetOffset()).Limit(metadata.Limit).Find(&customPagesDB).Error; err != nil {
		return []entities.CustomPage{}, errors.New("failed to get custom pages")
	}

	var customPages []entities.CustomPage
	for _, cp := range customPagesDB {
		customPages = append(customPages, entities.CustomPage{
			ID:        cp.ID,
			CustomURL: cp.CustomURL,
			Content:   cp.Content,
			User: entities.User{
				ID: cp.UserID,
				Username: cp.User.Username,
			},
		})
	}

	return customPages, nil
}

func (c CustomPageRepository) GetCustomPageByID(id uint) (entities.CustomPage, error) {
	var customPageDB models.CustomPage

	if err := c.DB.Where("id = ?", id).Preload("User").First(&customPageDB).Error; err != nil {
		return entities.CustomPage{}, utils.ErrInvalidCustomPageID
	}

	return entities.CustomPage{
		ID:        customPageDB.ID,
		CustomURL: customPageDB.CustomURL,
		Content:   customPageDB.Content,
		User: entities.User{
			ID: customPageDB.UserID,
			Username: customPageDB.User.Username,
		},
	}, nil
}

func (c CustomPageRepository) UpdateCustomPage(customPage entities.CustomPage) (entities.CustomPage, error) {
	customPageDB := models.NewCustomPage(customPage.CustomURL, customPage.Content, customPage.User.ID)

	if err := c.DB.Where("id = ?", customPage.ID).First(&customPageDB).Error; err != nil {
		return entities.CustomPage{}, utils.ErrInvalidCustomPageID
	}

	if customPageDB.UserID != customPage.User.ID {
		return entities.CustomPage{}, utils.ErrUnauthorized
	}

	customPageDB.CustomURL = customPage.CustomURL
	customPageDB.Content = customPage.Content
	customPageDB.UserID = customPage.User.ID
	if err := c.DB.Save(&customPageDB).Error; err != nil {
		return entities.CustomPage{}, errors.New("failed to update custom page")
	}

	return entities.CustomPage{
		ID:        customPageDB.ID,
		CustomURL: customPageDB.CustomURL,
		Content:   customPageDB.Content,
		User: entities.User{
			ID: customPageDB.UserID,
		},
	}, nil
}

func (c CustomPageRepository) DeleteCustomPage(id uint, userID uint) error {
	var customPageDB models.CustomPage

	if err := c.DB.Where("id = ?", id).First(&customPageDB).Error; err != nil {
		return utils.ErrInvalidCustomPageID
	}

	if customPageDB.UserID != userID {
		return utils.ErrUnauthorized
	}

	if err := c.DB.Delete(&customPageDB).Error; err != nil {
		return errors.New("failed to delete custom page")
	}

	return nil
}