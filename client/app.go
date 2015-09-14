package main

import (
	"github.com/influx6/relay/engine"
)

// App represents a basic application
type App struct {
	*engine.ClientConfig
}

// NewApp returns a new instance of the app
func NewApp(c *engine.ClientConfig) *App {
	return &App{c}
}

func main() {

	config := engine.NewClientConfig()
	config.Name = "app.client"
	config.Templates.Dir = "../templates"

	app := NewApp(config)

	_ = app
}
