package main

import (
	"test-indonesia-cakap-digital/configs"
	"test-indonesia-cakap-digital/controllers"
	"test-indonesia-cakap-digital/repositories"
	"test-indonesia-cakap-digital/routes"
	"test-indonesia-cakap-digital/usecases"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.LoadEnv()

	db := configs.ConnectDB()

	userRepositories := repositories.NewUserRepository(db)
	userUsecases := usecases.NewUserUsecase(userRepositories)
	userController := controllers.NewUserController(userUsecases)

	categoryRepositories := repositories.NewCategoryRepository(db)
	categoryUsecases := usecases.NewCategoryUsecase(categoryRepositories)
	categoryController := controllers.NewCategoryController(categoryUsecases)

	newsRepositories := repositories.NewNewsRepository(db)
	newsUsecases := usecases.NewNewsUsecase(newsRepositories)
	newsController := controllers.NewNewsController(newsUsecases)

	commentRepositories := repositories.NewCommentRepository(db)
	commentUsecases := usecases.NewCommentUsecase(commentRepositories)
	commentController := controllers.NewCommentController(commentUsecases)

	customPageRepositories := repositories.NewCustomPageRepository(db)
	customPageUsecases := usecases.NewCustomPageUsecase(customPageRepositories)
	customPageController := controllers.NewCustomPageController(customPageUsecases)

	tagRepositories := repositories.NewTagRepository(db)
	tagUsecases := usecases.NewTagUsecase(tagRepositories)
	tagController := controllers.NewTagController(tagUsecases)

	route := routes.NewRoute(userController, categoryController, newsController, commentController, customPageController, tagController)

	e := echo.New()
	route.InitializeRoute(e)
	e.Logger.Fatal(e.Start(":8080"))
}