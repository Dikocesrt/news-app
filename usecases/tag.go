package usecases

import (
	"test-indonesia-cakap-digital/entities"
	"test-indonesia-cakap-digital/utils"
)

type TagUsecase struct {
	tagRepository entities.TagRepositoryInterface
}

func NewTypeUsecase(tagRepository entities.TagRepositoryInterface) TagUsecase {
	return TagUsecase{
		tagRepository: tagRepository,
	}
}

func (c TagUsecase) CreateTag(tag entities.Tag) (entities.Tag, error) {
	if tag.Name == "" {
		return entities.Tag{}, utils.ErrEmptyField
	}
	return c.tagRepository.CreateTag(tag)
}