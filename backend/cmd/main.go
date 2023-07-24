package main

import (
	"log"

	"github.com/codescalersinternships/todoapp-Asmaa/internal"
)

func main() {

	app := internal.NewApp(3000)

	err := app.Run()
	if err != nil {
		log.Fatalf("Error:%s", err)
	}

}
