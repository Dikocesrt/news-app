package main

import (
	"makanan-app/configs"
	"makanan-app/controllers"
	"makanan-app/repositories"
	"makanan-app/routes"
	"makanan-app/usecases"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.LoadEnv()

	db := configs.ConnectDB()

	userRepositories := repositories.NewUserRepository(db)
	userUsecases := usecases.NewUserUsecase(userRepositories)
	userController := controllers.NewUserController(userUsecases)

	route := routes.NewRoute(userController)

	e := echo.New()
	route.InitializeRoute(e)
	e.Logger.Fatal(e.Start(":8080"))
}