package entities

import (
	"time"
)

type Task struct {
	ID          uint       `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	DueDate     *time.Time `json:"due_date"`
	ProjectID   uint       `json:"project_id" gorm:"column:project_id;not null"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
