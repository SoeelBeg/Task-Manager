package models

import (
	"database/sql"
)

type Task struct {
	ID   int
	Name string
}

var db *sql.DB

func SetDatabase(database *sql.DB) {
	db = database
}

func GetTasks() ([]Task, error) {
	rows, err := db.Query("SELECT id, name FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Name); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

func AddTask(name string) error {
	_, err := db.Exec("INSERT INTO tasks (name) VALUES (?)", name)
	return err
}

func DeleteTask(id int) error {
	_, err := db.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}
