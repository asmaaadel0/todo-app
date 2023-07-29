package internal

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Task is a struct representing a task
type Task struct {
	// Id uniqe for each task
	Id int `json:"id"`
	// Title descripe each task
	Title string `json:"title"`
	// Completed if it's complete or not
	Completed bool `json:"completed"`
}

// @Summary Get all tasks
// @Description Get a list of all tasks
// @Tags tasks
// @Produce json
// @Success 200 {array} Task
// @Router /tasks [get]
func (app *App) GetTasks(context *gin.Context) {

	context.Writer.Header().Set("Content-Type", "application/json")
	context.Next()
	tasks, err := app.getTasks()
	if err != nil {
		context.JSON(http.StatusOK, Task{})
		return
	}
	context.JSON(http.StatusAccepted, &tasks)
}

// @Summary Add a new task
// @Description Add a new task to the list
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body Task true "New task object"
// @Success 200 {object} Task
// @Router /tasks [post]
func (app *App) AddTask(context *gin.Context) {

	var newTask Task
	if err := context.BindJSON(&newTask); err != nil {
		context.JSON(http.StatusBadRequest, err)
	}

	respose, err := app.addTask(newTask.Title, newTask.Completed)
	if err != nil {
		context.JSON(http.StatusInternalServerError, Task{})
		return
	}
	context.JSON(http.StatusOK, respose)
}

// @Summary Delete a task by ID
// @Description Delete a task by its ID
// @Produce json
// @Param id path int true "Task ID"
// @Success 200 {string} string "Successfully deleted"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /tasks/{id} [delete]
func (app *App) DeleteTask(context *gin.Context) {

	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	respose, err := app.deleteTask(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, Task{})
		return
	}
	context.JSON(http.StatusOK, respose)
}

// UpdateTask updates a task.
// @Summary Update a task
// @Description Update a task with new information
// @Produce json
// @Param task body Task true "Task object that needs to be updated"
// @Success 201 {string} string "Successfully updated"
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /tasks [put]
func (app *App) UpdateTask(context *gin.Context) {

	var updateTask Task
	if err := context.BindJSON(&updateTask); err != nil {
		context.JSON(http.StatusBadRequest, err)
	}
	respose, err := app.updateTask(updateTask)
	if err != nil {
		context.JSON(http.StatusInternalServerError, Task{})
		return
	}
	context.JSON(http.StatusCreated, respose)
}
