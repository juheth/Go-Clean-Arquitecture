package usecases

import (
	"errors"

	db "github.com/juheth/Go-Clean-Arquitecture/src/infrastructure/db/adapter"
	entities "github.com/juheth/Go-Clean-Arquitecture/src/modules/domain/entities/user"
)

type UserUseCase interface {
	Execute(user *entities.User) error
}

type CreateUserUseCase struct {
	db *db.DBConnection
}

func NewCreateUserUseCase(db *db.DBConnection) UserUseCase {
	return &CreateUserUseCase{db: db}
}

func (uc *CreateUserUseCase) Execute(user *entities.User) error {
	_, err := uc.db.DB.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", user.Name, user.Email, user.Password)
	if err != nil {
		return errors.New("failed to create user")
	}
	return nil
}
