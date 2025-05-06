package controllers

import (
	"errors"
	"net/http"
	"test-indonesia-cakap-digital/entities"
	"test-indonesia-cakap-digital/utils"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUsecase entities.UserUsecaseInterface
}

func NewUserController(userUsecase entities.UserUsecaseInterface) UserController {
	return UserController{
		userUsecase: userUsecase,
	}
}

type userRequest struct {
	Password string `json:"password" form:"password"`
	Username string `json:"username" form:"username"`
}

type userResponseToken struct {
	Token string `json:"token"`
}

func (u UserController) Register(c echo.Context) error {
	var req userRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(errors.New("internal server error").Error()))
	}

	user := entities.User{
		Username: req.Username,
		Password: req.Password,
	}

	newUser, err := u.userUsecase.Register(user)
	if err != nil {
		return c.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(err.Error()))
	}

	userResponse := userResponseToken{
		Token: newUser.Token,
	}

	return c.JSON(http.StatusCreated, utils.NewBaseSuccessResponse("success register user", userResponse))
}

func (u UserController) Login(c echo.Context) error {
	var req userRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(errors.New("internal server error").Error()))
	}

	user := entities.User{
		Username: req.Username,
		Password: req.Password,
	}

	userLogin, err := u.userUsecase.Login(user)
	if err != nil {
		return c.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(err.Error()))
	}

	userResponse := userResponseToken{
		Token: userLogin.Token,
	}

	return c.JSON(http.StatusOK, utils.NewBaseSuccessResponse("success login user", userResponse))
}