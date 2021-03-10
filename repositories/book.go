package repositories

import (
	"fmt"

	"github.com/magicwarms/fiberGOv1/config"
	"github.com/magicwarms/fiberGOv1/models"
)

// GetAllBooks is to get all books data
func GetAllBooks() []models.Books {
	var books []models.Books
	err := config.DB.Find(&books).Error
	if err != nil {
		fmt.Println(err)
		return books
	}
	return books
}
