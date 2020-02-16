package main

import (
	"github.com/mikzuit/api/config"
	"github.com/mikzuit/api/app"
)

func main() {
	config := config.GetConfig()

	// Initialize app in specific port
	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
