package usecases

import (
	"makanan-app/entities"
	"makanan-app/utils"
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
	if news.Content == "" || news.CategoryID == 0 || news.UserID == 0 {
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
	if news.Content == "" || news.CategoryID == 0 || news.UserID == 0 {
		return entities.News{}, utils.ErrEmptyField
	}
	
	return c.newsRepository.UpdateNews(news)
}

func (c NewsUsecase) DeleteNews(newsID uint) error {
	return c.newsRepository.DeleteNews(newsID)
}