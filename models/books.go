package models

import (
	"gorm.io/gorm"
)

// Books model
type Books struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}
