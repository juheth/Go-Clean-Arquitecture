package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	common "github.com/juheth/Go-Clean-Arquitecture/src/common/response"
	"github.com/juheth/Go-Clean-Arquitecture/src/modules/projects/domain/dto"
	"github.com/juheth/Go-Clean-Arquitecture/src/modules/projects/domain/entities"
	"github.com/juheth/Go-Clean-Arquitecture/src/modules/projects/usecases"
)

type ProjectController struct {
	ProjectController usecases.UCProjectsUseCase
}

func NewProjectController(projectController usecases.UCProjectsUseCase) *ProjectController {
	return &ProjectController{
		ProjectController: projectController,
	}
}

func (pc *ProjectController) CreateProject(c *fiber.Ctx) error {
	result := common.NewResult()
	var request dto.CreateProjectRequest

	if err := c.BodyParser(&request); err != nil {
		return result.Bad(c, "Invalid request body")
	}

	if request.Name == "" || request.UserID == 0 {
		return result.Bad(c, "Name and user ID are required")
	}

	project := &entities.Project{
		Name:   request.Name,
		UserID: int(request.UserID),
	}

	if createdProject, err := pc.ProjectController.ExecuteCreateProject(project); err != nil {
		return result.Error(c, "Could not create project")
	} else {
		user, err := pc.ProjectController.ExecuteGetUserByID(int64(request.UserID))
		if err != nil {
			return result.Error(c, "User not found")
		}

		projectResponse := fiber.Map{
			"id":     createdProject.ID,
			"name":   createdProject.Name,
			"userId": createdProject.UserID,
			"user": fiber.Map{
				"id":    user.ID,
				"name":  user.Name,
				"email": user.Email,
			},
			"status":    createdProject.Status,
			"createdAt": createdProject.CreatedAt,
			"updatedAt": createdProject.UpdatedAt,
		}

		return result.Ok(c, fiber.Map{
			"message": "Project created successfully",
			"project": projectResponse,
		})
	}

}

func (pc *ProjectController) GetProjectByID(c *fiber.Ctx) error {
	result := common.NewResult()
	id := c.Params("id")

	if id == "" {
		return result.Bad(c, "Project ID is required")
	}

	projectID, err := strconv.Atoi(id)
	if err != nil {
		return result.Bad(c, "Invalid project ID")
	}

	project, err := pc.ProjectController.ExecuteGetProjectByID(int64(projectID))
	if err != nil {
		return result.Error(c, "Could not retrieve project")
	}

	user, err := pc.ProjectController.ExecuteGetUserByID(int64(project.UserID))
	if err != nil {
		return result.Error(c, "User not found")
	}

	projectResponse := fiber.Map{
		"id":     project.ID,
		"name":   project.Name,
		"userId": project.UserID,
		"user": fiber.Map{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
		},
		"status":    project.Status,
		"createdAt": project.CreatedAt,
		"updatedAt": project.UpdatedAt,
	}

	return result.Ok(c, fiber.Map{
		"message": "Project retrieved successfully",
		"project": projectResponse,
	})
}

func (pc *ProjectController) GetAllProjects(c *fiber.Ctx) error {
	result := common.NewResult()

	projects, err := pc.ProjectController.ExecuteGetAllProjects()
	if err != nil {
		return result.Error(c, "Could not retrieve projects")
	}

	var projectResponses []fiber.Map
	for _, project := range projects {
		user, err := pc.ProjectController.ExecuteGetUserByID(int64(project.UserID))
		if err != nil {
			return result.Error(c, "User not found")
		}

		projectResponse := fiber.Map{
			"id":     project.ID,
			"name":   project.Name,
			"userId": project.UserID,
			"user": fiber.Map{
				"id":    user.ID,
				"name":  user.Name,
				"email": user.Email,
			},
			"status":    project.Status,
			"createdAt": project.CreatedAt,
			"updatedAt": project.UpdatedAt,
		}
		projectResponses = append(projectResponses, projectResponse)
	}

	return result.Ok(c, fiber.Map{
		"message":  "Projects retrieved successfully",
		"projects": projectResponses,
	})
}

func (pc *ProjectController) UpdateProject(c *fiber.Ctx) error {
	result := common.NewResult()
	id := c.Params("id")

	if id == "" {
		return result.Bad(c, "Project ID is required")
	}

	projectID, err := strconv.Atoi(id)
	if err != nil {
		return result.Bad(c, "Invalid project ID")
	}

	var request dto.UpdateProjectRequest
	if err := c.BodyParser(&request); err != nil {
		return result.Bad(c, "Invalid request body")
	}

	if request.Name == "" || request.UserID == 0 {
		return result.Bad(c, "Name and user ID are required")
	}

	project := &entities.Project{
		ID:     int64(projectID),
		Name:   request.Name,
		UserID: int(request.UserID),
	}

	if updatedProject, err := pc.ProjectController.ExecuteUpdateProject(project.ID, project); err != nil {
		return result.Error(c, "Could not update project")
	} else {
		user, err := pc.ProjectController.ExecuteGetUserByID(int64(request.UserID))
		if err != nil {
			return result.Error(c, "User not found")
		}

		projectResponse := fiber.Map{
			"id":     updatedProject.ID,
			"name":   updatedProject.Name,
			"userId": updatedProject.UserID,
			"user": fiber.Map{
				"id":    user.ID,
				"name":  user.Name,
				"email": user.Email,
			},
			"status":    updatedProject.Status,
			"createdAt": updatedProject.CreatedAt,
			"updatedAt": updatedProject.UpdatedAt,
		}

		return result.Ok(c, fiber.Map{
			"message": "Project updated successfully",
			"project": projectResponse,
		})
	}
}

func (pc *ProjectController) DeleteProject(c *fiber.Ctx) error {
	result := common.NewResult()
	id := c.Params("id")

	if id == "" {
		return result.Bad(c, "Project ID is required")
	}

	projectID, err := strconv.Atoi(id)
	if err != nil {
		return result.Bad(c, "Invalid project ID")
	}

	if err := pc.ProjectController.ExecuteDeleteProject(int64(projectID)); err != nil {
		return result.Error(c, "Could not delete project")
	}

	return result.Ok(c, fiber.Map{
		"message": "Project deleted successfully",
	})
}
