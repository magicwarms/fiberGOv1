package controllers

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/magicwarms/fiberGOv1/config"
	"github.com/magicwarms/fiberGOv1/models"
	"github.com/magicwarms/fiberGOv1/repositories"
)

type authorLocal struct {
	ID        string
	Fullname  string
	CreatedAt time.Time
	UpdatedAt time.Time
	Books     []bookLocal
}

type bookLocal struct {
	ID        string
	Title     string
	Rating    float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

// GetAuthors is to get all Authors data
func GetAllAuthors(c *fiber.Ctx) error {
	var getAllData []authorLocal
	getAllAuthors := repositories.GetAllAuthors()
	if len(getAllAuthors) < 0 {
		return c.JSON(config.AppResponse{
			Code:    http.StatusOK,
			Message: "NOT-FOUND",
			Data:    nil,
		})
	}
	for _, item := range getAllAuthors {
		var booksData []bookLocal
		for _, book := range item.Books {
			booksData = append(booksData, bookLocal{
				ID:        book.ID,
				Title:     book.Title,
				Rating:    book.Rating,
				CreatedAt: book.CreatedAt,
				UpdatedAt: book.UpdatedAt,
			})
		}
		getAllData = append(getAllData, authorLocal{
			ID:        item.ID,
			Fullname:  item.Fullname,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
			Books:     booksData,
		})
	}
	return c.JSON(config.AppResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    getAllData,
	})
}

// GetAuthor is to get one Author data
func GetAuthor(c *fiber.Ctx) error {
	return c.SendString("Single Author")
}

// CreateAuthor is create new Author data
func CreateAuthor(c *fiber.Ctx) error {
	author := new(models.Authors)
	if err := c.BodyParser(author); err != nil {
		return c.JSON(config.AppResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "INVALID-PARAMS",
			Data:    nil,
		})
	}
	createAuthor := repositories.CreateAuthor(author)
	return c.JSON(config.AppResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    createAuthor,
	})
}

// DeleteAuthor is to delete Author data
func DeleteAuthor(c *fiber.Ctx) error {
	return c.SendString("Delete Author")
}

// UpdateAuthor is to update Author data
func UpdateAuthor(c *fiber.Ctx) error {
	return c.SendString("Update Author")
}
