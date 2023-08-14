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
)

func TestGetTasks(t *testing.T) {
	t.Run("Get request", func(t *testing.T) {
		router := gin.Default()

		app, err := NewApp("./database.db", 3000)
		if err != nil {
			t.Fatalf("Error: %v", err)
		}

		router.GET("/tasks", app.getTasks)

		req, err := http.NewRequest("GET", "/tasks", nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		if recorder.Code != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
		}

		var tasks []Task
		if err := json.Unmarshal(recorder.Body.Bytes(), &tasks); err != nil {
			t.Fatalf("Failed to unmarshal response JSON: %v", err)
		}
	})
}

func TestAddTask(t *testing.T) {
	t.Run("Add task", func(t *testing.T) {
		router := gin.Default()

		app, err := NewApp("./database.db", 3000)
		if err != nil {
			t.Fatalf("Error: %v", err)
		}

		router.POST("/tasks", app.addTask)

		newTask := Task{
			Title:     "Test Task",
			Completed: false,
		}
		requestBody, err := json.Marshal(newTask)
		if err != nil {
			t.Fatalf("Failed to marshal request JSON: %v", err)
		}

		req, err := http.NewRequest("POST", "/tasks", bytes.NewBuffer(requestBody))
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		if recorder.Code != http.StatusCreated {
			t.Errorf("Expected status code %d, but got %d", http.StatusCreated, recorder.Code)
		}
	})
}

func TestDeleteTask(t *testing.T) {
	t.Run("delete task", func(t *testing.T) {
		router := gin.Default()

		app, err := NewApp("./database.db", 3000)
		if err != nil {
			t.Fatalf("Error: %v", err)
		}

		router.DELETE("/tasks/:id", app.deleteTask)

		taskID := 123
		req, err := http.NewRequest("DELETE", "/tasks/"+strconv.Itoa(taskID), nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		if recorder.Code != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
		}
	})
}

func TestUpdateTask(t *testing.T) {
	t.Run("update task", func(t *testing.T) {
		router := gin.Default()

		app, err := NewApp("./database.db", 3000)
		if err != nil {
			t.Fatalf("Error: %v", err)
		}

		router.PUT("/tasks", app.updateTask)

		updateTask := Task{
			Id:        123,
			Title:     "Updated Task",
			Completed: true,
		}
		requestBody, err := json.Marshal(updateTask)
		if err != nil {
			t.Fatalf("Failed to marshal request JSON: %v", err)
		}

		req, err := http.NewRequest("PUT", "/tasks", bytes.NewBuffer(requestBody))
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		if recorder.Code != http.StatusCreated {
			t.Errorf("Expected status code %d, but got %d", http.StatusCreated, recorder.Code)
		}
	})
	t.Run("test bad request ", func(t *testing.T) {
		router := gin.Default()

		app, err := NewApp("./database.db", 3000)
		if err != nil {
			t.Fatalf("Error: %v", err)
		}
		defer os.Remove("./database.db")

		router.PUT("/tasks", app.updateTask)

		req, err := http.NewRequest("PUT", "/tasks", nil)
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")

		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		if recorder.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, but got %d", http.StatusCreated, recorder.Code)
		}
	})
}
