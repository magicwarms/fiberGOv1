package repositories

import (
	config "fiberGOv1/config"
	entity "fiberGOv1/entities"
)

var db = config.DBConn

// GetAllBooks is to get all books data
func GetAllBooks() entity.Book {
	var books entity.Book
	db.Find(&books)
	return books
}
