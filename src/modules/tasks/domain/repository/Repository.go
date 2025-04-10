package repository

import (
	"errors"

	"github.com/juheth/Go-Clean-Arquitecture/src/modules/tasks/domain/entities"
	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task *entities.Task) error
	GetAllTasks() ([]*entities.Task, error)
	GetTaskById(id uint) (*entities.Task, error)
	UpdateTask(task *entities.Task) error
	DeleteTask(id uint) error
}

type taskRepo struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepo{db: db}
}

func (r *taskRepo) CreateTask(task *entities.Task) error {
	if task == nil {
		return errors.New("cannot insert nil task")
	}
	return r.db.Create(task).Error
}

func (r *taskRepo) GetAllTasks() ([]*entities.Task, error) {
	var tasks []*entities.Task
	if err := r.db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskRepo) GetTaskById(id uint) (*entities.Task, error) {
	var task entities.Task
	if err := r.db.First(&task, id).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *taskRepo) UpdateTask(task *entities.Task) error {
	if task == nil {
		return errors.New("cannot update nil task")
	}
	return r.db.Save(task).Error
}

func (r *taskRepo) DeleteTask(id uint) error {
	return r.db.Delete(&entities.Task{}, id).Error
}
