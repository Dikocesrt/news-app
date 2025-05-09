package repositories

import (
	"errors"
	"test-indonesia-cakap-digital/entities"
	"test-indonesia-cakap-digital/models"
	"test-indonesia-cakap-digital/utils"

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
	newsDB := models.NewNews(news.Content, news.Category.ID, news.User.ID)

	if err := c.DB.Create(&newsDB).Error; err != nil {
		return entities.News{}, errors.New("failed to create news")
	}

	return entities.News{
		ID:         newsDB.ID,
		Content:    newsDB.Content,
		Category: entities.Category{
			ID: newsDB.CategoryID,
		},
		User: entities.User{
			ID: newsDB.UserID,
		},
	}, nil
}

func (c NewsRepository) GetAllNews(metadata entities.Metadata) ([]entities.News, error) {
	var newsDB []models.News

	if err := c.DB.Preload("Category").Preload("User").Limit(metadata.Limit).Offset(metadata.GetOffset()).Find(&newsDB).Error; err != nil {
		return []entities.News{}, errors.New("failed to get news")
	}

	var news []entities.News
	for _, n := range newsDB {
		var latestComments []models.Comment
		if err := c.DB.Where("news_id = ?", n.ID).Limit(3).Order("created_at desc").Find(&latestComments).Error; err != nil {
			return []entities.News{}, errors.New("failed to get latest comments")
		}

		var comments []entities.Comment
		for _, c := range latestComments {
			comments = append(comments, entities.Comment{
				ID:      c.ID,
				Name:    c.Name,
				Comment: c.Comment,
			})
		}

		news = append(news, entities.News{
			ID:         n.ID,
			Content:    n.Content,
			Category: entities.Category{
				ID: n.CategoryID,
				Name: n.Category.Name,
			},
			User: entities.User{
				ID: n.UserID,
				Username: n.User.Username,
			},
			Comments: comments,
		})
	}

	return news, nil
}

func (c NewsRepository) GetNewsByID(newsID uint) (entities.News, error) {
	var newsDB models.News

	if err := c.DB.Where("id = ?", newsID).Preload("Category").Preload("User").First(&newsDB).Error; err != nil {
		return entities.News{}, utils.ErrInvalidNewsID
	}

	var commentsDB []models.Comment
	if err := c.DB.Where("news_id = ?", newsID).Limit(10).Order("created_at asc").Find(&commentsDB).Error; err != nil {
		return entities.News{}, errors.New("failed to get comments")
	}

	var comments []entities.Comment
	for _, c := range commentsDB {
		comments = append(comments, entities.Comment{
			ID:      c.ID,
			Name:    c.Name,
			Comment: c.Comment,
		})
	}

	var newsTagsDB []models.NewsTag
	if err := c.DB.Where("news_id = ?", newsID).Preload("Tag").Find(&newsTagsDB).Error; err != nil {
		return entities.News{}, errors.New("failed to get tags")
	}

	var tags []entities.Tag
	for _, t := range newsTagsDB {
		tags = append(tags, entities.Tag{
			ID:   t.TagID,
			Name: t.Tag.Name,
		})
	}

	return entities.News{
		ID:         newsDB.ID,
		Content:    newsDB.Content,
		Category: entities.Category{
			ID: newsDB.CategoryID,
			Name: newsDB.Category.Name,
		},
		User: entities.User{
			ID: newsDB.UserID,
			Username: newsDB.User.Username,
		},
		Tags: tags,
		Comments: comments,
	}, nil
}

func (c NewsRepository) UpdateNews(news entities.News) (entities.News, error) {
	var newsDB models.News

	if err := c.DB.Where("id = ?", news.ID).First(&newsDB).Error; err != nil {
		return entities.News{}, utils.ErrInvalidNewsID
	}

	if newsDB.UserID != news.User.ID {
		return entities.News{}, utils.ErrUnauthorized
	}

	newsDB.Content = news.Content
	newsDB.CategoryID = news.Category.ID
	newsDB.UserID = news.User.ID
	if err := c.DB.Save(&newsDB).Error; err != nil {
		return entities.News{}, errors.New("failed to update news")
	}

	if err := c.DB.Where("news_id = ?", news.ID).Delete(&models.NewsTag{}).Error; err != nil {
		return entities.News{}, errors.New("failed to delete old tags")
	}

	for _, tag := range news.Tags {
		newsTag := models.NewsTag{
			NewsID: news.ID,
			TagID:  tag.ID,
		}
		if err := c.DB.Create(&newsTag).Error; err != nil {
			return entities.News{}, errors.New("failed to associate new tags")
		}
	}

	return entities.News{
		ID:         newsDB.ID,
		Content:    newsDB.Content,
		Category: entities.Category{
			ID: newsDB.CategoryID,
		},
		User: entities.User{
			ID: newsDB.UserID,
		},
		Tags: news.Tags,
	}, nil
}

func (c NewsRepository) DeleteNews(newsID uint, userID uint) error {
	var newsDB models.News

	if err := c.DB.Where("id = ?", newsID).First(&newsDB).Error; err != nil {
		return utils.ErrInvalidNewsID
	}

	if newsDB.UserID != userID {
		return utils.ErrUnauthorized
	}

	if err := c.DB.Delete(&newsDB).Error; err != nil {
		return errors.New("failed to delete news")
	}

	return nil
}