package todo

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

// data-source-name = "USER:PASSWORD@tcp(HOST_NAME:3306)/DB_NAME?charset=utf8mb4&parseTime=True&loc=Local"
const dsn = "root:@tcp(127.0.0.1:3306)/gotuts?charset=utf8mb4&parseTime=True&loc=Local"

type Todo struct {
	gorm.Model

	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func IntialMigration() {
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
		panic("Database connection error")
	}

	DB.AutoMigrate(&Todo{})
}

// save new todo
func SaveTodo(c *fiber.Ctx) error {
	todo := new(Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	DB.Create(&todo)
	return c.JSON(todo)
}

// get all todos
func All(c *fiber.Ctx) error {
	var todos []Todo
	DB.Find(&todos)

	return c.JSON(todos)
}

// get a todo detail
func Detail(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo Todo
	DB.Find(&todo, id)
	return c.JSON(todo)
}

// get all todos
func Update(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo Todo

	DB.First(&todo, id)
	if todo.Title == "" {
		return c.Status(500).SendString("Todo Not Available")
	}

	if err := c.BodyParser(&todo); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	DB.Save(&todo)
	return c.SendString("Todo Updated")
}

// delete a todo
func Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo Todo

	DB.First(&todo, id)
	if todo.Title == "" {
		return c.Status(500).SendString("Todo Not Available")
	}

	DB.Delete(&todo)
	return c.SendString("Todo Deleted")
}
