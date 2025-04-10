package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	common "github.com/juheth/Go-Clean-Arquitecture/src/common/response"
	"github.com/juheth/Go-Clean-Arquitecture/src/modules/users/domain/dto"
	entities "github.com/juheth/Go-Clean-Arquitecture/src/modules/users/domain/entities/user"
	"github.com/juheth/Go-Clean-Arquitecture/src/modules/users/usecases"
)

type UserController struct {
	useCase usecases.UserUseCase
}

func NewUserController(useCase usecases.UserUseCase) *UserController {
	return &UserController{useCase: useCase}
}

func (uc *UserController) CreateUser(c *fiber.Ctx) error {
	result := common.NewResult()
	var request dto.CreateUserRequest

	if err := c.BodyParser(&request); err != nil {
		return result.Bad(c, "Invalid request body")
	}

	if request.Name == "" || request.Email == "" || request.Password == "" {
		return result.Bad(c, "Name, email, and password are required")
	}

	if len(request.Password) < 6 {
		return result.Bad(c, "Password must be at least 6 characters long")
	}

	user := &entities.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	if err := uc.useCase.ExecuteCreateUser(user); err != nil {
		return result.Error(c, "Could not create user")
	}

	return result.Ok(c, fiber.Map{"message": "User created successfully"})
}

func (uc *UserController) GetAllUsers(c *fiber.Ctx) error {
	result := common.NewResult()
	users, err := uc.useCase.ExecuteGetAllUsers()
	if err != nil {
		return result.Error(c, "Could not retrieve users")
	}

	var response []dto.UserResponse
	for _, u := range users {
		response = append(response, dto.UserResponse{
			ID:    u.ID,
			Name:  u.Name,
			Email: u.Email,
		})
	}

	return result.Ok(c, response)
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

	user, err := uc.useCase.ExecuteGetUserByID(intID)
	if err != nil {
		return result.Error(c, "Could not retrieve user")
	}

	response := dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	return result.Ok(c, response)
}

func (uc *UserController) UpdateUser(c *fiber.Ctx) error {
	result := common.NewResult()
	var request dto.UpdateUserRequest

	if err := c.BodyParser(&request); err != nil {
		return result.Bad(c, "Invalid request body")
	}

	if request.ID == 0 {
		return result.Bad(c, "ID is required")
	}

	if request.Name == "" || request.Email == "" || request.Password == "" {
		return result.Bad(c, "Name, email, and password are required")
	}

	if len(request.Password) < 6 {
		return result.Bad(c, "Password must be at least 6 characters long")
	}

	user := &entities.User{
		ID:       request.ID,
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}

	if err := uc.useCase.ExecuteUpdateUser(user); err != nil {
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

	intID, err := strconv.Atoi(id)
	if err != nil {
		return result.Bad(c, "ID must be a valid integer")
	}

	if err := uc.useCase.ExecuteDeleteUser(intID); err != nil {
		return result.Error(c, "Could not delete user")
	}

	return result.Ok(c, fiber.Map{"message": "User deleted successfully"})
}
