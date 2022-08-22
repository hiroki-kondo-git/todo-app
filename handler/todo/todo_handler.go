package todo_handler

import (
	"fmt"
	"net/http"
	"strconv"
	"todo-app/usecase"

	"github.com/labstack/echo"
)

type TodoHandler interface {
	CreateTodo(echo.Context) error
	UpdateTodo(echo.Context) error
	DeleteTodo(echo.Context) error
	GetAllTodo(echo.Context) error
	GetTodoFromId(echo.Context) error
}

type todoHandler struct {
	todoUsecase usecase.TodoUseCase
}

func NewTodoHandler(tu usecase.TodoUseCase) TodoHandler {
	return &todoHandler{
		todoUsecase: tu,
	}
}

type TodoResponse struct {
	Id      int    `json:"id"`
	Content string `json:"content`
	Status  string `json:"status"`
}

type TodoRequest struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Status  string `json:"status"`
}

func (th todoHandler) CreateTodo(c echo.Context) error {
	todo := &TodoRequest{}
	if err := c.Bind(todo); err != nil {
		return err
	}

	res, err := th.todoUsecase.CreateTodo(todo.Content, convertStatusAtoI(todo.Status))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

func (th todoHandler) UpdateTodo(c echo.Context) error {
	todo := &TodoRequest{}
	if err := c.Bind(todo); err != nil {
		return err
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	todo.Id = id

	res, err := th.todoUsecase.UpdateTodo(todo.Content, convertStatusAtoI(todo.Status), todo.Id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (th todoHandler) DeleteTodo(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Sprintln("failed to get query id")
	}
	res, err := th.todoUsecase.DeleteTodo(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

func (th todoHandler) GetAllTodo(c echo.Context) error {
	rows, err := th.todoUsecase.GetAllTodo()
	if err != nil {
		return err
	}
	var res []*TodoResponse
	for _, v := range rows {
		r := &TodoResponse{}
		r.Id = v.Id
		r.Content = v.Content
		r.Status = convertStatusItoA(v.Status)
		res = append(res, r)
	}

	return c.JSON(http.StatusOK, res)
}

func (th todoHandler) GetTodoFromId(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	res := &TodoResponse{}
	row, err := th.todoUsecase.GetTodoFromId(id)
	if err != nil {
		return err
	}
	res.Id = row.Id
	res.Content = row.Content
	res.Status = convertStatusItoA(row.Status)

	return c.JSON(http.StatusOK, res)
}

func convertStatusItoA(status int) string {
	var res string
	switch status {
	case 0:
		res = "todo"
	case 1:
		res = "doing"
	case 2:
		res = "done"
	}
	return res
}

func convertStatusAtoI(status string) int {
	var res int
	switch status {
	case "todo":
		res = 0
	case "doing":
		res = 1
	case "done":
		res = 2
	}
	return res
}
