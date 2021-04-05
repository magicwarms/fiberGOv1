package routes

import (
	"strings"

	controller "github.com/magicwarms/fiberGOv1/controllers"
	"github.com/magicwarms/fiberGOv1/middlewares"

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
	v1.Post("/create", controller.CreateUser)
	v1.Post("/login", controller.LoginUser)

	books := v1.Group("/book", middlewares.CheckAuthorization)
	books.Get("/list", controller.GetAllBooks)
	books.Get("/get", controller.GetBook)
	books.Post("/create", controller.CreateBook)
	books.Delete("/delete", controller.DeleteBook)
	books.Put("/update", controller.UpdateBook)

	authors := v1.Group("/author", middlewares.CheckAuthorization)
	authors.Get("/list", controller.GetAllAuthors)
	authors.Get("/get", controller.GetAuthor)
	authors.Post("/create", controller.CreateAuthor)
	authors.Delete("/delete", controller.DeleteAuthor)
	authors.Put("/update", controller.UpdateAuthor)

	users := v1.Group("/user", middlewares.CheckAuthorization)
	users.Get("/list", controller.GetAllUsers)
	users.Get("/get", controller.GetUser)
	users.Delete("/delete", controller.DeleteUser)
	users.Put("/update", controller.UpdateUser)

}
