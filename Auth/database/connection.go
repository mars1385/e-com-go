package database

import (
	"AUTH/ent"
	logger "AUTH/helper"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var (
	Database *ent.Client
)

func DbConnection() error {

	dbHost := viper.GetString("DB_HOST")
	dbUser := viper.GetString("DB_USER")
	dbPassword := viper.GetString("DB_PASSWORD")
	dbName := viper.GetString("DB_NAME")
	dbSSl := viper.GetString("DB_SSL")
	dbDebug := viper.GetBool("DB_DEBUG")

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
		dbUser, dbPassword, dbName, dbSSl)

	client, err := ent.Open(dbHost, connectionString)

	if err != nil {

		logger.Errorf(err.Error())
		logger.Fatalf("Db connection error")
		return err
	}
	if dbDebug {
		client = client.Debug()
	}
	Database = client
	logger.Infof("Db connection success")

	return nil
}
