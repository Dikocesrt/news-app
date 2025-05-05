package routes

import (
	"makanan-app/controllers"

	"github.com/labstack/echo/v4"
)

type Route struct {
	userController controllers.UserController
}

func NewRoute(userController controllers.UserController) Route {
	return Route{
		userController: userController,
	}
}

func (r Route) InitializeRoute(e *echo.Echo) {
	auth := e.Group("/auth")
	auth.POST("/register", r.userController.Register)
	auth.POST("/login", r.userController.Login)
}