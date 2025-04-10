package usecases

import (
	"errors"
	"fmt"

	entities "github.com/juheth/Go-Clean-Arquitecture/src/modules/users/domain/entities/user"
	"github.com/juheth/Go-Clean-Arquitecture/src/modules/users/domain/repository"
)

type UserUseCase interface {
	ExecuteCreateUser(user *entities.User) error
	ExecuteGetAllUsers() ([]*entities.User, error)
	ExecuteGetUserByID(id int) (*entities.User, error)
	ExecuteUpdateUser(user *entities.User) error
	ExecuteDeleteUser(id int) error
}

type userUseCase struct {
	repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}

func (uc *userUseCase) ExecuteCreateUser(user *entities.User) error {
	if user == nil {
		return errors.New("cannot insert nil user")
	}
	return uc.repo.CreateUser(user)
}

func (uc *userUseCase) ExecuteGetAllUsers() ([]*entities.User, error) {
	return uc.repo.GetAllUsers()
}

func (uc *userUseCase) ExecuteGetUserByID(id int) (*entities.User, error) {
	user, err := uc.repo.GetUserByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by ID %d: %w", id, err)
	}
	return user, nil
}

func (uc *userUseCase) ExecuteUpdateUser(user *entities.User) error {
	if user == nil {
		return errors.New("cannot update nil user")
	}
	return uc.repo.UpdateUser(user)
}

func (uc *userUseCase) ExecuteDeleteUser(id int) error {
	return uc.repo.DeleteUser(id)
}
