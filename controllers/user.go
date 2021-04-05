package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/magicwarms/fiberGOv1/config"
	"github.com/magicwarms/fiberGOv1/helpers"
	"github.com/magicwarms/fiberGOv1/models"
	"github.com/magicwarms/fiberGOv1/repositories"
)

type loginInput struct {
	Email    string
	Password string
}

type loginResponse struct {
	Token     string `json:"token"`
	ExpiresAt string `json:"expiresAt"`
}

// GetAllUsers is to get all users data
func GetAllUsers(c *fiber.Ctx) error {
	getAllUsers := repositories.GetAllUsers()
	if len(getAllUsers) < 0 {
		return c.JSON(config.AppResponse{
			Code:    http.StatusOK,
			Message: "NOT-FOUND",
			Data:    nil,
		})
	}
	return c.JSON(config.AppResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    getAllUsers,
	})
}

// GetUser is to get one user data
func GetUser(c *fiber.Ctx) error {
	userId := c.Query("id")
	getUser := repositories.GetUser(userId)
	if getUser.Email == "" {
		return c.JSON(config.AppResponse{
			Code:    http.StatusOK,
			Message: "NOT-FOUND",
			Data:    nil,
		})
	}
	return c.JSON(config.AppResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    getUser,
	})
}

// CreateUser is create new user data
func CreateUser(c *fiber.Ctx) error {
	user := new(models.Users)
	if err := c.BodyParser(user); err != nil {
		return c.JSON(config.AppResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "INVALID-PARAMS",
			Data:    nil,
		})
	}
	user.Password = helpers.HashPassword(user.Password)
	createUser := repositories.CreateUser(user)
	return c.JSON(config.AppResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    createUser,
	})
}

// DeleteUser is to delete user data
func DeleteUser(c *fiber.Ctx) error {
	user := new(models.Users)
	if err := c.BodyParser(user); err != nil {
		return c.JSON(config.AppResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "INVALID-PARAMS",
			Data:    nil,
		})
	}
	getUser := repositories.GetUser(user.ID)
	if getUser.Email == "" {
		return c.JSON(config.AppResponse{
			Code:    http.StatusOK,
			Message: "NOT-FOUND",
			Data:    nil,
		})
	}
	repositories.DeleteUser(user)
	return c.JSON(config.AppResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    nil,
	})
}

// UpdateUser is to update user data
func UpdateUser(c *fiber.Ctx) error {
	user := new(models.Users)
	if err := c.BodyParser(user); err != nil {
		return c.JSON(config.AppResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "INVALID-PARAMS",
			Data:    nil,
		})
	}
	getUser := repositories.GetUser(user.ID)
	if getUser.Email == "" {
		return c.JSON(config.AppResponse{
			Code:    http.StatusOK,
			Message: "NOT-FOUND",
			Data:    nil,
		})
	}
	updateUser := repositories.UpdateUser(user)
	return c.JSON(config.AppResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data:    updateUser,
	})
}

// LoginUser is to login to app using user email and password
func LoginUser(c *fiber.Ctx) error {
	loginData := new(loginInput)
	if err := c.BodyParser(loginData); err != nil {
		return c.JSON(config.AppResponse{
			Code:    http.StatusUnprocessableEntity,
			Message: "INVALID-PARAMS",
			Data:    nil,
		})
	}
	checkUserDataByEmail := repositories.GetUserByEmail(loginData.Email)
	if checkUserDataByEmail.Email == "" {
		return c.JSON(config.AppResponse{
			Code:    http.StatusOK,
			Message: "EMAIL-NOT-FOUND",
			Data:    nil,
		})
	}
	comparePassword := helpers.CheckPasswordHash(loginData.Password, checkUserDataByEmail.Password)
	if !comparePassword {
		return c.JSON(config.AppResponse{
			Code:    http.StatusUnauthorized,
			Message: "WRONG-PASSWORD",
			Data:    nil,
		})
	}
	token := helpers.GenerateToken(checkUserDataByEmail.ID, checkUserDataByEmail.Email)
	return c.JSON(config.AppResponse{
		Code:    http.StatusOK,
		Message: "OK",
		Data: loginResponse{
			Token:     token,
			ExpiresAt: "3 Days",
		},
	})
}
