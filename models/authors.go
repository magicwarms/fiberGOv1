package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Authors struct {
	ID        string         `gorm:"default:uuid_generate_v4();primaryKey" json:"id"`
	Fullname  string         `gorm:"index" json:"fullname"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
	Books     []Books        `gorm:"foreignKey:AuthorID" json:"books"`
}

//DEFINE HOOKS

func (author *Authors) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Before create data", author)
	return
}

func (author *Authors) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println("After create data", author)
	return
}

func (author *Authors) BeforeUpdate(tx *gorm.DB) (err error) {
	fmt.Println("Before update data", author)
	return
}

func (author *Authors) AfterUpdate(tx *gorm.DB) (err error) {
	fmt.Println("After update data", author)
	return
}

func (author *Authors) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Println("Before delete data", author)
	return
}

func (author *Authors) AfterDelete(tx *gorm.DB) (err error) {
	fmt.Println("After delete data", author)
	return
}
