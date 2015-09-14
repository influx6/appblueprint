package main

import (
	"log"

	"github.com/influx6/appblueprint/app"
)

func main() {

	if err := app.App.Load("./config/app.yaml"); err != nil {
		if app.App.Env == "dev" {
			panic(err.Error())
		}
	}

	log.Printf("%s is now up", app.App.Name)
	app.App.Serve()
}
