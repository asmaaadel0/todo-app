package internal

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	_ "github.com/codescalersinternships/todoapp-Asmaa/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// ErrOutOfRange if user enter invalid port
var ErrOutOfRange = errors.New("port number out of range, range should be between [1, 65535]")

type App struct {
	router *gin.Engine
	client DBClient
	port   int
}

// NewApp to create and initialize app
func NewApp(databasePath string, port int) (*App, error) {
	app := &App{}

	err := app.client.connectDatabase(databasePath)
	if err != nil {
		return nil, err
	}

	err = app.client.createTable()
	if err != nil {
		return nil, err
	}

	if port < 1 || port > 65535 {
		return nil, ErrOutOfRange
	}

	app.port = port

	app.router = gin.Default()
	return app, nil
}

// Run runs functions of app
func (app *App) Run() error {

	app.router.Use(cors.Default())

	app.router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	app.router.GET("/tasks", app.getTasks)
	app.router.POST("/tasks", app.addTask)
	app.router.DELETE("/tasks/:id", app.deleteTask)
	app.router.PUT("/tasks/:id", app.updateTask)

	portListner := fmt.Sprintf(":%d", app.port)

	if err := app.router.Run(portListner); err != nil {
		return err
	}

	log.Println("Server started on port", portListner)

	return http.ListenAndServe(portListner, nil)
}
