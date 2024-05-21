package main

import (
	"log"

	"github.com/MJU-Capstone-6/devmark-frontend/internal/app"
)

func main() {
	application, err := app.InitApplication()
	if err != nil {
		log.Fatal(err)
	}
	application.Run()
}
