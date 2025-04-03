package usecases

import (
	"errors"

	db "github.com/juheth/Go-Clean-Arquitecture/src/infrastructure/db/adapter"
	entities "github.com/juheth/Go-Clean-Arquitecture/src/modules/domain/entities/user"
)

type UserUseCase interface {
	ExecuteCreateUser(user *entities.User) error
	ExecuteGetAllUsers() ([]*entities.User, error)
}

type CreateUserUseCase struct {
	db *db.DBConnection
}

func NewCreateUserUseCase(db *db.DBConnection) UserUseCase {
	return &CreateUserUseCase{db: db}
}

func (uc *CreateUserUseCase) ExecuteCreateUser(user *entities.User) error {
	_, err := uc.db.DB.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", user.Name, user.Email, user.Password)
	if err != nil {
		return errors.New("failed to create user")
	}
	return nil
}

func (uc *CreateUserUseCase) ExecuteGetAllUsers() ([]*entities.User, error) {

	var users []*entities.User
	rows, err := uc.db.DB.Query("SELECT id, name, email, password FROM users")
	if err != nil {
		return nil, errors.New("failed to retrieve users")
	}
	defer rows.Close()

	for rows.Next() {
		var user entities.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
			return nil, errors.New("failed to scan user")
		}
		users = append(users, &user)
	}

	return users, nil
}
