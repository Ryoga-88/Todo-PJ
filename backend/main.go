package main

import (
	"log"

	"github.com/Ryoga-88/Todo-PJ/backend/config"
	"github.com/Ryoga-88/Todo-PJ/backend/controller"
	"github.com/Ryoga-88/Todo-PJ/backend/repository"
	"github.com/Ryoga-88/Todo-PJ/backend/router"
	"github.com/Ryoga-88/Todo-PJ/backend/usecase"
	"github.com/Ryoga-88/Todo-PJ/backend/validator"
)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatalf("failed to read config: %v\n", err)
	}

	db, err := repository.NewDB(config.Conf)
	if err != nil {
		log.Fatalln(err)
	}

	if err := repository.Migrate(db); err != nil {
		log.Fatalln(err)
	}

	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()

	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)

	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)

	e := router.NewRouter(userController, taskController)

	e.Logger.Fatal(e.Start(":8080"))
}
