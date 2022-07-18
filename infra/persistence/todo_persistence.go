package persistence

import (
	"database/sql"
	"fmt"
	"todo-app/domain/model"
	"todo-app/domain/repository"
)

type todoPersistence struct {
	db *sql.DB
}

func NewTodoPersistence(db *sql.DB) repository.TodoRepository {
	return &todoPersistence{
		db: db,
	}
}

func (tp todoPersistence) Insert(t *model.Todo) (string, error) {
	res, err := tp.db.Exec(
		"INSERT INTO todo (text, status) VALUES (?, ?)",
		t.Text,
		t.Status,
	)
	if err != nil {
		return "", err
	}
	rows, _ := res.RowsAffected()
	if rows != 1 {
		return "", fmt.Errorf("faled to insert to todo")
	}

	return fmt.Sprint("successfully insert to todo"), nil
}
