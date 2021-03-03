package services

import (
	entity "fiberGOv1/entities"
	repository "fiberGOv1/repositories"
)

// GetAllBooks is to get all books data
func GetAllBooks() entity.Book {
	return repository.GetAllBooks()
}
