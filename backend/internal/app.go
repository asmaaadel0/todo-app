package internal

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
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

	router := gin.Default()

	// err := app.createTable()
	// if err != nil {
	// 	return err
	// }

	// Configure CORS middleware with desired options
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true, // Replace with false to restrict allowed origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	router.GET("/tasks", app.GetTasks)
	router.POST("/tasks", app.AddTask)
	router.DELETE("/tasks/:id", app.DeleteTask)
	router.PUT("/tasks", app.UpdateTask)

	router.Run(":3000")

	fmt.Println("Server started on port: 3000")
	err := http.ListenAndServe(":3000", nil)
	return err
}
