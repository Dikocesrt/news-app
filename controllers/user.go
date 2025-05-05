package controllers

import (
	"errors"
	"makanan-app/entities"
	"makanan-app/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUsecase entities.UsecaseInterface
}

func NewUserController(userUsecase entities.UsecaseInterface) UserController {
	return UserController{
		userUsecase: userUsecase,
	}
}

type userRequest struct {
	Password string `json:"password" form:"password"`
	Username string `json:"username" form:"username"`
}

type userResponse struct {
	Token string `json:"token"`
}

func (u UserController) Register(c echo.Context) error {
	var req userRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(utils.ConvertErrorCode(err), errors.New("internal server error"))
	}

	user := entities.User{
		Username: req.Username,
		Password: req.Password,
	}

	newUser, err := u.userUsecase.Register(user)
	if err != nil {
		return c.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(err.Error()))
	}

	userResponse := userResponse{
		Token: newUser.Token,
	}

	return c.JSON(http.StatusCreated, utils.NewBaseSuccessResponse("success register user", userResponse))
}

func (u UserController) Login(c echo.Context) error {
	var req userRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(utils.ConvertErrorCode(err), errors.New("internal server error"))
	}

	user := entities.User{
		Username: req.Username,
		Password: req.Password,
	}

	userLogin, err := u.userUsecase.Login(user)
	if err != nil {
		return c.JSON(utils.ConvertErrorCode(err), utils.NewBaseErrorResponse(err.Error()))
	}

	userResponse := userResponse{
		Token: userLogin.Token,
	}

	return c.JSON(http.StatusOK, utils.NewBaseSuccessResponse("success login user", userResponse))
}