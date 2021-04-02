package main

import (
	"fmt"
	"log"
	"os"

	"github.com/magicwarms/fiberGOv1/config"
	"github.com/magicwarms/fiberGOv1/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Print current process
	if fiber.IsChild() {
		fmt.Printf("[%d] CHILD\n", os.Getppid())
	} else {
		fmt.Printf("[%d] MASTER\n", os.Getppid())
	}
	enablePrefork := false
	if config.GoDotEnvVariable("APP_ENV") == "production" {
		enablePrefork = true
	}
	app := fiber.New(fiber.Config{
		Prefork: enablePrefork,
		// Enables the Server HTTP header with the given value.
		ServerHeader: "FiberGOV1",
		// Override default error handler
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err != nil {
				return c.Status(fiber.StatusInternalServerError).JSON(config.AppResponse{
					Code:    fiber.StatusInternalServerError,
					Message: "Internal Server Error - " + err.Error(),
					Data:    nil,
				})
			}
			// Return from handler
			return nil
		},
	})
	// will compress the response using gzip, deflate and brotli compression depending on the Accept-Encoding header.
	app.Use(compress.New())
	// to enable Cross-Origin Resource Sharing with various options.
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET,POST,PUT,DELETE",
	}))
	// Logger middleware for Fiber that logs HTTP request/response details.
	file, err := os.OpenFile("./logs/app-logging.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()
	app.Use(logger.New(logger.Config{
		Output:     file,
		Format:     "${pid} [${time}] | [${host} - ${ip}] | ${status} - ${latency} - ${method} | ${path} || ${error}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "Asia/Jakarta",
	}))
	// To recover from a panic thrown by any handler in the stack
	app.Use(recover.New())
	// for Fiber to let's caches be more efficient and save bandwidth,
	// as a web server does not need to resend a full response if the content has not changed.
	app.Use(etag.New())
	//start DB connection
	config.InitDatabase()
	// Custom Timer middleware
	app.Use(config.Timer())
	// setup routes list
	routes.AppRoutes(app)
	// setup not found 404 response
	config.NotFoundConfig(app)
	// start listen app
	log.Fatal(app.Listen(":" + config.GoDotEnvVariable("APP_PORT")))
}
