package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

var tasks []Task
var idCounter int

func main() {
	app := fiber.New()

	// Serve static files for HTML
	app.Static("/", "./public")

	// Route to display the form and list of tasks
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./public/index.html")
	})

	// Add a task
	app.Post("/add", func(c *fiber.Ctx) error {
		title := c.FormValue("title")
		if title != "" {
			idCounter++
			task := Task{ID: idCounter, Title: title}
			tasks = append(tasks, task)
		}
		return c.Redirect("/")
	})

	// Delete a task
	app.Delete("/delete/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).SendString("Invalid Task ID")
		}
		for i, task := range tasks {
			if task.ID == id {
				tasks = append(tasks[:i], tasks[i+1:]...)
				return c.JSON(task)
			}
		}
		return c.Status(404).SendString("Task Not Found")
	})

	// Update a task
	app.Put("/update/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).SendString("Invalid Task ID")
		}
		newTitle := c.FormValue("title")
		for i, task := range tasks {
			if task.ID == id {
				tasks[i].Title = newTitle
				return c.JSON(tasks[i])
			}
		}
		return c.Status(404).SendString("Task Not Found")
	})

	// List all tasks
	app.Get("/tasks", func(c *fiber.Ctx) error {
		return c.JSON(tasks)
	})

	log.Fatal(app.Listen(":3000"))
}
