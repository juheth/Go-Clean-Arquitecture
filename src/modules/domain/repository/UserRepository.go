package repository

import (
	"errors"

	entities "github.com/juheth/Go-Clean-Arquitecture/src/modules/domain/entities/user"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *entities.User) error
	// GetUserByID(id string) (*entities.User, error)
	// UpdateUser(user *entities.User) error
	// DeleteUser(id string) error
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

// func (r *userRepo) GetUserByID(id string) (*entities.User, error) {
// 	var user entities.User
// 	err := r.db.First(&user, "id = ?", id).Error
// 	return &user, err
// }

// func (r *userRepo) UpdateUser(user *entities.User) error {
// 	return r.db.Save(user).Error
// }

// func (r *userRepo) DeleteUser(id string) error {
// 	return r.db.Delete(&entities.User{}, "id = ?", id).Error
// }
