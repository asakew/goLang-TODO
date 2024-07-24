package main

import (
	"todo-app/database"
	"todo-app/handlers"
	"todo-app/models"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Initialize Database
	database.ConnectDB()
	database.DB.AutoMigrate(&models.Todo{})

	// Static files
	app.Static("/", "./views")

	// Routes
	app.Get("/todos", handlers.GetTodos)
	app.Post("/todos", handlers.CreateTodo)
	app.Put("/todos/:id", handlers.UpdateTodo)
	app.Delete("/todos/:id", handlers.DeleteTodo)

	app.Listen(":3000")
}
