package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// Users model
type Users struct {
	ID        string         `gorm:"default:uuid_generate_v4();primaryKey" json:"id"`
	Email     string         `gorm:"index" json:"email"`
	Password  string         `json:"-"`
	IsActive  bool           `gorm:"default:false" json:"isActive"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

//DEFINE HOOKS

func (user *Users) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("Before create data", user)
	return
}

func (user *Users) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println("After create data", user)
	return
}

func (user *Users) BeforeUpdate(tx *gorm.DB) (err error) {
	fmt.Println("Before update data", user)
	return
}

func (user *Users) AfterUpdate(tx *gorm.DB) (err error) {
	fmt.Println("After update data", user)
	return
}

func (user *Users) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Println("Before delete data", user)
	return
}

func (user *Users) AfterDelete(tx *gorm.DB) (err error) {
	fmt.Println("After delete data", user)
	return
}
