package main

import (
	"AUTH/config"
	logger "AUTH/helper"
	"AUTH/router"
	"net/http"
)

func main() {

	if err := config.SetupConfig(); err != nil {
		logger.Fatalf("config SetupConfig() error: %s", err)
	}

	router := router.SetupRoute()
	logger.Fatalf("%v", http.ListenAndServe(config.ServerConfig(), router))

}
