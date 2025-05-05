package routes

import (
	"makanan-app/controllers"

	"github.com/labstack/echo/v4"
)

type Route struct {
	userController controllers.UserController
	categoryController controllers.CategoryController
	newsController controllers.NewsController
	commentController controllers.CommentController
	customPageController controllers.CustomPageController
}

func NewRoute(userController controllers.UserController, categoryController controllers.CategoryController, newsController controllers.NewsController, commentController controllers.CommentController, customPageController controllers.CustomPageController) Route {
	return Route{
		userController: userController,
		categoryController: categoryController,
		newsController: newsController,
		commentController: commentController,
		customPageController: customPageController,
	}
}

func (r Route) InitializeRoute(e *echo.Echo) {
	auth := e.Group("/v1/auth")
	auth.POST("/register", r.userController.Register)
	auth.POST("/login", r.userController.Login)

	Route := e.Group("/v1")
	Route.POST("/categories", r.categoryController.CreateCategory)
	Route.GET("/categories", r.categoryController.GetAllCategories)
	Route.GET("/categories/:id", r.categoryController.GetCategoryByID)
	Route.PUT("/categories/:id", r.categoryController.UpdateCategory)
	Route.DELETE("/categories/:id", r.categoryController.DeleteCategory)

	Route.POST("/news", r.newsController.CreateNews)
	Route.GET("/news", r.newsController.GetAllNews)
	Route.GET("/news/:id", r.newsController.GetNewsByID)
	Route.PUT("/news/:id", r.newsController.UpdateNews)
	Route.DELETE("/news/:id", r.newsController.DeleteNews)

	Route.POST("/comments", r.commentController.CreateComment)

	Route.POST("/custom-page", r.customPageController.CreateCustomPage)
	Route.GET("/custom-page", r.customPageController.GetAllCustomPages)
	Route.GET("/custom-page/:id", r.customPageController.GetCustomPageByID)
	Route.PUT("/custom-page/:id", r.customPageController.UpdateCustomPage)
	Route.DELETE("/custom-page/:id", r.customPageController.DeleteCustomPage)
}