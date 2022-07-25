package todo_handler

import (
	"fmt"
	"net/http"
	"strconv"
	"todo-app/domain/model"
	"todo-app/usecase"

	"github.com/labstack/echo"
)

type TodoHandler interface {
	CreateTodo(echo.Context) error
	UpdateTodo(echo.Context) error
	DeleteTodo(echo.Context) error
	GetAllTodo() ([]*model.Todo, error)
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
	status, err := strconv.Atoi(c.QueryParam("status"))
	if err != nil {
		return err
	}

	res, err := th.todoUsecase.CreateTodo(text, status)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, res)
}

func (tu todoHandler) UpdateTodo(c echo.Context) error {
	text := c.QueryParam("text")
	status, err := strconv.Atoi(c.QueryParam("status"))
	if err != nil {
		return err
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return err
	}

	res, err := tu.todoUsecase.UpdateTodo(text, status, id)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, res)
}

func (tu todoHandler) DeleteTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Sprintln("failed to get query id")
	}
	res, err := tu.todoUsecase.DeleteTodo(id)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, res)
}

func (tu todoHandler) GetAllTodo() ([]*model.Todo, error) {
	res, err := tu.todoUsecase.GetAllTodo()
	if err != nil {
		return nil, err
	}

	return res, err
}
