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
	GetAllTodo() ([]*TodoResponse, error)
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
	Status  int    `json:"status"`
}

func (th todoHandler) CreateTodo(c echo.Context) error {
	content := c.FormValue("content")
	status, err := strconv.Atoi(c.FormValue("status"))
	if err != nil {
		return err
	}

	res, err := th.todoUsecase.CreateTodo(content, status)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return c.Redirect(http.StatusFound, "/")
}

func (tu todoHandler) UpdateTodo(c echo.Context) error {
	content := c.QueryParam("content")
	status, err := strconv.Atoi(c.FormValue("status"))
	if err != nil {
		return err
	}
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return err
	}

	res, err := tu.todoUsecase.UpdateTodo(content, status, id)
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

func (tu todoHandler) GetAllTodo() ([]*TodoResponse, error) {
	rows, err := tu.todoUsecase.GetAllTodo()
	if err != nil {
		return nil, err
	}
	var res []*TodoResponse
	for _, v := range rows {
		r := &TodoResponse{}
		r.Id = v.Id
		r.Content = v.Content
		r.Status = v.Status
		res = append(res, r)
	}

	return res, err
}
