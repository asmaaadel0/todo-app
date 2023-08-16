package internal

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetTasks(t *testing.T) {
	t.Run("Get request", func(t *testing.T) {
		router := gin.Default()

		app, err := NewApp("./database.db", 3000)
		assert.Nil(t, err)

		router.GET("/tasks", app.getTasks)

		req, err := http.NewRequest("GET", "/tasks", nil)
		assert.Nil(t, err)

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code, "failed in response code")

		var tasks []Task
		err = json.Unmarshal(recorder.Body.Bytes(), &tasks)

		assert.Nil(t, err)
	})
}

func TestAddTask(t *testing.T) {
	t.Run("Add task", func(t *testing.T) {
		router := gin.Default()

		app, err := NewApp("./database.db", 3000)
		assert.Nil(t, err)

		router.POST("/tasks", app.addTask)

		newTask := Task{
			Title:     "Test Task",
			Completed: false,
		}
		requestBody, err := json.Marshal(newTask)
		assert.Nil(t, err)

		req, err := http.NewRequest("POST", "/tasks", bytes.NewBuffer(requestBody))
		assert.Nil(t, err)

		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusCreated, recorder.Code, "failed in response code")
	})
}

func TestDeleteTask(t *testing.T) {
	t.Run("delete task", func(t *testing.T) {
		router := gin.Default()

		app, err := NewApp("./database.db", 3000)
		assert.Nil(t, err)

		router.DELETE("/tasks/:id", app.deleteTask)

		taskID := 123
		req, err := http.NewRequest("DELETE", "/tasks/"+strconv.Itoa(taskID), nil)
		assert.Nil(t, err)

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code, "failed in response code")
	})
}

func TestUpdateTask(t *testing.T) {
	t.Run("update task", func(t *testing.T) {
		router := gin.Default()

		app, err := NewApp("./database.db", 3000)
		assert.Nil(t, err)

		router.PUT("/tasks/:id", app.updateTask)

		updateTask := Task{
			Title:     "Updated Task",
			Completed: true,
		}
		requestBody, err := json.Marshal(updateTask)

		assert.Nil(t, err)

		req, err := http.NewRequest("PUT", "/tasks/1", bytes.NewBuffer(requestBody))
		assert.Nil(t, err)

		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusOK, recorder.Code, "failed in response code")
	})

	t.Run("test bad request ", func(t *testing.T) {
		router := gin.Default()

		app, err := NewApp("./database.db", 3000)
		assert.Nil(t, err)
		defer os.Remove("./database.db")

		router.PUT("/tasks/:id", app.updateTask)

		req, err := http.NewRequest("PUT", "/tasks/1", nil)

		assert.Nil(t, err)

		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusBadRequest, recorder.Code, "failed in response code")
	})

	t.Run("test Not found", func(t *testing.T) {
		router := gin.Default()

		app, err := NewApp("./database.db", 3000)
		assert.Nil(t, err)
		defer os.Remove("./database.db")

		router.PUT("/tasks", app.updateTask)

		updateTask := Task{
			Title:     "Updated Task",
			Completed: true,
		}
		requestBody, err := json.Marshal(updateTask)

		assert.Nil(t, err)

		req, err := http.NewRequest("PUT", "/tasks/11111", bytes.NewBuffer(requestBody))

		assert.Nil(t, err)

		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		assert.Equal(t, http.StatusNotFound, recorder.Code, "failed in response code")
	})
}
