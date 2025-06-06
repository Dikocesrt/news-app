package usecases

import (
	"fmt"
	"test-indonesia-cakap-digital/entities"
	"test-indonesia-cakap-digital/utils"
)

type NewsUsecase struct {
	newsRepository entities.NewsRepositoryInterface
}

func NewNewsUsecase(newsRepository entities.NewsRepositoryInterface) NewsUsecase {
	return NewsUsecase{
		newsRepository: newsRepository,
	}
}

func (c NewsUsecase) CreateNews(news entities.News) (entities.News, error) {
	if news.Content == "" || news.Category.ID == 0 || news.User.ID == 0 {
		fmt.Println("content:", news.Content, "categoryID:", news.Category.ID, "userID:", news.User.ID)
		return entities.News{}, utils.ErrEmptyField
	}

	return c.newsRepository.CreateNews(news)
}

func (c NewsUsecase) GetAllNews(metadata entities.Metadata) ([]entities.News, error) {
	return c.newsRepository.GetAllNews(metadata)
}

func (c NewsUsecase) GetNewsByID(newsID uint) (entities.News, error) {
	return c.newsRepository.GetNewsByID(newsID)
}

func (c NewsUsecase) UpdateNews(news entities.News) (entities.News, error) {
	if news.Content == "" || news.Category.ID == 0 || news.User.ID == 0 {
		return entities.News{}, utils.ErrEmptyField
	}
	
	return c.newsRepository.UpdateNews(news)
}

func (c NewsUsecase) DeleteNews(newsID uint, userID uint) error {
	return c.newsRepository.DeleteNews(newsID, userID)
}