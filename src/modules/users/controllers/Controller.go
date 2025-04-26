package controllers

import (
	"log"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/juheth/Go-Clean-Arquitecture/src/common/auth"
	common "github.com/juheth/Go-Clean-Arquitecture/src/common/response"
	"github.com/juheth/Go-Clean-Arquitecture/src/modules/users/domain/dto"
	entities "github.com/juheth/Go-Clean-Arquitecture/src/modules/users/domain/entities/user"
	"github.com/juheth/Go-Clean-Arquitecture/src/modules/users/usecases"
)

type UserController struct {
	useCase usecases.UserUseCase
	jwt     *auth.JWT
}

func NewUserController(useCase usecases.UserUseCase, jwt *auth.JWT) *UserController {
	return &UserController{
		useCase: useCase,
		jwt:     jwt,
	}
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

	if _, err := uc.useCase.ExecuteCreateUser(user); err != nil {
		return result.Error(c, "Could not create user")
	}

	jwtService := auth.NewJWT("JWT_SECRET")

	token, err := jwtService.GenerateToken(user.ID, user.Email)
	if err != nil {
		log.Printf("Error generating JWT token: %v", err)
		return result.Error(c, "Could not generate token")
	}

	return result.Ok(c, fiber.Map{
		"message": "User created successfully",
		"token":   token,
	})
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
			ID:    int(u.ID),
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
		ID:    int(user.ID),
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
		ID:       int(request.ID),
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

func (uc *UserController) JwtUser(c *fiber.Ctx) error {
	result := common.NewResult()
	var request dto.JwtRequest

	if err := c.BodyParser(&request); err != nil {
		return result.Bad(c, "Invalid request body")
	}

	if request.Email == "" || request.Password == "" {
		return result.Bad(c, "Email and password are required")
	}

	token, err := uc.useCase.ExecuteAuthenticateUser(request.Email, request.Password)
	if err != nil {
		return result.Error(c, "Authentication failed")
	}

	return result.Ok(c, fiber.Map{
		"token": token,
	})
}

func (uc *UserController) IsAuthenticated(c *fiber.Ctx) error {
	result := common.NewResult()

	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return result.Bad(c, "Token de autorización requerido")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return result.Bad(c, "Formato de token inválido")
	}

	token := parts[1]

	claims, err := uc.jwt.ValidateToken(token)
	if err != nil {
		return result.Bad(c, "Token inválido o expirado")
	}

	return result.Ok(c, fiber.Map{
		"message": "Usuario autenticado",
		"user": fiber.Map{
			"id":    claims.ID,
			"email": claims.Email,
		},
	})
}
