package internal

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	_ "github.com/codescalersinternships/todoapp-Asmaa/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// ErrorOutOfRange if user enter invalid port
var ErrorOutOfRange = errors.New("port number out of range, range should be between [1, 65535]")

type App struct {
	router *gin.Engine
	db     *sql.DB
	port   int
}

// NewApp to create and initialize app
func NewApp(databasePath string, port int) (*App, error) {
	app := &App{}

	err := app.connectDatabase(databasePath)
	if err != nil {
		return nil, err
	}
	if port < 1 || port > 65535 {
		return nil, ErrorOutOfRange
	}

	app.port = port

	app.router = gin.Default()
	return app, nil
}

// Run runs functions of app
func (app *App) Run() error {

	app.router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	app.router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	app.router.GET("/tasks", app.getTasks)
	app.router.POST("/tasks", app.addTask)
	app.router.DELETE("/tasks/:id", app.deleteTask)
	app.router.PUT("/tasks", app.updateTask)

	portListner := fmt.Sprintf(":%d", app.port)

	err := app.router.Run(portListner)
	if err != nil {
		return err
	}

	fmt.Println("Server started on port", portListner)
	err = http.ListenAndServe(portListner, nil)
	return err
}
