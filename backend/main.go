package main

import (
	"flag"
	"log"

	"github.com/codescalersinternships/todoapp-Asmaa/internal"
)

func main() {

	var databasePath string
	flag.StringVar(&databasePath, "db", "database.db", "database path")

	app, err := internal.NewApp(databasePath)
	if err != nil {
		log.Fatalf("Error:%s", err)
	}

	err = app.Run()
	if err != nil {
		log.Fatalf("Error:%s", err)
	}

}
