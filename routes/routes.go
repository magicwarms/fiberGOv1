package routes

import (
	controller "fiberGOv1/controllers"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func helloWorld(c *fiber.Ctx) error {

	paramLala := c.Params("name")
	if paramLala == "" {
		return c.SendString("Where is your parameter?")
	}
	return c.SendString("Hi " + strings.ToTitle(paramLala))
}

// RoutesList is all routes in app
func RoutesAppList(app *fiber.App) {
	v1 := app.Group("/api/v1")
	v1.Get("/hello/:name?", helloWorld)
	v1.Get("/stack", func(c *fiber.Ctx) error {
		return c.JSON(c.App().Stack())
	})

	books := v1.Group("/book")
	books.Get("/list", controller.GetBooks)
	books.Get("/get", controller.GetBook)
	books.Get("/create", controller.NewBook)
	books.Get("/delete", controller.DeleteBook)

	persons := v1.Group("/person")
	persons.Get("/list", controller.GetPersons)
	persons.Get("/get", controller.GetPerson)
	persons.Get("/create", controller.NewPerson)
	persons.Get("/delete", controller.DeletePerson)
}
