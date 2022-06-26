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

func IsCompleted(id string, db *sql.DB) (bool, error) {
	rows, err := db.Query(`SELECT completed FROM todos WHERE id = ?`, id)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	completed := 0
	rows.Next()
	rows.Scan(&completed)
	return completed == 1, nil
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

func Update(id string, newTitle string, db *sql.DB) error {
	_, err := db.Exec(`
	UPDATE todos
	SET title = ?
	WHERE id = ?;
	`, newTitle, id)

	return checkErr(err)
}

func CompleteTask(id string, db *sql.DB) error {
	_, err := db.Exec(`
	UPDATE todos
	SET completed = ?
	WHERE id = ?;
	`, 1, id)
	return checkErr(err)
}

func GetUncompletedTasks(db *sql.DB) ([]models.Todo, error) {
	result := make([]models.Todo, 0)

	rows, err := db.Query(`SELECT * FROM todos WHERE completed = ?`, 0)
	if err != nil {
		return result, err
	}
	defer rows.Close()
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

func GetCompletedTasks(db *sql.DB) ([]models.Todo, error) {
	result := make([]models.Todo, 0)

	rows, err := db.Query(`SELECT * FROM todos WHERE completed = ?`, 1)
	if err != nil {
		return result, err
	}
	defer rows.Close()
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
