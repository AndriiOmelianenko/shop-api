package main

import (
	"log"

	"github.com/AndriiOmelianenko/shop-api/actions"
)

// main is the starting point to shop-api application.
func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
