package usecases

import (
	"errors"

	db "github.com/juheth/Go-Clean-Arquitecture/src/infrastructure/db/adapter"
	entities "github.com/juheth/Go-Clean-Arquitecture/src/modules/domain/entities/user"
)

type UserUseCase interface {
	ExecuteCreateUser(user *entities.User) error
	ExecuteGetAllUsers() ([]*entities.User, error)
	ExecuteGetUserByID(id int) (*entities.User, error)
	ExecuteUpdateUser(user *entities.User) error
	ExecuteDeleteUser(id string) error
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

func (uc *CreateUserUseCase) ExecuteGetUserByID(id int) (*entities.User, error) {
	var user entities.User
	err := uc.db.DB.QueryRow("SELECT id, name, email, password FROM users WHERE id = ?", id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, errors.New("failed to retrieve user")
	}
	return &user, nil
}

func (uc *CreateUserUseCase) ExecuteUpdateUser(user *entities.User) error {
	_, err := uc.db.DB.Exec("UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?", user.Name, user.Email, user.Password, user.ID)
	if err != nil {
		return errors.New("failed to update user")
	}
	return nil
}

func (uc *CreateUserUseCase) ExecuteDeleteUser(id string) error {
	_, err := uc.db.DB.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return errors.New("failed to delete user")
	}
	return nil
}
