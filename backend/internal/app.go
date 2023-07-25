package internal

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/cors"
)

type App struct {
	db *sql.DB
}

// NewApp to create and initialize app
func NewApp() (*App, error) {
	database, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		return nil, err
	}

	return &App{db: database}, nil
}

func (app *App) Run() error {

	router := gin.New()

	// err := app.createTable()
	// if err != nil {
	// 	return err
	// }

	router.GET("/tasks", app.GetTasks)
	router.POST("/tasks", app.AddTask)
	router.DELETE("/tasks/:id", app.DeleteTask)
	router.PUT("/tasks", app.UpdateTask)

	c := cors.Default()

	handler := c.Handler(router)
	http.Handle("/", handler)

	fmt.Println("Server started on port: 3000")
	err := http.ListenAndServe(":3000", nil)
	return err
}
