package main

import (
	"log"

	"github.com/codescalersinternships/todoapp-Asmaa/internal"
)

func main() {

	app, err := internal.NewApp()
	if err != nil {
		log.Fatalf("Error:%s", err)
	}

	err = app.Run()
	if err != nil {
		log.Fatalf("Error:%s", err)
	}

}
