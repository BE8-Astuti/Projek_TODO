package main

import (
	"projek/todo/config"

	"projek/todo/delivery/routes"

	controllertask "projek/todo/delivery/controller/task"
	taskRepo "projek/todo/repository/task"

	controllerprojek "projek/todo/delivery/controller/projek"
	projekRepo "projek/todo/repository/projek"

	controllerus "projek/todo/delivery/controller/user"
	userRepo "projek/todo/repository/user"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func main() {
	// Get Access Database
	database := config.InitDB()
	// config.Migrate()

	userRepo := userRepo.New(database)
	userControl := controllerus.New(userRepo, validator.New())

	taskRepo := taskRepo.NewDB(database)
	tControl := controllertask.NewControlTask(taskRepo, validator.New())

	projekRepo := projekRepo.NewDB(database)
	pControl := controllerprojek.NewControlPro(projekRepo, validator.New())

	// Initiate Echo
	e := echo.New()
	// Akses Path Addressss
	routes.Path(e, userControl, tControl, pControl)
	e.Logger.Fatal(e.Start(":8005"))
}
