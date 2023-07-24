package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Task to store tasks
type Task struct {
	// Id uniqe for each task
	Id int `json:"id"`
	// Title descripe each task
	Title string `json:"title"`
	// IsChecked if it's complete or not
	IsChecked bool `json:"isChecked"`
}

var tasks []Task

// getTasks to retrieve tasks
func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// addTask to add new task
func addTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newTask Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	newTask.Id = len(tasks)
	tasks = append(tasks, newTask)
	json.NewEncoder(w).Encode(newTask)
}

// deleteTask to delete task
func deleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	taskIDStr := vars["id"]
	taskID, err := strconv.Atoi(taskIDStr)

	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var taskIndex int = -1
	for i, task := range tasks {
		if task.Id == taskID {
			taskIndex = i
			break
		}
	}

	if taskIndex == -1 {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	tasks = append(tasks[:taskIndex], tasks[taskIndex+1:]...)

	json.NewEncoder(w).Encode(map[string]string{"message": "Task deleted successfully"})
}

// updateTask to update task
func updateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("asssssssssssss")
	w.Header().Set("Content-Type", "application/json")

	var updatedTask Task
	err := json.NewDecoder(r.Body).Decode(&updatedTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	taskIDStr := vars["id"]
	taskID, err := strconv.Atoi(taskIDStr)
	fmt.Println(taskID)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var taskToUpdate *Task
	for i := range tasks {
		if tasks[i].Id == taskID {
			taskToUpdate = &tasks[i]
			break
		}
	}

	if taskToUpdate == nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	taskToUpdate.Title = updatedTask.Title
	taskToUpdate.IsChecked = updatedTask.IsChecked

	json.NewEncoder(w).Encode(taskToUpdate)
}
