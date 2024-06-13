package main

import (
	"github.com/mars1385/e-com-go/auth/config"
	"github.com/mars1385/e-com-go/auth/database"
	logger "github.com/mars1385/e-com-go/auth/helper"
)

func main() {

	if err := config.SetupConfig(); err != nil {
		logger.Fatalf("config SetupConfig() error: %s", err)
	}

	client, err := database.DbConnection()

	if err != nil {
		logger.Fatalf("database DbConnection error: %s", err)
	}

	app := config.Application{
		DB: client,
	}

	// listen for signals
	go app.ListenForShutdown()

	app.Serve()
	// defer database.Database.Close()

	// logger.Fatalf("%v", http.ListenAndServe(config.ServerConfig(), router))

}
