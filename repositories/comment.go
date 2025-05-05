package repositories

import (
	"errors"
	"makanan-app/entities"
	"makanan-app/models"
	"makanan-app/utils"

	"gorm.io/gorm"
)

type CommentRepository struct {
	DB *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return CommentRepository{
		DB: db,
	}
}

func (c CommentRepository) CreateComment(comment entities.Comment) (entities.Comment, error) {
	commentDB := models.NewComment(comment.Name, comment.Comment, comment.News.ID)

	if err := c.DB.Where("id = ?", comment.News.ID).First(&models.News{}).Error; err != nil {
		return entities.Comment{}, utils.ErrInvalidNewsID
	}

	if err := c.DB.Create(&commentDB).Error; err != nil {
		return entities.Comment{}, errors.New("failed to create comment")
	}

	return entities.Comment{
		ID:      commentDB.ID,
		Name:    commentDB.Name,
		Comment: commentDB.Comment,
		News: entities.News{
			ID: commentDB.NewsID,
		},
	}, nil
}