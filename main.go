package main

import (
	"github.com/devmigi/go-todo-app/todo"
	"github.com/gofiber/fiber/v2"
)

func main() {
	todo.IntialMigration()
	app := fiber.New()
	Routers(app)
	app.Listen(":3000")
}

func Routers(app *fiber.App) {
	// create todo
	app.Post("/todos", todo.SaveTodo)

	// get all todos
	app.Get("/todos", todo.All)

	// get single todo
	app.Get("/todos/:id", todo.Detail)

	// update todo
	app.Put("/todos/:id", todo.Update)

	// delete todo
	app.Delete("/todos/:id", todo.Delete)

}
