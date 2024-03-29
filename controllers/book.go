package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/magicwarms/fiberGOv1/config"
	"github.com/magicwarms/fiberGOv1/models"
	"github.com/magicwarms/fiberGOv1/repositories"
)

// GetAllBooks is to get all books data
func GetAllBooks(c *fiber.Ctx) error {
	getAllBooks := repositories.GetAllBooks()
	if len(getAllBooks) < 0 {
		return c.JSON(config.AppResponse{
			Code:    http.StatusOK,
			Message: "NOT-FOUND",
			Data:    nil,
		})
	}

	return c.JSON(config.AppResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    getAllBooks,
	})
}

// GetBook is to get one book data
func GetBook(c *fiber.Ctx) error {
	bookId := c.Query("id")
	getBook := repositories.GetBook(bookId)
	if getBook.Title == "" {
		return c.JSON(config.AppResponse{
			Code:    http.StatusOK,
			Message: "NOT-FOUND",
			Data:    nil,
		})
	}
	return c.JSON(config.AppResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    getBook,
	})
}

// CreateBook is create new book data
func CreateBook(c *fiber.Ctx) error {
	book := new(models.Books)
	if err := c.BodyParser(book); err != nil {
		return c.JSON(config.AppResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "INVALID-PARAMS",
			Data:    nil,
		})
	}
	createBook := repositories.CreateBook(book)
	return c.JSON(config.AppResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    createBook,
	})
}

// DeleteBook is to delete book data
func DeleteBook(c *fiber.Ctx) error {
	book := new(models.Books)
	if err := c.BodyParser(book); err != nil {
		return c.JSON(config.AppResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "INVALID-PARAMS",
			Data:    nil,
		})
	}
	getBook := repositories.GetBook(book.ID)
	if getBook.Title == "" {
		return c.JSON(config.AppResponse{
			Code:    http.StatusOK,
			Message: "NOT-FOUND",
			Data:    nil,
		})
	}
	go repositories.DeleteBook(book)
	return c.JSON(config.AppResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    nil,
	})
}

// UpdateBook is to update book data
func UpdateBook(c *fiber.Ctx) error {
	book := new(models.Books)
	if err := c.BodyParser(book); err != nil {
		return c.JSON(config.AppResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "INVALID-PARAMS",
			Data:    nil,
		})
	}
	getBook := repositories.GetBook(book.ID)
	if getBook.Title == "" {
		return c.JSON(config.AppResponse{
			Code:    http.StatusOK,
			Message: "NOT-FOUND",
			Data:    nil,
		})
	}
	updateBook := repositories.UpdateBook(book)
	return c.JSON(config.AppResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    updateBook,
	})
}
