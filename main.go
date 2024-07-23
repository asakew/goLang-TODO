package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

type Task struct {
	Title string `json:"title"`
}

var tasks []Task

func main() {
	app := fiber.New()

	// Serve static files for HTML
	app.Static("/", "./public")

	// Define the get route to display the form and list of tasks
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./public/index.html")
	})

	// Handle form submission
	app.Post("/add", func(c *fiber.Ctx) error {
		title := c.FormValue("title")
		if title != "" {
			tasks = append(tasks, Task{Title: title})
		}
		return c.Redirect("/")
	})

	app.Get("/tasks", func(c *fiber.Ctx) error {
		return c.JSON(tasks)
	})

	log.Fatal(app.Listen(":3000"))
}
