package usecase

import (
	"github.com/Ryoga-88/Todo-PJ/backend/entity"
	"github.com/Ryoga-88/Todo-PJ/backend/repository"
	"github.com/Ryoga-88/Todo-PJ/backend/validator"
)

type ITaskUsecase interface {
	GetAllTasks(userId uint) ([]entity.TaskResponse, error)
	GetTaskById(userId uint, taskId uint) (entity.TaskResponse, error)
	CreateTask(task entity.Task) (entity.TaskResponse, error)
	UpdateTask(task entity.Task, userId uint, taskId uint) (entity.TaskResponse, error)
	DeleteTask(userId uint, taskId uint) error
}

type taskUsecase struct {
	tr repository.ITaskRepository
	tv validator.ITaskValidator
}

func NewTaskUsecase(tr repository.ITaskRepository, tv validator.ITaskValidator) ITaskUsecase {
	return &taskUsecase{tr, tv}
}

func (tu *taskUsecase) GetAllTasks(userId uint) ([]entity.TaskResponse, error) {
	tasks := []entity.Task{}
	if err := tu.tr.GetAllTasks(&tasks, userId); err != nil {
		return nil, err
	}
	resTasks := []entity.TaskResponse{}
	for _, v := range tasks {
		t := entity.TaskResponse{
			ID:        v.ID,
			Title:     v.Title,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
		resTasks = append(resTasks, t)
	}
	return resTasks, nil
}

func (tu *taskUsecase) GetTaskById(userId uint, taskId uint) (entity.TaskResponse, error) {
	task := entity.Task{}
	if err := tu.tr.GetTaskById(&task, userId, taskId); err != nil {
		return entity.TaskResponse{}, err
	}
	resTask := entity.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return resTask, nil
}

func (tu *taskUsecase) CreateTask(task entity.Task) (entity.TaskResponse, error) {
	if err := tu.tv.TaskValidate(task); err != nil {
		return entity.TaskResponse{}, err
	}

	if err := tu.tr.CreateTask(&task); err != nil {
		return entity.TaskResponse{}, err
	}
	resTask := entity.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return resTask, nil
}

func (tu *taskUsecase) UpdateTask(task entity.Task, userId uint, taskId uint) (entity.TaskResponse, error) {
	if err := tu.tv.TaskValidate(task); err != nil {
		return entity.TaskResponse{}, err
	}

	if err := tu.tr.UpdateTask(&task, userId, taskId); err != nil {
		return entity.TaskResponse{}, err
	}
	resTask := entity.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}
	return resTask, nil
}

func (tu *taskUsecase) DeleteTask(userId uint, taskId uint) error {
	if err := tu.tr.DeleteTask(userId, taskId); err != nil {
		return err
	}
	return nil
}
