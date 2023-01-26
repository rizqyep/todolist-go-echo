package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	ID        int64          `json:"id"`
	Task      string         `json:"task"`
	CreatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
