package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Books model
type Books struct {
	ID        string         `gorm:"default:uuid_generate_v4();primaryKey" json:"id"`
	Title     string         `gorm:"index" json:"title"`
	Author    string         `json:"author"`
	Rating    float64        `json:"rating"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

func (book *Books) BeforeUpdate(tx *gorm.DB) (err error) {
	fmt.Println("Updated data", book.Title)
	return
}
