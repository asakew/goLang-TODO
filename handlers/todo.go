package handlers

import (
	"todo-app/database"
	"todo-app/models"

	"github.com/gofiber/fiber/v2"
)

func GetTodos(c *fiber.Ctx) error {
	var todos []models.Todo
	database.DB.Find(&todos)
	return c.JSON(todos)
}

func CreateTodo(c *fiber.Ctx) error {
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	database.DB.Create(&todo)
	return c.JSON(todo)
}

func UpdateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo models.Todo
	database.DB.First(&todo, id)
	if todo.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Todo not found"})
	}

	if err := c.BodyParser(&todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}
	database.DB.Save(&todo)
	return c.JSON(todo)
}

func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo models.Todo
	database.DB.First(&todo, id)
	if todo.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Todo not found"})
	}
	database.DB.Delete(&todo)
	return c.JSON(fiber.Map{"message": "Todo deleted"})
}
