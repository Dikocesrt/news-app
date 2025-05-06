package usecases

import (
	"test-indonesia-cakap-digital/entities"
	"test-indonesia-cakap-digital/utils"
)

type CategoryUsecase struct {
	categoryRepository entities.CategoryRepositoryInterface
}

func NewCategoryUsecase(categoryRepository entities.CategoryRepositoryInterface) CategoryUsecase {
	return CategoryUsecase{
		categoryRepository: categoryRepository,
	}
}

func (c CategoryUsecase) CreateCategory(category entities.Category) (entities.Category, error) {
	if category.Name == "" {
		return entities.Category{}, utils.ErrEmptyField
	}

	return c.categoryRepository.CreateCategory(category)
}

func (c CategoryUsecase) GetAllCategories(metadata entities.Metadata) ([]entities.Category, error) {
	return c.categoryRepository.GetAllCategories(metadata)
}

func (c CategoryUsecase) GetCategoryByID(categoryID uint) (entities.Category, error) {
	return c.categoryRepository.GetCategoryByID(categoryID)
}

func (c CategoryUsecase) UpdateCategory(category entities.Category) (entities.Category, error) {
	if category.Name == "" {
		return entities.Category{}, utils.ErrEmptyField
	}

	return c.categoryRepository.UpdateCategory(category)
}

func (c CategoryUsecase) DeleteCategory(categoryID uint) error {
	return c.categoryRepository.DeleteCategory(categoryID)
}