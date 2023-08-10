package main

import (
	"flag"
	"log"

	"github.com/codescalersinternships/todoapp-Asmaa/internal"
)

func main() {

	var databasePath string
	var port int
	flag.StringVar(&databasePath, "db", "database.db", "database path")
	flag.IntVar(&port, "p", 3000, "port")

	app, err := internal.NewApp(databasePath, port)
	if err != nil {
		log.Fatalf("Error:%s", err)
	}

	err = app.Run()
	if err != nil {
		log.Fatalf("Error:%s", err)
	}

}
