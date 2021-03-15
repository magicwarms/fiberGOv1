package models

import (
	"gorm.io/gorm"
)

// Books model
type Books struct {
	gorm.Model
	Title  string `gorm:"column:title"`
	Author string `gorm:"column:author"`
	Rating int    `gorm:"column:rating"`
}
