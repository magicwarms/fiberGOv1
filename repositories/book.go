package repositories

import (
	"fmt"

	database "github.com/magicwarms/fiberGOv1/database"
	model "github.com/magicwarms/fiberGOv1/models"
)

// GetAllBooks is to get all books data
func GetAllBooks() []model.Books {
	db := database.DBConn
	var books []model.Books
	fmt.Print("mulai sini")
	fmt.Print(db, "<--- ini nih hasil nya nil, kayak db nya itu gak ada koneksi apa-apa")
	fmt.Println(&books, "<--- ini nih books")
	db.Find(&books)
	fmt.Print("gak kesini")
	return books
}
