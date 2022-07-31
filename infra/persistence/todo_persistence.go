package persistence

import (
	"database/sql"
	"fmt"
	"log"
	"todo-app/domain/model"
	"todo-app/domain/repository"
)

//todo transaction https://www.wakuwakubank.com/posts/869-go-database-sql/s

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
		"INSERT INTO todo (content, status) VALUES (?, ?)",
		t.Content,
		t.Status,
	)
	if err != nil {
		return "", fmt.Errorf("failed to insert todo sql error")
	}
	rows, _ := res.RowsAffected()
	if rows != 1 {
		return "", fmt.Errorf("faled to insert to todo")
	}

	return fmt.Sprint("successfully insert to todo"), nil
}

func (tp todoPersistence) Upsert(t *model.Todo) (string, error) {
	res, err := tp.db.Exec(
		"UPDATE todo SET content = ?, status = ? WHERE id = ?",
		t.Content,
		t.Status,
		t.Id,
	)
	if err != nil {
		return "", fmt.Errorf("failed to update todo sql error")
	}
	rows, _ := res.RowsAffected()
	if rows != 1 {
		return "", fmt.Errorf("failed to upsert todo")
	}

	return fmt.Sprintf("successfully update to todo id=%d", t.Id), nil
}

func (tp todoPersistence) Delete(id int) (string, error) {
	res, err := tp.db.Exec(
		"DELETE FROM todo WHERE id = ?",
		id,
	)
	if err != nil {
		return "", fmt.Errorf("failed to delete todo sql error")
	}
	rows, _ := res.RowsAffected()
	if rows != 1 {
		return "", fmt.Errorf("failed to delete todo")
	}

	return fmt.Sprintf("seccessfully delete from todo id=%d", id), nil
}

func (tp todoPersistence) Getall() ([]*model.Todo, error) {
	rows, err := tp.db.Query("SELECT * FROM todo")
	if err != nil {
		return nil, fmt.Errorf("failed to get all todo")
	}
	defer rows.Close()

	// 返却するtodoList作成
	todoList := []*model.Todo{}
	for rows.Next() {
		t := &model.Todo{}
		if err := rows.Scan(&t.Id, &t.Content, &t.Status, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan record")
		}
		todoList = append(todoList, t)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal("getrows err err:%V", err)
	}

	return todoList, nil
}

// func (tp todoPersistence) GetTodo(id int) (*model.Todo) {
// 	rows, err := db.Query("")
// }
