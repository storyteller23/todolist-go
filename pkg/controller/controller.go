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
	return checkErr(err)
}

func Add(title string, db *sql.DB) error {
	_, err := db.Exec(`
	INSERT INTO todos (title, completed)
	VALUES(?, 0)
	`, title)
	return checkErr(err)
}

func Delete(id string, db *sql.DB) error {
	_, err := db.Exec(`
	DELETE FROM todos WHERE id = ?
	`, id)
	return checkErr(err)
}

func GetTodoList(db *sql.DB) ([]models.Todo, error) {
	result := make([]models.Todo, 0)

	rows, err := db.Query(`SELECT * FROM todos`)
	if err != nil {
		return result, err
	}
	var id, completed int
	var title string
	for rows.Next() {
		err = rows.Scan(&id, &title, &completed)
		if err != nil {
			return []models.Todo{}, err
		}
		todo := models.Todo{
			Id:        id,
			Title:     title,
			Completed: completed,
		}
		result = append(result, todo)
	}

	return result, nil
}

func checkErr(err error) error {
	if err != nil {
		return err
	}
	return nil
}
