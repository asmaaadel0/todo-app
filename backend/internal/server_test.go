package internal

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// )

// func TestGetTasks(t *testing.T) {
// 	t.Run("Get tasks", func(t *testing.T) {
// 		req, err := http.NewRequest("GET", "/tasks", nil)
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		rr := httptest.NewRecorder()

// 		handler := http.HandlerFunc(GetTasks)
// 		handler.ServeHTTP(rr, req)

// 		if status := rr.Code; status != http.StatusOK {
// 			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
// 		}
// 		var responseTasks []Task
// 		err = json.NewDecoder(rr.Body).Decode(&responseTasks)
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		expectedTasks := []Task{}

// 		if len(responseTasks) != len(expectedTasks) {
// 			t.Errorf("handler returned wrong number of tasks: got %d want %d", len(responseTasks), len(expectedTasks))
// 			return
// 		}
// 	})
// }

// func TestAddTask(t *testing.T) {
// 	t.Run("add task", func(t *testing.T) {
// 		newTask := Task{
// 			Id:        1,
// 			Title:     "New Task",
// 			Completed: false,
// 		}
// 		newTaskJSON, err := json.Marshal(newTask)
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		req, err := http.NewRequest("POST", "/tasks", bytes.NewBuffer(newTaskJSON))
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		rr := httptest.NewRecorder()

// 		handler := http.HandlerFunc(AddTask)
// 		handler.ServeHTTP(rr, req)

// 		if status := rr.Code; status != http.StatusOK {
// 			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
// 		}

// 		var responseTask Task
// 		err = json.NewDecoder(rr.Body).Decode(&responseTask)
// 		if err != nil {
// 			t.Fatal(err)
// 		}

// 		if responseTask.Title != newTask.Title {
// 			t.Errorf("handler returned wrong task title: got %s want %s", responseTask.Title, newTask.Title)
// 		}
// 	})
// }
