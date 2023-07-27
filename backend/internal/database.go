package internal

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
)

func (app *App) connectDatabase(path string) error {
	var err error
	fmt.Printf("trying to connect to database with path: %s\n", path)

	if app.db, err = sql.Open("sqlite3", path); err != nil {
		return err
	}

	app.createTable()
	return nil
}

func (app *App) createTable() error {
	tasks_table := `CREATE TABLE tasks (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        "title" TEXT,
        "completed" boolean);`
	query, err := app.db.Prepare(tasks_table)
	if err != nil {
		return err
	}
	query.Exec()
	fmt.Println("Table created successfully!")
	return nil
}

func (app *App) getTasks() ([]Task, error) {

	var tasks []Task

	record, err := app.db.Query("SELECT * FROM tasks")
	if err != nil {
		return nil, err
	}
	defer record.Close()

	for record.Next() {
		var id int
		var title string
		var completed bool
		record.Scan(&id, &title, &completed)

		tasks = append(tasks, Task{
			Id:        id,
			Title:     title,
			Completed: completed})
	}
	return tasks, nil
}

func (app *App) addTask(title string, completed bool) ([]byte, error) {
	records := `INSERT INTO tasks(title, completed) VALUES (?, ?)`
	query, err := app.db.Prepare(records)
	if err != nil {
		return nil, err
	}
	response, err := query.Exec(title, completed)
	if err != nil {
		log.Fatal(err)
	}
	data, err := json.Marshal(&response)
	return data, err
}

func (app *App) deleteTask(id int) ([]byte, error) {

	records := `DELETE FROM tasks WHERE id = ?;`
	query, err := app.db.Prepare(records)
	if err != nil {
		return nil, err
	}
	response, err := query.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
	data, err := json.Marshal(&response)
	return data, err
}

func (app *App) updateTask(task Task) ([]byte, error) {

	records := `UPDATE tasks SET title = ?, completed = ? WHERE id = ?;`
	query, err := app.db.Prepare(records)
	if err != nil {
		return nil, err
	}
	response, err := query.Exec(task.Title, task.Completed, task.Id)
	if err != nil {
		return nil, err
	}
	data, err := json.Marshal(&response)
	return data, err

}
