package main

import (
	"AUTH/config"
	"AUTH/database"
	logger "AUTH/helper"
	"AUTH/router"
	"net/http"
)

func main() {

	if err := config.SetupConfig(); err != nil {
		logger.Fatalf("config SetupConfig() error: %s", err)
	}

	if err := database.DbConnection(); err != nil {
		logger.Fatalf("database DbConnection error: %s", err)
	}

	defer database.Database.Close()
	router := router.SetupRoute()
	logger.Fatalf("%v", http.ListenAndServe(config.ServerConfig(), router))

}
