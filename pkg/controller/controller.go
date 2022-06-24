package controller

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/storyteller23/todolist-go/pkg/models"
)

func CreateTable(db *sql.DB) error {
	sql_table := `
	CREATE TABLE IF NOT EXISTS todos(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		completed INTEGER
	);
	`
	_, err := db.Exec(sql_table)
	if err != nil {
		return err
	}
	return nil
}

func Add(title string, db *sql.DB) {
	db.Exec(`
	INSERT INTO todos (title, completed)
	VALUES(?, 0)
	`, title)
}

func GetTodoList(db *sql.DB) (models.TodoList, error) {
	rows, err := db.Query(`SELECT * FROM todos`)
	if err != nil {
		return models.TodoList{}, err
	}
	result := make([]models.Todo, 0)
	for rows.Next() {
		todo := models.Todo{}
		err = rows.Scan(&todo.Id, &todo.Title, &todo.Completed)
		if err != nil {
			return models.TodoList{}, err
		}
		result = append(result, todo)
	}

	return models.TodoList{Todos: result}, nil
}
