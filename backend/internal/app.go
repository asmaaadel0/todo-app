package internal

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/codescalersinternships/todoapp-Asmaa/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type App struct {
	router      *gin.Engine
	db          *sql.DB
	sqlCommands []string
}

// NewApp to create and initialize app
func NewApp(databasePath, schemaPath string) (*App, error) {
	app := &App{}

	err := app.readSqlCommands(schemaPath)
	if err != nil {
		return nil, err
	}

	err = app.connectDatabase(databasePath)
	if err != nil {
		return nil, err
	}

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
	// docs route
	app.router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	app.router.GET("/tasks", app.GetTasks)
	app.router.POST("/tasks", app.AddTask)
	app.router.DELETE("/tasks/:id", app.DeleteTask)
	app.router.PUT("/tasks", app.UpdateTask)

	err := app.router.Run(":3000")
	if err != nil {
		return err
	}

	fmt.Println("Server started on port: 3000")
	err = http.ListenAndServe(":3000", nil)
	return err
}
