package internal

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetTasks(t *testing.T) {
	router := gin.Default()

	app, err := NewApp("database.db")
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	router.GET("/tasks", app.GetTasks)

	req, err := http.NewRequest("GET", "/tasks", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusAccepted {
		t.Errorf("Expected status code %d, but got %d", http.StatusAccepted, recorder.Code)
	}

	var tasks []Task
	if err := json.Unmarshal(recorder.Body.Bytes(), &tasks); err != nil {
		t.Fatalf("Failed to unmarshal response JSON: %v", err)
	}
}

func TestAddTask(t *testing.T) {
	router := gin.Default()

	app, err := NewApp("database.db")
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	router.POST("/tasks", app.AddTask)

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

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
	}
}

func TestDeleteTask(t *testing.T) {
	router := gin.Default()

	app, err := NewApp("database.db")
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	router.DELETE("/tasks/:id", app.DeleteTask)

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
}

func TestUpdateTask(t *testing.T) {
	router := gin.Default()

	app, err := NewApp("database.db")
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	router.PUT("/tasks", app.UpdateTask)

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
}
