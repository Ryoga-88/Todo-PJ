package controller

import (
	"net/http"
	"strconv"

	"github.com/Ryoga-88/Todo-PJ/backend/entity"
	"github.com/Ryoga-88/Todo-PJ/backend/usecase"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type ITaskController interface {
	GetAllTasks(c echo.Context) error
	GetTaskById(c echo.Context) error
	CreateTask(c echo.Context) error
	UpdateTask(c echo.Context) error
	DeleteTask(c echo.Context) error
}

type taskController struct {
	tu usecase.ITaskUsecase
}

func NewTaskController(tu usecase.ITaskUsecase) ITaskController {
	return &taskController{tu}
}

func (tc *taskController) GetAllTasks(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	tasksRes, err := tc.tu.GetAllTasks(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, tasksRes)
}

func (tc *taskController) GetTaskById(e echo.Context) error {
	user := e.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	id := e.Param("taskId")
	taskId, _ := strconv.Atoi(id)
	taskRes, err := tc.tu.GetTaskById(uint(userId.(float64)), uint(taskId))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, taskRes)
}

func (tc *taskController) CreateTask(e echo.Context) error {
	user := e.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	task := entity.Task{}
	if err := e.Bind(&task); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	task.UserID = uint(userId.(float64))
	taskRes, err := tc.tu.CreateTask(task)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusCreated, taskRes)
}

func (tc *taskController) UpdateTask(e echo.Context) error {
	user := e.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	id := e.Param("taskId")
	taskId, _ := strconv.Atoi(id)
	task := entity.Task{}
	if err := e.Bind(&task); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}
	taskRes, err := tc.tu.UpdateTask(task, uint(userId.(float64)), uint(taskId))
	if err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, taskRes)
}

func (tc *taskController) DeleteTask(e echo.Context) error {
	user := e.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	id := e.Param("taskId")
	taskId, _ := strconv.Atoi(id)
	if err := tc.tu.DeleteTask(uint(userId.(float64)), uint(taskId)); err != nil {
		return e.JSON(http.StatusInternalServerError, err.Error())
	}
	return e.NoContent(http.StatusOK)
}
