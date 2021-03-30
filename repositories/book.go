package repositories

import (
	"fmt"

	"github.com/magicwarms/fiberGOv1/config"
	"github.com/magicwarms/fiberGOv1/models"
)

// GetAllBooks is to get all books data
func GetAllBooks() []models.Books {
	var books []models.Books
	results := config.DB.Preload("Authors").Find(&books)
	if results.Error != nil {
		fmt.Println(results.Error)
		return books
	}
	return books
}

// CreateBook is to create book data based on input body
func CreateBook(book *models.Books) *models.Books {
	result := config.DB.Create(&book)
	if result.Error != nil {
		panic(result.Error)
	}
	return book
}

// GetBook is to get only one book data
func GetBook(bookId string) models.Books {
	var book models.Books
	result := config.DB.Preload("Authors").First(&book, "id = ?", bookId)
	if result.Error != nil {
		fmt.Println(result.Error)
		return book
	}
	return book
}

// DeleteBook is to delete one book data
func DeleteBook(bookId string) models.Books {
	var book models.Books
	result := config.DB.Delete(&book, "id = ?", bookId)
	if result.Error != nil {
		panic(result.Error)
	}
	return book
}

func UpdateBook(book *models.Books) *models.Books {
	result := config.DB.Model(&book).Select("title", "author", "rating").Updates(map[string]interface{}{"title": book.Title, "author": book.AuthorID, "rating": book.Rating})
	if result.Error != nil {
		panic(result.Error)
	}
	return book
}
