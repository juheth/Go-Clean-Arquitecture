package dto

import (
	"time"

	"github.com/juheth/Go-Clean-Arquitecture/src/modules/users/domain/dto"
)

type CreateProjectRequest struct {
	Name   string `json:"name" binding:"required"`
	UserID int64  `json:"userId" binding:"required"`
}

type UpdateProjectRequest struct {
	Name   string `json:"name" binding:"required"`
	Status string `json:"status" binding:"required"`
	UserID int    `json:"userId" binding:"required"`
}

type ProjectResponse struct {
	ID        int64            `json:"id"`
	Name      string           `json:"name"`
	UserID    int64            `json:"userId"`
	User      dto.UserResponse `json:"user"`
	CreatedAt time.Time        `json:"createdAt"`
	UpdatedAt time.Time        `json:"updatedAt"`
	DeletedAt *time.Time       `json:"deletedAt,omitempty"`
	Status    string           `json:"status"`
}
