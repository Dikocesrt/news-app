package usecases

import (
	"makanan-app/entities"
	"makanan-app/utils"
	"strings"
)

type CustomPageUsecase struct {
	customPageRepository entities.CustomPageRepositoryInterface
}

func NewCustomPageUsecase(customPageRepository entities.CustomPageRepositoryInterface) CustomPageUsecase {
	return CustomPageUsecase{
		customPageRepository: customPageRepository,
	}
}

func (c CustomPageUsecase) CreateCustomPage(customPage entities.CustomPage) (entities.CustomPage, error) {
	if customPage.CustomURL == "" || customPage.Content == "" || customPage.User.ID == 0 {
		return entities.CustomPage{}, utils.ErrEmptyField
	}

	if strings.Contains(customPage.CustomURL, " ") {
		return entities.CustomPage{}, utils.ErrInvalidCustomURL
	}

	return c.customPageRepository.CreateCustomPage(customPage)
}

func (c CustomPageUsecase) GetAllCustomPages(metadata entities.Metadata) ([]entities.CustomPage, error) {
	return c.customPageRepository.GetAllCustomPages(metadata)
}

func (c CustomPageUsecase) GetCustomPageByID(id uint) (entities.CustomPage, error) {
	return c.customPageRepository.GetCustomPageByID(id)
}

func (c CustomPageUsecase) UpdateCustomPage(customPage entities.CustomPage) (entities.CustomPage, error) {
	if customPage.ID == 0 || customPage.CustomURL == "" || customPage.Content == "" || customPage.User.ID == 0 {
		return entities.CustomPage{}, utils.ErrEmptyField
	}

	if strings.Contains(customPage.CustomURL, " ") {
		return entities.CustomPage{}, utils.ErrInvalidCustomURL
	}

	return c.customPageRepository.UpdateCustomPage(customPage)
}

func (c CustomPageUsecase) DeleteCustomPage(id uint, userID uint) error {
	return c.customPageRepository.DeleteCustomPage(id, userID)
}