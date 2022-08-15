package model

import "time"

// status 1=todo, 2=done
type Todo struct {
	Id        int    `json:"id"`
	Content   string `json:"content`
	Status    int    `json:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTodo(content string, status int) *Todo {
	todo := new(Todo)
	todo.Content = content
	todo.Status = status
	return todo
}
