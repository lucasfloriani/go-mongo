package main

import (
	"fmt"

	"github.com/lucasfloriani/go-mongo/app"
	"github.com/lucasfloriani/go-mongo/db"
	"github.com/lucasfloriani/go-mongo/router"
)

func main() {
	// Loads configuration data
	if err := app.LoadConfig("./config"); err != nil {
		panic(fmt.Errorf("Invalid application configuration: %s", err))
	}

	// Connects to the database
	database := db.Connect()

	// Runs the server
	routers := router.Setup(database)
	routers.Start(fmt.Sprintf(":%v", app.Config.ServerPort))
}
