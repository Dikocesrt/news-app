package repositories

import (
	"errors"
	"makanan-app/entities"
	"makanan-app/models"
	"makanan-app/utils"

	"gorm.io/gorm"
)

type NewsRepository struct {
	DB *gorm.DB
}

func NewNewsRepository(db *gorm.DB) NewsRepository {
	return NewsRepository{
		DB: db,
	}
}

func (c NewsRepository) CreateNews(news entities.News) (entities.News, error) {
	newsDB := models.NewNews(news.Content, news.CategoryID, news.UserID)

	if err := c.DB.Create(&newsDB).Error; err != nil {
		return entities.News{}, errors.New("failed to create news")
	}

	return entities.News{
		ID:         newsDB.ID,
		Content:    newsDB.Content,
		CategoryID: newsDB.CategoryID,
		UserID:     newsDB.UserID,
	}, nil
}

func (c NewsRepository) GetAllNews(metadata entities.Metadata) ([]entities.News, error) {
	var newsDB []models.News

	if err := c.DB.Limit(metadata.Limit).Offset(metadata.GetOffset()).Find(&newsDB).Error; err != nil {
		return []entities.News{}, errors.New("failed to get news")
	}

	var news []entities.News
	for _, newsDB := range newsDB {
		news = append(news, entities.News{
			ID:         newsDB.ID,
			Content:    newsDB.Content,
			CategoryID: newsDB.CategoryID,
			UserID:     newsDB.UserID,
		})
	}

	return news, nil
}

func (c NewsRepository) GetNewsByID(newsID uint) (entities.News, error) {
	var newsDB models.News

	if err := c.DB.Where("id = ?", newsID).First(&newsDB).Error; err != nil {
		return entities.News{}, utils.ErrInvalidNewsID
	}

	return entities.News{
		ID:         newsDB.ID,
		Content:    newsDB.Content,
		CategoryID: newsDB.CategoryID,
		UserID:     newsDB.UserID,
	}, nil
}

func (c NewsRepository) UpdateNews(news entities.News) (entities.News, error) {
	var newsDB models.News

	if err := c.DB.Where("id = ?", news.ID).First(&newsDB).Error; err != nil {
		return entities.News{}, utils.ErrInvalidNewsID
	}

	newsDB.Content = news.Content
	newsDB.CategoryID = news.CategoryID
	newsDB.UserID = news.UserID
	if err := c.DB.Save(&newsDB).Error; err != nil {
		return entities.News{}, errors.New("failed to update news")
	}

	return entities.News{
		ID:         newsDB.ID,
		Content:    newsDB.Content,
		CategoryID: newsDB.CategoryID,
		UserID:     newsDB.UserID,
	}, nil
}

func (c NewsRepository) DeleteNews(newsID uint) error {
	var newsDB models.News

	if err := c.DB.Where("id = ?", newsID).First(&newsDB).Error; err != nil {
		return utils.ErrInvalidNewsID
	}

	if err := c.DB.Delete(&newsDB).Error; err != nil {
		return errors.New("failed to delete news")
	}

	return nil
}