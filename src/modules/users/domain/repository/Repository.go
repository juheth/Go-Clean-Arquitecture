package repository

import (
	"errors"

	entities "github.com/juheth/Go-Clean-Arquitecture/src/modules/users/domain/entities/user"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *entities.User) error
	GetAllUsers() ([]*entities.User, error)
	GetUserByID(id int) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
	UpdateUser(user *entities.User) error
	DeleteUser(id int) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) CreateUser(user *entities.User) error {
	if user == nil {
		return errors.New("cannot insert nil user")
	}
	return r.db.Create(user).Error
}

func (r *userRepo) GetAllUsers() ([]*entities.User, error) {
	var users []*entities.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepo) GetUserByID(id int) (*entities.User, error) {
	var user entities.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepo) UpdateUser(user *entities.User) error {
	if user == nil {
		return errors.New("cannot update nil user")
	}
	return r.db.Save(user).Error
}

func (r *userRepo) DeleteUser(id int) error {
	return r.db.Delete(&entities.User{}, id).Error
}

func (r *userRepo) GetUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
