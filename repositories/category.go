package repositories

import (
	"errors"
	"test-indonesia-cakap-digital/entities"
	"test-indonesia-cakap-digital/models"
	"test-indonesia-cakap-digital/utils"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return CategoryRepository{
		DB: db,
	}
}

func (c CategoryRepository) CreateCategory(category entities.Category) (entities.Category, error) {
	categoryDB := models.NewCategory(category.Name)

	if err := c.DB.Where("name = ?", category.Name).First(&categoryDB).Error; err == nil {
		return entities.Category{}, utils.ErrCategoryAlreadyExists
	}

	if err := c.DB.Create(&categoryDB).Error; err != nil {
		return entities.Category{}, errors.New("failed to create category")
	}

	return entities.Category{
		ID:   categoryDB.ID,
		Name: categoryDB.Name,
	}, nil
}

func (c CategoryRepository) GetAllCategories(metadata entities.Metadata) ([]entities.Category, error) {
	var categoriesDB []models.Category

	if err := c.DB.Limit(metadata.Limit).Offset(metadata.GetOffset()).Find(&categoriesDB).Error; err != nil {
		return []entities.Category{}, errors.New("failed to get categories")
	}

	var categories []entities.Category
	for _, categoryDB := range categoriesDB {
		categories = append(categories, entities.Category{
			ID:   categoryDB.ID,
			Name: categoryDB.Name,
		})
	}

	return categories, nil
}

func (c CategoryRepository) GetCategoryByID(categoryID uint) (entities.Category, error) {
	var categoryDB models.Category

	if err := c.DB.Where("id = ?", categoryID).First(&categoryDB).Error; err != nil {
		return entities.Category{}, utils.ErrInvalidCategoryID
	}

	return entities.Category{
		ID:   categoryDB.ID,
		Name: categoryDB.Name,
	}, nil
}

func (c CategoryRepository) UpdateCategory(category entities.Category) (entities.Category, error) {
	var categoryDB models.Category

	if err := c.DB.Where("name = ?", category.Name).First(&categoryDB).Error; err == nil {
		return entities.Category{}, utils.ErrCategoryAlreadyExists
	}

	if err := c.DB.Where("id = ?", category.ID).First(&categoryDB).Error; err != nil {
		return entities.Category{}, utils.ErrInvalidCategoryID
	}

	categoryDB.Name = category.Name
	if err := c.DB.Save(&categoryDB).Error; err != nil {
		return entities.Category{}, errors.New("failed to update category")
	}

	return entities.Category{
		ID:   categoryDB.ID,
		Name: categoryDB.Name,
	}, nil
}

func (c CategoryRepository) DeleteCategory(categoryID uint) error {
	var categoryDB models.Category

	if err := c.DB.Where("id = ?", categoryID).First(&categoryDB).Error; err != nil {
		return utils.ErrInvalidCategoryID
	}

	if err := c.DB.Delete(&categoryDB).Error; err != nil {
		return errors.New("failed to delete category")
	}

	return nil
}