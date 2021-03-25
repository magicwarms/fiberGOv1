package repositories

import (
	"fmt"

	"github.com/magicwarms/fiberGOv1/config"
	"github.com/magicwarms/fiberGOv1/models"
)

// GetAllBooks is to get all books data
func GetAllBooks() []models.Books {
	var books []models.Books
	result := config.DB.Find(&books)
	if result.Error != nil {
		fmt.Println(result.Error)
		return books
	}
	return books
}

// CreateBook is to create book data based on input body
func CreateBook(book *models.Books) *models.Books {
	err := config.DB.Create(&book).Error
	if err != nil {
		panic(err)
	}
	return book
}

// GetBook is to get only one book data
func GetBook(bookId int) models.Books {
	var book models.Books
	err := config.DB.First(&book, bookId).Error
	if err != nil {
		fmt.Println(err)
		return book
	}
	return book
}

// DeleteBook is to delete one book data
func DeleteBook(bookId int) models.Books {
	var book models.Books
	err := config.DB.Delete(&book, bookId).Error
	if err != nil {
		panic(err)
	}
	return book
}
