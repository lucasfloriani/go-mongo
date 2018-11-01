package db

import (
	"context"
	"log"

	"github.com/lucasfloriani/go-mongo/app"

	"github.com/mongodb/mongo-go-driver/mongo"
)

// Connect to database and returns the connection
func Connect() *mongo.Database {
	databaseName, connectionString := selectDatabase()
	client, err := mongo.NewClient(connectionString)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return client.Database(databaseName)
}

// selectDatabase by connection environment from flag option
func selectDatabase() (string, string) {
	switch app.Config.Environment {
	case "test":
		test := app.Config.Database.Test
		return test.Database, test.Connection
	case "development":
		dev := app.Config.Database.Development
		return dev.Database, dev.Connection
	default:
		prod := app.Config.Database.Production
		return prod.Database, prod.Connection
	}
}
