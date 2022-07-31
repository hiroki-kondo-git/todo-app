package usecase

import (
	"todo-app/domain/model"
	"todo-app/domain/repository"
)

type TodoUseCase interface {
	CreateTodo(string, int) (string, error)
	UpdateTodo(string, int, int) (string, error)
	DeleteTodo(int) (string, error)
	GetAllTodo() ([]*model.Todo, error)
}

type todoUseCase struct {
	todoRepository repository.TodoRepository
}

func NewTodoUseCase(tr repository.TodoRepository) TodoUseCase {
	return &todoUseCase{
		todoRepository: tr,
	}
}

func (tu todoUseCase) CreateTodo(content string, status int) (string, error) {
	todo := model.NewTodo(content, status)
	// validation
	res, err := tu.todoRepository.Insert(todo)
	if err != nil {
		return "", err
	}
	return res, nil
}

func (tu todoUseCase) UpdateTodo(content string, status int, id int) (string, error) {
	todo := model.NewTodo(content, status)
	todo.Id = id
	// validation
	res, err := tu.todoRepository.Upsert(todo)
	if err != nil {
		return "", nil
	}
	return res, nil
}

func (tu todoUseCase) DeleteTodo(id int) (string, error) {
	res, err := tu.todoRepository.Delete(id)
	if err != nil {
		return "", err
	}

	return res, nil
}

func (tu todoUseCase) GetAllTodo() ([]*model.Todo, error) {
	rows, err := tu.todoRepository.Getall()
	if err != nil {
		return nil, err
	}
	return rows, nil
}
