package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"test-indonesia-cakap-digital/entities"
	"test-indonesia-cakap-digital/utils"

	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	categoryUsecase entities.CategoryUsecaseInterface
}

func NewCategoryController(categoryUsecase entities.CategoryUsecaseInterface) CategoryController {
	return CategoryController{
		categoryUsecase: categoryUsecase,
	}
}

type categoryRequest struct {
	Name string `json:"name" form:"name"`
}

type categoryResponseID struct {
	ID   uint `json:"id"`
}

type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (c CategoryController) CreateCategory(ctx echo.Context) error {
	token := ctx.Request().Header.Get("Authorization")
	if token == "" {
		return ctx.JSON(http.StatusUnauthorized, utils.NewBaseErrorResponse("unauthorized"))
	}
	_, err := utils.GetIDFromToken(token)
	if err != nil {
		return ctx.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(err.Error()))
	}

	req := categoryRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(errors.New("internal server error").Error()))
	}

	category := entities.Category{
		Name: req.Name,
	}

	newCategory, err := c.categoryUsecase.CreateCategory(category)
	if err != nil {
		return ctx.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(err.Error()))
	}

	categoryResponse := categoryResponseID{
		ID: newCategory.ID,
	}

	return ctx.JSON(http.StatusCreated, utils.NewBaseSuccessResponse("success create category", categoryResponse))
}

func (c CategoryController) GetAllCategories(ctx echo.Context) error {
	token := ctx.Request().Header.Get("Authorization")
	if token == "" {
		return ctx.JSON(http.StatusUnauthorized, utils.NewBaseErrorResponse("unauthorized"))
	}
	_, err := utils.GetIDFromToken(token)
	if err != nil {
		return ctx.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(err.Error()))
	}

	page := ctx.QueryParam("page")
	limit := ctx.QueryParam("limit")

	metadata := entities.GetMetadata(page, limit)

	categories, err := c.categoryUsecase.GetAllCategories(metadata)
	if err != nil {
		return ctx.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(err.Error()))
	}

	var categoriesResponse []CategoryResponse
	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, CategoryResponse{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	return ctx.JSON(http.StatusOK, utils.NewBaseMetadataSuccessResponse("success get all categories", metadata, categoriesResponse))
}

func (c CategoryController) GetCategoryByID(ctx echo.Context) error {
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

	category, err := c.categoryUsecase.GetCategoryByID(uint(id))
	if err != nil {
		return ctx.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(err.Error()))
	}

	categoryResponse := CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}

	return ctx.JSON(http.StatusOK, utils.NewBaseSuccessResponse("success get category by id", categoryResponse))
}

func (c CategoryController) UpdateCategory(ctx echo.Context) error {
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

	req := categoryRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(errors.New("internal server error").Error()))
	}

	category := entities.Category{
		ID:   uint(id),
		Name: req.Name,
	}

	_, err = c.categoryUsecase.UpdateCategory(category)
	if err != nil {
		return ctx.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, utils.NewBaseSuccessResponse("success update category", struct{}{}))
}

func (c CategoryController) DeleteCategory(ctx echo.Context) error {
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

	err = c.categoryUsecase.DeleteCategory(uint(id))
	if err != nil {
		return ctx.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, utils.NewBaseSuccessResponse("success delete category", struct{}{}))
}
