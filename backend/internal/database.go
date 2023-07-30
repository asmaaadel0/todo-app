package internal

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
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
	fmt.Printf("trying to connect to database with path: %s\n", path)

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

	fmt.Println("Table created successfully!")
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

func (app *App) deleteTask(id int) ([]byte, error) {

	records := app.sqlCommands[3]
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

	records := app.sqlCommands[4]
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
