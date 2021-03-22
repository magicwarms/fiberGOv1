package config

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// NotFoundConfig is to handle route 404 not found exception
func NotFoundConfig(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(AppResponse{
			Code:    fiber.StatusNotFound,
			Message: "Route not found",
			Data:    nil,
		})
	})
}

// GoDotEnvVariable at godot package to load/read the .env file and
// return the value of the key
func GoDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
		panic("Error loading .env file")
	}

	return os.Getenv(key)
}

// AppResponse is for response config show to Frontend side
type AppResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
