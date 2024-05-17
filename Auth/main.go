package main

import (
	"net/http"

	"github.com/mars1385/e-com-go/auth/config"
	"github.com/mars1385/e-com-go/auth/database"
	logger "github.com/mars1385/e-com-go/auth/helper"
	"github.com/mars1385/e-com-go/auth/router"
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
