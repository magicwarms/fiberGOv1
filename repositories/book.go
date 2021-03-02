package repositories

import (
	database "fiberGOv1/config/database"
	model "fiberGOv1/models"
)

func getAllBooks() {
	db := database.DBConn
	var books []model.Book
	db
}
