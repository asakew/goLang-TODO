package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title  string `json:"title"`
	Status string `json:"status"`
}

var DB *gorm.DB

func initDatabase() {
	dsn := "host=localhost user=postgres password=superUser7 dbname=marina port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	DB.AutoMigrate(&Todo{})
}

func setupRoutes(app *fiber.App) {
	app.Get("/todos", getTodos)
	app.Post("/todos", addTodo)
	app.Put("/todos/:id", updateTodo)
	app.Delete("/todos/:id", deleteTodo)
	app.Static("/", "./public")
}

func main() {
	app := fiber.New()

	initDatabase()
	setupRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))
}

func getTodos(c *fiber.Ctx) error {
	var todos []Todo
	DB.Find(&todos)
	return c.JSON(todos)
}

func addTodo(c *fiber.Ctx) error {
	todo := new(Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	DB.Create(&todo)
	return c.JSON(todo)
}

func updateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	todo := new(Todo)
	if err := DB.First(&todo, id).Error; err != nil {
		return c.Status(404).SendString("Todo not found")
	}
	if err := c.BodyParser(todo); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	DB.Save(&todo)
	return c.JSON(todo)
}

func deleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := DB.Delete(&Todo{}, id).Error; err != nil {
		return c.Status(404).SendString("Todo not found")
	}
	return c.SendString("Todo successfully deleted")
}
