package models

import (
	"time"

	"gorm.io/gorm"
)

// Books model
type Books struct {
	gorm.Model
	ID        uint           `json:"id"`
	Title     string         `json:"title"`
	Author    string         `json:"author"`
	Rating    int            `json:"rating"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}
