package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	common "github.com/juheth/Go-Clean-Arquitecture/src/common/response"
	"github.com/juheth/Go-Clean-Arquitecture/src/modules/tasks/domain/dto"
	"github.com/juheth/Go-Clean-Arquitecture/src/modules/tasks/domain/entities"
	"github.com/juheth/Go-Clean-Arquitecture/src/modules/tasks/usecases"
)

type TaskController struct {
	TaskUseCase usecases.UCTaskUseCase
}

func NewTaskController(taskUseCase usecases.UCTaskUseCase) *TaskController {
	return &TaskController{TaskUseCase: taskUseCase}
}

func (tc *TaskController) CreateTask(c *fiber.Ctx) error {
	result := common.NewResult()
	var request dto.CreateTaskRequest

	if err := c.BodyParser(&request); err != nil {
		return result.Bad(c, "Invalid request body")
	}

	if request.Title == "" || request.Description == "" {
		return result.Bad(c, "Title and description are required")
	}

	now := time.Now()
	if request.DueDate == nil {
		request.DueDate = &now
	}

	task := &entities.Task{
		Title:       request.Title,
		Description: request.Description,
		DueDate:     request.DueDate,
		Status:      "pending",
	}

	if err := tc.TaskUseCase.ExecuteCreateTask(task); err != nil {
		return result.Error(c, "Could not create task")
	}

	return result.Ok(c, fiber.Map{"message": "Task created successfully"})
}

func (tc *TaskController) GetAllTasks(c *fiber.Ctx) error {
	result := common.NewResult()
	tasks, err := tc.TaskUseCase.ExecuteGetAllTask()
	if err != nil {
		return result.Error(c, "Could not retrieve tasks")
	}

	var response []dto.TaskResponse
	for _, t := range tasks {
		response = append(response, dto.TaskResponse{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			Status:      t.Status,
			CreatedAt:   t.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   t.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return result.Ok(c, response)
}

func (tc *TaskController) GetTaskByID(c *fiber.Ctx) error {
	result := common.NewResult()
	id, err := c.ParamsInt("id")
	if err != nil {
		return result.Bad(c, "Invalid ID")
	}

	task, err := tc.TaskUseCase.ExecuteGetTaskByID(uint(id))
	if err != nil {
		return result.Error(c, "Task not found")
	}

	response := dto.TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		CreatedAt:   task.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   task.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return result.Ok(c, response)
}

func (tc *TaskController) UpdateTask(c *fiber.Ctx) error {
	result := common.NewResult()
	id, err := c.ParamsInt("id")
	if err != nil {
		return result.Bad(c, "Invalid ID")
	}

	var request dto.UpdateTaskRequest
	if err := c.BodyParser(&request); err != nil {
		return result.Bad(c, "Invalid request body")
	}

	task, err := tc.TaskUseCase.ExecuteGetTaskByID(uint(id))
	if err != nil {
		return result.Error(c, "Task not found")
	}

	task.Title = request.Title
	task.Description = request.Description
	now := time.Now()
	task.DueDate = &now

	if err := tc.TaskUseCase.ExecuteUpdateTask(task); err != nil {
		return result.Error(c, "Failed to update task")
	}

	updatedTask, err := tc.TaskUseCase.ExecuteGetTaskByID(uint(id))
	if err != nil {
		return result.Error(c, "Failed to fetch updated task")
	}

	response := dto.TaskResponse{
		ID:          updatedTask.ID,
		Title:       updatedTask.Title,
		Description: updatedTask.Description,
		Status:      updatedTask.Status,
		CreatedAt:   updatedTask.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   updatedTask.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return result.Ok(c, response)
}

func (tc *TaskController) DeleteTask(c *fiber.Ctx) error {
	result := common.NewResult()
	id, err := c.ParamsInt("id")
	if err != nil {
		return result.Bad(c, "Invalid ID")
	}

	if err := tc.TaskUseCase.ExecuteDeleteTask(uint(id)); err != nil {
		return result.Error(c, "Failed to delete task")
	}

	return result.Ok(c, fiber.Map{"message": "Task deleted successfully"})
}

func (tc *TaskController) UpdateTaskStatus(c *fiber.Ctx) error {
	result := common.NewResult()
	id, err := c.ParamsInt("id")
	if err != nil {
		return result.Bad(c, "Invalid ID")
	}

	var request dto.UpdateTaskStatusRequest
	if err := c.BodyParser(&request); err != nil {
		return result.Bad(c, "Invalid request body")
	}

	if err := tc.TaskUseCase.ExecuteUpdateTaskStatus(uint(id), request.Status); err != nil {
		return result.Error(c, err.Error())
	}

	return result.Ok(c, fiber.Map{"message": "Task status updated successfully"})
}

func (tc *TaskController) GetTasksByProjectID(c *fiber.Ctx) error {
	result := common.NewResult()
	projectID, err := c.ParamsInt("project_id")
	if err != nil {
		return result.Bad(c, "Invalid project ID")
	}

	tasks, err := tc.TaskUseCase.ExecuteGetTasksByProjectID(uint(projectID))
	if err != nil {
		return result.Error(c, "Failed to retrieve tasks for the project")
	}

	var response []dto.TaskResponse
	for _, t := range tasks {
		response = append(response, dto.TaskResponse{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,

			Status:    t.Status,
			CreatedAt: t.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: t.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return result.Ok(c, response)
}
