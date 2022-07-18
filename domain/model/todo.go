package model

import "time"

type Todo struct {
	Id        int    `json:"id"`
	Text      string `json:"text`
	Status    string `json:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTodo(text string, status string) *Todo {
	todo := new(Todo)
	todo.Text = text
	todo.Status = status
	return todo
}
