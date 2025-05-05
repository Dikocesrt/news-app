package controllers

import (
	"makanan-app/entities"
	"makanan-app/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CommentController struct {
	commentUsecase entities.CommentUsecaseInterface
}

func NewCommentController(commentUsecase entities.CommentUsecaseInterface) CommentController {
	return CommentController{
		commentUsecase: commentUsecase,
	}
}

type commentRequest struct {
	Name    string `json:"name" form:"name"`
	Comment string `json:"comment" form:"comment"`
	NewsID  uint   `json:"news_id" form:"news_id"`
}

func (c CommentController) CreateComment(ctx echo.Context) error {
	comment := commentRequest{}
	if err := ctx.Bind(&comment); err != nil {
		return err
	}

	_, err := c.commentUsecase.CreateComment(entities.Comment{
		Name:    comment.Name,
		Comment: comment.Comment,
		News: entities.News{
			ID: comment.NewsID,
		},
	})
	if err != nil {
		return ctx.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusCreated, utils.NewBaseSuccessResponse("success create comment", struct{}{}))
}
