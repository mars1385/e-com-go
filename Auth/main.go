package main

import (
	"E-COM/config"
	logger "E-COM/helper"
	"E-COM/router"
	"net/http"
)

func main() {

	if err := config.SetupConfig(); err != nil {
		logger.Fatalf("config SetupConfig() error: %s", err)
	}

	router := router.SetupRoute()
	logger.Fatalf("%v", http.ListenAndServe(config.ServerConfig(), router))

}
