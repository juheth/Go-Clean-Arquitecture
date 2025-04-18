package repository

import (
	"github.com/juheth/Go-Clean-Arquitecture/src/modules/projects/domain/entities"
	Entities "github.com/juheth/Go-Clean-Arquitecture/src/modules/users/domain/entities/user"
	"gorm.io/gorm"
)

type ProjectRepository interface {
	CreateProject(project *entities.Project) error
	GetAllProjects() ([]*entities.Project, error)
	GetProjectById(id int64) (*entities.Project, error)
	GetUserById(id int64) (*Entities.User, error)
	UpdateProject(project *entities.Project) error
	DeleteProject(id int64) error
}

type ProjectRepo struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &ProjectRepo{db: db}
}

func (r *ProjectRepo) CreateProject(project *entities.Project) error {
	if project == nil {
		return gorm.ErrInvalidData
	}
	return r.db.Create(project).Error
}

func (r *ProjectRepo) GetAllProjects() ([]*entities.Project, error) {
	var projects []*entities.Project
	if err := r.db.Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

func (r *ProjectRepo) GetProjectById(id int64) (*entities.Project, error) {
	var project entities.Project
	if err := r.db.First(&project, id).Error; err != nil {
		return nil, err
	}
	return &project, nil
}

func (r *ProjectRepo) UpdateProject(project *entities.Project) error {
	if project == nil || project.ID == 0 {
		return gorm.ErrInvalidData
	}
	existingProject := &entities.Project{}
	if err := r.db.First(existingProject, project.ID).Error; err != nil {
		return err
	}
	return r.db.Model(existingProject).Updates(project).Error
}

func (r *ProjectRepo) DeleteProject(id int64) error {
	var project entities.Project
	if err := r.db.First(&project, id).Error; err != nil {
		return err
	}
	return r.db.Delete(&project).Error
}

func (r *ProjectRepo) GetUserById(id int64) (*Entities.User, error) {
	var user Entities.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
