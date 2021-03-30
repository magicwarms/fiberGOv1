package repositories

import (
	"fmt"

	"github.com/magicwarms/fiberGOv1/config"
	"github.com/magicwarms/fiberGOv1/models"
)

// GetAllAuthors is to get all authors data
func GetAllAuthors() []models.Authors {
	var authors []models.Authors
	results := config.DB.Preload("Books").Find(&authors)
	if results.Error != nil {
		fmt.Println(results.Error)
		return authors
	}
	return authors
}

// CreateAuthor is to create author data based on input body
func CreateAuthor(author *models.Authors) *models.Authors {
	result := config.DB.Create(&author)
	if result.Error != nil {
		panic(result.Error)
	}
	return author
}

// GetAuthor is to get only one author data
func GetAuthor(authorId string) models.Authors {
	var author models.Authors
	result := config.DB.First(&author, "id = ?", authorId)
	if result.Error != nil {
		fmt.Println(result.Error)
		return author
	}
	return author
}
