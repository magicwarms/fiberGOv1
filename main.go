package main

import (
	appConfig "fiberGOv1/config"
	database "fiberGOv1/config"
	routes "fiberGOv1/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	// will compress the response using gzip, deflate and brotli compression depending on the Accept-Encoding header.
	app.Use(compress.New())
	// to enable Cross-Origin Resource Sharing with various options.
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET,POST,PUT,DELETE",
	}))
	// to intercept responses and cache them
	app.Use(cache.New())
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
	// for Fiber to lets caches be more efficient and save bandwidth,
	// as a web server does not need to resend a full response if the content has not changed.
	app.Use(etag.New())
	// Initiate DB connection
	database.InitDatabase()
	// setup routes list
	routes.RoutesAppList(app)
	// setup not found 404 response
	appConfig.NotFoundConfig(app)
	// start listen app
	log.Fatal(app.Listen(":9000"))
}
