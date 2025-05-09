package controllers

import (
	"errors"
	"net/http"
	"test-indonesia-cakap-digital/entities"
	"test-indonesia-cakap-digital/utils"

	"github.com/labstack/echo/v4"
)

type TagController struct {
	tagUsecase entities.TagUsecaseInterface
}

func NewTagController(tagUsecase entities.TagUsecaseInterface) TagController {
	return TagController{
		tagUsecase: tagUsecase,
	}
}

type tagRequest struct {
	Name string `json:"name" form:"name"`
}

type tagResponseID struct {
	ID uint `json:"id"`
}

func (tagController TagController) CreateTag(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	if token == "" {
		return c.JSON(http.StatusUnauthorized, utils.NewBaseErrorResponse("unauthorized"))
	}
	_, err := utils.GetIDFromToken(token)
	if err != nil {
		return c.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(err.Error()))
	}

	req := tagRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(errors.New("internal server error").Error()))
	}

	tag := entities.Tag{
		Name: req.Name,
	}

	newTag, err := tagController.tagUsecase.CreateTag(tag)
	if err != nil {
		return c.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(err.Error()))
	}

	tagResponse := tagResponseID{
		ID: newTag.ID,
	}

	return c.JSON(http.StatusCreated, utils.NewBaseSuccessResponse("success create tag", tagResponse))
}