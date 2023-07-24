package internal

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type App struct {
	port int
}

func NewApp(port int) *App {
	return &App{port: port}
}

func (app *App) Run() error {

	r := mux.NewRouter()

	r.HandleFunc("/tasks", getTasks).Methods("GET")
	r.HandleFunc("/tasks", addTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
	r.HandleFunc("/tasks/{id}", updateTask).Methods("PUT")

	c := cors.Default()

	handler := c.Handler(r)
	http.Handle("/", handler)

	portListner := fmt.Sprintf(":%d", app.port)
	fmt.Println("Server started on port", portListner)
	err := http.ListenAndServe(portListner, nil)
	return err
}
