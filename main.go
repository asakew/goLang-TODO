package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	ConnectDatabase()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("index.html")
	})

	app.Get("/todos", GetTodos)
	app.Post("/todos", CreateTodo)
	app.Get("/todos/:id", GetTodo)
	app.Put("/todos/:id", UpdateTodo)
	app.Delete("/todos/:id", DeleteTodo)

	app.Listen(":3000")
}

func GetTodos(c *fiber.Ctx) error {
	var todos []Todo
	DB.Find(&todos)
	return c.JSON(todos)
}

func CreateTodo(c *fiber.Ctx) error {
	todo := new(Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	DB.Create(&todo)
	return c.JSON(todo)
}

func GetTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo Todo
	if result := DB.First(&todo, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, "Todo not found")
	}
	return c.JSON(todo)
}

func UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo Todo
	if result := DB.First(&todo, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, "Todo not found")
	}

	updatedTodo := new(Todo)
	if err := c.BodyParser(updatedTodo); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	todo.Title = updatedTodo.Title
	todo.Description = updatedTodo.Description
	todo.Completed = updatedTodo.Completed
	DB.Save(&todo)
	return c.JSON(todo)
}

func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo Todo
	if result := DB.First(&todo, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, "Todo not found")
	}

	DB.Delete(&todo)
	return c.SendString("Todo successfully deleted")
}
