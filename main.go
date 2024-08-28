package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"log"

	"repo/database"
	"repo/routes"
	"repo/handlers"
)

func setUpRoutes(app *fiber.App) {
	//route := a.Group("/api/v1")
	app.Get("/", routes.Hello)
	app.Get("/metadata", routes.GetFileMetadata)
	app.Post("/upload", handlers.UploadFile)
}


func main() {

	database.ConnectDB()
	app := fiber.New()

	setUpRoutes(app)

	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) 
	})

	log.Fatal(app.Listen(":3001"))
}


