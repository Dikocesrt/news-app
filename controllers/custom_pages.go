package controllers

import (
	"errors"
	"makanan-app/entities"
	"makanan-app/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CustomPageController struct {
	customPageUsecase entities.CustomPageUsecaseInterface
}

func NewCustomPageController(customPageUsecase entities.CustomPageUsecaseInterface) CustomPageController {
	return CustomPageController{
		customPageUsecase: customPageUsecase,
	}
}

type CustomPageRequest struct {
	CustomURL string `json:"custom_url"`
	Content string `json:"content"`
}

type CustomPageResponseID struct {
	ID uint `json:"id"`
}

type CustomPageResponse struct {
	ID uint `json:"id"`
	CustomURL string `json:"custom_url"`
	Content string `json:"content"`
	User UserReponse `json:"user"`
}

func (c CustomPageController) CreateCustomPage(ctx echo.Context) error {
	token := ctx.Request().Header.Get("Authorization")
	if token == "" {
		return ctx.JSON(http.StatusUnauthorized, utils.NewBaseErrorResponse("unauthorized"))
	}
	userID, err := utils.GetIDFromToken(token)
	if err != nil {
		return ctx.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(err.Error()))
	}

	req := CustomPageRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(errors.New("internal server error").Error()))
	}

	customPage := entities.CustomPage{
		User: entities.User{ID: userID},
		CustomURL: req.CustomURL,
		Content: req.Content,
	}

	newCustomPage, err := c.customPageUsecase.CreateCustomPage(customPage)
	if err != nil {
		return ctx.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(err.Error()))
	}

	customPageResponse := CustomPageResponseID{
		ID: newCustomPage.ID,
	}

	return ctx.JSON(http.StatusCreated, utils.NewBaseSuccessResponse("success create custom page", customPageResponse))
}

func (c CustomPageController) GetAllCustomPages(ctx echo.Context) error {
	page := ctx.QueryParam("page")
	limit := ctx.QueryParam("limit")

	metadata := entities.GetMetadata(page, limit)

	customPages, err := c.customPageUsecase.GetAllCustomPages(metadata)
	if err != nil {
		return ctx.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(err.Error()))
	}

	var customPagesResponse []CustomPageResponse
	for _, customPage := range customPages {
		customPagesResponse = append(customPagesResponse, CustomPageResponse{
			ID: customPage.ID,
			CustomURL: customPage.CustomURL,
			Content: customPage.Content,
			User: UserReponse{
				ID: customPage.User.ID,
				Username: customPage.User.Username,
			},
		})
	}

	return ctx.JSON(http.StatusOK, utils.NewBaseMetadataSuccessResponse("success get all custom pages", metadata, customPagesResponse))
}

func (c CustomPageController) GetCustomPageByID(ctx echo.Context) error {
	IDParam := ctx.Param("id")
	id, _ := strconv.Atoi(IDParam)

	customPage, err := c.customPageUsecase.GetCustomPageByID(uint(id))
	if err != nil {
		return ctx.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(err.Error()))
	}

	customPageResponse := CustomPageResponse{
		ID: customPage.ID,
		CustomURL: customPage.CustomURL,
		Content: customPage.Content,
		User: UserReponse{
			ID: customPage.User.ID,
			Username: customPage.User.Username,
		},
	}

	return ctx.JSON(http.StatusOK, utils.NewBaseSuccessResponse("success get custom page by id", customPageResponse))
}

func (c CustomPageController) UpdateCustomPage(ctx echo.Context) error {
	token := ctx.Request().Header.Get("Authorization")
	if token == "" {
		return ctx.JSON(http.StatusUnauthorized, utils.NewBaseErrorResponse("unauthorized"))
	}
	userID, err := utils.GetIDFromToken(token)
	if err != nil {
		return ctx.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(err.Error()))
	}

	IDParam := ctx.Param("id")
	id, _ := strconv.Atoi(IDParam)

	req := CustomPageRequest{}
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(errors.New("internal server error").Error()))
	}

	customPage := entities.CustomPage{
		ID: uint(id),
		CustomURL: req.CustomURL,
		Content: req.Content,
		User: entities.User{
			ID: userID,
		},
	}

	_, err = c.customPageUsecase.UpdateCustomPage(customPage)
	if err != nil {
		return ctx.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, utils.NewBaseSuccessResponse("success update custom page", struct{}{}))
}

func (c CustomPageController) DeleteCustomPage(ctx echo.Context) error {
	token := ctx.Request().Header.Get("Authorization")
	if token == "" {
		return ctx.JSON(http.StatusUnauthorized, utils.NewBaseErrorResponse("unauthorized"))
	}
	userID, err := utils.GetIDFromToken(token)
	if err != nil {
		return ctx.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(err.Error()))
	}

	IDParam := ctx.Param("id")
	id, _ := strconv.Atoi(IDParam)

	err = c.customPageUsecase.DeleteCustomPage(uint(id), userID)
	if err != nil {
		return ctx.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(err.Error()))
	}

	return ctx.JSON(http.StatusOK, utils.NewBaseSuccessResponse("success delete custom page", struct{}{}))
}