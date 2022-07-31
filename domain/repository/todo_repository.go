package repository

import "todo-app/domain/model"

type TodoRepository interface {
	Insert(*model.Todo) (string, error)
	Upsert(*model.Todo) (string, error)
	Delete(int) (string, error)
	Getall() ([]*model.Todo, error)
	GetTodoFromId(id int) (*model.Todo, error)
}
