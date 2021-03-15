package controllers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/magicwarms/fiberGOv1/config"
	"github.com/magicwarms/fiberGOv1/models"
	"github.com/magicwarms/fiberGOv1/services"
)

// GetAllBooks is to get all books data
func GetAllBooks(c *fiber.Ctx) error {
	getAllBooks := services.GetAllBooks()
	return c.JSON(config.AppResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   getAllBooks,
	})
}

// GetBook is to get one book data
func GetBook(c *fiber.Ctx) error {
	bookId, _ := strconv.Atoi(c.Query("id"))
	getBook := services.GetBook(bookId)
	if (getBook == models.Books{}) {
		return c.JSON(config.AppResponse{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   nil,
		})
	}
	return c.JSON(config.AppResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   getBook,
	})
}

// CreateBook is create new book data
func CreateBook(c *fiber.Ctx) error {
	book := new(models.Books)
	if err := c.BodyParser(book); err != nil {
		return c.JSON(config.AppResponse{
			Code:   http.StatusUnprocessableEntity,
			Status: "UNPROCESSABLE-ENTITY",
			Data:   nil,
		})
	}
	createBook := services.CreateBook(book)
	return c.JSON(config.AppResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   createBook,
	})

}

// DeleteBook is to delete book data
func DeleteBook(c *fiber.Ctx) error {
	return c.SendString("Delete Book")
}
