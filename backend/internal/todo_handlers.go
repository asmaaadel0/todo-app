package internal

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
// @Success 200 {array} Task
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /tasks [get]
func (app *App) getTasks(context *gin.Context) {

	context.Writer.Header().Set("Content-Type", "application/json")
	context.Next()
	tasks, err := app.client.getTasksDB()
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusNotFound, err)
		return
	}
	context.JSON(http.StatusOK, &tasks)
}

// AddTask adds a new task to the list.
// @Summary Add a new task
// @Description Add a new task to the list
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body Task true "New task object"
// @Success 201 {object} Task
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Internal Server Error"
// @Router /tasks [post]
func (app *App) addTask(context *gin.Context) {

	var newTask Task
	if err := context.BindJSON(&newTask); err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, err)
		return
	}

	respose, err := app.client.addTaskDB(newTask.Title, newTask.Completed)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	context.JSON(http.StatusCreated, respose)
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
func (app *App) deleteTask(context *gin.Context) {

	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, err)
		return
	}
	err = app.client.deleteTaskDB(id)
	if err != nil {
		log.Println(err)
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
func (app *App) updateTask(context *gin.Context) {

	var updateTask Task
	if err := context.BindJSON(&updateTask); err != nil {
		log.Println(err)
		context.JSON(http.StatusBadRequest, err)
		return
	}
	err := app.client.updateTaskDB(updateTask)
	if err != nil {
		log.Println(err)
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	context.JSON(http.StatusCreated, "Successfully updated")
}
