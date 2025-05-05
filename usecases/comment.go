package usecases

import (
	"makanan-app/entities"
	"makanan-app/utils"
)

type CommentUsecase struct {
	commentRepository entities.CommentRepositoryInterface
}

func NewCommentUsecase(commentRepository entities.CommentRepositoryInterface) CommentUsecase {
	return CommentUsecase{
		commentRepository: commentRepository,
	}
}

func (c CommentUsecase) CreateComment(comment entities.Comment) (entities.Comment, error) {
	if comment.Comment == "" || comment.News.ID == 0 {
		return entities.Comment{}, utils.ErrEmptyField
	}

	return c.commentRepository.CreateComment(comment)
}