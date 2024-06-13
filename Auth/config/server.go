package config

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/mars1385/e-com-go/auth/ent"
	logger "github.com/mars1385/e-com-go/auth/helper"
	"github.com/mars1385/e-com-go/auth/router"
	"github.com/spf13/viper"
)

type ServerConfiguration struct {
	Port   string
	Secret string

	JWT_ACCESS_SECRET  string
	JWT_REFRESH_SECRET string
	JWT_ACCESS_TIMER   string
	JWT_REFRESH_TIMER  string

	DB_NAME     string
	DB_USER     string
	DB_PASSWORD string
	DB_HOST     string
	DB_SSL      string
	DB_DEBUG    bool
}

func ServerConfig() string {
	viper.SetDefault("SERVER_HOST", "0.0.0.0")
	viper.SetDefault("SERVER_PORT", "8000")

	appServer := fmt.Sprintf("%s:%s", viper.GetString("SERVER_HOST"), viper.GetString("SERVER_PORT"))
	log.Print("Server Running at :", appServer)
	return appServer
}

type Application struct {
	DB *ent.Client
}

func (app *Application) Serve() {
	// start http server

	router := router.SetupRoute()
	srv := &http.Server{
		Addr:    ServerConfig(),
		Handler: router,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func (app *Application) ListenForShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	app.shutdown()
	os.Exit(0)
}

func (app *Application) shutdown() {
	// perform any cleanup tasks

	app.DB.Close()
	logger.Infof("shutting down application..")

}
