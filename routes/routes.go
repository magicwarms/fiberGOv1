package routes

import (
	"strings"

	controller "github.com/magicwarms/fiberGOv1/controllers"

	"github.com/gofiber/fiber/v2"
)

func helloWorld(c *fiber.Ctx) error {

	paramLala := c.Params("name")
	if paramLala == "" {
		return c.SendString("Where is your parameter?")
	}
	return c.SendString("Hi " + strings.ToTitle(paramLala))
}

// AppRoutes is all routes in app
func AppRoutes(app *fiber.App) {
	v1 := app.Group("/api/v1")
	v1.Get("/hello/:name?", helloWorld)
	v1.Get("/stack", func(c *fiber.Ctx) error {
		return c.JSON(c.App().Stack())
	})

	books := v1.Group("/book")
	books.Get("/list", controller.GetAllBooks)
	books.Get("/get", controller.GetBook)
	books.Post("/create", controller.CreateBook)
	books.Delete("/delete", controller.DeleteBook)

	persons := v1.Group("/person")
	persons.Get("/list", controller.GetPersons)
	persons.Get("/get", controller.GetPerson)
	persons.Get("/create", controller.NewPerson)
	persons.Get("/delete", controller.DeletePerson)
}
