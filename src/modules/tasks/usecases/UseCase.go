package usecases

import (
	"errors"

	"github.com/juheth/Go-Clean-Arquitecture/src/modules/tasks/domain/entities"
	"github.com/juheth/Go-Clean-Arquitecture/src/modules/tasks/domain/repository"
)

type UCTaskUseCase interface {
	ExecuteCreateTask(task *entities.Task) error
	ExecuteGetAllTask() ([]*entities.Task, error)
	ExecuteGetTaskByID(id uint) (*entities.Task, error)
	ExecuteUpdateTask(task *entities.Task) error
	ExecuteDeleteTask(id uint) error
	ExecuteUpdateTaskStatus(id uint, status string) error
	ExecuteGetTasksByProjectID(projectID uint) ([]*entities.Task, error)
}

type TaskUseCase struct {
	repo repository.TaskRepository
}

func NewTaskUseCase(repo repository.TaskRepository) UCTaskUseCase {
	return &TaskUseCase{repo: repo}
}

func (uc *TaskUseCase) ExecuteCreateTask(task *entities.Task) error {
	if !uc.IsValidStatus(task.Status) {
		return errors.New("invalid status")
	}
	if err := uc.repo.CreateTask(task); err != nil {
		return errors.New("failed to create task")
	}
	return nil
}

func (uc *TaskUseCase) ExecuteGetAllTask() ([]*entities.Task, error) {
	tasks, err := uc.repo.GetAllTasks()
	if err != nil {
		return nil, errors.New("failed to get tasks")
	}
	return tasks, nil
}

func (uc *TaskUseCase) ExecuteGetTaskByID(id uint) (*entities.Task, error) {
	task, err := uc.repo.GetTaskById(id)
	if err != nil {
		return nil, errors.New("failed to get task by ID")
	}
	return task, nil
}

func (uc *TaskUseCase) ExecuteUpdateTask(task *entities.Task) error {
	if !uc.IsValidStatus(task.Status) {
		return errors.New("invalid status")
	}
	if err := uc.repo.UpdateTask(task); err != nil {
		return errors.New("failed to update task")
	}
	return nil
}

func (uc *TaskUseCase) ExecuteDeleteTask(id uint) error {
	if err := uc.repo.DeleteTask(id); err != nil {
		return errors.New("failed to delete task")
	}
	return nil
}

func (uc *TaskUseCase) IsValidStatus(status string) bool {
	validStatuses := []string{"pending", "in_progress", "completed", "cancelled"}
	for _, s := range validStatuses {
		if s == status {
			return true
		}
	}
	return false
}

func (uc *TaskUseCase) ExecuteUpdateTaskStatus(id uint, status string) error {
	if !uc.IsValidStatus(status) {
		return errors.New("invalid status")
	}
	task, err := uc.repo.GetTaskById(id)
	if err != nil {
		return errors.New("failed to get task by ID")
	}
	task.Status = status
	if err := uc.repo.UpdateTask(task); err != nil {
		return errors.New("failed to update task status")
	}
	return nil
}

func (uc *TaskUseCase) ExecuteGetTasksByProjectID(projectID uint) ([]*entities.Task, error) {
	tasks, err := uc.repo.GetTasksByProjectID(projectID)
	if err != nil {
		return nil, errors.New("failed to get tasks by project ID")
	}
	return tasks, nil
}
