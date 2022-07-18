package repository

import "todo-app/domain/model"

type TodoRepository interface {
	Insert(*model.Todo) (string, error)
}
