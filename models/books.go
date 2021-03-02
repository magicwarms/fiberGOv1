package models

import (
	"time"

	"gorm.io/gorm"
)

// Books model
type Books struct {
	gorm.Model
	ID        int
	Title     string `json:"name"`
	Author    string `json:"author"`
	Rating    int    `json:"rating"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
