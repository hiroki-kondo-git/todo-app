package usecase

import (
	"todo-app/domain/model"
	"todo-app/domain/repository"
)

type TodoUseCase interface {
	CreateTodo(string, string) (string, error)
}

type todoUseCase struct {
	todoRepository repository.TodoRepository
}

func NewTodoUseCase(tr repository.TodoRepository) TodoUseCase {
	return &todoUseCase{
		todoRepository: tr,
	}
}

func (tu todoUseCase) CreateTodo(text string, status string) (string, error) {
	todo := model.NewTodo(text, status)
	// validation
	res, err := tu.todoRepository.Insert(todo)
	if err != nil {
		return "", err
	}
	return res, nil
}
