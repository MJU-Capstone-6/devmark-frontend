package main

import (
	"log"

	"github.com/MJU-Capstone-6/devmark-frontend/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
