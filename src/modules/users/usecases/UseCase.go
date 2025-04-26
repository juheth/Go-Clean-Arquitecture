package usecases

import (
	"errors"
	"fmt"

	jwt "github.com/juheth/Go-Clean-Arquitecture/src/common/auth"
	user "github.com/juheth/Go-Clean-Arquitecture/src/modules/users/domain/entities/user"
	"github.com/juheth/Go-Clean-Arquitecture/src/modules/users/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	ExecuteCreateUser(user *user.User) (string, error)
	ExecuteGetAllUsers() ([]*user.User, error)
	ExecuteGetUserByID(id int) (*user.User, error)
	ExecuteUpdateUser(user *user.User) error
	ExecuteDeleteUser(id int) error
	ExecuteAuthenticateUser(email, password string) (string, error)
}

type userUseCase struct {
	repo repository.UserRepository
	jwt  *jwt.JWT
}

func NewUserUseCase(repo repository.UserRepository, jwt *jwt.JWT) UserUseCase {
	return &userUseCase{
		repo: repo,
		jwt:  jwt,
	}
}

func (uc *userUseCase) ExecuteCreateUser(user *user.User) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error hashing password: %w", err)
	}
	user.Password = string(hashedPassword)

	if user == nil {
		return "", errors.New("cannot insert nil user")
	}

	if err := uc.repo.CreateUser(user); err != nil {
		return "", err
	}

	token, err := uc.jwt.GenerateToken(user.ID, user.Email)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (uc *userUseCase) ExecuteGetAllUsers() ([]*user.User, error) {
	return uc.repo.GetAllUsers()
}

func (uc *userUseCase) ExecuteGetUserByID(id int) (*user.User, error) {
	user, err := uc.repo.GetUserByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by ID %d: %w", id, err)
	}
	return user, nil
}

func (uc *userUseCase) ExecuteUpdateUser(user *user.User) error {
	if user == nil {
		return errors.New("cannot update nil user")
	}
	return uc.repo.UpdateUser(user)
}

func (uc *userUseCase) ExecuteDeleteUser(id int) error {
	return uc.repo.DeleteUser(id)
}

func (uc *userUseCase) ExecuteAuthenticateUser(email, password string) (string, error) {
	user, err := uc.repo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := uc.jwt.GenerateToken(user.ID, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
