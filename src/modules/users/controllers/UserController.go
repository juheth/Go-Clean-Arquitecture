package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	common "github.com/juheth/Go-Clean-Arquitecture/src/common/response"
	entities "github.com/juheth/Go-Clean-Arquitecture/src/modules/users/domain/entities/user"
	"github.com/juheth/Go-Clean-Arquitecture/src/modules/users/usecases"
)

type UserController struct {
	createUserUseCase usecases.UserUseCase
}

func NewUserController(createUserUseCase usecases.UserUseCase) *UserController {
	return &UserController{createUserUseCase: createUserUseCase}
}

func (uc *UserController) CreateUser(c *fiber.Ctx) error {
	result := common.NewResult()
	var user entities.User

	if err := c.BodyParser(&user); err != nil {
		return result.Bad(c, "Invalid request body")
	}

	if user.Name == "" || user.Email == "" || user.Password == "" {
		return result.Bad(c, "Name, email, and password are required")
	}

	if len(user.Password) < 6 {
		return result.Bad(c, "Password must be at least 6 characters long")
	}

	if err := uc.createUserUseCase.ExecuteCreateUser(&user); err != nil {
		return result.Error(c, "Could not create user")
	}

	return result.Ok(c, fiber.Map{"message": "User created successfully"})
}

func (uc *UserController) GetAllUsers(c *fiber.Ctx) error {
	result := common.NewResult()
	users, err := uc.createUserUseCase.ExecuteGetAllUsers()
	if err != nil {
		return result.Error(c, "Could not retrieve users")
	}
	return result.Ok(c, users)
}

func (uc *UserController) GetUserByID(c *fiber.Ctx) error {
	result := common.NewResult()
	id := c.Params("id")
	if id == "" {
		return result.Bad(c, "ID is required")
	}

	intID, err := strconv.Atoi(id)
	if err != nil {
		return result.Bad(c, "ID must be a valid integer")
	}

	user, err := uc.createUserUseCase.ExecuteGetUserByID(intID)
	if err != nil {
		return result.Error(c, "Could not retrieve user")
	}
	return result.Ok(c, user)
}

func (uc *UserController) UpdateUser(c *fiber.Ctx) error {
	result := common.NewResult()
	var user entities.User
	if err := c.BodyParser(&user); err != nil {
		return result.Bad(c, "Invalid request body")
	}
	if user.ID == 0 {
		return result.Bad(c, "ID is required")
	}
	if user.Name == "" || user.Email == "" || user.Password == "" {
		return result.Bad(c, "Name, email, and password are required")
	}
	if len(user.Password) < 6 {
		return result.Bad(c, "Password must be at least 6 characters long")
	}
	if err := uc.createUserUseCase.ExecuteUpdateUser(&user); err != nil {

		return result.Error(c, "Could not update user")
	}
	return result.Ok(c, fiber.Map{"message": "User updated successfully"})
}

func (uc *UserController) DeleteUser(c *fiber.Ctx) error {
	result := common.NewResult()
	id := c.Params("id")
	if id == "" {
		return result.Bad(c, "ID is required")
	}
	if err := uc.createUserUseCase.ExecuteDeleteUser(id); err != nil {
		return result.Error(c, "Could not delete user")
	}
	return result.Ok(c, fiber.Map{"message": "User deleted successfully"})
}
