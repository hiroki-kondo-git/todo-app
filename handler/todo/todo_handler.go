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
	Content string `json:"text`
	Status  string `json:"status"`
}

func (th todoHandler) CreateTodo(c echo.Context) error {
	content := c.QueryParam("content")
	status, err := strconv.Atoi(c.QueryParam("status"))
	if err != nil {
		return err
	}

	res, err := th.todoUsecase.CreateTodo(content, status)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return c.String(http.StatusOK, res)
}

func (th todoHandler) UpdateTodo(c echo.Context) error {
	content := c.QueryParam("content")
	status, err := strconv.Atoi(c.QueryParam("status"))
	if err != nil {
		return err
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	res, err := th.todoUsecase.UpdateTodo(content, status, id)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, res)
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

	return c.String(http.StatusOK, res)
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
		switch v.Status {
		case 0:
			r.Status = "todo"
		case 1:
			r.Status = "doing"
		case 2:
			r.Status = "done"
		}
		res = append(res, r)
	}

	return c.String(http.StatusOK, "a")
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
	switch row.Status {
	case 0:
		res.Status = "todo"
	case 1:
		res.Status = "doing"
	case 2:
		res.Status = "done"
	}
	return c.String(http.StatusOK, "i")
}
