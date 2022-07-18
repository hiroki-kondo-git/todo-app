package todo_handler

import (
	"net/http"
	"todo-app/usecase"

	"github.com/labstack/echo"
)

type TodoHandler interface {
	CreateTodo(c echo.Context) error
}

type todoHandler struct {
	todoUsecase usecase.TodoUseCase
}

func NewTodoHandler(tu usecase.TodoUseCase) TodoHandler {
	return &todoHandler{
		todoUsecase: tu,
	}
}

func (th todoHandler) CreateTodo(c echo.Context) error {
	text := c.QueryParam("text")
	status := c.QueryParam("status")
	res, err := th.todoUsecase.CreateTodo(text, status)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, res)
}
