package entities

import "time"

type Project struct {
	ID        int64      `json:"id" gorm:"primaryKey"`
	Name      string     `json:"name" gorm:"column:name;not null"`
	UserID    int        `json:"userId" gorm:"column:user_id;not null"`
	CreatedAt time.Time  `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time  `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt *time.Time `json:"deletedAt,omitempty" gorm:"index"`
	Status    string     `json:"status" gorm:"column:status;default:'active'"`
}
