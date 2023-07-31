package internal

import (
	"database/sql"
	_ "embed"
)

//go:embed sql/createTable.sql
var createTable string

//go:embed sql/getTasks.sql
var getTasks string

//go:embed sql/addTask.sql
var addTask string

//go:embed sql/deleteTask.sql
var deleteTask string

//go:embed sql/updateTask.sql
var updateTask string

func (app *App) connectDatabase(path string) error {
	var err error

	if app.db, err = sql.Open("sqlite3", path); err != nil {
		return err
	}

	return app.createTable()
}

func (app *App) createTable() error {
	query, err := app.db.Prepare(createTable)
	if err != nil {
		return err
	}

	_, err = query.Exec()
	if err != nil {
		return err
	}

	return nil
}

func (app *App) getTasks() ([]Task, error) {

	var tasks []Task

	record, err := app.db.Query(getTasks)
	if err != nil {
		return nil, err
	}
	defer record.Close()

	for record.Next() {
		var id int
		var title string
		var completed bool
		err = record.Scan(&id, &title, &completed)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, Task{
			Id:        id,
			Title:     title,
			Completed: completed})
	}
	return tasks, nil
}

func (app *App) addTask(title string, completed bool) (int64, error) {

	query, err := app.db.Prepare(addTask)
	if err != nil {
		return 0, err
	}
	response, err := query.Exec(title, completed)
	if err != nil {
		return 0, err
	}

	id, err := response.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, err
}

func (app *App) deleteTask(id int) error {

	query, err := app.db.Prepare(deleteTask)
	if err != nil {
		return err
	}
	_, err = query.Exec(id)
	return err
}

func (app *App) updateTask(task Task) error {

	query, err := app.db.Prepare(updateTask)
	if err != nil {
		return err
	}
	_, err = query.Exec(task.Title, task.Completed, task.Id)
	return err
}
