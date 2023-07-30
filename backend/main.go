package main

import (
	"flag"
	"log"

	"github.com/codescalersinternships/todoapp-Asmaa/internal"
)

func main() {

	var path string
	flag.StringVar(&path, "db", "database.db", "database path")

	app, err := internal.NewApp(path)
	if err != nil {
		log.Fatalf("Error:%s", err)
	}

	err = app.Run()
	if err != nil {
		log.Fatalf("Error:%s", err)
	}

}
