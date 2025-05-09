package routes

import (
	"test-indonesia-cakap-digital/controllers"

	"github.com/labstack/echo/v4"
)

type Route struct {
	userController controllers.UserController
	categoryController controllers.CategoryController
	newsController controllers.NewsController
	commentController controllers.CommentController
	customPageController controllers.CustomPageController
	tagController controllers.TagController
}

func NewRoute(userController controllers.UserController, categoryController controllers.CategoryController, newsController controllers.NewsController, commentController controllers.CommentController, customPageController controllers.CustomPageController, tagController controllers.TagController) Route {
	return Route{
		userController: userController,
		categoryController: categoryController,
		newsController: newsController,
		commentController: commentController,
		customPageController: customPageController,
		tagController: tagController,
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

	Route.POST("/custom-pages", r.customPageController.CreateCustomPage)
	Route.GET("/custom-pages", r.customPageController.GetAllCustomPages)
	Route.GET("/custom-pages/:id", r.customPageController.GetCustomPageByID)
	Route.PUT("/custom-pages/:id", r.customPageController.UpdateCustomPage)
	Route.DELETE("/custom-pages/:id", r.customPageController.DeleteCustomPage)

	Route.POST("/tags", r.tagController.CreateTag)
}