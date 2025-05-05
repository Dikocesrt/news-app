package controllers

import (
	"errors"
	"makanan-app/entities"
	"makanan-app/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type NewsController struct {
	newsUsecase entities.NewsUsecaseInterface
}

func NewNewsController(newsUsecase entities.NewsUsecaseInterface) NewsController {
	return NewsController{
		newsUsecase: newsUsecase,
	}
}

type newsRequest struct {
	Content     string `json:"content"`
	CategoryID  uint   `json:"category_id"`
}

type newsResponseID struct {
	ID uint `json:"id"`
}

type newsResponse struct {
	ID         uint   `json:"id"`
	Category CategoryResponse `json:"category"`
	User       userReponse `json:"user"`
	Content    string `json:"content"`
	Comment    []CommentResponse `json:"comments"`
}

type userReponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

type CommentResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

func (c NewsController) CreateNews(ctx echo.Context) error {
	token := ctx.Request().Header.Get("Authorization")
	if token == "" {
		return ctx.JSON(http.StatusUnauthorized, utils.NewBaseErrorResponse("unauthorized"))
	}
	userID, err := utils.GetIDFromToken(token)
	if err != nil {
		return ctx.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(err.Error()))
	}

	req := newsRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(errors.New("internal server error").Error()))
	}

	news := entities.News{
		User:       entities.User{ID: userID},
		Content:    req.Content,
		Category:   entities.Category{ID: req.CategoryID},
	}

	newNews, err := c.newsUsecase.CreateNews(news)
	if err != nil {
		return ctx.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(err.Error()))
	}

	newsResponse := newsResponseID{
		ID: newNews.ID,
	}

	return ctx.JSON(http.StatusCreated, utils.NewBaseSuccessResponse("success create news", newsResponse))
}

func (c NewsController) GetAllNews(ctx echo.Context) error {
	page := ctx.QueryParam("page")
	limit := ctx.QueryParam("limit")

	metadata := entities.GetMetadata(page, limit)

	news, err := c.newsUsecase.GetAllNews(metadata)
	if err != nil {
		return ctx.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(err.Error()))
	}

	var newssResponse []newsResponse
	for _, new := range news {
		var comments []CommentResponse
		for _, comment := range new.Comments {
			comments = append(comments, CommentResponse{
				ID:      comment.ID,
				Name:    comment.Name,
				Comment: comment.Comment,
			})
		}

		newssResponse = append(newssResponse, newsResponse{
			ID:         new.ID,
			Content:    new.Content,
			Category: CategoryResponse{
				ID:   new.Category.ID,
				Name: new.Category.Name,
			},
			User: userReponse{
				ID:       new.User.ID,
				Username: new.User.Username,
			},
			Comment: comments,
		})
	}

	return ctx.JSON(http.StatusOK, utils.NewBaseMetadataSuccessResponse("success get all news", metadata, newssResponse))
}

func (c NewsController) GetNewsByID(ctx echo.Context) error {
	IDParam := ctx.Param("id")
	id, _ := strconv.Atoi(IDParam)

	news, err := c.newsUsecase.GetNewsByID(uint(id))
	if err != nil {
		return ctx.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(err.Error()))
	}

	var comments []CommentResponse
	for _, comment := range news.Comments {
		comments = append(comments, CommentResponse{
			ID:      comment.ID,
			Name:    comment.Name,
			Comment: comment.Comment,
		})
	}

	newsResponse := newsResponse{
		ID:         news.ID,
		Content:    news.Content,
		Category: CategoryResponse{
			ID:   news.Category.ID,
			Name: news.Category.Name,
		},
		User: userReponse{
			ID:       news.User.ID,
			Username: news.User.Username,
		},
		Comment: comments,
	}

	return ctx.JSON(http.StatusOK, utils.NewBaseSuccessResponse("success get news by id", newsResponse))
}

func (c NewsController) UpdateNews(ctx echo.Context) error {
	token := ctx.Request().Header.Get("Authorization")
	if token == "" {
		return ctx.JSON(http.StatusUnauthorized, utils.NewBaseErrorResponse("unauthorized"))
	}
	_, err := utils.GetIDFromToken(token)
	if err != nil {
		return ctx.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(err.Error()))
	}

	IDParam := ctx.Param("id")
	id, _ := strconv.Atoi(IDParam)

	req := newsRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(errors.New("internal server error").Error()))
	}

	news := entities.News{
		ID:         uint(id),
		Content:    req.Content,
		Category: entities.Category{
			ID: req.CategoryID,
		},
	}

	_, err = c.newsUsecase.UpdateNews(news)
	if err != nil {
		return ctx.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, utils.NewBaseSuccessResponse("success update news", struct{}{}))
}

func (c NewsController) DeleteNews(ctx echo.Context) error {
	token := ctx.Request().Header.Get("Authorization")
	if token == "" {
		return ctx.JSON(http.StatusUnauthorized, utils.NewBaseErrorResponse("unauthorized"))
	}
	_, err := utils.GetIDFromToken(token)
	if err != nil {
		return ctx.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(err.Error()))
	}

	IDParam := ctx.Param("id")
	id, _ := strconv.Atoi(IDParam)

	err = c.newsUsecase.DeleteNews(uint(id))
	if err != nil {
		return ctx.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, utils.NewBaseSuccessResponse("success delete news", struct{}{}))
}