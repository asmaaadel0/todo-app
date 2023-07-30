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

// ErrorResponse represents an error response.
type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// GetTasks retrieves a list of all tasks.
// @Summary Get all tasks
// @Description Get a list of all tasks
// @Tags tasks
// @Produce json
// @Success 202 {array} Task
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /tasks [get]
func (app *App) GetTasks(context *gin.Context) {

	context.Writer.Header().Set("Content-Type", "application/json")
	context.Next()
	tasks, err := app.getTasks()
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	context.JSON(http.StatusAccepted, &tasks)
}

// AddTask adds a new task to the list.
// @Summary Add a new task
// @Description Add a new task to the list
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body Task true "New task object"
// @Success 200 {object} Task
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /tasks [post]
func (app *App) AddTask(context *gin.Context) {

	var newTask Task
	if err := context.BindJSON(&newTask); err != nil {
		context.JSON(http.StatusBadRequest, err)
	}

	respose, err := app.addTask(newTask.Title, newTask.Completed)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	context.JSON(http.StatusOK, respose)
}

// DeleteTask delete a task from list.
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
		context.JSON(http.StatusBadRequest, err)
		return
	}
	err = app.deleteTask(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	context.JSON(http.StatusOK, "Successfully deleted")
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
	err := app.updateTask(updateTask)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	context.JSON(http.StatusCreated, "Successfully updated")
}
