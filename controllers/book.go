package controllers

import (
	"net/http"

	config "github.com/magicwarms/fiberGOv1/config"
	service "github.com/magicwarms/fiberGOv1/services"

	"github.com/gofiber/fiber/v2"
)

// GetAllBooks is to get all books data
func GetAllBooks(c *fiber.Ctx) error {
	getAllBooks := service.GetAllBooks()
	return c.JSON(config.AppResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   getAllBooks,
	})
}

// GetBook is to get one book data
func GetBook(c *fiber.Ctx) error {
	return c.SendString("Single Book")
}

// NewBook is create new book data
func NewBook(c *fiber.Ctx) error {
	return c.SendString("New Book")
}

// DeleteBook is to delete book data
func DeleteBook(c *fiber.Ctx) error {
	return c.SendString("Delete Book")
}
