package internal

import (
	"database/sql"
	"os"
	"strings"
)

func (app *App) readSqlCommands(schemaPath string) error {
	sqlFile, err := os.ReadFile(schemaPath)
	if err != nil {
		return err
	}
	app.sqlCommands = strings.Split(string(sqlFile), ";")
	return nil
}

func (app *App) connectDatabase(path string) error {
	var err error

	if app.db, err = sql.Open("sqlite3", path); err != nil {
		return err
	}

	return app.createTable()
}

func (app *App) createTable() error {
	tasks_table := app.sqlCommands[0]
	query, err := app.db.Prepare(tasks_table)
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

	record, err := app.db.Query(app.sqlCommands[1])
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
	records := app.sqlCommands[2]
	query, err := app.db.Prepare(records)
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

	records := app.sqlCommands[3]
	query, err := app.db.Prepare(records)
	if err != nil {
		return err
	}
	_, err = query.Exec(id)
	return err
}

func (app *App) updateTask(task Task) error {

	records := app.sqlCommands[4]
	query, err := app.db.Prepare(records)
	if err != nil {
		return err
	}
	_, err = query.Exec(task.Title, task.Completed, task.Id)
	return err
}
