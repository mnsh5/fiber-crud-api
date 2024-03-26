package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/mnsh5/fiber-crud-api/database"
	"github.com/mnsh5/fiber-crud-api/routes"
)

func main() {
	database.ConnectDB()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowMethods: "GET, POST, PUT, DELETE",
		AllowHeaders: "Origin, Content-Type, Accept",
		// AllowCredentials: true,
	}))

	routes.TaskRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
