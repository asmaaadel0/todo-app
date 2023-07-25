package internal

import (
	"fmt"
	"log"
)

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

func (app *App) getTasks() []Task {

	var tasks []Task

	record, err := app.db.Query("SELECT * FROM tasks")
	if err != nil {
		log.Fatal(err)
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
	return tasks
}

func (app *App) addTask(title string, completed bool) {
	records := `INSERT INTO tasks(title, completed) VALUES (?, ?)`
	query, err := app.db.Prepare(records)
	if err != nil {
		log.Fatal(err)
	}
	_, err = query.Exec(title, completed)
	if err != nil {
		log.Fatal(err)
	}
}

func (app *App) deleteTask(id int) {

	records := `DELETE FROM tasks WHERE id = ?;`
	query, err := app.db.Prepare(records)
	if err != nil {
		log.Fatal(err)
	}
	_, err = query.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
}

func (app *App) updateTask(task Task) {

	records := `UPDATE tasks SET title = ?, completed = ? WHERE id = ?;`
	query, err := app.db.Prepare(records)
	if err != nil {
		log.Fatal(err)
	}
	_, err = query.Exec(task.Title, task.Completed, task.Id)
	if err != nil {
		log.Fatal(err)
	}
}
