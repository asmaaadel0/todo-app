package main

import (
	"fmt"
	"net/http"

	"github.com/codescalersinternships/todoapp-Asmaa/internal"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/tasks", internal.GetTasks).Methods("GET")
	r.HandleFunc("/tasks", internal.AddTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", internal.DeleteTask).Methods("DELETE")
	r.HandleFunc("/tasks/{id}", internal.UpdateTask).Methods("PUT")

	c := cors.Default()

	handler := c.Handler(r)
	http.Handle("/", handler)

	fmt.Println("Server started on port 3000")
	http.ListenAndServe(":3000", nil)
}
