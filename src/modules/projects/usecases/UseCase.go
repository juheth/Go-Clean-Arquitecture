package usecases

import (
	"github.com/juheth/Go-Clean-Arquitecture/src/modules/projects/domain/entities"
	"github.com/juheth/Go-Clean-Arquitecture/src/modules/projects/domain/repository"
	Entities "github.com/juheth/Go-Clean-Arquitecture/src/modules/users/domain/entities/user"
)

type UCProjectsUseCase interface {
	ExecuteCreateProject(project *entities.Project) (*entities.Project, error)
	ExecuteUpdateProject(id int64, project *entities.Project) (*entities.Project, error)
	ExecuteGetProjectByID(id int64) (*entities.Project, error)
	ExecuteGetUserByID(id int64) (*Entities.User, error)
	ExecuteDeleteProject(id int64) error

	ExecuteGetAllProjects() ([]*entities.Project, error)
}

type projectUseCase struct {
	repo repository.ProjectRepository
}

func NewProjectUseCase(repo repository.ProjectRepository) UCProjectsUseCase {
	return &projectUseCase{repo: repo}
}

func (uc *projectUseCase) ExecuteCreateProject(project *entities.Project) (*entities.Project, error) {
	if err := uc.repo.CreateProject(project); err != nil {
		return nil, err
	}
	return project, nil
}

func (uc *projectUseCase) ExecuteUpdateProject(id int64, project *entities.Project) (*entities.Project, error) {
	project.ID = id
	if err := uc.repo.UpdateProject(project); err != nil {
		return nil, err
	}
	return project, nil
}

func (uc *projectUseCase) ExecuteGetProjectByID(id int64) (*entities.Project, error) {
	project, err := uc.repo.GetProjectById(id)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (uc *projectUseCase) ExecuteDeleteProject(id int64) error {
	if err := uc.repo.DeleteProject(id); err != nil {
		return err
	}
	return nil
}

func (uc *projectUseCase) ExecuteGetUserByID(id int64) (*Entities.User, error) {
	user, err := uc.repo.GetUserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *projectUseCase) ExecuteGetAllProjects() ([]*entities.Project, error) {
	projects, err := uc.repo.GetAllProjects()
	if err != nil {
		return nil, err
	}
	return projects, nil
}
