package repositories

import (
	"fmt"

	"github.com/magicwarms/fiberGOv1/config"
	"github.com/magicwarms/fiberGOv1/models"
)

// GetAllUsers is to get all users data
func GetAllUsers() []models.Users {
	var users []models.Users
	results := config.DB.Find(&users)
	if results.Error != nil {
		fmt.Println(results.Error)
		return users
	}
	return users
}

// CreateUser is to create user data based on input body
func CreateUser(user *models.Users) *models.Users {
	result := config.DB.Create(&user)
	if result.Error != nil {
		panic(result.Error)
	}
	return user
}

// GetUser is to get only one user data
func GetUser(userId string) models.Users {
	var user models.Users
	result := config.DB.First(&user, "id = ?", userId)
	if result.Error != nil {
		fmt.Println(result.Error)
		return user
	}
	return user
}

// DeleteUser is to delete one user data
func DeleteUser(user *models.Users) *models.Users {
	result := config.DB.Delete(&user)
	if result.Error != nil {
		panic(result.Error)
	}
	return user
}

// UpdateUser is to update one user data
func UpdateUser(user *models.Users) *models.Users {
	result := config.DB.Model(&user).Select("email", "is_active").Updates(map[string]interface{}{"email": user.Email, "is_active": user.IsActive})
	if result.Error != nil {
		panic(result.Error)
	}
	return user
}

// GetUserByEmail is to get only one user data by email
func GetUserByEmail(email string) models.Users {
	var user models.Users
	result := config.DB.First(&user, "email = ?", email)
	if result.Error != nil {
		fmt.Println(result.Error)
		return user
	}
	return user
}
