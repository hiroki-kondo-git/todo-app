package template_handler

import (
	"html/template"
	"io"
	"net/http"
	todo_handler "todo-app/handler/todo"

	"github.com/labstack/echo"
)

type Template struct {
	templates *template.Template
}

func NewTemplate(path string) *Template {
	t := &Template{
		templates: template.Must(template.ParseGlob(path)),
	}
	return t
}

type TemplateHandler interface {
	Index(echo.Context) error
}
type templateHandler struct {
	todoHandler todo_handler.TodoHandler
}

func NewTemplateHandler(th todo_handler.TodoHandler) TemplateHandler {
	return &templateHandler{todoHandler: th}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// status一旦int
type dataStruct struct {
	Todo   string
	Status int
}

func (th templateHandler) Index(c echo.Context) error {
	res, err := th.todoHandler.GetAllTodo()
	if err != nil {
		return err
	}
	data := []dataStruct{}
	for _, v := range res {
		ds := dataStruct{
			Todo:   v.Text,
			Status: v.Status,
		}
		data = append(data, ds)
	}
	return c.Render(http.StatusOK, "index", data)
}
