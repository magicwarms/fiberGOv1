package services

import (
	model "github.com/magicwarms/fiberGOv1/models"
	repository "github.com/magicwarms/fiberGOv1/repositories"
)

// GetAllBooks is to get all books data
func GetAllBooks() []model.Books {
	getAllBooks := repository.GetAllBooks()
	return getAllBooks
}
