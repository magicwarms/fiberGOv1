package controllers

import (
	"github.com/gofiber/fiber/v2"
)

// GetBooks is to get all books data
func GetBooks(c *fiber.Ctx) error {
	return c.SendString("All Books")
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
