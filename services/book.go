package services

import (
	"github.com/magicwarms/fiberGOv1/models"
	"github.com/magicwarms/fiberGOv1/repositories"
)

// GetAllBooks is to get all books data
func GetAllBooks() []models.Books {
	getAllBooks := repositories.GetAllBooks()
	return getAllBooks
}

// CreateBook is to get all books data
func CreateBook(book *models.Books) *models.Books {
	createBook := repositories.CreateBook(book)
	return createBook
}

// GetBook is to get only one book data
func GetBook(bookId int) models.Books {
	getBook := repositories.GetBook(bookId)
	return getBook
}

// DeleteBook is to delete one book data
func DeleteBook(bookId int) models.Books {
	deleteBook := repositories.DeleteBook(bookId)
	return deleteBook
}
