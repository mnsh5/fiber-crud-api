package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mnsh5/fiber-crud-api/handlers"
)

func TaskRoutes(app *fiber.App) {
	// Grupo principal de la API
	api := app.Group("/api")

	// Subgrupo para la versión 1 de la API
	v1 := api.Group("/v1")

	// Rutas específicas de la versión 1 de la API

	// http://localhost:3000/api/v1/tasks
	v1.Get("/tasks", handlers.GetTasks)

	// http://localhost:3000/api/v1/task/5
	v1.Get("/task/:id", handlers.GetTask)

	// http://localhost:3000/api/v1/task
	v1.Post("/task", handlers.CreateTask)

	// http://localhost:3000/api/v1/task/5
	v1.Put("/task/:id", handlers.EditTask)

	// http://localhost:3000/api/v1/task/5
	v1.Delete("/task/:id", handlers.DeleteTask)
}
