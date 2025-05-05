package controllers

import (
	"errors"
	"makanan-app/entities"
	"makanan-app/utils"
	"net/http"
	"strconv"

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

type categoryResponse struct {
	ID   uint `json:"id"`
	Name string `json:"name"`
	Status string `json:"status"`
}

func (c CategoryController) CreateCategory(ctx echo.Context) error {
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

	categoryResponse := categoryResponse{
		ID: newCategory.ID,
	}

	return ctx.JSON(http.StatusCreated, utils.NewBaseSuccessResponse("success create category", categoryResponse))
}

func (c CategoryController) GetAllCategories(ctx echo.Context) error {
	page := ctx.QueryParam("page")
	limit := ctx.QueryParam("limit")

	metadata := entities.GetMetadata(page, limit)

	token := ctx.Request().Header.Get("Authorization")
	_, err := utils.GetIDFromToken(token)
	if err != nil {
		return ctx.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(err.Error()))
	}

	categories, err := c.categoryUsecase.GetAllCategories(metadata)
	if err != nil {
		return ctx.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(err.Error()))
	}

	var categoriesResponse []categoryResponse
	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, categoryResponse{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	return ctx.JSON(http.StatusOK, utils.NewBaseMetadataSuccessResponse("success get all categories", metadata, categoriesResponse))
}

func (c CategoryController) GetCategoryByID(ctx echo.Context) error {
	IDParam := ctx.Param("id")
	id, _ := strconv.Atoi(IDParam)

	token := ctx.Request().Header.Get("Authorization")
	_, err := utils.GetIDFromToken(token)
	if err != nil {
		return ctx.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(err.Error()))
	}

	category, err := c.categoryUsecase.GetCategoryByID(uint(id))
	if err != nil {
		return ctx.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(err.Error()))
	}

	categoryResponse := categoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}

	return ctx.JSON(http.StatusOK, utils.NewBaseSuccessResponse("success get category by id", categoryResponse))
}

func (c CategoryController) UpdateCategory(ctx echo.Context) error {
	IDParam := ctx.Param("id")
	id, _ := strconv.Atoi(IDParam)

	token := ctx.Request().Header.Get("Authorization")
	_, err := utils.GetIDFromToken(token)
	if err != nil {
		return ctx.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(err.Error()))
	}

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

	return ctx.JSON(http.StatusOK, utils.NewBaseSuccessResponse("success update category", categoryResponse{}))
}

func (c CategoryController) DeleteCategory(ctx echo.Context) error {
	IDParam := ctx.Param("id")
	id, _ := strconv.Atoi(IDParam)

	token := ctx.Request().Header.Get("Authorization")
	_, err := utils.GetIDFromToken(token)
	if err != nil {
		return ctx.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(err.Error()))
	}

	err = c.categoryUsecase.DeleteCategory(uint(id))
	if err != nil {
		return ctx.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, utils.NewBaseSuccessResponse("success delete category", categoryResponse{}))
}
