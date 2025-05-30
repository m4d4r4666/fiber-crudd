package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/m4d4r4666/fiber-crud/handlers"
)

func RegisterRoutes(app *fiber.App) {
	api := app.Group("")

	api.Get("/tasks", handlers.GetTasks)
	api.Post("/tasks", handlers.CreateTask)
	api.Put("/tasks/:id", handlers.UpdateTask)
	api.Delete("/tasks/:id", handlers.DeleteTask)
}
