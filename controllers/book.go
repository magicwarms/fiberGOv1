package controllers

import (
	"net/http"

	"github.com/magicwarms/fiberGOv1/config"
	"github.com/magicwarms/fiberGOv1/services"

	"github.com/gofiber/fiber/v2"
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
