package routes

import (
	book "fiberGOv1/controllers/book"
	person "fiberGOv1/controllers/person"
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
func RoutesList(app *fiber.App) {
	v1 := app.Group("/api/v1")
	v1.Get("/hello/:name?", helloWorld)
	v1.Get("/stack", func(c *fiber.Ctx) error {
		return c.JSON(c.App().Stack())
	})

	books := v1.Group("/book")
	books.Get("/list", book.GetBooks)
	books.Get("/get", book.GetBook)
	books.Get("/create", book.NewBook)
	books.Get("/delete", book.DeleteBook)

	persons := v1.Group("/person")
	persons.Get("/list", person.GetPersons)
	persons.Get("/get", person.GetPerson)
	persons.Get("/create", person.NewPerson)
	persons.Get("/delete", person.DeletePerson)
}
