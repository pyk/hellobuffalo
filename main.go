package main

import (
	"log"

	"github.com/pyk/hellobuffalo/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
