package dto

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type UpdateUserRequest struct {
	ID       int    `json:"id" validate:"required"`
	Name     string `json:"name"`
	Email    string `json:"email" validate:"omitempty,email"`
	Password string `json:"password,omitempty" validate:"omitempty,min=6"`
}

type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type JwtRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
