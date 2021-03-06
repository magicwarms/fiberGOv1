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
func DeleteBook(book *models.Books) *models.Books {
	result := config.DB.Delete(&book)
	if result.Error != nil {
		panic(result.Error)
	}
	return book
}

// UpdateBook is to update one book data
func UpdateBook(book *models.Books) *models.Books {
	result := config.DB.Model(&book).Select("title", "author_id", "rating").Updates(map[string]interface{}{"title": book.Title, "author_id": book.AuthorID, "rating": book.Rating})
	if result.Error != nil {
		panic(result.Error)
	}
	return book
}
