package internal

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Task to store tasks
type Task struct {
	// Id uniqe for each task
	Id int `json:"id"`
	// Title descripe each task
	Title string `json:"title"`
	// Completed if it's complete or not
	Completed bool `json:"completed"`
}

// GetTasks to retrieve tasks
func (app *App) GetTasks(context *gin.Context) {

	context.Writer.Header().Set("Content-Type", "application/json")
	context.Next()
	tasks := app.getTasks()
	context.JSON(http.StatusAccepted, &tasks)
}

// AddTask to add new task
func (app *App) AddTask(context *gin.Context) {

	var newTask Task
	if err := context.BindJSON(&newTask); err != nil {
		context.JSON(http.StatusBadRequest, err)
	}

	app.addTask(newTask.Title, newTask.Completed)
	context.JSON(http.StatusCreated, newTask)
}

// DeleteTask to delete task
func (app *App) DeleteTask(context *gin.Context) {

	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	app.deleteTask(id)
	context.JSON(http.StatusOK, id)
}

// UpdateTask to add new task
func (app *App) UpdateTask(context *gin.Context) {

	var updateTask Task
	if err := context.BindJSON(&updateTask); err != nil {
		context.JSON(http.StatusBadRequest, err)
	}
	app.updateTask(updateTask)
	context.JSON(http.StatusCreated, updateTask)
}
