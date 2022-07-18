package main

import (
	db "todo-app/config"
	template_handler "todo-app/handler/template"
	todo_handler "todo-app/handler/todo"
	"todo-app/infra/persistence"
	"todo-app/usecase"

	"github.com/labstack/echo"
)

func main() {
	db, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	todoPersistence := persistence.NewTodoPersistence(db)
	todoUseCase := usecase.NewTodoUseCase(todoPersistence)
	todoHandler := todo_handler.NewTodoHandler(todoUseCase)

	e := echo.New()
	e.Renderer = template_handler.NewTemplate("templates/index.html")

	e.GET("/", template_handler.Index)
	e.POST("/new", todoHandler.CreateTodo)
	e.Logger.Fatal(e.Start(":8080"))
}
