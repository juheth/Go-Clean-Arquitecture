package controllers

import (
	"github.com/gofiber/fiber/v2"
	common "github.com/juheth/Go-Clean-Arquitecture/src/common/response"
	entities "github.com/juheth/Go-Clean-Arquitecture/src/modules/domain/entities/user"
	"github.com/juheth/Go-Clean-Arquitecture/src/modules/usecases"
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

	if err := uc.createUserUseCase.Execute(&user); err != nil {
		return result.Error(c, "Could not create user")
	}

	return result.Ok(c, fiber.Map{"message": "User created successfully"})
}
