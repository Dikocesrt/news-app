package routes

import (
	"makanan-app/controllers"

	"github.com/labstack/echo/v4"
)

type Route struct {
	userController controllers.UserController
	categoryController controllers.CategoryController
}

func NewRoute(userController controllers.UserController, categoryController controllers.CategoryController) Route {
	return Route{
		userController: userController,
		categoryController: categoryController,
	}
}

func (r Route) InitializeRoute(e *echo.Echo) {
	auth := e.Group("/v1/auth")
	auth.POST("/register", r.userController.Register)
	auth.POST("/login", r.userController.Login)

	userRoute := e.Group("/v1")
	userRoute.POST("/categories", r.categoryController.CreateCategory)
	userRoute.GET("/categories", r.categoryController.GetAllCategories)
	userRoute.GET("/categories/:id", r.categoryController.GetCategoryByID)
	userRoute.PUT("/categories/:id", r.categoryController.UpdateCategory)
	userRoute.DELETE("/categories/:id", r.categoryController.DeleteCategory)
}