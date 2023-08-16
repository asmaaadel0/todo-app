package internal

import (
	"database/sql"
	"os"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestConnectDatabase(t *testing.T) {
	t.Run("connect to database", func(t *testing.T) {
		client := DBClient{}
		var err error
		client.db, err = sql.Open("sqlite3", "memory")
		if err != nil {
			t.Fatalf("Failed to create memory database: %v", err)
		}
		defer client.db.Close()

		app := &App{}

		err = app.client.connectDatabase("memory")
		if err != nil {
			t.Fatalf("Failed to connect to the database: %v", err)
		}
	})
}

func TestCreateTable(t *testing.T) {
	t.Run("test create table", func(t *testing.T) {
		client := DBClient{}
		var err error
		client.db, err = sql.Open("sqlite3", "memory")
		if err != nil {
			t.Fatalf("Failed to create memory database: %v", err)
		}
		defer client.db.Close()

		app := &App{client: client}

		err = app.client.createTable()
		if err != nil {
			t.Fatalf("Failed to create table: %v", err)
		}
		defer os.Remove("memory")
	})
}

func TestGetTasksDB(t *testing.T) {
	t.Run("test get tasks", func(t *testing.T) {
		client := DBClient{}
		var err error
		var mock sqlmock.Sqlmock
		client.db, mock, err = sqlmock.New()
		if err != nil {
			t.Fatalf("Failed to create mock database: %v", err)
		}
		defer client.db.Close()

		app := &App{client: client}

		rows := sqlmock.NewRows([]string{"id", "title", "completed"}).
			AddRow(1, "Task 1", true).
			AddRow(2, "Task 2", false)

		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM tasks")).WillReturnRows(rows)

		tasks, err := app.client.getTasksDB()
		if err != nil {
			t.Fatalf("Failed to get tasks: %v", err)
		}

		if len(tasks) != 2 {
			t.Errorf("Expected 2 tasks, but got %d", len(tasks))
		}

		if tasks[0].Id != 1 || tasks[0].Title != "Task 1" || tasks[0].Completed != true {
			t.Errorf("Unexpected content for task 1")
		}

		if tasks[1].Id != 2 || tasks[1].Title != "Task 2" || tasks[1].Completed != false {
			t.Errorf("Unexpected content for task 2")
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %s", err)
		}
	})
}

func TestAddTasksDB(t *testing.T) {
	t.Run("test add task", func(t *testing.T) {
		client := DBClient{}
		var err error
		var mock sqlmock.Sqlmock
		client.db, mock, err = sqlmock.New()
		if err != nil {
			t.Fatalf("Failed to create mock database: %v", err)
		}
		defer client.db.Close()

		app := &App{client: client}

		result := sqlmock.NewResult(1, 1)

		mock.ExpectPrepare("INSERT INTO tasks(.+)").ExpectExec().WillReturnResult(result)

		title := "New Task"
		completed := true
		id, err := app.client.addTaskDB(title, completed)
		if err != nil {
			t.Fatalf("Failed to add task: %v", err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("Unfulfilled expectations: %s", err)
		}

		expectedID := int64(1)

		if id != expectedID {
			t.Errorf("expected id: %d, but found %d", expectedID, id)
		}
	})
}

func TestDeleteTaskDB(t *testing.T) {
	t.Run("test delete task", func(t *testing.T) {
		client := DBClient{}
		var err error
		var mock sqlmock.Sqlmock
		client.db, mock, err = sqlmock.New()
		if err != nil {
			t.Fatalf("Failed to create mock database: %v", err)
		}
		defer client.db.Close()

		app := &App{client: client}
		result := sqlmock.NewResult(1, 1)

		mock.ExpectPrepare("DELETE FROM tasks WHERE id = \\?;").ExpectExec().WillReturnResult(result)

		id := 123
		err = app.client.deleteTaskDB(id)
		if err != nil {
			t.Fatalf("Failed to delete task: %v", err)
		}

		err = app.client.deleteTaskDB(id)
		if err == nil {
			t.Fatalf("error delete non existed task: %v", err)
		}
	})
}

func TestUpdateTaskDB(t *testing.T) {
	t.Run("test update task", func(t *testing.T) {
		client := DBClient{}
		var err error
		var mock sqlmock.Sqlmock
		client.db, mock, err = sqlmock.New()
		if err != nil {
			t.Fatalf("Failed to create mock database: %v", err)
		}
		defer client.db.Close()

		app := &App{client: client}

		task := Task{
			Id:        123,
			Title:     "Updated Task",
			Completed: true,
		}

		expectedAffectedRows := int64(1)
		result := sqlmock.NewResult(expectedAffectedRows, expectedAffectedRows)

		mock.ExpectPrepare("UPDATE tasks SET title = \\?, completed = \\? WHERE id = \\?;").ExpectExec().WillReturnResult(result)

		err = app.client.updateTaskDB(task)
		if err != nil {
			t.Fatalf("Failed to update task: %v", err)
		}

		task = Task{
			Id:        1233333333,
			Title:     "Updated Task",
			Completed: true,
		}
		err = app.client.updateTaskDB(task)
		if err == nil {
			t.Fatalf("error update non existent task: %v", err)
		}
	})
}
